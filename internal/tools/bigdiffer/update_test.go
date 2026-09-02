// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

// TestNormalizeWithDecisions applies a freeze to an existing block and a
// suppression (with reason) to a newly added block, and checks the attributes
// land on exactly those blocks while an unrelated hand-annotated block survives
// byte-for-byte.
func TestNormalizeWithDecisions(t *testing.T) {
	t.Parallel()

	decisions := map[string]policyDecision{
		"AWS::EC2::Beta": {setAttrs: map[string]string{attrFrozenSince: "2026-08-28"}},
		"AWS::EC2::Newthing": {
			addBlock: true,
			setAttrs: map[string]string{attrSuppressResource: "true"},
			reasons:  map[string]string{attrSuppressionReasonResource: "generation failed: boom"},
		},
	}

	out, _, err := normalizeWithDecisions(testOverlay(), testBase(), testPrevious(), map[string]bool{}, decisions)
	if err != nil {
		t.Fatalf("normalizeWithDecisions: %v", err)
	}

	// Hand-annotated block with no decision is preserved verbatim.
	if !strings.Contains(out, complexBlock) {
		t.Errorf("complex block not preserved:\n%s", out)
	}

	// The frozen block gained frozen_since.
	beta := blockFor(out, "aws_ec2_beta")
	if !strings.Contains(beta, "frozen_since") || !strings.Contains(beta, `"2026-08-28"`) {
		t.Errorf("beta block missing frozen_since:\n%s", beta)
	}

	// The newly added block gained suppress_resource_generation + reason.
	newthing := blockFor(out, "aws_ec2_newthing")
	if !strings.Contains(newthing, "suppress_resource_generation") || !strings.Contains(newthing, "= true") {
		t.Errorf("newthing missing suppress flag:\n%s", newthing)
	}
	if !strings.Contains(newthing, `suppression_reason_resource`) || !strings.Contains(newthing, `"generation failed: boom"`) {
		t.Errorf("newthing missing suppression_reason_resource:\n%s", newthing)
	}
}

// blockFor returns the resource_schema block text for a label from overlay text.
func blockFor(overlay, label string) string {
	start := strings.Index(overlay, `resource_schema "`+label+`"`)
	if start < 0 {
		return ""
	}
	end := strings.Index(overlay[start:], "\n}")
	if end < 0 {
		return overlay[start:]
	}
	return overlay[start : start+end+2]
}

// TestNormalizeWithDecisionsSkipsCommentedOutCollision is a regression test: a
// fully commented-out block can carry the same cloudformation_type_name as a
// live block elsewhere in the overlay (classifyItem's comment fallback still
// assigns it a key, with live=false). A decision keyed by that CloudFormation
// type name must only mutate the live block; applying setBlockAttributes to the
// commented-out text has zero resource_schema blocks and previously aborted the
// whole run.
func TestNormalizeWithDecisionsSkipsCommentedOutCollision(t *testing.T) {
	t.Parallel()

	overlay := testHead +
		"# 1 CloudFormation resource types schemas are available for use with the Cloud Control API.\n\n" +
		`resource_schema "aws_ec2_live" {
  cloudformation_type_name = "AWS::EC2::Live"
}

# resource_schema "aws_ec2_old" {
#   cloudformation_type_name = "AWS::EC2::Live"
# }` + "\n"

	base := []resourceRow{
		{ResourceTypeName: "aws_ec2_live", CloudFormationTypeName: "AWS::EC2::Live"},
	}
	decisions := map[string]policyDecision{
		"AWS::EC2::Live": {setAttrs: map[string]string{attrFrozenSince: "2026-08-28"}},
	}

	out, _, err := normalizeWithDecisions(overlay, base, nil, map[string]bool{}, decisions)
	if err != nil {
		t.Fatalf("normalizeWithDecisions: %v", err)
	}
	live := blockFor(out, "aws_ec2_live")
	if !strings.Contains(live, "frozen_since") {
		t.Errorf("live block should be frozen, got:\n%s", live)
	}
	if !strings.Contains(out, "# resource_schema \"aws_ec2_old\"") {
		t.Errorf("commented-out block should survive untouched:\n%s", out)
	}
}

func TestBuildCandidates(t *testing.T) {
	t.Parallel()

	results := []changeResult{
		{cfType: "AWS::EC2::New", status: statusNew},
		{cfType: "AWS::EC2::Chg", status: statusChanged},
		{cfType: "AWS::EC2::Same", status: statusUnchanged},
		{cfType: "AWS::EC2::Froze", status: statusFrozen},
	}
	discByCFN := map[string]discovered{
		"AWS::EC2::New": {row: resourceRow{ResourceTypeName: "aws_ec2_new", CloudFormationTypeName: "AWS::EC2::New"}, schema: []byte("new")},
		"AWS::EC2::Chg": {schema: []byte("chg")},
	}
	overlayByCFN := map[string]resourceRow{
		"AWS::EC2::Chg": {ResourceTypeName: "aws_ec2_chg", CloudFormationTypeName: "AWS::EC2::Chg", SuppressResourceGeneration: true},
	}

	cands := buildCandidates(results, discByCFN, overlayByCFN)
	if len(cands) != 2 {
		t.Fatalf("want 2 candidates, got %d", len(cands))
	}

	byCFN := map[string]candidate{}
	for _, c := range cands {
		byCFN[c.cfType] = c
	}

	if c := byCFN["AWS::EC2::New"]; c.class != classNew || string(c.schema) != "new" || c.row.ResourceTypeName != "aws_ec2_new" {
		t.Errorf("New candidate wrong: %+v", c)
	}
	// Changed uses the overlay row (so its suppress flag is honored) and is Present.
	if c := byCFN["AWS::EC2::Chg"]; c.class != classPresent || !c.row.SuppressResourceGeneration || string(c.schema) != "chg" {
		t.Errorf("Changed candidate wrong: %+v", c)
	}
}

