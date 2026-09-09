// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codegen

import (
	"golang.org/x/tools/imports"
)

// goTabWidth is the standard Go source tab width, matching gofmt's own
// convention. Named to satisfy the mnd (magic number) linter.
const goTabWidth = 8

// FormatGo applies goimports to generated Go source: it gofmt-formats and, unlike
// go/format.Source, prunes unused imports and adds missing ones. This replicates
// the legacy generation pipeline's `goimports -w` post-step (see GNUmakefile),
// which bare gofmt does not: a template that unconditionally emits an import a
// given schema does not use (e.g. the framework "types" package) yields a
// non-compiling file under gofmt but a clean one under goimports. filename only
// identifies the source for diagnostics and import grouping; a synthetic name is
// used when it is empty.
func FormatGo(filename string, src []byte) ([]byte, error) {
	if filename == "" {
		filename = "generated.go"
	}
	return imports.Process(filename, src, &imports.Options{
		Comments:  true,
		TabIndent: true,
		TabWidth:  goTabWidth,
	})
}
