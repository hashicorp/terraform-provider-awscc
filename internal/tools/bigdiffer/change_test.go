// SPDX-License-Identifier: MPL-2.0

package main

import (
	"os"
	"testing"
)

func TestClassifyChange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		discovered   []byte
		cached       []byte
		cachedExists bool
		frozen       bool
		want         byteStatus
	}{
		{"frozen wins even when bytes differ", []byte("new"), []byte("old"), true, true, statusFrozen},
		{"frozen with no cache", []byte("new"), nil, false, true, statusFrozen},
		{"discovery failure", nil, []byte("old"), true, false, statusMissing},
		{"no cache is new", []byte("new"), nil, false, false, statusNew},
		{"identical is unchanged", []byte("same"), []byte("same"), true, false, statusUnchanged},
		{"different is changed", []byte("new"), []byte("old"), true, false, statusChanged},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := classifyChange(tt.discovered, tt.cached, tt.cachedExists, tt.frozen); got != tt.want {
				t.Errorf("classifyChange = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDetectChanges(t *testing.T) {
	t.Parallel()

	cacheDir := t.TempDir()

	// Seed the cache: identical (unchanged) and stale (changed) copies.
	if err := os.WriteFile(schemaCachePath(cacheDir, "AWS::Svc::Same"), []byte("SCHEMA-A"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(schemaCachePath(cacheDir, "AWS::Svc::Diff"), []byte("OLD-BYTES"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(schemaCachePath(cacheDir, "AWS::Svc::Frozen"), []byte("OLD-FROZEN"), 0o644); err != nil {
		t.Fatal(err)
	}

	row := func(cfn string) resourceRow { return resourceRow{CloudFormationTypeName: cfn} }
	disc := []discovered{
		{row: row("AWS::Svc::Same"), schema: []byte("SCHEMA-A")},     // identical -> unchanged
		{row: row("AWS::Svc::Diff"), schema: []byte("NEW-BYTES")},    // differs -> changed
		{row: row("AWS::Svc::New")},                                  // no cache, has... set below
		{row: row("AWS::Svc::Frozen"), schema: []byte("NEW-FROZEN")}, // frozen -> skipped
		{row: row("AWS::Svc::Broke"), err: os.ErrDeadlineExceeded},   // discovery failed -> missing
	}
	// Give the "New" type bytes but no cache file.
	disc[2].schema = []byte("BRAND-NEW")

	frozen := map[string]bool{"AWS::Svc::Frozen": true}

	results, err := detectChanges(disc, frozen, cacheDir)
	if err != nil {
		t.Fatalf("detectChanges: %v", err)
	}

	want := map[string]byteStatus{
		"AWS::Svc::Same":   statusUnchanged,
		"AWS::Svc::Diff":   statusChanged,
		"AWS::Svc::New":    statusNew,
		"AWS::Svc::Frozen": statusFrozen,
		"AWS::Svc::Broke":  statusMissing,
	}
	got := make(map[string]byteStatus, len(results))
	for _, r := range results {
		got[r.cfType] = r.status
	}
	for cfn, w := range want {
		if got[cfn] != w {
			t.Errorf("%s: status = %q, want %q", cfn, got[cfn], w)
		}
	}

	s := summarize(results)
	if s.New != 1 || s.Changed != 1 || s.Unchanged != 1 || s.Frozen != 1 || s.Missing != 1 {
		t.Errorf("summary = %+v, want 1 of each", s)
	}

	cands := gateCandidates(results)
	if len(cands) != 2 {
		t.Errorf("gateCandidates = %v, want New+Changed (2)", cands)
	}
}
