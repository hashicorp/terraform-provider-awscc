// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// writeFile is a small test helper: creates path (and its parent dirs) with
// the given content.
func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("creating parent dirs for %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("writing %s: %v", path, err)
	}
}

// TestDiffStagedTreesNoDiff confirms a staged tree that is byte-identical to
// the committed tree it mirrors produces zero diffs — the common case for
// check on a non-engine PR, where regeneration is a no-op.
func TestDiffStagedTreesNoDiff(t *testing.T) {
	t.Parallel()

	stagingDir := t.TempDir()
	outputRoot := t.TempDir()
	cacheDir := t.TempDir()

	writeFile(t, filepath.Join(stagingDir, "out", "svc", "resource_gen.go"), "package svc\n")
	writeFile(t, filepath.Join(outputRoot, "svc", "resource_gen.go"), "package svc\n")

	cfg := config{outputRoot: outputRoot, cacheDir: cacheDir, repoRoot: t.TempDir()}
	diffs, err := diffStagedTrees(cfg, stagingDir)
	if err != nil {
		t.Fatalf("diffStagedTrees: %v", err)
	}
	if len(diffs) != 0 {
		t.Errorf("want 0 diffs, got %v", diffs)
	}
}

// TestDiffStagedTreesDetectsByteDiff confirms a staged file whose bytes
// differ from the committed file it mirrors is reported — the "engine change
// not yet -reconcile'd and committed" case check exists to catch.
func TestDiffStagedTreesDetectsByteDiff(t *testing.T) {
	t.Parallel()

	stagingDir := t.TempDir()
	outputRoot := t.TempDir()
	cacheDir := t.TempDir()

	writeFile(t, filepath.Join(stagingDir, "out", "svc", "resource_gen.go"), "package svc // new template output\n")
	writeFile(t, filepath.Join(outputRoot, "svc", "resource_gen.go"), "package svc // old template output\n")

	cfg := config{outputRoot: outputRoot, cacheDir: cacheDir, repoRoot: filepath.Dir(outputRoot)}
	diffs, err := diffStagedTrees(cfg, stagingDir)
	if err != nil {
		t.Fatalf("diffStagedTrees: %v", err)
	}
	if len(diffs) != 1 {
		t.Fatalf("want exactly 1 diff, got %v", diffs)
	}
}

// TestDiffStagedTreesDetectsMissingCommitted confirms a staged file with no
// committed counterpart at all (a type recovering from suppression via a
// machinery fix, with nothing previously committed to compare against) is
// still reported as a diff — there is genuinely new output to reconcile and
// commit, not a false "no diff."
func TestDiffStagedTreesDetectsMissingCommitted(t *testing.T) {
	t.Parallel()

	stagingDir := t.TempDir()
	outputRoot := t.TempDir()
	cacheDir := t.TempDir()

	writeFile(t, filepath.Join(stagingDir, "out", "svc", "newly_recovered_gen.go"), "package svc\n")
	// Deliberately nothing written under outputRoot for this file.

	cfg := config{outputRoot: outputRoot, cacheDir: cacheDir, repoRoot: filepath.Dir(outputRoot)}
	diffs, err := diffStagedTrees(cfg, stagingDir)
	if err != nil {
		t.Fatalf("diffStagedTrees: %v", err)
	}
	if len(diffs) != 1 {
		t.Fatalf("want exactly 1 diff (the file with no committed counterpart), got %v", diffs)
	}
}

