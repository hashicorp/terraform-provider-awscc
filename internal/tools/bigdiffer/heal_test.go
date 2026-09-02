// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestSuppressionComment(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		text string
		want string
	}{
		{
			name: "inline text after colon",
			text: `resource_schema "x" {
  cloudformation_type_name = "AWS::X::Y"

  # Suppression Reason: duplicate attribute name mapping for CloudFormation property Id
  suppress_resource_generation = true
}`,
			want: "duplicate attribute name mapping for CloudFormation property Id",
		},
		{
			name: "multi-line comment block",
			text: `resource_schema "x" {
  cloudformation_type_name = "AWS::X::Y"

  # Suppression Reason:
  # Recursive Attribute Definitions
  # https://github.com/hashicorp/terraform-provider-awscc/issues/95
  suppress_resource_generation = true
}`,
			want: "Recursive Attribute Definitions https://github.com/hashicorp/terraform-provider-awscc/issues/95",
		},
		{
			name: "case-insensitive, different casing",
			text: `resource_schema "x" {
  # suppression reason: recursive object definitions
  suppress_resource_generation = true
}`,
			want: "recursive object definitions",
		},
		{
			name: "no comment at all",
			text: `resource_schema "x" {
  cloudformation_type_name = "AWS::X::Y"
  suppress_resource_generation = true
}`,
			want: "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := suppressionComment(tc.text); got != tc.want {
				t.Errorf("suppressionComment() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestCommentOrUnknown(t *testing.T) {
	t.Parallel()

	t.Run("migrates an existing comment to manual", func(t *testing.T) {
		t.Parallel()
		p := commentOrUnknown(healProposal{cfn: "AWS::X::Y"}, "recursive schema", false)
		if p.action != "reason" {
			t.Errorf("action = %q, want reason", p.action)
		}
		if p.reason != "manual: recursive schema" {
			t.Errorf("reason = %q, want manual: recursive schema", p.reason)
		}
	})

	t.Run("falls back to unknown with no comment", func(t *testing.T) {
		t.Parallel()
		p := commentOrUnknown(healProposal{cfn: "AWS::X::Y"}, "", false)
		if !strings.HasPrefix(p.reason, "unknown:") {
			t.Errorf("reason = %q, want unknown: prefix", p.reason)
		}
	})

	t.Run("multiPending offers the comment as a shared candidate, not a confirmed fact", func(t *testing.T) {
		t.Parallel()
		// Review resolution (suppressed-and-frozen.md, "-heal: re-probe and
		// fill gaps"): a row-level comment cannot be assumed to
		// describe more than one still-reason-less fact, so when more than
		// one fact is pending it must be visibly marked as a shared
		// candidate for a human to confirm per field, not silently
		// duplicated into each field as if confirmed.
		p := commentOrUnknown(healProposal{cfn: "AWS::X::Y"}, "recursive schema", true)
		if p.action != "reason" {
			t.Errorf("action = %q, want reason", p.action)
		}
		if !strings.HasPrefix(p.reason, "manual: recursive schema") {
			t.Errorf("reason = %q, want the manual: reason preserved verbatim as a prefix", p.reason)
		}
		if !strings.Contains(p.reason, "candidate") {
			t.Errorf("reason = %q, want it visibly marked as a shared candidate when multiPending", p.reason)
		}
	})

	t.Run("multiPending with no comment still falls back to unknown, unmarked", func(t *testing.T) {
		t.Parallel()
		p := commentOrUnknown(healProposal{cfn: "AWS::X::Y"}, "", true)
		if !strings.HasPrefix(p.reason, "unknown:") {
			t.Errorf("reason = %q, want unknown: prefix regardless of multiPending", p.reason)
		}
	})
}

// TestProbeArtifactIsolatesRecursiveSchema is the item-9 regression test for
// subprocess isolation: a schema that drives the emitter into unbounded
// recursion must be reported as a normal (if slow) generation_failed proposal,
// not crash the process running the probe. A real stack overflow can take
// tens of seconds to unwind, so this is intentionally the one slow bigdiffer
// test; it is worth the cost to prove containment against the exact failure
// mode discovered live (AWS::WAFv2::RuleGroup/WebACL).
//
// probeArtifact re-execs via os.Executable(), which under `go test` resolves
// to the test binary — a binary whose CLI is testing.Main, not bigdiffer's
// real main(), so it can't be probed directly here. This test instead builds
// a real bigdiffer binary (mirroring parity_test.go's `go run` pattern) and
// runs probeArtifactWithBinary against it, exercising the exact same
// exec.CommandContext/GOMEMLIMIT/timeout logic probeArtifact uses.
func TestProbeArtifactIsolatesRecursiveSchema(t *testing.T) {
	if testing.Short() {
		t.Skip("builds a binary and spawns a subprocess that can take tens of seconds to stack-overflow")
	}
	t.Parallel()
	cfg, rows := loadCorpus(t)

	var recursive resourceRow
	found := false
	for _, r := range rows {
		if r.CloudFormationTypeName == "AWS::WAFv2::WebACL" {
			recursive, found = r, true
			break
		}
	}
	if !found {
		t.Skip("AWS::WAFv2::WebACL not in overlay")
	}

	path := recursive.CloudFormationSchemaPath
	if path == "" {
		path = schemaCachePath(cfg.cacheDir, recursive.CloudFormationTypeName)
	}
	schema, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading schema: %v", err)
	}

	bin := filepath.Join(t.TempDir(), "bigdiffer")
	build := exec.Command("go", "build", "-o", bin, ".")
	build.Dir = "."
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("building bigdiffer for the isolation test: %v\n%s", err, out)
	}

	err = probeArtifactWithBinary(bin, cfg, recursive, artifactResource, schema)
	if err == nil {
		t.Fatal("expected the recursive schema to fail generation, not succeed")
	}
	if !strings.Contains(err.Error(), "stack overflow") && !strings.Contains(err.Error(), "crashed") {
		t.Errorf("expected a crash/stack-overflow signature in the error, got: %v", err)
	}
}

