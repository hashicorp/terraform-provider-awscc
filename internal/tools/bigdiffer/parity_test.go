// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// TestFullCorpusParity is the lynchpin: it generates every overlay type's
// artifacts through bigdiffer's owned engine (concurrently, via generateCorpus)
// and compares them to the committed *_gen.go. The true criterion is bigdiffer ==
// legacy, so committed is only the fast first pass: on a mismatch the harness runs
// the legacy generator for that artifact and compares against it — equal means the
// committed file is merely stale (logged, non-fatal), only a real
// bigdiffer-vs-legacy difference fails. Generating concurrently also exercises the
// engine under goroutines (run with -race to prove race-freedom). It writes
// nothing under version control.
//
// Slow, so it is skipped under -short. Set BIGDIFFER_PARITY_ONLY=AWS::Svc::Type to
// check a single type while iterating.
func TestFullCorpusParity(t *testing.T) {
	if testing.Short() {
		t.Skip("full-corpus parity is slow; run without -short")
	}
	only := os.Getenv("BIGDIFFER_PARITY_ONLY")

	overlayPath, err := filepath.Abs(filepath.Join("..", "..", "..", "internal", "provider", "all_schemas.hcl"))
	if err != nil {
		t.Fatal(err)
	}
	var f allSchemasFile
	if err := hclsimple.DecodeFile(overlayPath, nil, &f); err != nil {
		t.Fatalf("decoding overlay: %v", err)
	}
	cfg, err := newConfig(overlayPath, f.Defaults, f.Meta)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filepath.Join(cfg.overlayDir, "last_resource.txt"))

	rows := f.Resources
	if only != "" {
		var filtered []resourceRow
		for _, r := range rows {
			if r.CloudFormationTypeName == only {
				filtered = append(filtered, r)
			}
		}
		rows = filtered
	}

	// Measure serial vs parallel to record the speedup; check the parallel output.
	serialStart := time.Now()
	_ = generateCorpus(cfg, rows, 1)
	serialDur := time.Since(serialStart)

	parStart := time.Now()
	results := generateCorpus(cfg, rows, cfg.genConcurrency)
	parDur := time.Since(parStart)

	var checked, drift, stale, genErrors int
	for _, r := range results {
		if r.err != nil {
			genErrors++
			if genErrors <= 10 {
				t.Errorf("%s %s: generate: %v", r.p.cfType, r.a.kind, r.err)
			}
			continue
		}
		for _, pair := range []struct {
			file string
			got  []byte
		}{
			{r.a.codeFile, r.code},
			{r.a.testFile, r.test},
		} {
			checked++
			path := filepath.Join(cfg.outputRoot, r.a.pathSuffix, pair.file)
			want, rerr := os.ReadFile(path)
			if rerr != nil {
				drift++
				if drift <= 10 {
					t.Errorf("reading committed %s: %v", path, rerr)
				}
				continue
			}
			if bytes.Equal(pair.got, want) {
				continue
			}
			legacy, lerr := runLegacyArtifact(t, cfg, r.p, r.a)
			if lerr != nil {
				drift++
				if drift <= 10 {
					t.Errorf("MISMATCH %s (%s); legacy re-run failed: %v", path, r.a.kind, lerr)
				}
				continue
			}
			legacyGot := legacy.code
			if pair.file == r.a.testFile {
				legacyGot = legacy.test
			}
			if bytes.Equal(pair.got, legacyGot) {
				stale++
				t.Logf("STALE committed %s (%s): bigdiffer matches legacy; committed file is out of date", path, r.a.kind)
				continue
			}
			drift++
			if drift <= 10 {
				t.Errorf("DRIFT %s (%s): bigdiffer != legacy\n%s", path, r.a.kind, firstDiff(legacyGot, pair.got))
			}
		}
	}

	speedup := float64(serialDur) / float64(parDur)
	t.Logf("parity: %d types, %d files checked, %d drift, %d stale-committed, %d generate errors", len(rows), checked, drift, stale, genErrors)
	t.Logf("generation: serial %v, parallel(%d) %v, speedup %.1fx", serialDur.Round(time.Millisecond), cfg.genConcurrency, parDur.Round(time.Millisecond), speedup)
	if drift > 0 || genErrors > 0 {
		t.Fatalf("parity failed: %d drift, %d generate errors (%d stale-committed are non-fatal)", drift, genErrors, stale)
	}
}

type legacyOutput struct{ code, test []byte }

// runLegacyArtifact invokes the legacy generator (as the go:generate directives
// do) for one artifact into a temp dir and returns its bytes, so the test can
// compare bigdiffer against the actual current legacy output rather than a
// possibly-stale committed file. Run from internal/provider so the generators'
// CWD-relative reads (services.hcl) resolve.
func runLegacyArtifact(t *testing.T, cfg config, p plan, a genArtifact) (legacyOutput, error) {
	t.Helper()
	tmp := t.TempDir()
	codePath := filepath.Join(tmp, a.codeFile)
	testPath := filepath.Join(tmp, a.testFile)

	var args []string
	switch a.kind {
	case artifactResource:
		args = []string{"run", "generators/resource/main.go", "-resource", a.tfType, "-cfschema", p.schemaFile, "-package", a.packageName}
		if a.listResource {
			args = append(args, "-listresource")
		}
		args = append(args, "--", codePath, testPath)
	case artifactSingularDataSource:
		args = []string{"run", "generators/singular-data-source/main.go", "-data-source", a.tfType, "-cfschema", p.schemaFile, "-package", a.packageName, "--", codePath, testPath}
	case artifactPluralDataSource:
		args = []string{"run", "generators/plural-data-source/main.go", "-data-source", a.tfType, "-cftype", p.cfType, "-package", a.packageName, "--", codePath, testPath}
	}

	cmd := exec.Command("go", args...)
	cmd.Dir = cfg.overlayDir
	if out, err := cmd.CombinedOutput(); err != nil {
		return legacyOutput{}, fmt.Errorf("legacy generate: %w\n%s", err, out)
	}
	code, _ := os.ReadFile(codePath)
	test, _ := os.ReadFile(testPath)
	return legacyOutput{code: code, test: test}, nil
}

// firstDiff reports the first differing line between want and got.
func firstDiff(want, got []byte) string {
	wl := strings.Split(string(want), "\n")
	gl := strings.Split(string(got), "\n")
	for i := 0; i < len(wl) && i < len(gl); i++ {
		if wl[i] != gl[i] {
			return fmt.Sprintf("  first diff at line %d:\n    legacy:    %q\n    bigdiffer: %q", i+1, wl[i], gl[i])
		}
	}
	if len(wl) != len(gl) {
		return fmt.Sprintf("  line count differs: legacy=%d bigdiffer=%d", len(wl), len(gl))
	}
	return "  (identical by lines; trailing bytes differ)"
}