// TestAbsentTypes covers the filtering half of the absent-row probe
// (contributing/docs/absent-row-probe-design.md): which overlay rows need
// probing at all. A row in A is never absent; a row already frozen or
// checkout-pinned is already explained and needs no probe.
func TestAbsentTypes(t *testing.T) {
	t.Parallel()

	overlayByCFN := map[string]resourceRow{
		"AWS::EC2::Present":  {CloudFormationTypeName: "AWS::EC2::Present"},
		"AWS::EC2::Frozen":   {CloudFormationTypeName: "AWS::EC2::Frozen"},
		"AWS::EC2::Pinned":   {CloudFormationTypeName: "AWS::EC2::Pinned"},
		"AWS::EC2::NeedsPro": {CloudFormationTypeName: "AWS::EC2::NeedsPro"},
	}
	discByCFN := map[string]discovered{
		"AWS::EC2::Present": {row: resourceRow{CloudFormationTypeName: "AWS::EC2::Present"}},
	}
	frozenCFN := map[string]bool{"AWS::EC2::Frozen": true}
	checkout := map[string]bool{"AWS::EC2::Pinned": true}

	got := absentTypes(overlayByCFN, discByCFN, frozenCFN, checkout)
	want := []string{"AWS::EC2::NeedsPro"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("absentTypes() = %v, want %v", got, want)
	}
}

// TestAbsentRowDecisionClearsUnexplainedRetained is the reconciliation half of
// the absent-row probe design (contributing/docs/absent-row-probe-design.md):
// applying the decide() outcome for a probed absent row via
// normalizeWithDecisions must set the explaining attribute on the existing
// block, and a subsequent normalize pass over the rewritten overlay (as the
// next run's -check/-update would see it) must no longer flag that row as
// UnexplainedRetained.
//
// testOverlay/testBase's AWS::Old::Gone is the fixture: live in the overlay,
// absent from base, otherwise unexplained (main_test.go).
func TestAbsentRowDecisionClearsUnexplainedRetained(t *testing.T) {
	t.Parallel()

	today := "2026-09-02"

	tests := []struct {
		name       string
		class      changeClass
		wantAttr   string
		wantSuffix string // "= <value>" tail, tolerant of hclwrite's column alignment
	}{
		{name: "withdrawn", class: classWithdrawn, wantAttr: "frozen_since", wantSuffix: `= "` + today + `"`},
		{name: "non-provisionable", class: classNonProvisionable, wantAttr: "non_provisionable", wantSuffix: "= true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			decisions := map[string]policyDecision{
				"AWS::Old::Gone": decide(tt.class, gateResult{}, today),
			}
			out, report, err := normalizeWithDecisions(testOverlay(), testBase(), nil, map[string]bool{}, decisions)
			if err != nil {
				t.Fatalf("normalizeWithDecisions: %v", err)
			}

			gone := blockFor(out, "aws_old_gone")
			if !strings.Contains(gone, tt.wantAttr) || !strings.Contains(gone, tt.wantSuffix) {
				t.Errorf("aws_old_gone missing %s %s:\n%s", tt.wantAttr, tt.wantSuffix, gone)
			}

			// This run's own report is computed from the pre-edit overlay
			// text (crossValidate decodes the input, not the rewritten
			// output), so it still reports AWS::Old::Gone as
			// UnexplainedRetained here — the decision hasn't landed on disk
			// yet from its own point of view. What must actually hold is the
			// design doc's claim: the row is durably explained *starting next
			// run*, checked below against the rewritten overlay.
			found := false
			for _, ref := range report.UnexplainedRetained {
				if ref.cfn == "AWS::Old::Gone" {
					found = true
				}
			}
			if !found {
				t.Errorf("expected AWS::Old::Gone in this run's own UnexplainedRetained (report reflects the pre-edit overlay): %+v", report.UnexplainedRetained)
			}

			// A second pass over the rewritten overlay (simulating the next
			// run reading its own prior output) confirms the row is durably
			// explained, not just coincidentally absent from this run's
			// report ordering.
			_, report2, err := normalize(out, testBase(), nil, map[string]bool{})
			if err != nil {
				t.Fatalf("second normalize pass: %v", err)
			}
			for _, ref := range report2.UnexplainedRetained {
				if ref.cfn == "AWS::Old::Gone" {
					t.Errorf("AWS::Old::Gone still UnexplainedRetained on the next run: %+v", report2.UnexplainedRetained)
				}
			}
		})
	}
}

func TestGateResultFromGenResults(t *testing.T) {
	t.Parallel()

	ok := gateResultFromGenResults("AWS::X::Y", []genResult{
		{a: genArtifact{kind: artifactResource}},
		{a: genArtifact{kind: artifactSingularDataSource}},
	})
	if !ok.ok() {
		t.Errorf("expected ok gateResult")
	}

	broken := gateResultFromGenResults("AWS::X::Y", []genResult{
		{a: genArtifact{kind: artifactResource}},
		{a: genArtifact{kind: artifactSingularDataSource}, err: errors.New("boom")},
	})
	if broken.ok() {
		t.Errorf("expected broken gateResult")
	}
	// The failure maps to a singular-DS suppression via the policy.
	d := decide(classNew, broken, "2026-08-28")
	if d.setAttrs[attrSuppressSingular] != "true" {
		t.Errorf("expected singular suppression, got %+v", d.setAttrs)
	}
}