// TestProbeArtifactWithBinaryEmptyOutputCrash is the fix for the crash-vs-
// clean-failure classification review note: a process that dies with no
// self-reported output (e.g. a silent OOM/SIGKILL before it could print
// anything) must not produce an empty, useless error — it should be reported
// as a crash with a non-empty placeholder message, not errors.New("").
func TestProbeArtifactWithBinaryEmptyOutputCrash(t *testing.T) {
	t.Parallel()

	// A trivial script standing in for "exec succeeded, process ran, but
	// produced zero output and exited non-zero" — the exact shape a
	// signal-killed process with nothing printed yet would leave behind.
	dir := t.TempDir()
	script := filepath.Join(dir, "silent-fail.sh")
	if err := os.WriteFile(script, []byte("#!/bin/sh\nexit 1\n"), 0o755); err != nil {
		t.Fatal(err)
	}

	cfg := config{}
	row := resourceRow{ResourceTypeName: "aws_svc_thing", CloudFormationTypeName: "AWS::Svc::Thing"}
	err := probeArtifactWithBinary(script, cfg, row, artifactResource, []byte("{}"))
	if err == nil {
		t.Fatal("expected an error")
	}
	if err.Error() == "" {
		t.Error("error message must not be empty even when the process produced no output")
	}
	if !strings.Contains(err.Error(), "crashed") {
		t.Errorf("a non-ExitError-with-real-output failure should be classified as a crash, got: %v", err)
	}
}

// TestProbeArtifactWithBinaryCleanFailure confirms the common case is
// unaffected by the crash-classification fix: a process that exits non-zero
// but does report real output (a normal generation error) is classified as a
// clean, reported failure, not a crash.
func TestProbeArtifactWithBinaryCleanFailure(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	script := filepath.Join(dir, "clean-fail.sh")
	if err := os.WriteFile(script, []byte("#!/bin/sh\necho 'emitting schema code: boom' >&2\nexit 1\n"), 0o755); err != nil {
		t.Fatal(err)
	}

	cfg := config{}
	row := resourceRow{ResourceTypeName: "aws_svc_thing", CloudFormationTypeName: "AWS::Svc::Thing"}
	err := probeArtifactWithBinary(script, cfg, row, artifactResource, []byte("{}"))
	if err == nil {
		t.Fatal("expected an error")
	}
	if strings.Contains(err.Error(), "crashed") {
		t.Errorf("a reported failure with real output should not be classified as a crash, got: %v", err)
	}
	if !strings.Contains(err.Error(), "boom") {
		t.Errorf("expected the script's own error text to be preserved, got: %v", err)
	}
}

