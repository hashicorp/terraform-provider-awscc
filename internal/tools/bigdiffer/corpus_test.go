// SPDX-License-Identifier: MPL-2.0

package main

import (
	"path/filepath"
	"testing"
)

// loadCorpus decodes the overlay and builds the config for corpus tests.
func loadCorpus(t *testing.T) (config, []resourceRow) {
	t.Helper()
	overlayPath, err := filepath.Abs(filepath.Join("..", "..", "..", "internal", "provider", "all_schemas.hcl"))
	if err != nil {
		t.Fatal(err)
	}
	cfg, rows, err := loadOverlay(overlayPath)
	if err != nil {
		t.Fatalf("loading overlay: %v", err)
	}
	return cfg, rows
}

// TestGenerateCorpusRace exercises concurrent generation over a diverse sample of
// the corpus (every 4th type, spread across services) so the race detector has a
// wide, genuinely-parallel surface to inspect. Run with -race to prove the owned
// engine is safe to drive from goroutines. It asserts every artifact generates
// without error; byte-parity is TestFullCorpusParity's job.
func TestGenerateCorpusRace(t *testing.T) {
	if testing.Short() {
		t.Skip("generation is slow; run without -short")
	}
	cfg, rows := loadCorpus(t)

	var sample []resourceRow
	for i, r := range rows {
		if i%4 == 0 {
			sample = append(sample, r)
		}
	}

	results := generateCorpus(cfg, sample, cfg.genConcurrency)

	var errs int
	for _, r := range results {
		if r.err != nil {
			errs++
			if errs <= 10 {
				t.Errorf("%s %s: generate: %v", r.p.cfType, r.a.kind, r.err)
			}
		}
	}
	t.Logf("generated %d sampled types (%d artifacts) concurrently (%d workers), %d errors", len(sample), len(results), cfg.genConcurrency, errs)
}
