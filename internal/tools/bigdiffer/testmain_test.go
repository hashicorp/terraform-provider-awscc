// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/hashicorp/cli"
)

// TestMain points the candidate-generation executors at an in-process
// implementation for the whole test run. Production keeps the default
// subprocess-per-candidate isolation (generateArtifactIsolated /
// generateCandidateArtifacts re-exec via os.Executable); tests do not, because
// the whole-corpus tests regenerate the committed, already-triaged overlay
// where no type crashes — so isolation buys nothing there, and its ~1586
// subprocess launches per sweep are pure cost (and trip endpoint security on
// managed machines when several sweeps run back-to-back). A test that must
// verify real crash containment opts back into the subprocess path with
// realIsolation.
func TestMain(m *testing.M) {
	generateArtifactIsolated = inProcessArtifact
	generateCandidateArtifacts = inProcessCandidateArtifacts
	os.Exit(m.Run())
}

// inProcessArtifact generates one artifact in the current process — no
// subprocess, no crash containment — mirroring runRecheckProbeArtifact's
// single-artifact generation (the compile gate is driven separately by tests).
func inProcessArtifact(cfg config, row resourceRow, kind artifactKind, _ []byte) (code, test []byte, err error) {
	p, err := generationPlan(row, cfg.prefix, cfg.cacheDir)
	if err != nil {
		return nil, nil, err
	}
	ui := &cli.BasicUi{Writer: io.Discard, ErrorWriter: io.Discard}
	for _, a := range p.artifacts {
		if a.kind == kind {
			return generateArtifact(ui, cfg, p, a)
		}
	}
	return nil, nil, fmt.Errorf("artifact %s not derivable from this row", kind)
}

// inProcessCandidateArtifacts is inProcessArtifact's whole-candidate analogue,
// mirroring runProbeCandidate's in-band per-artifact error capture without the
// subprocess. Reads the schema from the plan's staged path, so the schema-byte
// argument the subprocess form ships is unused here.
func inProcessCandidateArtifacts(cfg config, row resourceRow, _ []genArtifact, _ []byte) ([]genResult, error) {
	p, err := generationPlan(row, cfg.prefix, cfg.cacheDir)
	if err != nil {
		return nil, err
	}
	ui := &cli.BasicUi{Writer: io.Discard, ErrorWriter: io.Discard}
	results := make([]genResult, len(p.artifacts))
	for i, a := range p.artifacts {
		code, test, genErr := generateArtifact(ui, cfg, p, a)
		results[i] = genResult{a: a, code: code, test: test, err: genErr}
	}
	return results, nil
}

// realIsolation swaps the executors to the actual re-exec'd subprocess path for
// a single test that must verify crash containment, restoring the in-process
// default on cleanup. Not parallel-safe: it mutates package-level vars, so
// callers must not use t.Parallel.
func realIsolation(t *testing.T) {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "bigdiffer")
	if out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput(); err != nil {
		t.Fatalf("building bigdiffer for an isolated-generation test: %v\n%s", err, out)
	}
	origArt, origCand := generateArtifactIsolated, generateCandidateArtifacts
	t.Cleanup(func() {
		generateArtifactIsolated = origArt
		generateCandidateArtifacts = origCand
	})
	generateArtifactIsolated = func(cfg config, row resourceRow, kind artifactKind, schema []byte) ([]byte, []byte, error) {
		return generateArtifactIsolatedWithBinary(bin, cfg, row, kind, schema)
	}
	generateCandidateArtifacts = func(cfg config, row resourceRow, artifacts []genArtifact, schema []byte) ([]genResult, error) {
		return generateCandidateArtifactsWithBinary(bin, cfg, row, artifacts, schema)
	}
}
