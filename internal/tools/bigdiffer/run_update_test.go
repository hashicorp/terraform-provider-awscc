// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/cli"
	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/codegen"
)

func logGroupRow(t *testing.T, rows []resourceRow) resourceRow {
	t.Helper()
	for _, r := range rows {
		if r.CloudFormationTypeName == "AWS::Logs::LogGroup" {
			return r
		}
	}
	t.Skip("AWS::Logs::LogGroup not in overlay")
	return resourceRow{}
}

// TestRefreshCandidateSuccess generates a real type from its cached bytes into
// a staging tree, promotes it, and confirms both the generated files and the
// schema cache land in the real (here, temp) tree.
func TestRefreshCandidateSuccess(t *testing.T) {
	t.Parallel()
	cfg, rows := loadCorpus(t)
	lg := logGroupRow(t, rows)

	schema, err := os.ReadFile(schemaCachePath(cfg.cacheDir, "AWS::Logs::LogGroup"))
	if err != nil {
		t.Fatalf("reading committed schema: %v", err)
	}

	tmp := t.TempDir()
	cfg.outputRoot = filepath.Join(tmp, "out")
	cfg.cacheDir = filepath.Join(tmp, "cache")
	staging := filepath.Join(tmp, "staging")
	if err := os.MkdirAll(staging, 0o755); err != nil {
		t.Fatal(err)
	}

	c := candidate{cfType: "AWS::Logs::LogGroup", class: classPresent, row: lg, schema: schema}
	gr, _, err := refreshCandidate(cfg, staging, c)
	if err != nil {
		t.Fatalf("refreshCandidate: %v", err)
	}
	if !gr.ok() {
		t.Fatalf("expected clean generation, got %+v", gr.artifacts)
	}
	if err := promoteStaged(cfg, staging); err != nil {
		t.Fatalf("promoteStaged: %v", err)
	}

	if _, err := os.Stat(filepath.Join(cfg.outputRoot, "aws", "logs", "log_group_resource_gen.go")); err != nil {
		t.Errorf("generated file not promoted: %v", err)
	}
	if _, err := os.Stat(schemaCachePath(cfg.cacheDir, "AWS::Logs::LogGroup")); err != nil {
		t.Errorf("schema cache not promoted: %v", err)
	}
}

// TestRefreshCandidateNeverRegress feeds unparseable schema bytes for a type
// whose plural data source is real (AWS::Logs::LogGroup). Per-artifact
// independence (generation-punchlist.md item 4) means this is a *partial*
// failure, not a total one: the plural data source generates from the
// CloudFormation type name alone (codegen.GeneratePluralDataSource never reads
// the schema file), so it succeeds regardless of the garbage bytes, while the
// resource and singular data source — which do need a parseable schema — fail.
// The plural artifact is promoted (and the cache advances to the new, if
// invalid-looking, bytes, since something real was generated against them);
// the resource is not written, and its last-good file (none exists here) is
// not regressed.
func TestRefreshCandidateNeverRegress(t *testing.T) {
	t.Parallel()
	cfg, rows := loadCorpus(t)
	lg := logGroupRow(t, rows)

	tmp := t.TempDir()
	cfg.outputRoot = filepath.Join(tmp, "out")
	cfg.cacheDir = filepath.Join(tmp, "cache")
	if err := os.MkdirAll(cfg.cacheDir, 0o755); err != nil {
		t.Fatal(err)
	}
	cachePath := schemaCachePath(cfg.cacheDir, "AWS::Logs::LogGroup")
	if err := os.WriteFile(cachePath, []byte("GOOD-LAST-KNOWN"), 0o644); err != nil {
		t.Fatal(err)
	}
	staging := filepath.Join(tmp, "staging")
	if err := os.MkdirAll(staging, 0o755); err != nil {
		t.Fatal(err)
	}

	c := candidate{cfType: "AWS::Logs::LogGroup", class: classPresent, row: lg, schema: []byte("{ not a valid schema }")}
	gr, _, err := refreshCandidate(cfg, staging, c)
	if err != nil {
		t.Fatalf("refreshCandidate returned error (failure should be in-band): %v", err)
	}
	if gr.ok() {
		t.Fatal("expected the resource and singular data source to fail on garbage bytes")
	}
	if !gr.anyOK() {
		t.Fatal("expected the plural data source to succeed independent of the garbage bytes")
	}
	if err := promoteStaged(cfg, staging); err != nil {
		t.Fatalf("promoteStaged: %v", err)
	}

	if _, err := os.Stat(filepath.Join(cfg.outputRoot, "aws", "logs", "log_group_resource_gen.go")); !os.IsNotExist(err) {
		t.Errorf("resource output written despite failed generation: %v", err)
	}
	if _, err := os.Stat(filepath.Join(cfg.outputRoot, "aws", "logs", "log_group_plural_data_source_gen.go")); err != nil {
		t.Errorf("plural data source output should have been promoted: %v", err)
	}
}

