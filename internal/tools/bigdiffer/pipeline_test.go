// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"errors"
	"reflect"
	"testing"
)

// TestMachineryFailures covers the load-bearing case a review caught: a
// machineryFailure decision can coexist with a settled batch that has no
// compileFixpoint error at all (a broken new artifact reverted to its
// still-compiling committed file, so the fixpoint itself goes green) — so
// every caller must scan decisions for machineryFailure explicitly, not just
// check compileFixpoint's own error return (contributing/docs/held-artifacts-design.md
// §3, "Why all-or-nothing, precisely").
func TestMachineryFailures(t *testing.T) {
	t.Parallel()

	t.Run("no machinery failures returns nil, even with other decisions present", func(t *testing.T) {
		t.Parallel()
		b := settledBatch{
			decisions: map[string]policyDecision{
				"AWS::Svc::Fine":   {summary: "present: refreshed OK"},
				"AWS::Svc::Frozen": {setAttrs: map[string]string{attrFrozenSince: "2026-01-01"}, summary: "present: generation broke, frozen at last-good bytes"},
			},
		}
		if got := machineryFailures(b); len(got) != 0 {
			t.Errorf("want no machinery failures, got %+v", got)
		}
	})

	t.Run("a flagged decision is reported even though it settled clean (the load-bearing case)", func(t *testing.T) {
		t.Parallel()
		gr := &gateResult{cfType: "AWS::Svc::Broken", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateFailedBuild, err: errors.New("undefined: Foo\nmore context")},
			{kind: artifactSingularDataSource, outcome: gateOK},
		}}
		b := settledBatch{
			decisions: map[string]policyDecision{
				// machineryFailure: true with no setAttrs/reasons at all — this
				// is exactly what decide()'s classPresentUnchanged branch
				// returns; nothing here signals a compileFixpoint-level error.
				"AWS::Svc::Broken": {machineryFailure: true, summary: "present (unchanged): generation or build broke — machinery regression, failing the run"},
			},
			gateResults: map[string]*gateResult{"AWS::Svc::Broken": gr},
		}
		got := machineryFailures(b)
		if len(got) != 1 {
			t.Fatalf("want 1 machinery failure, got %d: %+v", len(got), got)
		}
		if got[0].cfType != "AWS::Svc::Broken" {
			t.Errorf("wrong cfType: %+v", got[0])
		}
		want := []string{"resource: undefined: Foo"}
		if !reflect.DeepEqual(got[0].artifacts, want) {
			t.Errorf("artifacts = %v, want %v", got[0].artifacts, want)
		}
	})

	t.Run("multiple failures are all reported, sorted by type name, not just the first", func(t *testing.T) {
		t.Parallel()
		b := settledBatch{
			decisions: map[string]policyDecision{
				"AWS::Svc::Zebra": {machineryFailure: true},
				"AWS::Svc::Alpha": {machineryFailure: true},
				"AWS::Svc::Mid":   {machineryFailure: true},
				"AWS::Svc::Fine":  {},
			},
			gateResults: map[string]*gateResult{},
		}
		got := machineryFailures(b)
		var names []string
		for _, f := range got {
			names = append(names, f.cfType)
		}
		want := []string{"AWS::Svc::Alpha", "AWS::Svc::Mid", "AWS::Svc::Zebra"}
		if !reflect.DeepEqual(names, want) {
			t.Errorf("names = %v, want %v (sorted, all three, not just one)", names, want)
		}
	})

	t.Run("a baseDecisions entry with no gateResults entry is skipped safely, not a lookup panic", func(t *testing.T) {
		t.Parallel()
		// absentDecisions-style entries (runSync's absent-row probe) are never
		// machineryFailure and never have a gateResults entry — confirm the
		// lookup miss is handled, not just for a machineryFailure entry
		// (already covered above) but structurally: no entry at all here
		// should never be mistaken for one.
		b := settledBatch{
			decisions: map[string]policyDecision{
				"AWS::Svc::Absent": {setAttrs: map[string]string{attrNonProvisionable: "true"}, summary: "non-provisionable (live): annotated"},
			},
			gateResults: map[string]*gateResult{},
		}
		if got := machineryFailures(b); len(got) != 0 {
			t.Errorf("want no machinery failures, got %+v", got)
		}
	})

	t.Run("multiple broken artifacts on one type are all blamed, sorted", func(t *testing.T) {
		t.Parallel()
		gr := &gateResult{cfType: "AWS::Svc::DoubleBroken", artifacts: []artifactResult{
			{kind: artifactPluralDataSource, outcome: gateFailedGeneration, err: errors.New("plural boom")},
			{kind: artifactResource, outcome: gateFailedBuild, err: errors.New("resource boom")},
			{kind: artifactSingularDataSource, outcome: gateOK},
		}}
		b := settledBatch{
			decisions:   map[string]policyDecision{"AWS::Svc::DoubleBroken": {machineryFailure: true}},
			gateResults: map[string]*gateResult{"AWS::Svc::DoubleBroken": gr},
		}
		got := machineryFailures(b)
		if len(got) != 1 {
			t.Fatalf("want 1 machinery failure, got %d", len(got))
		}
		want := []string{"plural_data_source: plural boom", "resource: resource boom"}
		if !reflect.DeepEqual(got[0].artifacts, want) {
			t.Errorf("artifacts = %v, want %v (sorted by kind)", got[0].artifacts, want)
		}
	})
}
