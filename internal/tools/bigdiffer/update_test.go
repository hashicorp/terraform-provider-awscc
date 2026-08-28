// SPDX-License-Identifier: MPL-2.0

package main

import (
	"errors"
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
			reason:   "generation failed: boom",
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
	if !strings.Contains(newthing, `suppression_reason`) || !strings.Contains(newthing, `"generation failed: boom"`) {
		t.Errorf("newthing missing suppression_reason:\n%s", newthing)
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
