// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/cli"
)

// The incremental weekly pipeline: from a single discovery crawl, refresh only
// the types whose sanitized schema bytes changed (New or Changed), generate them
// for real from the fresh bytes, and let generation success/failure drive the
// policy decision written back to the overlay. Generation is the gate: if a type
// still generates, it is refreshed and its cache promoted; if it breaks, its
// last-good output and cache are kept (never regress) and the policy freezes or
// suppresses it so the release is never blocked.

// candidate is a type selected for regeneration because its discovered bytes
// differ from the cache (New or Changed). It carries the fresh bytes plus the
// row and class that drive generation and the policy decision.
type candidate struct {
	cfType string
	class  changeClass
	row    resourceRow
	schema []byte // freshly discovered, sanitized bytes to generate from
}

// buildCandidates selects the New+Changed types and attaches, for each, its
// policy class (new to the overlay vs already present), the row to generate from
// (the overlay row when present so its suppress_* flags are honored, else the
// discovered row), and the fresh discovered bytes.
func buildCandidates(results []changeResult, discByCFN map[string]discovered, overlayByCFN map[string]resourceRow) []candidate {
	var out []candidate
	for _, r := range results {
		if r.status != statusNew && r.status != statusChanged {
			continue
		}
		d := discByCFN[r.cfType]
		c := candidate{cfType: r.cfType, schema: d.schema}
		if row, ok := overlayByCFN[r.cfType]; ok {
			c.class = classPresent
			c.row = row
		} else {
			c.class = classNew
			c.row = d.row
		}
		out = append(out, c)
	}
	return out
}

// gateResultFromGenResults converts a candidate's real generation outcomes into
// a gateResult for the policy engine: any artifact that failed to generate marks
// the type as broken, so decide() freezes (Present) or suppresses (New) it.
// Plan-level errors (empty artifact kind) count as failures too.
func gateResultFromGenResults(cfType string, results []genResult) gateResult {
	gr := gateResult{cfType: cfType}
	for _, r := range results {
		outcome := gateOK
		if r.err != nil {
			outcome = gateFailedGeneration
		}
		gr.artifacts = append(gr.artifacts, artifactResult{
			kind:    r.a.kind,
			outcome: outcome,
			err:     r.err,
		})
	}
	return gr
}