// TestDiffStagedTreesChecksCacheToo confirms diffStagedTrees compares
// stagingDir/cache against cfg.cacheDir in addition to stagingDir/out against
// cfg.outputRoot — a cache-bytes drift (e.g. discovery normalizing schema
// bytes differently) is exactly as much an uncommitted diff as a generated
// code diff.
func TestDiffStagedTreesChecksCacheToo(t *testing.T) {
	t.Parallel()

	stagingDir := t.TempDir()
	outputRoot := t.TempDir()
	cacheDir := t.TempDir()

	writeFile(t, filepath.Join(stagingDir, "cache", "AWS_Svc_Thing.json"), `{"new":"bytes"}`)
	writeFile(t, filepath.Join(cacheDir, "AWS_Svc_Thing.json"), `{"old":"bytes"}`)

	cfg := config{outputRoot: outputRoot, cacheDir: cacheDir, repoRoot: filepath.Dir(outputRoot)}
	diffs, err := diffStagedTrees(cfg, stagingDir)
	if err != nil {
		t.Fatalf("diffStagedTrees: %v", err)
	}
	if len(diffs) != 1 {
		t.Fatalf("want exactly 1 diff (the cache file), got %v", diffs)
	}
}

// TestDiffStagedTreesEmptyStagingIsNoError confirms an empty (or absent)
// staging dir -- e.g. every candidate failed, so nothing was ever staged --
// produces zero diffs and no error, rather than a spurious failure.
func TestDiffStagedTreesEmptyStagingIsNoError(t *testing.T) {
	t.Parallel()

	stagingDir := t.TempDir() // nothing staged under out/ or cache/ at all
	outputRoot := t.TempDir()
	cacheDir := t.TempDir()

	cfg := config{outputRoot: outputRoot, cacheDir: cacheDir, repoRoot: filepath.Dir(outputRoot)}
	diffs, err := diffStagedTrees(cfg, stagingDir)
	if err != nil {
		t.Fatalf("diffStagedTrees: %v", err)
	}
	if len(diffs) != 0 {
		t.Errorf("want 0 diffs, got %v", diffs)
	}
}

// TestDiffImportExamplesNoDiff confirms an up-to-date committed
// import_examples_gen.json against the real, untouched corpus produces no
// diff -- the common case for check on a PR that didn't change anything
// affecting the aggregate.
func TestDiffImportExamplesNoDiff(t *testing.T) {
	if testing.Short() {
		t.Skip("recomputes the whole corpus's import examples; run without -short")
	}
	t.Parallel()
	cfg, rows := loadCorpus(t)
	overlayContent, err := os.ReadFile(cfg.overlayPath)
	if err != nil {
		t.Fatalf("reading overlay: %v", err)
	}
	checkout, err := parseCheckout(defaultCheckout)
	if err != nil {
		t.Fatalf("parsing checkout: %v", err)
	}

	diff, err := diffImportExamples(cfg, string(overlayContent), rows, checkout, nil)
	if err != nil {
		t.Fatalf("diffImportExamples: %v", err)
	}
	if diff != "" {
		t.Errorf("want no diff against the untouched, committed corpus, got: %s", diff)
	}
}

// TestDiffImportExamplesDetectsStaleCommittedFile confirms a committed
// import_examples_gen.json that no longer matches what the current overlay +
// cache would produce is reported -- the "engine (or the row set) changed,
// the aggregate was not regenerated" case item 2's cheap tier exists to
// catch. Points cfg at a temp importExamplesPath with deliberately stale
// content rather than touching the real committed file.
func TestDiffImportExamplesDetectsStaleCommittedFile(t *testing.T) {
	if testing.Short() {
		t.Skip("recomputes the whole corpus's import examples; run without -short")
	}
	t.Parallel()
	cfg, rows := loadCorpus(t)
	overlayContent, err := os.ReadFile(cfg.overlayPath)
	if err != nil {
		t.Fatalf("reading overlay: %v", err)
	}
	checkout, err := parseCheckout(defaultCheckout)
	if err != nil {
		t.Fatalf("parsing checkout: %v", err)
	}

	cfg.importExamplesPath = filepath.Join(t.TempDir(), "import_examples_gen.json")
	writeFile(t, cfg.importExamplesPath, `[{"resource": "this_is_deliberately_stale"}]`)

	diff, err := diffImportExamples(cfg, string(overlayContent), rows, checkout, nil)
	if err != nil {
		t.Fatalf("diffImportExamples: %v", err)
	}
	if diff == "" {
		t.Fatal("want a diff against a deliberately stale committed file, got none")
	}
	if !strings.Contains(diff, "differs from committed") {
		t.Errorf("want a 'differs from committed' problem description, got: %s", diff)
	}
}