// TestRefreshCandidateTotalFailureNeverRegress covers the total-failure case
// TestRefreshCandidateNeverRegress used to: a plan-level error (here, a
// malformed Terraform type name) fails before any artifact is even derived, so
// nothing is promoted and the cache is untouched.
func TestRefreshCandidateTotalFailureNeverRegress(t *testing.T) {
	t.Parallel()
	cfg, rows := loadCorpus(t)
	lg := logGroupRow(t, rows)
	lg.ResourceTypeName = "not_a_valid_type_name" // fails naming.ParseTerraformTypeName in generationPlan

	tmp := t.TempDir()
	cfg.outputRoot = filepath.Join(tmp, "out")
	cfg.cacheDir = filepath.Join(tmp, "cache")
	if err := os.MkdirAll(cfg.cacheDir, 0o755); err != nil {
		t.Fatal(err)
	}
	cachePath := schemaCachePath(cfg.cacheDir, "AWS::Logs::LogGroup")
	if err := os.WriteFile(cachePath, []byte("GOOD-LAST-KNOWN"), 0o644); err != nil {
		t.Fatal(err)
	}
	staging := filepath.Join(tmp, "staging")
	if err := os.MkdirAll(staging, 0o755); err != nil {
		t.Fatal(err)
	}

	c := candidate{cfType: "AWS::Logs::LogGroup", class: classPresent, row: lg, schema: []byte("irrelevant")}
	gr, _, err := refreshCandidate(cfg, staging, c)
	if err != nil {
		t.Fatalf("refreshCandidate returned error (failure should be in-band): %v", err)
	}
	if gr.anyOK() {
		t.Fatal("expected total failure: a malformed type name fails before any artifact is derived")
	}
	if err := promoteStaged(cfg, staging); err != nil {
		t.Fatalf("promoteStaged: %v", err)
	}

	b, err := os.ReadFile(cachePath)
	if err != nil || string(b) != "GOOD-LAST-KNOWN" {
		t.Errorf("cache regressed on total failure: %q (%v)", b, err)
	}
	if _, err := os.Stat(cfg.outputRoot); !os.IsNotExist(err) {
		t.Errorf("output written despite total failure: %v", err)
	}
}

// TestReconcileListResourceDropsListResourceOnFailedPlural exercises the
// resource↔plural ListResource coupling directly (generation-punchlist.md item
// 4's "Done when": a New/Present type whose resource succeeds but whose plural
// data source fails must not be promoted advertising a list resource with no
// working plural data source behind it).
func TestReconcileListResourceDropsListResourceOnFailedPlural(t *testing.T) {
	t.Parallel()
	cfg, rows := loadCorpus(t)
	lg := logGroupRow(t, rows)
	p, err := generationPlan(lg, cfg.prefix, cfg.cacheDir)
	if err != nil {
		t.Fatalf("generationPlan: %v", err)
	}

	var resArt, pluralArt genArtifact
	for _, a := range p.artifacts {
		switch a.kind {
		case artifactResource:
			resArt = a
		case artifactPluralDataSource:
			pluralArt = a
		}
	}
	if !resArt.listResource {
		t.Fatalf("expected AWS::Logs::LogGroup's resource to request ListResource, got %+v", resArt)
	}

	ui := &cli.BasicUi{Writer: io.Discard, ErrorWriter: io.Discard}
	code, test, err := codegen.GenerateResource(ui, p.schemaFile, resArt.tfType, resArt.packageName, cfg.servicesPath, true)
	if err != nil {
		t.Fatalf("GenerateResource: %v", err)
	}
	results := []genResult{
		{p: p, a: resArt, code: code, test: test},
		{p: p, a: pluralArt, err: errors.New("plural boom")}, // simulate a failed plural artifact
	}

	reconciled, err := reconcileListResource(cfg, results)
	if err != nil {
		t.Fatalf("reconcileListResource: %v", err)
	}
	if reconciled[0].a.listResource {
		t.Error("resource should be regenerated without ListResource once the plural artifact fails")
	}
	if reconciled[0].err != nil {
		t.Errorf("regenerated resource should still succeed, got %v", reconciled[0].err)
	}
	if bytes.Equal(reconciled[0].code, code) {
		t.Error("resource code should differ once GenerateListResource is dropped")
	}
}

