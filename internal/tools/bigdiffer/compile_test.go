// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseBuildErrors(t *testing.T) {
	t.Parallel()

	output := `# pathcheck2/a
a/broken.go:2:23: cannot use "x" (untyped string constant) as int value in return statement
# pathcheck2/b
b/broken.go:2:23: cannot use "y" (untyped string constant) as int value in return statement
`
	errs := parseBuildErrors(output, "/repo")
	if len(errs) != 2 {
		t.Fatalf("got %d errors, want 2: %+v", len(errs), errs)
	}
	if errs[0].file != filepath.Join("/repo", "a/broken.go") || errs[0].line != 2 || errs[0].col != 23 {
		t.Errorf("errs[0] = %+v", errs[0])
	}
	if errs[1].file != filepath.Join("/repo", "b/broken.go") {
		t.Errorf("errs[1] = %+v", errs[1])
	}
}

func TestParseBuildErrorsAbsolutePath(t *testing.T) {
	t.Parallel()

	// go build can also report an absolute path (e.g. when building a package
	// outside the module via a replace directive); it must be left as-is, not
	// re-joined onto repoRoot.
	output := "/elsewhere/broken.go:1:1: some error\n"
	errs := parseBuildErrors(output, "/repo")
	if len(errs) != 1 || errs[0].file != "/elsewhere/broken.go" {
		t.Errorf("got %+v, want the absolute path preserved", errs)
	}
}

func TestParseBuildErrorsIgnoresPackageHeaderAndUnrelatedLines(t *testing.T) {
	t.Parallel()

	output := "# pathcheck2/a\ngo: downloading pathcheck2 v0.0.0\nnote: module requires Go 1.25\n"
	errs := parseBuildErrors(output, "/repo")
	if len(errs) != 0 {
		t.Errorf("got %+v, want no matches for non-error lines", errs)
	}
}

func TestBlamedFiles(t *testing.T) {
	t.Parallel()

	errs := []buildError{
		{file: "/repo/a.go", line: 1, col: 1, message: "boom"},
		{file: "/repo/a.go", line: 2, col: 1, message: "boom again, same file"},
		{file: "/repo/b.go", line: 1, col: 1, message: "different file"},
	}
	got := blamedFiles(errs)
	if len(got) != 2 {
		t.Fatalf("got %d distinct files, want 2 (deduplicated): %v", len(got), got)
	}
	if _, ok := got["/repo/a.go"]; !ok {
		t.Error("missing a.go")
	}
	if _, ok := got["/repo/b.go"]; !ok {
		t.Error("missing b.go")
	}
}

func TestOverlayFilesAndRevertRestoresExisting(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	target := filepath.Join(dir, "existing.go")
	if err := os.WriteFile(target, []byte("original"), filePerm); err != nil {
		t.Fatal(err)
	}

	overlay, err := overlayFiles(map[string][]byte{target: []byte("overlaid")})
	if err != nil {
		t.Fatalf("overlayFiles: %v", err)
	}
	got, err := os.ReadFile(target)
	if err != nil || string(got) != "overlaid" {
		t.Fatalf("overlay did not write, got %q, err %v", got, err)
	}

	if err := overlay.revert(); err != nil {
		t.Fatalf("revert: %v", err)
	}
	got, err = os.ReadFile(target)
	if err != nil || string(got) != "original" {
		t.Fatalf("revert did not restore, got %q, err %v", got, err)
	}
}

func TestOverlayFilesAndRevertRemovesNew(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	target := filepath.Join(dir, "sub", "new.go") // doesn't exist yet, nor does its dir

	overlay, err := overlayFiles(map[string][]byte{target: []byte("new content")})
	if err != nil {
		t.Fatalf("overlayFiles: %v", err)
	}
	if _, err := os.Stat(target); err != nil {
		t.Fatalf("overlay did not create the file: %v", err)
	}

	if err := overlay.revert(); err != nil {
		t.Fatalf("revert: %v", err)
	}
	if _, err := os.Stat(target); !os.IsNotExist(err) {
		t.Fatalf("revert should have removed the file it created, stat err = %v", err)
	}
}