// TestRunHealProbeArtifactPassesCompileGateForARealType is the happy path:
// probing a real, already-working type must still report success once the
// compile gate step (added so -heal's "lift" proposals are trustworthy
// against both stages a real -update run enforces, not just generation) is
// exercised against the real module.
func TestRunHealProbeArtifactPassesCompileGateForARealType(t *testing.T) {
	if testing.Short() {
		t.Skip("runs a real go build ./... against the module")
	}
	// Not t.Parallel() — see TestRunHealProbeArtifactCatchesACompileGateFailure
	// below: both touch the real tree via buildOnce's overlay.

	cfg, rows := loadCorpus(t)
	lg := logGroupRow(t, rows)
	schemaPath := lg.CloudFormationSchemaPath
	if schemaPath == "" {
		schemaPath = schemaCachePath(cfg.cacheDir, lg.CloudFormationTypeName)
	}

	err := runHealProbeArtifact(lg.ResourceTypeName, lg.CloudFormationTypeName, string(artifactResource),
		schemaPath, cfg.prefix, cfg.cacheDir, cfg.servicesPath, cfg.repoRoot, cfg.outputRoot)
	if err != nil {
		t.Fatalf("a real, already-working type should pass generation and the compile gate, got: %v", err)
	}
}

// TestRunHealProbeArtifactCatchesACompileGateFailure proves the actual review
// item: a type that *generates* fine can still be rejected by the compile
// gate, and runHealProbeArtifact must report that as a failure rather than
// the generation-only success -heal reported before this change. Simulated by
// pre-placing a deliberately broken sibling file in the same real package the
// probed artifact would land in — buildOnce's overlay of the probed artifact's
// own (valid) code still fails the package build because of that sibling,
// which is mechanically the same "one bad file fails the whole package"
// reality compile-gate-design.md describes; it does not require constructing
// a schema that makes the owned engine itself render invalid code.
func TestRunHealProbeArtifactCatchesACompileGateFailure(t *testing.T) {
	if testing.Short() {
		t.Skip("runs a real go build ./... against the module")
	}
	// Not t.Parallel() — mutates the real tree via a real go build overlay,
	// same rule as every other real-build test in this package.

	cfg, rows := loadCorpus(t)
	lg := logGroupRow(t, rows)
	schemaPath := lg.CloudFormationSchemaPath
	if schemaPath == "" {
		schemaPath = schemaCachePath(cfg.cacheDir, lg.CloudFormationTypeName)
	}

	// Point outputRoot at a scratch package under the real internal/aws tree
	// (same technique compile_fixpoint_test.go uses) with a deliberately
	// broken sibling file already sitting in it, then probe the real,
	// well-formed LogGroup resource artifact "as if" it belonged to that
	// package by aliasing its row's ResourceTypeName's service to the
	// scratch one — generationPlan derives packageName purely from the
	// Terraform type name, so this is a legitimate way to redirect a real,
	// valid artifact's destination without needing a broken schema.
	row := lg
	row.ResourceTypeName = "aws_zzzhealcg_log_group"

	scratchDir := filepath.Join(cfg.outputRoot, "aws", "zzzhealcg")
	if err := os.MkdirAll(scratchDir, dirPerm); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(scratchDir) })
	broken := "package zzzhealcg\nfunc AlreadyBroken() int { return \"not an int\" }\n"
	if err := os.WriteFile(filepath.Join(scratchDir, "sibling_broken.go"), []byte(broken), filePerm); err != nil {
		t.Fatal(err)
	}

	err := runHealProbeArtifact(row.ResourceTypeName, row.CloudFormationTypeName, string(artifactResource),
		schemaPath, cfg.prefix, cfg.cacheDir, cfg.servicesPath, cfg.repoRoot, cfg.outputRoot)
	if err == nil {
		t.Fatal("expected the compile gate to reject the package due to its broken sibling file")
	}
	if !strings.Contains(err.Error(), "compile gate") && !strings.Contains(err.Error(), "fails the compile gate") {
		t.Errorf("expected a compile-gate-flavored error, got: %v", err)
	}
	// buildOnce must have reverted its overlay regardless of outcome — no
	// trace of the probed artifact's own file should remain.
	if _, statErr := os.Stat(filepath.Join(scratchDir, "log_group_resource_gen.go")); !os.IsNotExist(statErr) {
		t.Errorf("the compile gate's build overlay must be reverted; the probed artifact's file should not exist, stat err = %v", statErr)
	}
}

