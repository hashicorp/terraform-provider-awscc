// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
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

// buildCandidates selects every discovered type as a gate candidate — not just
// New/Changed (contributing/docs/bigdiffer-design.md §6, "The whole corpus is
// gated every run": gating only New/Changed left the ~1580 schema-unchanged
// types' committed output aging silently, with no guard once the legacy
// TestFullCorpusParity retires). For
// each, it attaches the policy class (new to the overlay vs already present,
// further split by whether its schema bytes actually moved) and the row to
// generate from (the overlay row when present so its suppress_* flags are
// honored, else the discovered row), and the fresh discovered bytes.
//
// Class assignment is keyed on two independent facts, not one: overlay
// presence decides New vs Present; byte status further splits Present into
// classPresent (statusChanged — the schema moved, so a failure freezes/
// suppresses as always) vs classPresentUnchanged (statusUnchanged — the
// schema is identical, so a failure can only be the machinery's fault; see
// decide()'s doc comment, policy.go). A statusChanged type absent from the
// overlay (rare — e.g. a previously checkout-pinned type) still lands on
// classNew exactly as before: byte status only refines Present, it never
// overrides New.
func buildCandidates(results []changeResult, discByCFN map[string]discovered, overlayByCFN map[string]resourceRow) []candidate {
	var out []candidate
	for _, r := range results {
		if r.status != statusNew && r.status != statusChanged && r.status != statusUnchanged {
			continue
		}
		d := discByCFN[r.cfType]
		c := candidate{cfType: r.cfType, schema: d.schema}
		if row, ok := overlayByCFN[r.cfType]; ok {
			c.class = classPresent
			if r.status == statusUnchanged {
				c.class = classPresentUnchanged
			}
			c.row = row
		} else {
			c.class = classNew
			c.row = d.row
		}
		out = append(out, c)
	}
	return out
}

// absentTypes returns the overlay's CloudFormation type names absent from this
// run's discovered set (O - A) that still need the probe: not already in A,
// and not already explained by a freeze or a checkout pin (a checkout pin may
// explain an absence the probe cannot see — a deliberately checked-out older
// version — so it is treated the same as an existing freeze). Pure and
// order-independent; the caller sorts if a stable probe order matters.
func absentTypes(overlayByCFN map[string]resourceRow, discByCFN map[string]discovered, frozenCFN map[string]bool, checkout map[string]bool) []string {
	var out []string
	for cfn := range overlayByCFN {
		if _, ok := discByCFN[cfn]; ok {
			continue // in A; not absent
		}
		if frozenCFN[cfn] || checkout[cfn] {
			continue // already explained; no probe needed
		}
		out = append(out, cfn)
	}
	return out
}