func TestBuildOnceGreenOnCleanPackage(t *testing.T) {
	if testing.Short() {
		t.Skip("runs a real go build ./... against the module (~3s warm)")
	}
	// Not t.Parallel(): this and TestBuildOnceCatchesAnInjectedTypeError both
	// build the real module tree via buildOnce's real overlay/revert; running
	// them concurrently would race each other's overlay against the same real
	// tree, which is exactly the class of hazard the design doc's "why not
	// per-candidate concurrent" section rules out for the gate itself — the
	// tests should not reintroduce it just to save a few seconds.
	ok, errs, err := buildOnce(repoRootForTest(t), map[string][]byte{})
	if err != nil {
		t.Fatalf("buildOnce: %v", err)
	}
	if !ok || len(errs) != 0 {
		t.Errorf("building the real module with no overlay should be green, got ok=%v errs=%v", ok, errs)
	}
}

func TestBuildOnceCatchesAnInjectedTypeError(t *testing.T) {
	if testing.Short() {
		t.Skip("runs a real go build ./... against the module (~3s warm)")
	}
	// Not t.Parallel() — see TestBuildOnceGreenOnCleanPackage.
	repoRoot := repoRootForTest(t)
	// Overlay a deliberately broken file into a real, harmless location under
	// internal/aws — a scratch subpackage-shaped path that doesn't collide with
	// any real generated file, reverted unconditionally by buildOnce.
	target := filepath.Join(repoRoot, "internal", "aws", "zzz_compile_gate_test_scratch", "broken.go")
	broken := []byte("package zzz_compile_gate_test_scratch\nfunc F() int { return \"not an int\" }\n")

	ok, errs, err := buildOnce(repoRoot, map[string][]byte{target: broken})
	if err != nil {
		t.Fatalf("buildOnce: %v", err)
	}
	if ok {
		t.Fatal("expected the injected type error to fail the build")
	}
	blamed := blamedFiles(errs)
	if _, found := blamed[target]; !found {
		t.Errorf("expected %s to be blamed, got %v", target, blamed)
	}
	if _, err := os.Stat(target); !os.IsNotExist(err) {
		t.Errorf("buildOnce must revert the overlay even on a red build; file still exists, stat err = %v", err)
	}
}

// repoRootForTest resolves the module root from the current test binary's
// working directory (internal/tools/bigdiffer) for tests that need to run a
// real `go build ./...` against the real module.
func repoRootForTest(t *testing.T) string {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	return filepath.Join(wd, "..", "..", "..")
}

// TestProjectRows covers the fixpoint's per-round row projection: it must
// reflect a decision's suppress_* flags on an existing row and add a brand new
// row for a classNew decision, matching what promotion will actually write —
// projectRows is a thin round-trip through the real normalizeWithDecisions, not
// a parallel reimplementation, so this mostly guards the round-trip plumbing
// (HCL encode/decode) rather than decision semantics already covered elsewhere.
func TestProjectRows(t *testing.T) {
	t.Parallel()

	overlay := testHead +
		"# 1 CloudFormation resource types schemas are available for use with the Cloud Control API.\n\n" +
		`resource_schema "aws_ec2_thing" {
  cloudformation_type_name = "AWS::EC2::Thing"
}
`

	base := []resourceRow{
		{ResourceTypeName: "aws_ec2_thing", CloudFormationTypeName: "AWS::EC2::Thing"},
	}
	decisions := map[string]policyDecision{
		"AWS::EC2::Thing": {
			setAttrs: map[string]string{attrSuppressResource: "true"},
			reason:   "build_failed: resource: undefined symbol",
		},
	}

	rows, err := projectRows(overlay, base, map[string]bool{}, decisions)
	if err != nil {
		t.Fatalf("projectRows: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("got %d rows, want 1: %+v", len(rows), rows)
	}
	if !rows[0].SuppressResourceGeneration {
		t.Errorf("expected the decision's suppress_resource_generation to be reflected, got %+v", rows[0])
	}
	if rows[0].SuppressionReason != "build_failed: resource: undefined symbol" {
		t.Errorf("expected the decision's reason to be reflected, got %+v", rows[0])
	}
}

func TestProjectRowsAddsNewRow(t *testing.T) {
	t.Parallel()

	overlay := testHead +
		"# 0 CloudFormation resource types schemas are available for use with the Cloud Control API.\n\n"

	base := []resourceRow{
		{ResourceTypeName: "aws_ec2_new_thing", CloudFormationTypeName: "AWS::EC2::NewThing"},
	}
	decisions := map[string]policyDecision{
		"AWS::EC2::NewThing": {addBlock: true},
	}

	rows, err := projectRows(overlay, base, map[string]bool{}, decisions)
	if err != nil {
		t.Fatalf("projectRows: %v", err)
	}
	if len(rows) != 1 || rows[0].CloudFormationTypeName != "AWS::EC2::NewThing" {
		t.Fatalf("got %+v, want the new row projected", rows)
	}
}
