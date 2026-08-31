// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
)

// TestNoLegacyGeneratorImports enforces the cutover assurance: bigdiffer (the
// tool and every subpackage) must not import the legacy generation machinery
// under internal/provider/generators/**. Owning naming and the codegen engine is
// what lets that machinery eventually be deleted; this test keeps a stray import
// from silently reintroducing the dependency.
func TestNoLegacyGeneratorImports(t *testing.T) {
	t.Parallel()

	const forbidden = "internal/provider/generators"
	fset := token.NewFileSet()

	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		f, perr := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if perr != nil {
			return perr
		}
		for _, imp := range f.Imports {
			p := strings.Trim(imp.Path.Value, `"`)
			if strings.Contains(p, forbidden) {
				t.Errorf("%s imports legacy generator package %q; bigdiffer must own its engine", path, p)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking bigdiffer sources: %v", err)
	}
}
