// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// runLint verifies all_schemas.hcl is normalized (sorted, canonical formatting,
// correct count header) and anomaly-free, using the overlay as its own base so no
// AWS query or snapshot file is needed. It writes nothing and is suitable for CI.
func runLint(allSchemasPath, checkoutPath string) error {
	cfg, rows, err := loadOverlay(allSchemasPath)
	if err != nil {
		return err
	}
	overlayContent, err := os.ReadFile(allSchemasPath)
	if err != nil {
		return fmt.Errorf("reading overlay %s: %w", allSchemasPath, err)
	}
	checkout, err := parseCheckout(checkoutPath)
	if err != nil {
		return fmt.Errorf("reading checkout %s: %w", checkoutPath, err)
	}

	// Reconcile the overlay against itself: no rows are added, so this only
	// re-sorts and re-formats, surfacing any hand-edit that left it un-normalized.
	out, report, err := normalize(string(overlayContent), rows, nil, checkout)
	if err != nil {
		return err
	}
	report.write()

	problems := report.anomalyProblems()
	// The count-header value counts schemas *available from AWS*, which an offline
	// run cannot know, so compare everything except that line: sorting, canonical
	// formatting, and byte-preservation of blocks.
	if countLineRE.ReplaceAllString(out, "#") != countLineRE.ReplaceAllString(string(overlayContent), "#") {
		problems = append([]string{"not normalized (sorting/formatting; re-run `-sync`, or fix by hand)"}, problems...)
	}
	if regProblem := checkRegistrationUpToDate(cfg, rows); regProblem != "" {
		problems = append(problems, regProblem)
	}
	if len(problems) > 0 {
		return fmt.Errorf("all_schemas.hcl check failed: %s", strings.Join(problems, "; "))
	}
	if n := len(report.UnexplainedRetained); n > 0 {
		_, _ = fmt.Fprintf(structOut, "bigdiffer: all_schemas.hcl is normalized; %d advisory anomaly line(s) reported above (not a check failure).\n", n)
		return nil
	}
	_, _ = fmt.Fprintln(structOut, "bigdiffer: all_schemas.hcl is normalized and anomaly-free.")
	return nil
}

// checkRegistrationUpToDate guards the committed registrations_gen.go against
// drift: it re-emits the registration file from the overlay and compares. A
// stale file (a resource added or removed without regenerating) would silently
// change the set of registered resources and data sources, so it must fail
// -lint. It returns a problem description, or "" when the file is up to date.
//
// Absence is not a failure: while the legacy directive files
// (resources.go/singular_data_sources.go/plural_data_sources.go) still register
// every type, this file is additive, and the documented legacy fallback deletes
// it. Once the legacy files are removed, presence should be required here.
func checkRegistrationUpToDate(cfg config, rows []resourceRow) string {
	committed, err := os.ReadFile(cfg.registrationPath)
	if errors.Is(err, os.ErrNotExist) {
		return ""
	}
	if err != nil {
		return fmt.Sprintf("reading %s: %v", filepath.Base(cfg.registrationPath), err)
	}
	want, err := emitRegistration(cfg, rows)
	if err != nil {
		return fmt.Sprintf("re-emitting registrations: %v", err)
	}
	if !bytes.Equal(committed, want) {
		return fmt.Sprintf("%s is stale (re-run `bigdiffer -reconcile` or `-sync`)", filepath.Base(cfg.registrationPath))
	}
	return ""
}
