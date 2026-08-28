// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/cli"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/codegen"
)

// TestFullCorpusParity is Brick 7's lynchpin: it generates every overlay type's
// artifacts through bigdiffer's owned in-process engine and compares them to the
// committed *_gen.go files. The true faithfulness criterion is "bigdiffer ==
// legacy", so a byte-compare against the committed corpus is only the fast first
// pass: when bigdiffer differs from the committed file, the test runs the legacy
// generator for that artifact and compares against that instead. If bigdiffer
// matches legacy, the committed file is simply stale (the release shipped code
// that no longer matches its own schema+engine) — not drift, and logged. Only a
// genuine bigdiffer-vs-legacy difference fails the test. It writes nothing under
// version control (legacy runs go to temp dirs; last_resource.txt litter is
// cleaned up).
//
// Slow (serial over the whole corpus), so it is skipped under -short. Set
// BIGDIFFER_PARITY_ONLY=AWS::Svc::Type to check a single type while iterating.
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
	// The legacy generators write a last_resource.txt debug file into their CWD
	// (internal/provider); clean it up afterward so the test leaves no trace.
	defer os.Remove(filepath.Join(cfg.overlayDir, "last_resource.txt"))

	ui := &cli.BasicUi{Writer: io.Discard, ErrorWriter: io.Discard}

	var checked, drift, stale, genErrors, types int
	for _, row := range f.Resources {
		if only != "" && row.CloudFormationTypeName != only {
			continue
		}
		types++
		p, err := generationPlan(row, cfg.prefix, cfg.cacheDir)
		if err != nil {
			t.Errorf("%s: plan: %v", row.CloudFormationTypeName, err)
			continue
		}
		for _, a := range p.artifacts {
			code, test, gerr := generateArtifact(ui, cfg, p, a)
			if gerr != nil {
				genErrors++
				if genErrors <= 10 {
					t.Errorf("%s %s: generate: %v", p.cfType, a.kind, gerr)
				}
				continue
			}
			for _, pair := range []struct {
				file string
				got  []byte
			}{
				{a.codeFile, code},
				{a.testFile, test},
			} {
				checked++
				path := filepath.Join(cfg.outputRoot, a.pathSuffix, pair.file)
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
				// Committed differs: is bigdiffer faithful to *legacy*, or drifting?
				legacy, lerr := runLegacyArtifact(t, cfg, p, a)
				if lerr != nil {
					drift++
					if drift <= 10 {
						t.Errorf("MISMATCH %s (%s); legacy re-run failed: %v", path, a.kind, lerr)
					}
					continue
				}
				legacyGot := legacy.code
				if pair.file == a.testFile {
					legacyGot = legacy.test
				}
				if bytes.Equal(pair.got, legacyGot) {
					stale++
					t.Logf("STALE committed %s (%s): bigdiffer matches legacy; committed file is out of date", path, a.kind)
					continue
				}
				drift++
				if drift <= 10 {
					t.Errorf("DRIFT %s (%s): bigdiffer != legacy\n%s", path, a.kind, firstDiff(legacyGot, pair.got))
				}
			}
		}
	}
	t.Logf("parity: %d types, %d files checked, %d drift, %d stale-committed, %d generate errors", types, checked, drift, stale, genErrors)
	if drift > 0 || genErrors > 0 {
		t.Fatalf("parity failed: %d drift, %d generate errors (%d stale-committed are non-fatal)", drift, genErrors, stale)
	}
}

// generateArtifact renders one artifact's code+test via bigdiffer's owned engine.
func generateArtifact(ui cli.Ui, cfg config, p plan, a genArtifact) (code, test []byte, err error) {
	switch a.kind {
	case artifactResource:
		return codegen.GenerateResource(ui, p.schemaFile, a.tfType, a.packageName, cfg.servicesPath, a.listResource)
	case artifactSingularDataSource:
		return codegen.GenerateSingularDataSource(ui, p.schemaFile, a.tfType, a.packageName, cfg.servicesPath)
	case artifactPluralDataSource:
		return codegen.GeneratePluralDataSource(p.cfType, a.tfType, a.packageName)
	}
	return nil, nil, fmt.Errorf("unknown artifact kind %q", a.kind)
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
