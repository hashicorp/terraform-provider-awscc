// SPDX-License-Identifier: MPL-2.0

package main

import (
	"os"
	"path/filepath"
	"testing"
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

// TestRefreshCandidateSuccess generates a real type from its cached bytes into a
// temp tree and confirms both the generated files and the schema cache are
// promoted.
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
	gr, err := refreshCandidate(cfg, staging, c)
	if err != nil {
		t.Fatalf("refreshCandidate: %v", err)
	}
	if !gr.ok() {
		t.Fatalf("expected clean generation, got %+v", gr.artifacts)
	}

	if _, err := os.Stat(filepath.Join(cfg.outputRoot, "aws", "logs", "log_group_resource_gen.go")); err != nil {
		t.Errorf("generated file not promoted: %v", err)
	}
	if _, err := os.Stat(schemaCachePath(cfg.cacheDir, "AWS::Logs::LogGroup")); err != nil {
		t.Errorf("schema cache not promoted: %v", err)
	}
}

// TestRefreshCandidateNeverRegress feeds unparseable bytes and confirms the
// broken generation neither overwrites the last-good cache nor writes output.
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
	gr, err := refreshCandidate(cfg, staging, c)
	if err != nil {
		t.Fatalf("refreshCandidate returned error (failure should be in-band): %v", err)
	}
	if gr.ok() {
		t.Fatal("expected broken generation for garbage bytes")
	}

	b, err := os.ReadFile(cachePath)
	if err != nil || string(b) != "GOOD-LAST-KNOWN" {
		t.Errorf("cache regressed on failure: %q (%v)", b, err)
	}
	if _, err := os.Stat(filepath.Join(cfg.outputRoot, "aws", "logs", "log_group_resource_gen.go")); !os.IsNotExist(err) {
		t.Errorf("output written despite failed generation: %v", err)
	}
}
