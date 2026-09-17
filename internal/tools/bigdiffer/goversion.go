// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

// goModVersionRE matches go.mod's "go 1.26.8" directive line, capturing the
// version number. Go has required a full three-component version there
// (matching what actions/setup-go's go-version-file and GNUmakefile's
// GOTOOLCHAIN_PIN both read) since the toolchain-selection feature shipped;
// bigdiffer only ever needs to compare against that value, not parse the
// rest of go.mod.
var goModVersionRE = regexp.MustCompile(`(?m)^go\s+(\S+)\s*$`)

// verifyGoToolchain confirms the running go toolchain (runtime.Version(),
// e.g. "go1.26.8") exactly matches go.mod's "go" directive, and returns a
// descriptive error if not. Go's own toolchain selection (GOTOOLCHAIN=auto,
// the default) only ever auto-*upgrades* a too-old local Go to satisfy a
// go.mod requirement — it never downgrades a newer one — so running
// `go run ./internal/tools/bigdiffer ...` directly (bypassing GNUmakefile's
// own GOTOOLCHAIN pin, documented there and export-ed to every make recipe)
// on a contributor's newer local Go silently succeeds today, but generates
// Go source formatted by that newer toolchain's gofmt/go-printer rules,
// which have changed between Go releases (e.g. map-literal key/value
// alignment width) — producing output that byte-diffs from what CI's
// pinned toolchain would produce or verify, even though nothing about the
// schema or bigdiffer's own logic changed. This exact failure mode broke
// PR #3334's -check CI step: six files (network_interface, delivery_stream,
// cloud_autonomous_vm_cluster resource + singular data source) were last
// regenerated and committed under a newer local Go than CI's then-pinned
// 1.26.6, bypassing GNUmakefile's guard by being run directly rather than
// via `make bigdiffer-reconcile`.
//
// Checked once, at startup, for every command that writes or verifies
// generated Go source (-sync, -reconcile, -check): -lint, -recheck, and
// -docs never touch Go source formatting (docs are prose/Terraform, not
// gofmt'd Go — codegen/importdocs.go) and so are exempt.
func verifyGoToolchain() error {
	modPath, err := findGoMod()
	if err != nil {
		// go.mod is unreadable/unfindable — do not block the run on that; the
		// compile gate will catch a genuinely broken environment regardless,
		// and this check is a formatting-drift guard, not a hard requirement
		// for bigdiffer to run at all.
		return nil
	}
	modBytes, err := os.ReadFile(modPath)
	if err != nil {
		return nil
	}
	m := goModVersionRE.FindSubmatch(modBytes)
	if m == nil {
		return nil
	}
	want := "go" + string(m[1])
	got := runtime.Version()
	if got == want {
		return nil
	}
	return fmt.Errorf("running under %s, but %s pins %q — gofmt/goimports output "+
		"(map-literal alignment and other formatting rules) can differ between Go "+
		"versions, which would generate Go source that byte-diffs from what CI's "+
		"pinned toolchain produces or verifies, independent of any real schema or "+
		"logic change (this broke PR #3334's -check step: six files were committed "+
		"under a newer local Go than CI pins). Run via `make bigdiffer-sync` / `make "+
		"bigdiffer-reconcile` (GNUmakefile pins GOTOOLCHAIN for you), or "+
		"`GOTOOLCHAIN=%s go run ./internal/tools/bigdiffer ...` directly",
		got, modPath, want, want)
}

// findGoMod walks upward from the current working directory looking for
// go.mod, the same resolution order `go` itself uses to find a module root.
// bigdiffer's documented usage is always `go run ./internal/tools/bigdiffer
// ...` from the repo root (every runbook example), so this typically finds
// it on the first try; the upward walk is defensive, not load-bearing.
func findGoMod() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		candidate := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found")
		}
		dir = parent
	}
}
