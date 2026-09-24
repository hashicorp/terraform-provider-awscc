// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// TestImportExamplesParity confirms bigdiffer emits import_examples_gen.json
// byte-identical to the committed aggregate. Re-derives every resource's
// identifiers, so it is slow — skipped under -short.
func TestImportExamplesParity(t *testing.T) {
	if testing.Short() {
		t.Skip("re-derives identifiers for the whole corpus; run without -short")
	}
	cfg, rows := loadCorpus(t)

	got, err := emitImportExamples(cfg, rows)
	if err != nil {
		t.Fatalf("emitImportExamples: %v", err)
	}
	want, err := os.ReadFile(filepath.Join(cfg.overlayDir, "import_examples_gen.json"))
	if err != nil {
		t.Fatalf("reading committed import_examples_gen.json: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("import_examples_gen.json parity mismatch:\n%s", firstDiff(want, got))
	}
}
