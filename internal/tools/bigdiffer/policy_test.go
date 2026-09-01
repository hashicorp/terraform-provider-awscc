// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"errors"
	"testing"
)

func TestDecide(t *testing.T) {
	t.Parallel()

	const today = "2026-08-27"

	okRes := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
		{kind: artifactResource, outcome: gateOK},
		{kind: artifactSingularDataSource, outcome: gateOK},
	}}
	brokenRes := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
		{kind: artifactResource, outcome: gateFailedGeneration, err: errors.New("recursive definition\nmore detail")},
		{kind: artifactSingularDataSource, outcome: gateOK},
	}}

	t.Run("new + ok adds a plain block", func(t *testing.T) {
		t.Parallel()
		d := decide(classNew, okRes, today)
		if !d.addBlock || len(d.setAttrs) != 0 {
			t.Errorf("got %+v, want addBlock with no attrs", d)
		}
	})

	t.Run("new + failed adds suppressed with reason", func(t *testing.T) {
		t.Parallel()
		d := decide(classNew, brokenRes, today)
		if !d.addBlock {
			t.Errorf("should still add the block (backlog)")
		}
		if d.setAttrs[attrSuppressResource] != "true" {
			t.Errorf("resource should be suppressed, got %+v", d.setAttrs)
		}
		if _, ok := d.setAttrs[attrSuppressSingular]; ok {
			t.Errorf("singular gated OK, should not be suppressed")
		}
		if d.reason == "" || d.reason != "generation_failed: resource: recursive definition" {
			t.Errorf("reason = %q, want generation_failed-tagged first-line failure detail", d.reason)
		}
	})

	t.Run("new + plural-only failure suppresses only the plural artifact", func(t *testing.T) {
		t.Parallel()
		brokenPlural := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateOK},
			{kind: artifactSingularDataSource, outcome: gateOK},
			{kind: artifactPluralDataSource, outcome: gateFailedGeneration, err: errors.New("plural boom")},
		}}
		d := decide(classNew, brokenPlural, today)
		if d.setAttrs[attrSuppressPlural] != "true" {
			t.Errorf("plural artifact failed, should be suppressed, got %+v", d.setAttrs)
		}
		if _, ok := d.setAttrs[attrSuppressResource]; ok {
			t.Errorf("resource gated OK, should not be suppressed, got %+v", d.setAttrs)
		}
		if _, ok := d.setAttrs[attrSuppressSingular]; ok {
			t.Errorf("singular gated OK, should not be suppressed, got %+v", d.setAttrs)
		}
	})

	t.Run("new + singular-only failure suppresses only the singular artifact", func(t *testing.T) {
		t.Parallel()
		brokenSingularNew := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateOK},
			{kind: artifactSingularDataSource, outcome: gateFailedGeneration, err: errors.New("singular boom")},
			{kind: artifactPluralDataSource, outcome: gateOK},
		}}
		d := decide(classNew, brokenSingularNew, today)
		if d.setAttrs[attrSuppressSingular] != "true" {
			t.Errorf("singular failed, should be suppressed, got %+v", d.setAttrs)
		}
		if _, ok := d.setAttrs[attrSuppressResource]; ok {
			t.Errorf("resource gated OK, should not be suppressed, got %+v", d.setAttrs)
		}
		if _, ok := d.setAttrs[attrSuppressPlural]; ok {
			t.Errorf("plural gated OK, should not be suppressed, got %+v", d.setAttrs)
		}
	})

	t.Run("present + ok keeps the block untouched", func(t *testing.T) {
		t.Parallel()
		d := decide(classPresent, okRes, today)
		if d.addBlock || len(d.setAttrs) != 0 {
			t.Errorf("got %+v, want no add and no attrs", d)
		}
	})

	t.Run("present + total failure freezes at last-good, no suppression", func(t *testing.T) {
		t.Parallel()
		totalFail := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateFailedGeneration, err: errors.New("boom")},
			{kind: artifactSingularDataSource, outcome: gateFailedGeneration, err: errors.New("boom")},
			{kind: artifactPluralDataSource, outcome: gateFailedGeneration, err: errors.New("boom")},
		}}
		d := decide(classPresent, totalFail, today)
		if d.addBlock {
			t.Errorf("present must not add a block")
		}
		if d.setAttrs[attrFrozenSince] != today {
			t.Errorf("should freeze at today, got %+v", d.setAttrs)
		}
		if _, ok := d.setAttrs[attrSuppressResource]; ok {
			t.Errorf("a total failure should not also suppress individual artifacts, got %+v", d.setAttrs)
		}
		if d.reason == "" {
			t.Errorf("freeze should record a reason")
		}
	})

	t.Run("present + resource-only failure freezes schema, suppresses only the resource", func(t *testing.T) {
		t.Parallel()
		d := decide(classPresent, brokenRes, today) // resource failed, singular ok
		if d.addBlock {
			t.Errorf("present must not add a block")
		}
		if d.setAttrs[attrFrozenSince] != today {
			t.Errorf("schema can't partially advance, should still freeze, got %+v", d.setAttrs)
		}
		if d.setAttrs[attrSuppressResource] != "true" {
			t.Errorf("resource failed, should be suppressed, got %+v", d.setAttrs)
		}
		if _, ok := d.setAttrs[attrSuppressSingular]; ok {
			t.Errorf("singular gated OK, should not be suppressed, got %+v", d.setAttrs)
		}
		if d.reason == "" {
			t.Errorf("partial failure should record a reason")
		}
	})

	t.Run("present + singular-only failure freezes schema, suppresses only singular", func(t *testing.T) {
		t.Parallel()
		brokenSingular := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateOK},
			{kind: artifactSingularDataSource, outcome: gateFailedGeneration, err: errors.New("singular boom")},
		}}
		d := decide(classPresent, brokenSingular, today)
		if d.setAttrs[attrFrozenSince] != today {
			t.Errorf("schema can't partially advance, should still freeze, got %+v", d.setAttrs)
		}
		if d.setAttrs[attrSuppressSingular] != "true" {
			t.Errorf("singular failed, should be suppressed, got %+v", d.setAttrs)
		}
		if _, ok := d.setAttrs[attrSuppressResource]; ok {
			t.Errorf("resource gated OK, should not be suppressed, got %+v", d.setAttrs)
		}
	})

	t.Run("present + plural-only failure freezes schema, suppresses only plural", func(t *testing.T) {
		t.Parallel()
		brokenPluralPresent := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateOK},
			{kind: artifactSingularDataSource, outcome: gateOK},
			{kind: artifactPluralDataSource, outcome: gateFailedGeneration, err: errors.New("plural boom")},
		}}
		d := decide(classPresent, brokenPluralPresent, today)
		if d.setAttrs[attrFrozenSince] != today {
			t.Errorf("schema can't partially advance, should still freeze, got %+v", d.setAttrs)
		}
		if d.setAttrs[attrSuppressPlural] != "true" {
			t.Errorf("plural failed, should be suppressed, got %+v", d.setAttrs)
		}
		if _, ok := d.setAttrs[attrSuppressResource]; ok {
			t.Errorf("resource gated OK, should not be suppressed, got %+v", d.setAttrs)
		}
		if _, ok := d.setAttrs[attrSuppressSingular]; ok {
			t.Errorf("singular gated OK, should not be suppressed, got %+v", d.setAttrs)
		}
	})

	t.Run("non-provisionable annotates", func(t *testing.T) {
		t.Parallel()
		d := decide(classNonProvisionable, gateResult{}, today)
		if d.setAttrs[attrNonProvisionable] != "true" || d.addBlock {
			t.Errorf("got %+v, want non_provisionable=true, no add", d)
		}
	})

	t.Run("withdrawn freezes", func(t *testing.T) {
		t.Parallel()
		d := decide(classWithdrawn, gateResult{}, today)
		if d.setAttrs[attrFrozenSince] != today || d.addBlock {
			t.Errorf("got %+v, want frozen_since=today, no add", d)
		}
	})
}