// TestDiffImportExamplesDetectsMissingCommittedFile confirms a missing
// import_examples_gen.json is reported as a diff, not a silent pass or an
// error -- there is genuinely new output to -reconcile and commit.
func TestDiffImportExamplesDetectsMissingCommittedFile(t *testing.T) {
	if testing.Short() {
		t.Skip("recomputes the whole corpus's import examples; run without -short")
	}
	t.Parallel()
	cfg, rows := loadCorpus(t)
	overlayContent, err := os.ReadFile(cfg.overlayPath)
	if err != nil {
		t.Fatalf("reading overlay: %v", err)
	}
	checkout, err := parseCheckout(defaultCheckout)
	if err != nil {
		t.Fatalf("parsing checkout: %v", err)
	}

	cfg.importExamplesPath = filepath.Join(t.TempDir(), "does-not-exist.json")

	diff, err := diffImportExamples(cfg, string(overlayContent), rows, checkout, nil)
	if err != nil {
		t.Fatalf("diffImportExamples: %v", err)
	}
	if diff == "" {
		t.Fatal("want a diff for a missing committed file, got none")
	}
	if !strings.Contains(diff, "does not exist") {
		t.Errorf("want a 'does not exist' problem description, got: %s", diff)
	}
}

// TestDiffDocsTreesNoDiff confirms a rendered tree byte-identical to the
// committed docs/ it mirrors produces zero diffs — diffDocsTrees' own logic,
// exercised on a cheap synthetic tree rather than a real provider build
// (that full exercise is TestDiffRenderedDocsNoDiff, below).
func TestDiffDocsTreesNoDiff(t *testing.T) {
	t.Parallel()

	renderedDir := t.TempDir()
	docsDir := t.TempDir()

	writeFile(t, filepath.Join(renderedDir, "resources", "thing.md"), "# thing\n")
	writeFile(t, filepath.Join(docsDir, "resources", "thing.md"), "# thing\n")

	cfg := config{docsDir: docsDir, repoRoot: filepath.Dir(docsDir)}
	diffs, err := diffDocsTrees(cfg, renderedDir)
	if err != nil {
		t.Fatalf("diffDocsTrees: %v", err)
	}
	if len(diffs) != 0 {
		t.Errorf("want 0 diffs, got %v", diffs)
	}
}

// TestDiffDocsTreesDetectsByteDiff confirms a rendered file whose bytes
// differ from the committed file it mirrors is reported — the "engine or
// docs-template change not yet -reconcile'd/-docs'd and committed" case the
// full tier exists to catch.
func TestDiffDocsTreesDetectsByteDiff(t *testing.T) {
	t.Parallel()

	renderedDir := t.TempDir()
	docsDir := t.TempDir()

	writeFile(t, filepath.Join(renderedDir, "resources", "thing.md"), "# thing (new template output)\n")
	writeFile(t, filepath.Join(docsDir, "resources", "thing.md"), "# thing (old template output)\n")

	cfg := config{docsDir: docsDir, repoRoot: filepath.Dir(docsDir)}
	diffs, err := diffDocsTrees(cfg, renderedDir)
	if err != nil {
		t.Fatalf("diffDocsTrees: %v", err)
	}
	if len(diffs) != 1 {
		t.Fatalf("want exactly 1 diff, got %v", diffs)
	}
}

