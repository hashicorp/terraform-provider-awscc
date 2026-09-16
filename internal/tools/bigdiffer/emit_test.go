// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// TestRegistrationParity asserts the single registration file bigdiffer emits
// imports exactly the set of service packages that the three committed directive
// files (resources.go, singular_data_sources.go, plural_data_sources.go) import
// between them — so registration coverage is identical.
func TestRegistrationParity(t *testing.T) {
	t.Parallel()
	cfg, rows := loadCorpus(t)

	reg, err := emitRegistration(cfg, rows)
	if err != nil {
		t.Fatalf("emitRegistration: %v", err)
	}
	got := blankImports(string(reg))

	want := map[string]struct{}{}
	for _, name := range []string{"resources.go", "singular_data_sources.go", "plural_data_sources.go"} {
		b, err := os.ReadFile(filepath.Join(cfg.overlayDir, name))
		if err != nil {
			t.Fatalf("reading %s: %v", name, err)
		}
		for p := range blankImports(string(b)) {
			want[p] = struct{}{}
		}
	}

	for p := range want {
		if _, ok := got[p]; !ok {
			t.Errorf("registration missing package imported by legacy: %s", p)
		}
	}
	for p := range got {
		if _, ok := want[p]; !ok {
			t.Errorf("registration imports package the legacy files do not: %s", p)
		}
	}
	t.Logf("registration: %d packages (legacy union %d)", len(got), len(want))
}

// blankImports extracts the set of blank-imported package paths under
// internal/aws/ from Go source text.
func blankImports(src string) map[string]struct{} {
	out := map[string]struct{}{}
	for line := range strings.SplitSeq(src, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "_ \"") {
			continue
		}
		if i := strings.IndexByte(line, '"'); i >= 0 {
			if p, err := strconv.Unquote(line[i:]); err == nil && strings.Contains(p, "/internal/aws/") {
				out[p] = struct{}{}
			}
		}
	}
	return out
}

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