// refreshCandidate generates a candidate from its freshly discovered bytes,
// refreshCandidate generates a candidate from its freshly discovered bytes and
// stages every artifact it produces under stagingDir — mirroring cfg.outputRoot
// and cfg.cacheDir, never writing to the real tree. Each artifact (resource,
// singular data source, plural data source) succeeds or fails independently:
// an artifact that generates cleanly is staged regardless of the others; one
// that fails is simply not staged — a Present type keeps its last-good real
// file for that artifact (untouched here), a New type simply never gets one.
// The resource/plural ListResource coupling is preserved: if the plural data
// source did not, in the end, succeed this pass, the resource is regenerated
// without GenerateListResource before being staged (reconcileListResource).
//
// The schema cache is staged only when at least one artifact succeeded — if
// every artifact failed, there is nothing new to keep, and the caller (decide,
// via classPresent) treats the type as fully broken. Nothing here touches the
// real output tree or schema cache; promoteStaged does that once, after the
// whole candidate batch has been staged successfully (generation-punchlist.md
// item 12: batch atomicity). It returns the gateResult for the policy engine,
// reflecting the final per-artifact outcome (after any ListResource-driven
// regeneration), plus the genResults that were actually staged (r.err == nil)
// — the compile gate's fixpoint (item 1) needs each staged artifact's
// pathSuffix/codeFile to map a blamed real file back to the candidate and
// artifact that produced it.
func refreshCandidate(cfg config, stagingDir string, c candidate) (gateResult, []genResult, error) {
	stagedSchema := schemaCachePath(filepath.Join(stagingDir, "input"), c.cfType)
	if err := os.MkdirAll(filepath.Dir(stagedSchema), dirPerm); err != nil {
		return gateResult{}, nil, fmt.Errorf("creating staging input dir: %w", err)
	}
	if err := os.WriteFile(stagedSchema, c.schema, filePerm); err != nil {
		return gateResult{}, nil, fmt.Errorf("staging %s: %w", c.cfType, err)
	}

	row := c.row
	row.CloudFormationSchemaPath = stagedSchema // generate from the staged bytes

	results := generateCorpus(cfg, []resourceRow{row}, 1, nil)
	results, err := reconcileListResource(cfg, results)
	if err != nil {
		return gateResult{}, nil, fmt.Errorf("regenerating resource for %s: %w", c.cfType, err)
	}
	gr := gateResultFromGenResults(c.cfType, results)

	// Stage successful artifacts under stagingDir/out, mirroring cfg.outputRoot.
	stageCfg := cfg
	stageCfg.outputRoot = filepath.Join(stagingDir, "out")
	var staged []genResult
	for _, r := range results {
		if r.err != nil {
			continue // leave last-good (Present) or nothing (New) for this artifact
		}
		if err := writeArtifact(stageCfg, r); err != nil {
			return gr, nil, fmt.Errorf("staging %s %s: %w", c.cfType, r.a.kind, err)
		}
		staged = append(staged, r)
	}
	if len(staged) == 0 {
		return gr, nil, nil // nothing succeeded; nothing to stage
	}

	// Stage the fresh schema bytes for promotion into the real cache too.
	stagedCache := filepath.Join(stagingDir, "cache")
	if err := os.MkdirAll(stagedCache, dirPerm); err != nil {
		return gr, staged, fmt.Errorf("creating staged cache dir: %w", err)
	}
	if err := os.WriteFile(schemaCachePath(stagedCache, c.cfType), c.schema, filePerm); err != nil {
		return gr, staged, fmt.Errorf("staging cache for %s: %w", c.cfType, err)
	}
	return gr, staged, nil
}

// promoteStaged copies every staged artifact and cached schema from stagingDir
// into the real tree (cfg.outputRoot, cfg.cacheDir). Called once, after every
// candidate in a batch has been staged successfully by refreshCandidate — this
// is the sole place -update writes outside stagingDir, so a hard error anywhere
// earlier in the batch (generation-punchlist.md item 12) leaves the real tree
// and the overlay untouched: nothing is promoted, and the overlay reconcile
// step that follows never even runs.
func promoteStaged(cfg config, stagingDir string) error {
	stagedOut := filepath.Join(stagingDir, "out")
	if err := copyTree(stagedOut, cfg.outputRoot); err != nil {
		return fmt.Errorf("promoting staged output: %w", err)
	}
	stagedCache := filepath.Join(stagingDir, "cache")
	if err := copyTree(stagedCache, cfg.cacheDir); err != nil {
		return fmt.Errorf("promoting staged cache: %w", err)
	}
	return nil
}

// copyTree copies every regular file under src into the identical relative
// path under dst, creating directories as needed. A missing src (nothing was
// staged there) is not an error.
func copyTree(src, dst string) error {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return nil
	}
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, dirPerm)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}
		if err := os.MkdirAll(filepath.Dir(target), dirPerm); err != nil {
			return err
		}
		return os.WriteFile(target, data, filePerm)
	})
}