// TestDiffDocsTreesDetectsMissingCommitted confirms a rendered file with no
// committed counterpart at all is still reported as a diff — there is
// genuinely new documentation to -reconcile/-docs and commit, not a false
// "no diff."
func TestDiffDocsTreesDetectsMissingCommitted(t *testing.T) {
	t.Parallel()

	renderedDir := t.TempDir()
	docsDir := t.TempDir()

	writeFile(t, filepath.Join(renderedDir, "resources", "newly_recovered.md"), "# newly recovered\n")
	// Deliberately nothing written under docsDir for this file.

	cfg := config{docsDir: docsDir, repoRoot: filepath.Dir(docsDir)}
	diffs, err := diffDocsTrees(cfg, renderedDir)
	if err != nil {
		t.Fatalf("diffDocsTrees: %v", err)
	}
	if len(diffs) != 1 {
		t.Fatalf("want exactly 1 diff (the file with no committed counterpart), got %v", diffs)
	}
}

// TestDiffRenderedDocsNoDiff confirms the full-tier docs check
// (contributing/docs/docs-pipeline-punchlist.md item 2) renders byte-identical
// docs against the real, untouched corpus: builds a real provider binary
// from the staged tree, extracts its schema via a real `terraform providers
// schema -json` call, and renders through tfplugindocs — exercising the
// whole extraction + render + diff path end to end, not a synthetic tree.
// Requires `terraform` and `tfplugindocs` on PATH (both installed by
// `make tools`); skipped if either is missing, and in -short mode regardless
// (a real go build of the whole provider plus a schema-extraction call is
// the heaviest single check in the whole bigdiffer suite).
func TestDiffRenderedDocsNoDiff(t *testing.T) {
	if testing.Short() {
		t.Skip("builds a real provider binary and extracts its schema; run without -short")
	}
	if _, err := exec.LookPath("terraform"); err != nil {
		t.Skip("terraform not on PATH")
	}
	if _, err := exec.LookPath("tfplugindocs"); err != nil {
		t.Skip("tfplugindocs not on PATH")
	}

	cfg, rows := loadCorpus(t)
	overlayContent, err := os.ReadFile(cfg.overlayPath)
	if err != nil {
		t.Fatalf("reading overlay: %v", err)
	}
	checkout, err := parseCheckout(defaultCheckout)
	if err != nil {
		t.Fatalf("parsing checkout: %v", err)
	}

	cands, _, _ := reconcileCandidates(rows, cfg.cacheDir)
	settled, err := settleBatch(context.Background(), cfg, cands, nil, string(overlayContent), rows, checkout, todayString())
	if err != nil {
		t.Fatalf("settleBatch: %v", err)
	}
	defer func() { _ = os.RemoveAll(settled.stagingDir) }()

	projected, err := projectRows(string(overlayContent), rows, checkout, settled.decisions)
	if err != nil {
		t.Fatalf("projectRows: %v", err)
	}

	start := time.Now()
	diffs, err := diffRenderedDocs(context.Background(), cfg, settled.stagingDir, projected)
	if err != nil {
		t.Fatalf("diffRenderedDocs: %v", err)
	}
	t.Logf("diffRenderedDocs against the untouched corpus took %s", time.Since(start).Round(time.Second))
	if len(diffs) != 0 {
		t.Errorf("want 0 diffs against the untouched, committed corpus, got %d:\n%s", len(diffs), strings.Join(diffs, "\n"))
	}
}

