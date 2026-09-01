// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"errors"
	"strings"
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
		if got := d.reasons[attrSuppressionReasonResource]; got != "generation_failed: recursive definition" {
			t.Errorf("reasons[%s] = %q, want generation_failed-tagged first-line failure detail", attrSuppressionReasonResource, got)
		}
		if _, ok := d.reasons[attrSuppressionReasonSingular]; ok {
			t.Errorf("singular gated OK, should not have a reason, got %+v", d.reasons)
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
		if d.reasons[attrFrozenReason] == "" {
			t.Errorf("freeze should record a frozen_reason, got %+v", d.reasons)
		}
		if _, ok := d.reasons[attrSuppressionReasonResource]; ok {
			t.Errorf("a total failure suppresses nothing, so it should not set a per-artifact reason either, got %+v", d.reasons)
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
		if d.reasons[attrFrozenReason] == "" {
			t.Errorf("partial failure should record a frozen_reason, got %+v", d.reasons)
		}
		if d.reasons[attrSuppressionReasonResource] == "" {
			t.Errorf("the failed resource should record its own reason too, got %+v", d.reasons)
		}
		if _, ok := d.reasons[attrSuppressionReasonSingular]; ok {
			t.Errorf("singular gated OK, should not have a reason, got %+v", d.reasons)
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

// TestReasonsForFailures covers the compile gate's taxonomy distinction (item
// 1) now rendered per-artifact (item 9b): a build-gate rejection
// (gateFailedBuild) must render build_failed, not generation_failed, since
// they are different stages with different remediation (a build failure means
// the compile gate's own go build rejected code the engine rendered fine; a
// generation failure means the engine itself errored).
func TestReasonsForFailures(t *testing.T) {
	t.Parallel()

	t.Run("generation failure tags generation_failed", func(t *testing.T) {
		t.Parallel()
		gr := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateFailedGeneration, err: errors.New("boom")},
		}}
		got := reasonsForFailures(gr)
		if got[attrSuppressionReasonResource] != "generation_failed: boom" {
			t.Errorf("got %+v, want generation_failed-tagged", got)
		}
	})

	t.Run("build failure tags build_failed", func(t *testing.T) {
		t.Parallel()
		gr := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateFailedBuild, err: errors.New("undefined: fwvalidators.Foo")},
		}}
		got := reasonsForFailures(gr)
		if got[attrSuppressionReasonResource] != "build_failed: undefined: fwvalidators.Foo" {
			t.Errorf("got %+v, want build_failed-tagged", got)
		}
	})

	t.Run("a mix of build and generation failures tags each artifact independently", func(t *testing.T) {
		t.Parallel()
		// Unlike the pre-9b single combined reason string, each artifact now
		// gets its own map entry keyed by its own reason attribute, so a
		// gateResult mixing both stages across its artifacts renders each
		// correctly rather than one string picking a single "more specific"
		// category for the whole type.
		gr := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateFailedGeneration, err: errors.New("gen boom")},
			{kind: artifactSingularDataSource, outcome: gateFailedBuild, err: errors.New("build boom")},
		}}
		got := reasonsForFailures(gr)
		if got[attrSuppressionReasonResource] != "generation_failed: gen boom" {
			t.Errorf("resource reason = %q, want generation_failed-tagged", got[attrSuppressionReasonResource])
		}
		if got[attrSuppressionReasonSingular] != "build_failed: build boom" {
			t.Errorf("singular reason = %q, want build_failed-tagged", got[attrSuppressionReasonSingular])
		}
	})

	t.Run("no attributable failure still yields a matching resource reason (the len==0 fallback)", func(t *testing.T) {
		t.Parallel()
		// suppressAttrsForFailures' own len(attrs)==0 fallback defensively
		// suppresses the resource when nothing in gr.artifacts is
		// attributable; reasonsForFailures must mirror that exact condition
		// so the fallback-suppressed resource still gets a reason — otherwise
		// -check's per-field anomaly would flag a suppression this same
		// decision just made as reason-less (review finding 3).
		gr := gateResult{cfType: "AWS::Svc::Thing"} // no artifacts at all
		attrs := suppressAttrsForFailures(gr)
		reasons := reasonsForFailures(gr)
		if attrs[attrSuppressResource] != "true" {
			t.Fatalf("expected the fallback to suppress the resource, got %+v", attrs)
		}
		if reasons[attrSuppressionReasonResource] == "" {
			t.Errorf("expected a matching fallback reason for the fallback-suppressed resource, got %+v", reasons)
		}
	})
}

// TestFrozenReasonForFailures covers review finding 4/5: the partial-failure
// branch's frozen_reason needs its own rationale text (not the per-artifact
// failure string), but its category prefix must still be derived from
// whether any failing artifact was rejected by the compile gate, not
// hardcoded — a partial failure can be triggered by a build failure just as
// easily as a generation failure.
func TestFrozenReasonForFailures(t *testing.T) {
	t.Parallel()

	t.Run("generation-triggered partial failure tags generation_failed", func(t *testing.T) {
		t.Parallel()
		gr := gateResult{artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateFailedGeneration, err: errors.New("boom")},
			{kind: artifactSingularDataSource, outcome: gateOK},
		}}
		got := frozenReasonForFailures(gr)
		if !strings.HasPrefix(got, "generation_failed:") {
			t.Errorf("got %q, want a generation_failed-tagged frozen_reason", got)
		}
	})

	t.Run("build-triggered partial failure tags build_failed, not generation_failed", func(t *testing.T) {
		t.Parallel()
		gr := gateResult{artifacts: []artifactResult{
			{kind: artifactPluralDataSource, outcome: gateFailedBuild, err: errors.New("undefined symbol")},
			{kind: artifactResource, outcome: gateOK},
		}}
		got := frozenReasonForFailures(gr)
		if !strings.HasPrefix(got, "build_failed:") {
			t.Errorf("got %q, want a build_failed-tagged frozen_reason — hardcoding generation_failed here would mislabel a build-triggered freeze", got)
		}
		if strings.HasPrefix(got, "generation_failed:") {
			t.Errorf("got %q, must not fall back to generation_failed for a build-triggered freeze", got)
		}
	})

	t.Run("rationale text describes the shared-JSON invariant, not any one artifact's error", func(t *testing.T) {
		t.Parallel()
		gr := gateResult{artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateFailedGeneration, err: errors.New("very specific artifact error text")},
		}}
		got := frozenReasonForFailures(gr)
		if strings.Contains(got, "very specific artifact error text") {
			t.Errorf("got %q, frozen_reason should not reuse the per-artifact failure detail", got)
		}
		if !strings.Contains(got, "cannot partially advance") {
			t.Errorf("got %q, want the shared-JSON rationale", got)
		}
	})
}

// TestTotalFailureReason covers the classPresent branch where every artifact
// failed: nothing is suppressed, so the combined, artifact-labeled reason
// (mirroring the pre-9b gateFailureReason) goes directly into frozen_reason.
func TestTotalFailureReason(t *testing.T) {
	t.Parallel()

	gr := gateResult{artifacts: []artifactResult{
		{kind: artifactResource, outcome: gateFailedGeneration, err: errors.New("gen boom")},
		{kind: artifactSingularDataSource, outcome: gateFailedBuild, err: errors.New("build boom")},
	}}
	got := totalFailureReason(gr)
	if got != "build_failed: resource: gen boom; singular_data_source: build boom" {
		t.Errorf("got %q, want build_failed-tagged with both details", got)
	}
}
