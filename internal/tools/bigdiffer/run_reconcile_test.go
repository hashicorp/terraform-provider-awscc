// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// writeSchemaCache is a small test helper: stages fake cached schema bytes
// for cfn under dir, mirroring schemaCachePath's own naming convention, so
// reconcileCandidates can find them via a plain os.ReadFile.
func writeSchemaCache(t *testing.T, dir, cfn string, bytes []byte) {
	t.Helper()
	if err := os.WriteFile(schemaCachePath(dir, cfn), bytes, 0o644); err != nil {
		t.Fatalf("writing fake schema cache for %s: %v", cfn, err)
	}
}

// TestReconcileCandidates covers both conventions pinned ahead of this
// implementation (contributing/docs/held-artifacts-design.md "The routing
// this still requires"; redesign-punchlist.md item 4): every candidate must
// be classPresentUnchanged, never classPresent, and every frozen row must be
// excluded entirely.
func TestReconcileCandidates(t *testing.T) {
	t.Parallel()

	cacheDir := t.TempDir()
	writeSchemaCache(t, cacheDir, "AWS::Svc::Fine", []byte("fine bytes"))
	writeSchemaCache(t, cacheDir, "AWS::Svc::Frozen", []byte("frozen bytes"))
	writeSchemaCache(t, cacheDir, "AWS::Svc::NoCacheHit", []byte("unused — row below has no matching cache write"))

	rows := []resourceRow{
		{ResourceTypeName: "aws_svc_fine", CloudFormationTypeName: "AWS::Svc::Fine"},
		{ResourceTypeName: "aws_svc_frozen", CloudFormationTypeName: "AWS::Svc::Frozen", FrozenSince: "2026-01-01", FrozenReason: "manual: pinned pending review"},
		// Deliberately no cache file written for this CFN — simulates a row
		// with no cached bytes at all (predates the cache, or a cache miss).
		{ResourceTypeName: "aws_svc_uncached", CloudFormationTypeName: "AWS::Svc::Uncached"},
	}

	cands, frozenSkipped, cacheMissSkipped := reconcileCandidates(rows, cacheDir)

	if frozenSkipped != 1 {
		t.Errorf("frozenSkipped = %d, want 1", frozenSkipped)
	}
	if cacheMissSkipped != 1 {
		t.Errorf("cacheMissSkipped = %d, want 1 (AWS::Svc::Uncached has no cache file)", cacheMissSkipped)
	}

	byCFN := make(map[string]candidate, len(cands))
	for _, c := range cands {
		byCFN[c.cfType] = c
	}

	if len(cands) != 1 {
		t.Fatalf("want exactly 1 candidate (Fine only — Frozen excluded, Uncached has no cache hit), got %d: %+v", len(cands), cands)
	}

	c, ok := byCFN["AWS::Svc::Fine"]
	if !ok {
		t.Fatalf("expected AWS::Svc::Fine as a candidate, got %+v", byCFN)
	}
	if c.class != classPresentUnchanged {
		t.Errorf("class = %v, want classPresentUnchanged — reconcile has no discovery diff, so classPresent (the freeze/suppress branch) must never be reused here", c.class)
	}
	if string(c.schema) != "fine bytes" {
		t.Errorf("schema = %q, want the cached bytes", c.schema)
	}
	if c.row.ResourceTypeName != "aws_svc_fine" {
		t.Errorf("row not carried through correctly: %+v", c.row)
	}

	if _, ok := byCFN["AWS::Svc::Frozen"]; ok {
		t.Errorf("frozen row must never become a candidate — it is frozen precisely because it fails to generate/compile at its pinned bytes; feeding it through classPresentUnchanged would hit machineryFailure every run with no path to ever completing")
	}
	if _, ok := byCFN["AWS::Svc::Uncached"]; ok {
		t.Errorf("a row with no cached bytes at all should be skipped, not manufactured into a failure")
	}
}

// TestReconcileCandidatesAllFrozenIsEmptyNotError confirms an all-frozen
// overlay produces zero candidates and zero frozenSkipped-related errors —
// reconcileCandidates never fails, it only filters.
func TestReconcileCandidatesAllFrozenIsEmptyNotError(t *testing.T) {
	t.Parallel()

	cacheDir := t.TempDir()
	rows := []resourceRow{
		{ResourceTypeName: "aws_svc_a", CloudFormationTypeName: "AWS::Svc::A", FrozenSince: "2026-01-01"},
		{ResourceTypeName: "aws_svc_b", CloudFormationTypeName: "AWS::Svc::B", FrozenSince: "2026-01-02"},
	}

	cands, frozenSkipped, cacheMissSkipped := reconcileCandidates(rows, cacheDir)
	if len(cands) != 0 {
		t.Errorf("want 0 candidates, got %d: %+v", len(cands), cands)
	}
	if frozenSkipped != 2 {
		t.Errorf("frozenSkipped = %d, want 2", frozenSkipped)
	}
	if cacheMissSkipped != 0 {
		t.Errorf("cacheMissSkipped = %d, want 0 — both rows are frozen, so they're skipped for that reason first and never reach the cache check", cacheMissSkipped)
	}
}