// TestRunCheckAgainstRealRepo is check's own end-to-end wiring test, run
// against the real committed corpus -- feasible here in a way
// TestRunReconcile never was (see run_reconcile_test.go's doc comment for
// why): runCheck is read-only by construction (it stages into a temp dir via
// settleBatch, scans it, then os.RemoveAll's it -- promoteStaged is never
// called), so unlike runReconcile there is no risk in pointing it at the real
// repo from within a test. It exercises the full real dispatch path,
// including the one piece diffStagedTrees' own unit tests above cannot
// (they operate on synthetic trees): that runCheck actually finds and
// compares the *right* real files.
//
// Two sub-cases against the live corpus: a clean pass (nothing modified, so
// check must exit nil), and a real, deliberately introduced byte diff on a
// real committed file. The target file's real path is sourced from a
// throwaway settleBatch call on one non-frozen candidate rather than
// re-deriving plan.go's package/path-suffix convention by hand in the test
// itself -- reusing the pipeline's own path logic is both less code and
// immune to drifting out of sync with a future naming-convention change (an
// earlier manual attempt at this same scenario picked AWS::Logs::LogGroup and
// AWS::S3::Bucket, both frozen and so silently excluded from the real
// reconcileCandidates/check path -- sourcing the target from the pipeline's
// own decisions, not a hand-picked type name, avoids repeating that mistake).
// The corrupted file is restored via t.Cleanup unconditionally, including on
// test failure or panic, so a failed assertion never leaves the working tree
// dirty.
func TestRunCheckAgainstRealRepo(t *testing.T) {
	if testing.Short() {
		t.Skip("full-corpus check run; skipped in -short mode (exercised in the full suite instead)")
	}

	cfg, rows := loadCorpus(t)
	allSchemasPath := cfg.overlayPath
	checkout, err := parseCheckout(defaultCheckout)
	if err != nil {
		t.Fatalf("parsing checkout: %v", err)
	}

	t.Run("clean pass", func(t *testing.T) {
		if err := runCheck(context.Background(), allSchemasPath, defaultCheckout, false); err != nil {
			t.Fatalf("runCheck on an untouched corpus: %v", err)
		}
	})

	t.Run("detects a real corruption", func(t *testing.T) {
		var target resourceRow
		for _, r := range rows {
			if r.FrozenSince == "" {
				target = r
				break
			}
		}
		if target.CloudFormationTypeName == "" {
			t.Skip("no non-frozen row found in the current overlay")
		}
		schema, err := os.ReadFile(schemaCachePath(cfg.cacheDir, target.CloudFormationTypeName))
		if err != nil {
			t.Skipf("no cached schema for %s: %v", target.CloudFormationTypeName, err)
		}
		cand := candidate{cfType: target.CloudFormationTypeName, class: classPresentUnchanged, row: target, schema: schema}
		today := todayString()
		overlayContent, err := os.ReadFile(allSchemasPath)
		if err != nil {
			t.Fatalf("reading overlay: %v", err)
		}
		settled, err := settleBatch(context.Background(), cfg, []candidate{cand}, nil, string(overlayContent), rows, checkout, today)
		if err != nil {
			t.Fatalf("settleBatch for %s: %v", target.CloudFormationTypeName, err)
		}
		defer func() { _ = os.RemoveAll(settled.stagingDir) }()
		if len(settled.stagedByDest) == 0 {
			t.Skipf("%s staged no artifacts (unexpected gate failure?)", target.CloudFormationTypeName)
		}
		var path string
		for dest := range settled.stagedByDest {
			path = dest
			break
		}

		original, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("reading committed artifact %s: %v", path, err)
		}
		t.Cleanup(func() {
			if err := os.WriteFile(path, original, filePerm); err != nil {
				t.Fatalf("restoring %s after corruption test: %v", path, err)
			}
		})
		corrupted := append(append([]byte{}, original...), []byte("\n// deliberate test corruption, restored by t.Cleanup\n")...)
		if err := os.WriteFile(path, corrupted, filePerm); err != nil {
			t.Fatalf("corrupting %s: %v", path, err)
		}

		err = runCheck(context.Background(), allSchemasPath, defaultCheckout, false)
		if err == nil {
			t.Fatalf("runCheck did not detect a corrupted, committed file at %s for %s", path, target.CloudFormationTypeName)
		}
		if !strings.Contains(err.Error(), "differ from committed") {
			t.Errorf("runCheck failed for an unexpected reason (want a diff-detection failure): %v", err)
		}
	})
}