// reconcileListResource re-generates the resource artifact without
// GenerateListResource if it was generated expecting a working plural data
// source that isn't there in the final result — either its generation failed
// this pass, or it was never attempted (already suppressed, so absent from
// results). A resource promoted with ListResource: true but no working plural
// data source would advertise a list resource with no backing data source.
func reconcileListResource(cfg config, results []genResult) ([]genResult, error) {
	resIdx := -1
	pluralOK := false
	for i, r := range results {
		switch r.a.kind {
		case artifactResource:
			resIdx = i
		case artifactPluralDataSource:
			pluralOK = r.err == nil
		}
	}
	if resIdx < 0 {
		return results, nil // no resource artifact this pass
	}
	r := results[resIdx]
	//nolint:nilerr // r.err here is the resource artifact's own outcome, not a
	// swallowed error: if the resource already failed on its own, there is
	// nothing to reconcile, so returning the results (including that error) is
	// correct, not an accidental nil-out.
	if !r.a.listResource || r.err != nil || pluralOK {
		return results, nil // didn't ask for ListResource, already failed, or plural is fine
	}

	// The resource generated expecting a working plural data source that isn't
	// there; regenerate it without ListResource before promoting.
	r.a.listResource = false
	ui := &cli.BasicUi{Writer: io.Discard, ErrorWriter: io.Discard}
	code, test, err := generateArtifact(ui, cfg, r.p, r.a)
	if err != nil {
		return nil, err
	}
	r.code, r.test, r.err = code, test, nil
	results[resIdx] = r
	return results, nil
}

