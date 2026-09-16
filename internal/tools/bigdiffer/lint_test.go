// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"os"
	"path/filepath"
	"testing"
)

// TestCheckRegistrationUpToDate exercises the -lint drift guard: a fresh emit
// passes, a drifted file is reported, and an absent file is allowed (the
// transition state, where the legacy directive files still register everything).
func TestCheckRegistrationUpToDate(t *testing.T) {
	cfg, rows := loadCorpus(t)

	reg, err := emitRegistration(cfg, rows)
	if err != nil {
		t.Fatalf("emitRegistration: %v", err)
	}
	cfg.registrationPath = filepath.Join(t.TempDir(), "registrations_gen.go")

	// Absent: allowed during the transition.
	if got := checkRegistrationUpToDate(cfg, rows); got != "" {
		t.Errorf("absent registration: got problem %q, want none", got)
	}

	// Fresh: up to date, no problem.
	if err := os.WriteFile(cfg.registrationPath, reg, 0o644); err != nil {
		t.Fatal(err)
	}
	if got := checkRegistrationUpToDate(cfg, rows); got != "" {
		t.Errorf("fresh registration: got problem %q, want none", got)
	}

	// Stale: a drifted committed file must be reported.
	if err := os.WriteFile(cfg.registrationPath, append(reg, "\n// drift\n"...), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := checkRegistrationUpToDate(cfg, rows); got == "" {
		t.Error("stale registration: got no problem, want a staleness report")
	}
}