// absentDecisions resolves every type absentTypes returns via one DescribeType
// probe each, per contributing/docs/absent-row-probe-design.md. A type the
// probe cannot resolve definitively (probeAbsent returns "") is left out of the
// returned map entirely, so it is neither annotated nor blocked — just
// deferred to a later run, exactly like any other transient discovery gap.
func absentDecisions(ctx context.Context, conn *cloudformation.Client, overlayByCFN map[string]resourceRow, discByCFN map[string]discovered, frozenCFN map[string]bool, checkout map[string]bool, today string) map[string]policyDecision {
	decisions := make(map[string]policyDecision)
	for _, cfn := range absentTypes(overlayByCFN, discByCFN, frozenCFN, checkout) {
		class := probeAbsent(ctx, conn, cfn)
		if class == "" {
			continue // probe inconclusive; leave the row as-is for a later run
		}
		decisions[cfn] = decide(class, gateResult{}, today)
	}
	return decisions
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
// whole candidate batch has been staged successfully (never-regress batch
// atomicity). It returns the gateResult for the policy engine,
// reflecting the final per-artifact outcome (after any ListResource-driven
// regeneration), plus the genResults that were actually staged (r.err == nil)
// — the compile gate's fixpoint (item 1) needs each staged artifact's
// pathSuffix/codeFile to map a blamed real file back to the candidate and
// artifact that produced it.
//
// Generates each candidate in a subprocess (generateCandidateIsolated), not
// in-process: codegen.Emitter has no recursion-depth guard, so a recursive
// schema is a Go stack overflow, a fatal error no in-process recover() can
// catch. "One failure never blocks the rest" was always the intent here —
// this closes the gap where it didn't hold.
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

	results, err := generateCandidateIsolated(cfg, row)
	if err != nil {
		return gateResult{}, nil, fmt.Errorf("regenerating %s: %w", c.cfType, err)
	}
	results, err = reconcileListResource(cfg, row, results)
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
// is the sole place -sync writes outside stagingDir, so a hard error anywhere
// earlier in the batch (never-regress batch atomicity) leaves the real tree
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

// generateCandidateIsolated builds the generation plan for row and generates
// every artifact in one subprocess call (generateCandidateArtifacts), not
// one per artifact — cheaper at whole-corpus scale, at the cost that a crash
// in one artifact marks its siblings failed too (see runRecheckProbeArtifact's
// manifest-mode comment). A plan error (e.g. an unparseable Terraform type
// name) is reported in-band as a failed genResult, matching generateCorpus's
// own planErrs handling: it's a row problem to freeze/suppress, not an
// infrastructure error that should abort the batch.
//
// repoRoot/outputRoot are cleared before calling out: those fields opt into
// runRecheckProbeArtifact's single-artifact ad-hoc compile gate
// (-recheck's own use case), which generation-only calls must not trigger —
// left set, it silently created cfg.outputRoot on disk via buildOnce's
// overlay even when nothing was promoted (TestSyncBatchAtomicity).
func generateCandidateIsolated(cfg config, row resourceRow) ([]genResult, error) {
	p, err := generationPlan(row, cfg.prefix, cfg.cacheDir)
	if err != nil {
		return []genResult{{
			p:   plan{cfType: row.CloudFormationTypeName},
			err: fmt.Errorf("plan: %w", err),
		}}, nil
	}
	schema, err := os.ReadFile(p.schemaFile)
	if err != nil {
		return nil, fmt.Errorf("reading staged schema %s: %w", p.schemaFile, err)
	}
	genCfg := cfg
	genCfg.repoRoot = ""
	genCfg.outputRoot = ""
	results, err := generateCandidateArtifacts(genCfg, row, p.artifacts, schema)
	if err != nil {
		return nil, err
	}
	for i := range results {
		results[i].p = p
	}
	return results, nil
}

// reconcileListResource re-generates the resource artifact without
// GenerateListResource if it was generated expecting a working plural data
// source that isn't there in the final result — either its generation failed
// this pass, or it was never attempted (already suppressed, so absent from
// results). A resource promoted with ListResource: true but no working plural
// data source would advertise a list resource with no backing data source.
// row must be the same row refreshCandidate generated results from (its
// CloudFormationSchemaPath already points at the staged schema bytes) so the
// re-generation below reads the identical schema, not a fresh cache lookup.
func reconcileListResource(cfg config, row resourceRow, results []genResult) ([]genResult, error) {
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
	// there; regenerate it without ListResource before promoting, via the
	// same subprocess isolation (same schema, same crash risk as the first
	// attempt).
	//
	// The subprocess re-derives its own plan from row, so mutating a local
	// copy of r.a.listResource has no effect on it — forcing
	// SuppressPluralDataSourceGeneration true on the row is what actually
	// makes generationPlan derive listResource: false there.
	r.a.listResource = false
	noListResourceRow := row
	noListResourceRow.SuppressPluralDataSourceGeneration = true
	schema, err := os.ReadFile(r.p.schemaFile)
	if err != nil {
		return nil, fmt.Errorf("reading staged schema %s: %w", r.p.schemaFile, err)
	}
	genCfg := cfg
	genCfg.repoRoot = ""
	genCfg.outputRoot = ""
	code, test, genErr := generateArtifactIsolated(genCfg, noListResourceRow, r.a.kind, schema)
	if genErr != nil {
		return nil, genErr
	}
	r.code, r.test, r.err = code, test, nil
	results[resIdx] = r
	return results, nil
}

// runSync is the live weekly incremental pipeline (-sync). One discovery
// crawl feeds both overlay reconciliation and change detection; only New/Changed
// types are regenerated from their fresh bytes; generation success/failure drives
// the policy written back to the overlay (never regressing broken types); and the
// aggregates are re-emitted. Documentation (import-example docs, terraform fmt,
// tfplugindocs) runs as the last step of the same tail by default — noDocs
// (-no-docs) skips it for a fast codegen inner loop
// (bigdiffer-design.md §6, "Documentation is part of the same gated
// pipeline"). Requires AWS credentials (queried in us-east-1).
func runSync(ctx context.Context, allSchemasPath, checkoutPath string, noDocs bool) error {
	cfg, overlayRows, err := loadOverlay(allSchemasPath)
	if err != nil {
		return err
	}

	disc, conn, err := discover(ctx)
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
	checkout, err := parseCheckout(checkoutPath)
	if err != nil {
		return fmt.Errorf("reading checkout %s: %w", checkoutPath, err)
	}

	changes, err := detectChanges(disc, frozenCFN, cfg.cacheDir)
	if err != nil {
		return err
	}
	summary := summarize(changes)
	stepf("Detected changes — new %d, changed %d, unchanged %d, frozen %d, missing %d.",
		summary.New, summary.Changed, summary.Unchanged, summary.Frozen, summary.Missing)

	cands := buildCandidates(changes, discByCFN, overlayByCFN)

	stepf("Regenerating %d changed type(s) from fresh bytes (never-regress)…", len(cands))
	today := todayString()

	// Absent-row probe
	// (contributing/docs/absent-row-probe-design.md): resolve every overlay row
	// this run's crawl didn't see at all (O - A) before the candidate loop, so
	// its decisions share the same map — the two sets are disjoint by
	// construction (a candidate came from `changes`, which only ever covers
	// discovered/A types). No staging or generation is needed for an absent
	// row: it is a pure annotation of the block that already exists.
	baseDecisions := absentDecisions(ctx, conn, overlayByCFN, discByCFN, frozenCFN, checkout, today)

	// Reconcile the overlay and apply the policy edits in one pass — read now,
	// rendered later (once the compile gate settles); recomputing decide()
	// from a fresher on-disk overlay mid-fixpoint would fight the batch it
	// already committed to when the candidate loop ran.
	overlayContent, err := os.ReadFile(allSchemasPath)
	if err != nil {
		return fmt.Errorf("reading overlay %s: %w", allSchemasPath, err)
	}

	// The shared generate -> stage -> compile-gate pipeline
	// (contributing/docs/bigdiffer-design.md §6, "One engine, three
	// callers"): everything from generating each candidate through the
	// compile gate settling and the list-resource coupling check is common
	// to sync/reconcile/check, so it lives in settleBatch, not here.
	settled, err := settleBatch(ctx, cfg, cands, baseDecisions, string(overlayContent), base, checkout, today)
	if err != nil {
		return err
	}
	defer func() { _ = os.RemoveAll(settled.stagingDir) }()

	// The whole-corpus gate's one real behavior change (§2, §3): a
	// schema-unchanged failure is self-inflicted (the machinery, not the
	// schema, broke), and a machinery bug that produces broken output for one
	// type can equally produce compiling-but-subtly-wrong output for another
	// type this same run — so nothing from this run is trustworthy, and
	// nothing is promoted, not even the schema-changed types that compiled
	// cleanly (§3, "Why all-or-nothing, precisely"). compileFixpoint can
	// legitimately reach green here (a broken new artifact reverted to its
	// still-compiling committed file settles the fixpoint cleanly) while
	// still leaving a machineryFailure decision behind — checking
	// compileFixpoint's own error return is not enough; every failure must be
	// scanned for and reported before promotion, not just the first one
	// (loud, complete surfacing is the point of gating the whole corpus).
	if failures := machineryFailures(settled); len(failures) > 0 {
		var b strings.Builder
		fmt.Fprintf(&b, "%d type(s) failed with the schema unchanged — machinery regression, promoting nothing this run:\n", len(failures))
		for _, f := range failures {
			fmt.Fprintf(&b, "  %s\n", f.cfType)
			for _, a := range f.artifacts {
				fmt.Fprintf(&b, "    %s\n", a)
			}
		}
		return errors.New(b.String())
	}

	decisions := settled.decisions
	stagedByDest := settled.stagedByDest

	// CHANGELOG delta: stagedByDest is now final — every remaining entry
	// survived the compile gate and will actually be promoted below — so the
	// changelog delta itself can be computed here, directly from stagedByDest
	// plus the pre-run overlay, with no dependency on promotion having
	// happened yet (contributing/docs/bigdiffer-design.md §8). The fragment
	// is not WRITTEN until after every real-tree write below has succeeded
	// (the fragment is written last on purpose): writing it earlier would leave an
	// orphan CHANGELOG.md entry for a change that never actually landed if
	// promotion, the overlay write, or the import-examples write failed.
	changelog := changelogEntries(stagedByDest, overlayByCFN)

	// Promote the compiled core — staged code, cache, the reconciled overlay,
	// and the gated registration file — together, only now that the whole
	// batch has staged successfully and gone green. A hard error anywhere
	// earlier (the candidate loop or the compile gate) returned before
	// reaching here, so the real tree and the overlay are left exactly as they
	// started (never-regress batch atomicity).
	stepf("Promoting staged output (%d artifact(s) refreshed)…", settled.okN+settled.brokeN)
	if err := promoteStaged(cfg, settled.stagingDir); err != nil {
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

	// Write the CHANGELOG fragment (bigdiffer-design.md §8), now that every
	// other real-tree write above has actually succeeded. writeChangelogFragment
	// inserts the FEATURES: bullets directly into CHANGELOG.md's top,
	// in-progress version block; a no-op when nothing was newly promoted.
	// Docs (below) trails even this, so the CHANGELOG entry for
	// already-promoted code is written regardless of whether docs succeed.
	if err := writeChangelogFragment(cfg.changelogPath, changelog); err != nil {
		return err
	}
	if len(changelog) > 0 {
		stepf("Added %d CHANGELOG entries to %s.", len(changelog), cfg.changelogPath)
	}

	// Documentation, last in the tail (bigdiffer-design.md §6,
	// "Documentation is part of the same gated pipeline"): everything above
	// has already promoted, so a docs failure here is never a broken commit
	// — bigdiffer never commits on its own, and the promoted code changes
	// are left in place either way. See runDocsTail's doc comment for the
	// recovery-oriented error this returns on failure.
	//
	// The change report and promotion summary print before this, not after:
	// if docs fail, the run still returns an error below, but the human
	// running it has already seen exactly what code landed, rather than
	// losing that summary behind a docs-only failure.
	report.write()
	// Recompute the summary counts from the settled decisions, not the
	// mid-loop tally above: the compile gate can downgrade a candidate that
	// generation reported OK, so okN/brokeN must reflect what was actually
	// promoted, not just what generation produced before the gate ran.
	finalOK, finalBroken := 0, 0
	for _, d := range decisions {
		if len(d.reasons) == 0 {
			finalOK++
		} else {
			finalBroken++
		}
	}
	stepf("Refreshed %d changed type(s) — %d generated OK, %d frozen/suppressed.", len(cands), finalOK, finalBroken)

	if err := runDocsTail(cfg, noDocs); err != nil {
		return err
	}

	stepf("Done.")
	infof("Review `git status`/`git diff`, then: `make build`, `make smoke`.")
	return nil
}
