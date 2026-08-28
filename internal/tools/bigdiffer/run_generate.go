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

	results := generateCorpus(cfg, rows, cfg.genConcurrency)
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

	n, err := writeCorpus(cfg, results)
	if err != nil {
		return err
	}

	reg, err := emitRegistration(cfg, rows)
	if err != nil {
		return err
	}
	if err := os.WriteFile(cfg.registrationPath, reg, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", cfg.registrationPath, err)
	}

	ie, err := emitImportExamples(cfg, rows)
	if err != nil {
		return err
	}
	if err := os.WriteFile(cfg.importExamplesPath, ie, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", cfg.importExamplesPath, err)
	}

	fmt.Fprintf(os.Stderr, "bigdiffer: generated %d files + registration + import_examples for %d types in %s\n",
		n, len(rows), time.Since(start).Round(time.Millisecond))
	return nil
}