// TestGateFailureReason covers the compile gate's taxonomy distinction (item 1):
// a build-gate rejection (gateFailedBuild) must render build_failed, not
// generation_failed, since they are different stages with different remediation
// (a build failure means the compile gate's own go build rejected code the
// engine rendered fine; a generation failure means the engine itself errored).
func TestGateFailureReason(t *testing.T) {
	t.Parallel()

	t.Run("generation failure tags generation_failed", func(t *testing.T) {
		t.Parallel()
		gr := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateFailedGeneration, err: errors.New("boom")},
		}}
		got := gateFailureReason(gr)
		if got != "generation_failed: resource: boom" {
			t.Errorf("got %q, want generation_failed-tagged", got)
		}
	})

	t.Run("build failure tags build_failed", func(t *testing.T) {
		t.Parallel()
		gr := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateFailedBuild, err: errors.New("undefined: fwvalidators.Foo")},
		}}
		got := gateFailureReason(gr)
		if got != "build_failed: resource: undefined: fwvalidators.Foo" {
			t.Errorf("got %q, want build_failed-tagged", got)
		}
	})

	t.Run("a mix of build and generation failures tags build_failed", func(t *testing.T) {
		t.Parallel()
		// A gateResult mixing both stages across its artifacts is an edge case
		// (a type whose resource failed to generate and whose singular DS
		// generated but failed the compile gate) — the render is one reason
		// string for the whole type, so it favors the more specific
		// build_failed tag rather than silently defaulting to
		// generation_failed. Each artifact's own suppress_* attribute is still
		// set correctly regardless (suppressAttrsForFailures is outcome-agnostic
		// beyond != gateOK), so nothing is mis-suppressed by this choice.
		gr := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateFailedGeneration, err: errors.New("gen boom")},
			{kind: artifactSingularDataSource, outcome: gateFailedBuild, err: errors.New("build boom")},
		}}
		got := gateFailureReason(gr)
		if got != "build_failed: resource: gen boom; singular_data_source: build boom" {
			t.Errorf("got %q, want build_failed-tagged with both details", got)
		}
	})
}
