// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestE2ELiveUpdate exercises the full live weekly pipeline against AWS —
// discover -> detectChanges (vs the real cache, read-only) -> refreshCandidate
// (generating New/Changed types from their fresh bytes) — but promotes generated
// files and cache into temp directories, so the real tree is never mutated. It is
// opt-in (needs AWS credentials and a slow throttled crawl) via BIGDIFFER_E2E=1.
//
//	BIGDIFFER_E2E=1 go test ./internal/tools/bigdiffer/ -run TestE2ELiveUpdate -v -timeout 30m
func TestE2ELiveUpdate(t *testing.T) {
	if os.Getenv("BIGDIFFER_E2E") == "" {
		t.Skip("set BIGDIFFER_E2E=1 to run the live AWS end-to-end test")
	}

	cfg, overlayRows := loadCorpus(t)

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Minute)
	defer cancel()

	start := time.Now()
	disc, err := discover(ctx)
	if err != nil {
		t.Fatalf("discover: %v", err)
	}
	t.Logf("discovered %d live types in %s", len(disc), time.Since(start).Round(time.Second))
	if len(disc) < 1000 {
		t.Fatalf("implausibly few discovered types (%d); credentials/region problem?", len(disc))
	}

	discByCFN := make(map[string]discovered, len(disc))
	var missing int
	for _, d := range disc {
		discByCFN[d.row.CloudFormationTypeName] = d
		if d.err != nil {
			missing++
		}
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
		t.Fatalf("detectChanges: %v", err)
	}
	s := summarize(changes)
	t.Logf("overlay rows %d; discovered %d (describe failures %d)", len(overlayRows), len(disc), missing)
	t.Logf("changes — new %d, changed %d, unchanged %d, frozen %d, missing %d",
		s.New, s.Changed, s.Unchanged, s.Frozen, s.Missing)

	cands := buildCandidates(changes, discByCFN, overlayByCFN)
	if len(cands) == 0 {
		t.Log("no New/Changed types this run — nothing to regenerate (expected shortly after a release)")
	}

	// Promote into temp dirs only; the real tree is never touched.
	tmp := t.TempDir()
	tmpCfg := cfg
	tmpCfg.outputRoot = filepath.Join(tmp, "out")
	tmpCfg.cacheDir = filepath.Join(tmp, "cache")
	staging := filepath.Join(tmp, "staging")
	if err := os.MkdirAll(staging, 0o755); err != nil {
		t.Fatal(err)
	}

	today := time.Now().Format(dateLayout)
	var okN, brokeN int
	for _, c := range cands {
		gr, err := refreshCandidate(tmpCfg, staging, c)
		if err != nil {
			t.Fatalf("refreshCandidate %s: %v", c.cfType, err)
		}
		d := decide(c.class, gr, today)
		if gr.ok() {
			okN++
			t.Logf("  %s (%s): %s", c.cfType, c.class, d.summary)
		} else {
			brokeN++
			t.Logf("  %s (%s): %s [%s]", c.cfType, c.class, d.summary, d.reason)
		}
	}
	t.Logf("regenerated %d candidate(s): %d OK, %d frozen/suppressed", len(cands), okN, brokeN)

	// Mirror runUpdate: promote the staged batch into the temp tree once,
	// after every candidate has staged successfully.
	if err := promoteStaged(tmpCfg, staging); err != nil {
		t.Fatalf("promoteStaged: %v", err)
	}
}
