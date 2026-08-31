// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"time"
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
// staged in a temp schema file so a failed generation never overwrites the
// good cached schema. Only when every artifact generates cleanly does it promote
// the outputs — write the generated files to their real paths and the fresh
// bytes to the schema cache. A failure leaves the last-good files and cache
// untouched (never regress). It returns the gateResult for the policy engine.
func refreshCandidate(cfg config, stagingDir string, c candidate) (gateResult, error) {
	stagedSchema := schemaCachePath(stagingDir, c.cfType)
	if err := os.WriteFile(stagedSchema, c.schema, 0o644); err != nil {
		return gateResult{}, fmt.Errorf("staging %s: %w", c.cfType, err)
	}

	row := c.row
	row.CloudFormationSchemaPath = stagedSchema // generate from the staged bytes

	results := generateCorpus(cfg, []resourceRow{row}, 1, nil)
	gr := gateResultFromGenResults(c.cfType, results)
	if !gr.ok() {
		return gr, nil // keep last-good output and cache
	}

	if _, err := writeCorpus(cfg, results); err != nil {
		return gr, fmt.Errorf("writing %s: %w", c.cfType, err)
	}
	if err := os.MkdirAll(cfg.cacheDir, 0o755); err != nil {
		return gr, fmt.Errorf("creating cache dir: %w", err)
	}
	if err := os.WriteFile(schemaCachePath(cfg.cacheDir, c.cfType), c.schema, 0o644); err != nil {
		return gr, fmt.Errorf("promoting cache for %s: %w", c.cfType, err)
	}
	return gr, nil
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
	defer os.RemoveAll(stagingDir)

	stepf("Regenerating %d changed type(s) from fresh bytes (never-regress)…", len(cands))
	candBar := newBar(len(cands), "regenerate")
	today := time.Now().Format(dateLayout)
	decisions := make(map[string]policyDecision, len(cands))
	var okN, brokeN int
	for _, c := range cands {
		gr, err := refreshCandidate(cfg, stagingDir, c)
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
		_ = candBar.Add(1)
	}
	_ = candBar.Finish()

	// Reconcile the overlay and apply the policy edits in one pass.
	stepf("Reconciling all_schemas.hcl and applying policy…")
	overlayContent, err := os.ReadFile(allSchemasPath)
	if err != nil {
		return fmt.Errorf("reading overlay %s: %w", allSchemasPath, err)
	}
	checkout, err := parseCheckout(checkoutPath)
	if err != nil {
		return fmt.Errorf("reading checkout %s: %w", checkoutPath, err)
	}
	out, report, err := normalizeWithDecisions(string(overlayContent), base, nil, checkout, decisions)
	if err != nil {
		return err
	}
	if err := os.WriteFile(allSchemasPath, []byte(out), 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", allSchemasPath, err)
	}

	// Re-emit the aggregates from the reconciled overlay.
	stepf("Emitting aggregates (registration + import examples)…")
	_, finalRows, err := loadOverlay(allSchemasPath)
	if err != nil {
		return err
	}
	reg, err := emitRegistration(cfg, finalRows)
	if err != nil {
		return err
	}
	if err := os.WriteFile(cfg.registrationPath, reg, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", cfg.registrationPath, err)
	}
	ie, err := emitImportExamples(cfg, finalRows)
	if err != nil {
		return err
	}
	if err := os.WriteFile(cfg.importExamplesPath, ie, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", cfg.importExamplesPath, err)
	}

	report.write(os.Stderr)
	stepf("Done: refreshed %d changed type(s) — %d generated OK, %d frozen/suppressed.", len(cands), okN, brokeN)
	infof("Review `git status`/`git diff`, then: `make build`, `make smoke`, `go run ./internal/tools/bigdiffer -docs`.")
	return nil
}
