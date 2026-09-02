// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestRunGenerateOffline runs the full offline -generate flow into a temp output
// root (never the real tree) and checks it produces the per-type files plus both
// aggregates, with import_examples byte-identical to the committed one.
func TestRunGenerateOffline(t *testing.T) {
	if testing.Short() {
		t.Skip("full generation; run without -short")
	}
	cfg, rows := loadCorpus(t)
	committedImportExamples := cfg.importExamplesPath

	tmp := t.TempDir()
	cfg.outputRoot = tmp
	cfg.registrationPath = filepath.Join(tmp, "registrations_gen.go")
	cfg.importExamplesPath = filepath.Join(tmp, "import_examples_gen.json")

	if err := runGenerate(cfg, rows); err != nil {
		t.Fatalf("runGenerate: %v", err)
	}

	reg, err := os.ReadFile(cfg.registrationPath)
	if err != nil {
		t.Fatalf("reading registration: %v", err)
	}
	if !strings.Contains(string(reg), "package provider") || !strings.Contains(string(reg), "/internal/aws/logs\"") {
		t.Errorf("registration file missing expected content")
	}

	got, err := os.ReadFile(cfg.importExamplesPath)
	if err != nil {
		t.Fatalf("reading generated import_examples: %v", err)
	}
	want, err := os.ReadFile(committedImportExamples)
	if err != nil {
		t.Fatalf("reading committed import_examples: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("generated import_examples differs from committed:\n%s", firstDiff(want, got))
	}

	sample := filepath.Join(tmp, "aws", "logs", "log_group_resource_gen.go")
	if _, err := os.Stat(sample); err != nil {
		t.Errorf("expected generated sample %s: %v", sample, err)
	}
}