// TestHealArtifactTagsBuildFailedDistinctFromGenerationFailed is the fix for
// the review finding: healArtifact must tag a proposal build_failed when the
// artifact generates cleanly but is rejected by the compile gate, and must
// keep tagging generation_failed when generation itself fails — the two
// stages must not collapse into the same category, which would defeat the
// whole point of wiring the compile gate into -heal (a human reviewing
// proposals needs to know which stage actually failed).
func TestHealArtifactTagsBuildFailedDistinctFromGenerationFailed(t *testing.T) {
	if testing.Short() {
		t.Skip("builds a binary and runs real go build ./... against the module")
	}
	// Not t.Parallel() — swaps the package-level probeArtifact var and
	// mutates the real tree via buildOnce's overlay in the gate-failure case.

	cfg, rows := loadCorpus(t)
	lg := logGroupRow(t, rows)
	schemaPath := lg.CloudFormationSchemaPath
	if schemaPath == "" {
		schemaPath = schemaCachePath(cfg.cacheDir, lg.CloudFormationTypeName)
	}
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		t.Fatalf("reading schema: %v", err)
	}

	bin := filepath.Join(t.TempDir(), "bigdiffer")
	build := exec.Command("go", "build", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("building bigdiffer: %v\n%s", err, out)
	}

	origProbe := probeArtifact
	t.Cleanup(func() { probeArtifact = origProbe })

	t.Run("gate failure tags build_failed", func(t *testing.T) {
		row := lg
		row.ResourceTypeName = "aws_zzzhealtag_log_group"

		scratchDir := filepath.Join(cfg.outputRoot, "aws", "zzzhealtag")
		if err := os.MkdirAll(scratchDir, dirPerm); err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() { _ = os.RemoveAll(scratchDir) })
		broken := "package zzzhealtag\nfunc AlreadyBroken() int { return \"not an int\" }\n"
		if err := os.WriteFile(filepath.Join(scratchDir, "sibling_broken.go"), []byte(broken), filePerm); err != nil {
			t.Fatal(err)
		}

		gateCfg := cfg
		probeArtifact = func(c config, r resourceRow, kind artifactKind, s []byte) error {
			return probeArtifactWithBinary(bin, c, r, kind, s)
		}

		proposal := healArtifact(gateCfg, row, healFact{kind: artifactResource, field: attrSuppressionReasonResource}, schema, nil, "", false)
		if proposal.action != "reason" {
			t.Fatalf("expected action=reason for a gate failure, got %q", proposal.action)
		}
		if !strings.HasPrefix(proposal.reason, string(reasonBuildFailed)+":") {
			t.Errorf("expected a build_failed-tagged reason, got %q", proposal.reason)
		}
		if strings.HasPrefix(proposal.reason, string(reasonGenerationFailed)+":") {
			t.Errorf("must not collapse a gate failure into generation_failed, got %q", proposal.reason)
		}
	})

	t.Run("generation failure still tags generation_failed", func(t *testing.T) {
		probeArtifact = func(config, resourceRow, artifactKind, []byte) error {
			return errors.New("boom: invalid schema")
		}
		proposal := healArtifact(cfg, lg, healFact{kind: artifactResource, field: attrSuppressionReasonResource}, schema, nil, "", false)
		if proposal.action != "reason" {
			t.Fatalf("expected action=reason, got %q", proposal.action)
		}
		if !strings.HasPrefix(proposal.reason, string(reasonGenerationFailed)+":") {
			t.Errorf("expected a generation_failed-tagged reason, got %q", proposal.reason)
		}
	})
}

// TestHealFactsForNeedsHealing is the per-artifact gating fix itself (item
// 9a/9b): a row's four facts (resource, singular DS, plural DS, freeze) must
// each be judged independently — a row with a real reason on one fact and no
// reason on another must surface only the still-reason-less one, not the
// whole row as a unit (the pre-9b row-level gate's bug) and not silently skip
// the still-broken one either.
func TestHealFactsForNeedsHealing(t *testing.T) {
	t.Parallel()

	t.Run("only the still-reason-less fact needs healing", func(t *testing.T) {
		t.Parallel()
		row := resourceRow{
			SuppressResourceGeneration:         true,
			SuppressionReasonResource:          "manual: recursive schema", // already reasoned
			SuppressPluralDataSourceGeneration: true,
			SuppressionReasonPluralDataSource:  "", // still reason-less — this is 9a's motivating case
		}
		var pending []string
		for _, f := range healFactsFor(row) {
			if f.needsHealing() {
				pending = append(pending, f.field)
			}
		}
		if len(pending) != 1 || pending[0] != attrSuppressionReasonPlural {
			t.Errorf("pending = %v, want only %s — the resource already has a real reason and must not be re-surfaced", pending, attrSuppressionReasonPlural)
		}
	})

	t.Run("a fact tagged unknown still needs healing", func(t *testing.T) {
		t.Parallel()
		row := resourceRow{
			SuppressResourceGeneration: true,
			SuppressionReasonResource:  "unknown: no schema cached and no existing comment to migrate; needs a human look",
		}
		facts := healFactsFor(row)
		if !facts[0].needsHealing() {
			t.Error("an unknown:-tagged reason is still in the backlog and must need healing")
		}
	})

	t.Run("an inactive fact never needs healing regardless of its reason field", func(t *testing.T) {
		t.Parallel()
		row := resourceRow{SuppressResourceGeneration: false, SuppressionReasonResource: ""}
		facts := healFactsFor(row)
		if facts[0].needsHealing() {
			t.Error("a fact that isn't active (not suppressed/frozen) must never be proposed for healing")
		}
	})

	t.Run("the freeze fact is independent of the three artifact facts", func(t *testing.T) {
		t.Parallel()
		row := resourceRow{
			FrozenSince:                today2026,
			FrozenReason:               "",
			SuppressResourceGeneration: true,
			SuppressionReasonResource:  "manual: recursive schema",
		}
		var pending []string
		for _, f := range healFactsFor(row) {
			if f.needsHealing() {
				pending = append(pending, f.field)
			}
		}
		if len(pending) != 1 || pending[0] != attrFrozenReason {
			t.Errorf("pending = %v, want only %s — the freeze has its own empty reason independent of the already-reasoned resource", pending, attrFrozenReason)
		}
	})
}

