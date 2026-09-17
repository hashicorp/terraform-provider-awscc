// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestVerifyGoToolchainAgainstRealGoMod is the real, load-bearing case: run
// from the actual repo (the test binary's own working directory), the
// running toolchain must exactly match this module's real go.mod "go"
// directive — this is the exact invariant PR #3334's CI failure violated
// (six files committed under a newer local Go than CI's pinned toolchain).
// If this test ever fails, it means whatever Go built/ran this test suite
// itself does not match go.mod, i.e. exactly the drift this check exists to
// catch — a real, not synthetic, signal.
func TestVerifyGoToolchainAgainstRealGoMod(t *testing.T) {
	if err := verifyGoToolchain(); err != nil {
		t.Fatalf("verifyGoToolchain: %v (the test binary's own toolchain, %s, does not "+
			"match this repo's go.mod)", err, runtime.Version())
	}
}

// TestVerifyGoToolchainMismatch exercises the failure path directly against
// a synthetic go.mod, without needing to actually run a different Go
// toolchain, by temporarily chdir'ing into a scratch directory.
func TestVerifyGoToolchainMismatch(t *testing.T) {
	withScratchGoMod(t, "go 99.99.99\n")

	err := verifyGoToolchain()
	if err == nil {
		t.Fatal("want an error when go.mod pins a version other than the running toolchain, got nil")
	}
	if !strings.Contains(err.Error(), "99.99.99") {
		t.Errorf("error should name the expected version: %v", err)
	}
	if !strings.Contains(err.Error(), runtime.Version()) {
		t.Errorf("error should name the running version: %v", err)
	}
}

// TestVerifyGoToolchainMatch confirms a go.mod pinning exactly the running
// toolchain's own version passes.
func TestVerifyGoToolchainMatch(t *testing.T) {
	withScratchGoMod(t, "go "+runtime.Version()[len("go"):]+"\n")

	if err := verifyGoToolchain(); err != nil {
		t.Errorf("want nil for a go.mod matching the running toolchain exactly, got: %v", err)
	}
}

// TestVerifyGoToolchainMalformedGoModTolerated confirms a go.mod with no
// parseable "go" directive is tolerated (returns nil) rather than erroring —
// this check is a formatting-drift guard, not a hard requirement for
// bigdiffer to run, so an environment quirk here should never block every
// other command.
func TestVerifyGoToolchainMalformedGoModTolerated(t *testing.T) {
	withScratchGoMod(t, "module scratch\n// no go directive at all\n")

	if err := verifyGoToolchain(); err != nil {
		t.Errorf("want nil for a go.mod with no parseable \"go\" directive, got: %v", err)
	}
}

// TestFindGoModWalksUpward confirms findGoMod locates go.mod from a nested
// subdirectory, not just the directory containing it directly — the same
// upward-walk resolution `go` itself uses.
func TestFindGoModWalksUpward(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("go 1.26.6\n"), filePerm); err != nil {
		t.Fatalf("writing scratch go.mod: %v", err)
	}
	nested := filepath.Join(dir, "a", "b", "c")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	t.Chdir(nested)

	got, err := findGoMod()
	if err != nil {
		t.Fatalf("findGoMod: %v", err)
	}
	want := filepath.Join(dir, "go.mod")
	if got != want {
		t.Errorf("findGoMod() = %q, want %q", got, want)
	}
}

// withScratchGoMod chdirs the test into a scratch directory containing only
// a go.mod with the given content, restoring the original working directory
// via t.Cleanup (t.Chdir).
func withScratchGoMod(t *testing.T, goModContent string) {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(goModContent), filePerm); err != nil {
		t.Fatalf("writing scratch go.mod: %v", err)
	}
	t.Chdir(dir)
}