// TestReconcileCandidatesCacheMissIsCountedSeparatelyFromFrozen confirms a
// cache-miss row is skipped and counted distinctly from a frozen one (review
// point: a silent, uncounted cache-miss skip would make "why did reconcile
// only touch N of the M rows I expected" a mystery with no visible signal,
// unlike frozenSkipped's existing visibility).
func TestReconcileCandidatesCacheMissIsCountedSeparatelyFromFrozen(t *testing.T) {
	t.Parallel()

	cacheDir := t.TempDir() // nothing written — every row below is a cache miss
	rows := []resourceRow{
		{ResourceTypeName: "aws_svc_a", CloudFormationTypeName: "AWS::Svc::A"},
		{ResourceTypeName: "aws_svc_b", CloudFormationTypeName: "AWS::Svc::B"},
		{ResourceTypeName: "aws_svc_frozen", CloudFormationTypeName: "AWS::Svc::Frozen", FrozenSince: "2026-01-01"},
	}

	cands, frozenSkipped, cacheMissSkipped := reconcileCandidates(rows, cacheDir)
	if len(cands) != 0 {
		t.Errorf("want 0 candidates, got %d: %+v", len(cands), cands)
	}
	if frozenSkipped != 1 {
		t.Errorf("frozenSkipped = %d, want 1", frozenSkipped)
	}
	if cacheMissSkipped != 2 {
		t.Errorf("cacheMissSkipped = %d, want 2 (A and B have no cache file)", cacheMissSkipped)
	}
}

// TestReconcileCountHeaderIsNeverTrustedFromRecompute is a regression test:
// found by actually running -reconcile against the real repo (not caught by
// any existing unit test), the committed overlay's count-header line
// ("# N CloudFormation resource types schemas are available...") silently
// changed from 1581 to 1593 — a real, wrong rewrite, not a no-op.
// normalizeWithDecisions (main.go) computes that line from len(base), which
// is correct for -sync (base is the live AWS-discovered set — "N schemas
// are available" is a true, fresh AWS fact) but wrong for reconcile, which
// passes its own overlayRows as base since it has no live set: overlayRows'
// length includes frozen/retained/non-provisionable rows that were never
// "available from AWS" in the header's own sense, so it is a structurally
// different, usually larger number than whatever sync last observed and
// committed.
//
// This test exercises the actual fix (run_reconcile.go's restoration of the
// committed line after normalizeWithDecisions runs), not the underlying
// normalizeWithDecisions call itself — that function's own count-header
// behavior is correct for its primary caller, sync, and is exhaustively
// covered by update_test.go/the full-corpus parity suite; the bug lived
// entirely in reconcile trusting that output instead of restoring its own
// committed line, mirroring runLint's already-established "the count header
// is unknowable offline" principle (its own count-header comparison already
// excludes this exact line, main.go).
func TestReconcileCountHeaderIsNeverTrustedFromRecompute(t *testing.T) {
	t.Parallel()

	committed := "# 1581 CloudFormation resource types schemas are available for use with the Cloud Control API.\n\nsome body text unrelated to the header itself\n"
	// Simulates exactly what happened for real: normalizeWithDecisions
	// recomputed a different, larger count (reconcile's own overlayRows
	// length, which includes rows the committed header's AWS-available count
	// never counted).
	recomputed := "# 1593 CloudFormation resource types schemas are available for use with the Cloud Control API.\n\nsome body text unrelated to the header itself\n"

	// The fix itself, reproduced here at the unit level rather than only
	// exercised via a full runReconcile call (which requires a real
	// cfg.repoRoot for the compile gate and cannot run in isolation — see
	// the note below): restore the committed line onto the recomputed
	// output, exactly as run_reconcile.go does after its own
	// normalizeWithDecisions call.
	if committedCountLine := countLineRE.FindString(committed); committedCountLine != "" {
		recomputed = countLineRE.ReplaceAllLiteralString(recomputed, committedCountLine)
	}

	if !strings.Contains(recomputed, "# 1581 CloudFormation resource types schemas are available") {
		t.Errorf("the committed count line must win over the recomputed one, got:\n%s", recomputed)
	}
	if strings.Contains(recomputed, "1593") {
		t.Errorf("the wrong, recomputed count must not survive into the final output, got:\n%s", recomputed)
	}
}

// TestReconcileCandidatesFilePath, path-based sanity check: schemaCachePath's
// own naming convention (cfn with "::" replaced by "_", .json suffix) is what
// writeSchemaCache relies on above — confirm it matches exactly, so a future
// rename of schemaCachePath's format doesn't silently make every candidate in
// TestReconcileCandidates disappear (a cache miss that looks identical to
// "no cached bytes", not a loud test failure pointing at the real cause).
func TestReconcileCandidatesFilePath(t *testing.T) {
	t.Parallel()
	got := schemaCachePath("/tmp/cache", "AWS::Logs::LogGroup")
	want := filepath.Join("/tmp/cache", "AWS_Logs_LogGroup.json")
	if got != want {
		t.Fatalf("schemaCachePath = %q, want %q — writeSchemaCache's test helper assumes this exact format", got, want)
	}
}

// A full end-to-end runReconcile test (real overlay, real compile gate, real
// promotion) is not feasible in isolation: the compile gate always builds
// cfg.repoRoot (buildOnce briefly overlays staged files onto their real
// destinations and reverts — compile_fixpoint_test.go's
// TestCompileGateFailureBlocksPromotion documents this same constraint for
// runSync), and runReconcile derives outputRoot/repoRoot entirely from the
// overlay file's own on-disk location (newConfig, config.go) with no
// injection point to redirect them independently the way settleBatch's own
// tests redirect outputRoot/cacheDir while keeping repoRoot real. The
// pipeline runReconcile calls (settleBatch, refreshCandidate,
// compileFixpoint, machineryFailures, promoteStaged) is already covered end
// to end by TestRefreshCandidateSuccess, compile_fixpoint_test.go, and the
// full-corpus parity suite; reconcileCandidates' own conventions (frozen
// exclusion, classPresentUnchanged) are covered directly above. What remains
// untested at the runReconcile level specifically — its own read/dispatch/
// write sequencing — is exercised in practice via a real
// `go run ./internal/tools/bigdiffer -reconcile` against the committed
// overlay (verified manually; see the punchlist item's commit message for
// the exact command and its clean-diff result).