// runUpdate is the live weekly incremental pipeline (-update). One discovery
// crawl feeds both overlay reconciliation and change detection; only New/Changed
// types are regenerated from their fresh bytes; generation success/failure drives
// the policy written back to the overlay (never regressing broken types); and the
// aggregates are re-emitted. Requires AWS credentials (queried in us-east-1).
func runUpdate(ctx context.Context, allSchemasPath, checkoutPath string) error {
	cfg, overlayRows, err := loadOverlay(allSchemasPath)
	if err != nil {
		return err
	}

	disc, err := discover(ctx)
	if err != nil {
		return fmt.Errorf("discovery: %w", err)
	}

	discByCFN := make(map[string]discovered, len(disc))
	base := make([]resourceRow, 0, len(disc))
	for _, d := range disc {
		discByCFN[d.row.CloudFormationTypeName] = d
		base = append(base, d.row)
	}
	overlayByCFN := make(map[string]resourceRow, len(overlayRows))
	frozenCFN := make(map[string]bool)
	for _, r := range overlayRows {
		overlayByCFN[r.CloudFormationTypeName] = r
		if r.FrozenSince != "" {
			frozenCFN[r.CloudFormationTypeName] = true
		}
	}

	changes, err := detectChanges(disc, frozenCFN, cfg.cacheDir)
	if err != nil {
		return err
	}
	summary := summarize(changes)
	stepf("Detected changes — new %d, changed %d, unchanged %d, frozen %d, missing %d.",
		summary.New, summary.Changed, summary.Unchanged, summary.Frozen, summary.Missing)

	cands := buildCandidates(changes, discByCFN, overlayByCFN)

	stagingDir, err := os.MkdirTemp("", "bigdiffer-staging-")
	if err != nil {
		return fmt.Errorf("creating staging dir: %w", err)
	}
	defer func() { _ = os.RemoveAll(stagingDir) }()

	stepf("Regenerating %d changed type(s) from fresh bytes (never-regress)…", len(cands))
	candBar := newBar(len(cands), "regenerate")
	today := time.Now().Format(dateLayout)
	decisions := make(map[string]policyDecision, len(cands))
	stagedByDest := make(map[string]stagedArtifact)
	var okN, brokeN int
	for _, c := range cands {
		gr, staged, err := refreshCandidate(cfg, stagingDir, c)
		if err != nil {
			return err
		}
		decisions[c.cfType] = decide(c.class, gr, today)
		if gr.ok() {
			okN++
		} else {
			brokeN++
			infof("%s: %s", c.cfType, decisions[c.cfType].summary)
		}
		// gr is shared by pointer with every staged artifact of this
		// candidate: the compile gate (below) downgrades one artifact's
		// outcome in place and recomputes decide() from the same gateResult,
		// so a build-gate rejection of, say, the plural data source does not
		// lose the fact that the resource and singular data source are still
		// fine.
		grPtr := gr
		for _, r := range staged {
			dest := filepath.Join(cfg.outputRoot, r.a.pathSuffix, r.a.codeFile)
			testDest := filepath.Join(cfg.outputRoot, r.a.pathSuffix, r.a.testFile)
			// Find this artifact's index within grPtr.artifacts (staged only
			// contains artifacts that generated OK, i.e. gateOK in grPtr).
			for ai := range grPtr.artifacts {
				if grPtr.artifacts[ai].kind == r.a.kind {
					stagedByDest[dest] = stagedArtifact{class: c.class, gr: &grPtr, artifact: ai, testDest: testDest}
					break
				}
			}
		}
		_ = candBar.Add(1)
	}
	_ = candBar.Finish()

	// Reconcile the overlay and apply the policy edits in one pass — read now,
	// rendered later (once the compile gate settles); recomputing decide()
	// from a fresher on-disk overlay mid-fixpoint would fight the batch it
	// already committed to when the candidate loop ran.
	overlayContent, err := os.ReadFile(allSchemasPath)
	if err != nil {
		return fmt.Errorf("reading overlay %s: %w", allSchemasPath, err)
	}
	checkout, err := parseCheckout(checkoutPath)
	if err != nil {
		return fmt.Errorf("reading checkout %s: %w", checkoutPath, err)
	}

	// Compile gate (generation-punchlist.md item 1): build the staged code plus
	// the registration file it implies against the real module until it
	// compiles clean, downgrading whatever the compiler rejects. Runs before
	// promotion so a build failure changes what gets promoted rather than
	// promoting something broken; never blocks the release (design doc
	// contributing/docs/compile-gate-design.md).
	stepf("Compile-gating %d staged artifact(s)…", len(stagedByDest))
	if err := compileFixpoint(cfg, stagingDir, string(overlayContent), base, checkout, decisions, stagedByDest, today); err != nil {
		return fmt.Errorf("compile gate: %w", err)
	}

	// Promote the compiled core — staged code, cache, the reconciled overlay,
	// and the gated registration file — together, only now that the whole
	// batch has staged successfully and gone green. A hard error anywhere
	// earlier (the candidate loop or the compile gate) returned before
	// reaching here, so the real tree and the overlay are left exactly as they
	// started (generation-punchlist.md item 12).
	stepf("Promoting staged output (%d artifact(s) refreshed)…", okN+brokeN)
	if err := promoteStaged(cfg, stagingDir); err != nil {
		return err
	}

	stepf("Reconciling all_schemas.hcl and applying policy…")
	out, report, err := normalizeWithDecisions(string(overlayContent), base, nil, checkout, decisions)
	if err != nil {
		return err
	}
	if err := os.WriteFile(allSchemasPath, []byte(out), filePerm); err != nil {
		return fmt.Errorf("writing %s: %w", allSchemasPath, err)
	}

	// Emit import_examples_gen.json from the promoted state: unlike the
	// registration file, it reads each resource's cached schema bytes to
	// recover primary-identifier names, so it must run after the cache is
	// promoted — it is not part of the compiled/gated/atomic core, it trails
	// it as a deterministic projection of the final, promoted overlay + cache.
	stepf("Emitting import examples…")
	_, finalRows, err := loadOverlay(allSchemasPath)
	if err != nil {
		return err
	}
	ie, err := emitImportExamples(cfg, finalRows)
	if err != nil {
		return err
	}
	if err := os.WriteFile(cfg.importExamplesPath, ie, filePerm); err != nil {
		return fmt.Errorf("writing %s: %w", cfg.importExamplesPath, err)
	}

	report.write()
	// Recompute the summary counts from the settled decisions, not the
	// mid-loop tally above: the compile gate can downgrade a candidate that
	// generation reported OK, so okN/brokeN must reflect what was actually
	// promoted, not just what generation produced before the gate ran.
	finalOK, finalBroken := 0, 0
	for _, d := range decisions {
		if d.reason == "" {
			finalOK++
		} else {
			finalBroken++
		}
	}
	stepf("Done: refreshed %d changed type(s) — %d generated OK, %d frozen/suppressed.", len(cands), finalOK, finalBroken)
	infof("Review `git status`/`git diff`, then: `make build`, `make smoke`, `go run ./internal/tools/bigdiffer -docs`.")
	return nil
}
