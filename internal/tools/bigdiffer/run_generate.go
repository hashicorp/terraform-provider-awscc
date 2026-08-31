// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"
	"time"
)

// runGenerate performs a full, offline regeneration of the provider from the
// committed overlay + schema cache: generate every artifact in parallel, write
// them, and emit the aggregates (the single registration file and
// import_examples_gen.json). This is bigdiffer's -generate mode — the parallel,
// in-process replacement for the legacy `make resources/…/schemas` directive
// crawl. It does not touch AWS.
func runGenerate(cfg config, rows []resourceRow) error {
	start := time.Now()

	stepf("Generating the provider from %d types (concurrency %d)…", len(rows), cfg.genConcurrency)
	genBar := newBar(countArtifacts(cfg, rows), "generate")
	results := generateCorpus(cfg, rows, cfg.genConcurrency, func() { _ = genBar.Add(1) })
	_ = genBar.Finish()

	var genErrs int
	for _, r := range results {
		if r.err != nil {
			genErrs++
			if genErrs <= 10 {
				fmt.Fprintf(os.Stderr, "bigdiffer: %s %s: %v\n", r.p.cfType, r.a.kind, r.err)
			}
		}
	}
	if genErrs > 0 {
		return fmt.Errorf("generation failed for %d artifact(s)", genErrs)
	}

	stepf("Writing generated files for %d artifacts…", len(results))
	n, err := writeCorpus(cfg, results)
	if err != nil {
		return err
	}

	stepf("Emitting aggregates (registration + import examples)…")
	reg, err := emitRegistration(cfg, rows)
	if err != nil {
		return err
	}
	if err := os.WriteFile(cfg.registrationPath, reg, filePerm); err != nil {
		return fmt.Errorf("writing %s: %w", cfg.registrationPath, err)
	}

	ie, err := emitImportExamples(cfg, rows)
	if err != nil {
		return err
	}
	if err := os.WriteFile(cfg.importExamplesPath, ie, filePerm); err != nil {
		return fmt.Errorf("writing %s: %w", cfg.importExamplesPath, err)
	}

	stepf("Done: generated %d files for %d types in %s.", n, len(rows), time.Since(start).Round(time.Millisecond))
	infof("Next: `make build` to compile-check, then `go run ./internal/tools/bigdiffer -docs`.")
	return nil
}

// countArtifacts totals the artifacts across all rows' plans, for a progress bar
// bound. Plan errors are ignored here; they surface as generation errors.
func countArtifacts(cfg config, rows []resourceRow) int {
	total := 0
	for _, row := range rows {
		if p, err := generationPlan(row, cfg.prefix, cfg.cacheDir); err == nil {
			total += len(p.artifacts)
		}
	}
	return total
}