// TestRefreshCandidateLeavesWorkingPairAlone confirms the common case is a
// no-op: a resource whose plural data source succeeded is untouched.
func TestReconcileListResourceLeavesWorkingPairAlone(t *testing.T) {
	t.Parallel()
	cfg, rows := loadCorpus(t)
	lg := logGroupRow(t, rows)
	results := generateCorpus(cfg, []resourceRow{lg}, 1, nil)

	reconciled, err := reconcileListResource(cfg, results)
	if err != nil {
		t.Fatalf("reconcileListResource: %v", err)
	}
	for i, r := range reconciled {
		if r.a.kind == artifactResource && !r.a.listResource {
			t.Errorf("resource should keep ListResource when the plural artifact succeeded, got %+v", reconciled[i])
		}
	}
}

// TestUpdateBatchAtomicity is the item-12 regression test
// (generation-punchlist.md "Never-regress cross-type atomicity"): it drives the
// same stage-everything-then-promote-once sequence runUpdate uses across a
// multi-candidate batch, injects a hard failure partway through (a candidate
// whose staging write fails outright, mirroring what a disk error mid-batch
// looks like to the loop), and asserts that once that error aborts the batch
// before promoteStaged is ever called, the real tree is left exactly as it
// started — nothing from the earlier, successfully staged candidates leaks
// through, and there is nothing to reconcile into the overlay.
func TestUpdateBatchAtomicity(t *testing.T) {
	t.Parallel()
	cfg, rows := loadCorpus(t)
	lg := logGroupRow(t, rows)
	schema, err := os.ReadFile(schemaCachePath(cfg.cacheDir, "AWS::Logs::LogGroup"))
	if err != nil {
		t.Fatalf("reading committed schema: %v", err)
	}

	tmp := t.TempDir()
	cfg.outputRoot = filepath.Join(tmp, "out")
	cfg.cacheDir = filepath.Join(tmp, "cache")
	staging := filepath.Join(tmp, "staging")

	// Candidate 1 (of what would be a larger batch) stages successfully.
	c1 := candidate{cfType: "AWS::Logs::LogGroup", class: classPresent, row: lg, schema: schema}
	if _, _, err := refreshCandidate(cfg, staging, c1); err != nil {
		t.Fatalf("refreshCandidate (candidate 1): %v", err)
	}

	// Candidate 2 hits a hard I/O error while staging — simulated by occupying
	// the staging input directory's own path with a plain file, so
	// os.MkdirAll fails outright, the same shape of failure a disk error
	// mid-batch would produce. This mirrors runUpdate's
	// `if err != nil { return err }`: the batch loop aborts immediately.
	badStagingDir := filepath.Join(staging, "input")
	if err := os.MkdirAll(filepath.Dir(badStagingDir), dirPerm); err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll(badStagingDir); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(badStagingDir, []byte("occupied"), filePerm); err != nil {
		t.Fatal(err)
	}
	c2 := candidate{cfType: "AWS::Logs::LogGroup", class: classPresent, row: lg, schema: schema}
	_, _, err = refreshCandidate(cfg, staging, c2)
	if err == nil {
		t.Fatal("expected refreshCandidate to fail: the staging input dir path is occupied by a file")
	}

	// The batch loop aborts here (runUpdate: `if err != nil { return err }`),
	// so promoteStaged is never reached, and the overlay reconcile step after
	// it never runs. Confirm the real tree — which candidate 1 alone would have
	// left non-empty, had it been promoted — is untouched.
	if _, statErr := os.Stat(cfg.outputRoot); !os.IsNotExist(statErr) {
		t.Errorf("real output tree should not exist: promotion must not have happened, got %v", statErr)
	}
	if _, statErr := os.Stat(cfg.cacheDir); !os.IsNotExist(statErr) {
		t.Errorf("real cache dir should not exist: promotion must not have happened, got %v", statErr)
	}
}
