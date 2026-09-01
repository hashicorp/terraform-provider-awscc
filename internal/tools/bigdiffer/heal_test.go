// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
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
		p := commentOrUnknown(healProposal{cfn: "AWS::X::Y"}, "recursive schema")
		if p.action != "reason" {
			t.Errorf("action = %q, want reason", p.action)
		}
		if p.reason != "manual: recursive schema" {
			t.Errorf("reason = %q, want manual: recursive schema", p.reason)
		}
	})

	t.Run("falls back to unknown with no comment", func(t *testing.T) {
		t.Parallel()
		p := commentOrUnknown(healProposal{cfn: "AWS::X::Y"}, "")
		if !strings.HasPrefix(p.reason, "unknown:") {
			t.Errorf("reason = %q, want unknown: prefix", p.reason)
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
