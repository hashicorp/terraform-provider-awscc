// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// TestMain builds a real bigdiffer binary once for the whole test binary run
// and points generateArtifactIsolated/generateCandidateArtifacts at it.
// Both default to os.Executable(), which resolves to the `go test` binary
// under test, not bigdiffer's real main() — so they can't re-exec into
// -recheck-probe-artifact without this.
//
// probeArtifact/probeArtifactWithBinary's own tests (recheck_test.go) build
// and swap their own binary explicitly per test; this TestMain does not
// touch that var.
func TestMain(m *testing.M) {
	os.Exit(runTestMain(m))
}

// runTestMain is TestMain's body, split out so the scratch dir's defer
// actually runs before the process exits — os.Exit from directly inside
// TestMain would skip it (gocritic: exitAfterDefer).
func runTestMain(m *testing.M) int {
	dir, err := os.MkdirTemp("", "bigdiffer-testmain-*")
	if err != nil {
		fmt.Fprintln(os.Stderr, "TestMain: creating scratch dir:", err)
		return 1
	}
	defer func() { _ = os.RemoveAll(dir) }()

	bin := filepath.Join(dir, "bigdiffer")
	build := exec.Command("go", "build", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		fmt.Fprintf(os.Stderr, "TestMain: building bigdiffer for isolated-generation tests: %v\n%s\n", err, out)
		return 1
	}

	generateArtifactIsolated = func(cfg config, row resourceRow, kind artifactKind, schema []byte) (code, test []byte, err error) {
		return generateArtifactIsolatedWithBinary(bin, cfg, row, kind, schema)
	}
	generateCandidateArtifacts = func(cfg config, row resourceRow, artifacts []genArtifact, schema []byte) ([]genResult, error) {
		return generateCandidateArtifactsWithBinary(bin, cfg, row, artifacts, schema)
	}

	return m.Run()
}