const today2026 = "2026-01-01"

// TestHealRowMultiPendingCommentCandidate is a healRow-level (not just
// healArtifact-level) proof of the propose-don't-auto-split resolution: a row
// with more than one still-reason-less fact and no cached schema (so every
// fact falls through to the comment/unknown fallback) must offer the same
// comment text to each pending fact, each visibly marked as a shared
// candidate — not silently duplicated as if independently confirmed.
func TestHealRowMultiPendingCommentCandidate(t *testing.T) {
	t.Parallel()

	row := resourceRow{
		CloudFormationTypeName:             "AWS::Svc::NoSchema",
		ResourceTypeName:                   "aws_svc_noschema",
		CloudFormationSchemaPath:           "/nonexistent/schema.json",
		SuppressResourceGeneration:         true,
		SuppressPluralDataSourceGeneration: true,
	}
	pending := []healFact{
		{kind: artifactResource, field: attrSuppressionReasonResource, active: true},
		{kind: artifactPluralDataSource, field: attrSuppressionReasonPlural, active: true},
	}

	proposals := healRow(config{cacheDir: "/nonexistent/cache"}, row, pending, "some legacy free-form reason")
	if len(proposals) != 2 {
		t.Fatalf("got %d proposals, want 2 (one per pending fact)", len(proposals))
	}
	for _, p := range proposals {
		if p.action != "reason" {
			t.Errorf("proposal %+v: action should be reason (no schema cached, nothing to probe)", p)
		}
		if !strings.Contains(p.reason, "some legacy free-form reason") {
			t.Errorf("proposal %+v: should carry the migrated comment text", p)
		}
		if !strings.Contains(p.reason, "candidate") {
			t.Errorf("proposal %+v: with 2 pending facts, the comment must be marked as a shared candidate, not a confirmed fact", p)
		}
	}
	fields := map[string]bool{}
	for _, p := range proposals {
		fields[p.field] = true
	}
	if !fields[attrSuppressionReasonResource] || !fields[attrSuppressionReasonPlural] {
		t.Errorf("expected one proposal per pending field, got %+v", proposals)
	}
}

// TestHealRowSinglePendingNoComment confirms the single-fact path is
// unaffected by the multiPending wording change — a row with only one
// still-reason-less fact gets a plain, unmarked reason.
func TestHealRowSinglePendingNoComment(t *testing.T) {
	t.Parallel()

	row := resourceRow{
		CloudFormationTypeName:     "AWS::Svc::NoSchema2",
		ResourceTypeName:           "aws_svc_noschema2",
		CloudFormationSchemaPath:   "/nonexistent/schema.json",
		SuppressResourceGeneration: true,
	}
	pending := []healFact{
		{kind: artifactResource, field: attrSuppressionReasonResource, active: true},
	}

	proposals := healRow(config{cacheDir: "/nonexistent/cache"}, row, pending, "")
	if len(proposals) != 1 {
		t.Fatalf("got %d proposals, want 1", len(proposals))
	}
	if !strings.HasPrefix(proposals[0].reason, "unknown:") {
		t.Errorf("proposal reason = %q, want unknown: prefix (no comment, no schema)", proposals[0].reason)
	}
	if strings.Contains(proposals[0].reason, "candidate") {
		t.Errorf("proposal reason = %q, a single pending fact must not be marked as a shared candidate", proposals[0].reason)
	}
}
