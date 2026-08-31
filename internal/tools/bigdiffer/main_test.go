// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"strings"
	"testing"
)

const testHead = `# Copyright IBM Corp. 2021, 2026
# SPDX-License-Identifier: MPL-2.0

defaults {
  schema_cache_directory     = "../service/cloudformation/schemas"
  terraform_type_name_prefix = "awscc"
}

meta_schema {
  path = "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"
}

`

// A complex block: manual suppression plus free-form comments that a gohcl
// round-trip would destroy. It must survive byte-for-byte.
const complexBlock = `resource_schema "aws_ec2_complex" {
  cloudformation_type_name = "AWS::EC2::Complex"

  # Suppression Reason: something unsupported
  # https://github.com/hashicorp/terraform-provider-awscc/issues/9999
  suppress_resource_generation           = true
  suppress_plural_data_source_generation = true
}`

func testOverlay() string {
	// Deliberately out of order (Beta before Alpha) and containing a retained
	// row (Old::Gone) plus a wrong-label row (Widget).
	blocks := []string{
		`resource_schema "aws_ec2_beta" {
  cloudformation_type_name = "AWS::EC2::Beta"
}`,
		`resource_schema "aws_ec2_alpha" {
  cloudformation_type_name = "AWS::EC2::Alpha"
}`,
		complexBlock,
		`resource_schema "aws_old_gone" {
  cloudformation_type_name = "AWS::Old::Gone"
}`,
		`resource_schema "aws_ec2_wrong" {
  cloudformation_type_name = "AWS::EC2::Widget"
}`,
	}
	return testHead +
		"# 3 CloudFormation resource types schemas are available for use with the Cloud Control API.\n\n" +
		strings.Join(blocks, "\n\n") + "\n"
}

func testBase() []resourceRow {
	return []resourceRow{
		{ResourceTypeName: "aws_ec2_alpha", CloudFormationTypeName: "AWS::EC2::Alpha"},
		{ResourceTypeName: "aws_ec2_beta", CloudFormationTypeName: "AWS::EC2::Beta"},
		{ResourceTypeName: "aws_ec2_complex", CloudFormationTypeName: "AWS::EC2::Complex", SuppressPluralDataSourceGeneration: true},
		{ResourceTypeName: "aws_ec2_widget", CloudFormationTypeName: "AWS::EC2::Widget"},
		{ResourceTypeName: "aws_ec2_newthing", CloudFormationTypeName: "AWS::EC2::Newthing", SuppressPluralDataSourceGeneration: true},
		{ResourceTypeName: "aws_ec2_backlog", CloudFormationTypeName: "AWS::EC2::Backlog"},
	}
}

func testPrevious() []resourceRow {
	// Backlog was available last week; Newthing was not.
	return []resourceRow{
		{ResourceTypeName: "aws_ec2_alpha", CloudFormationTypeName: "AWS::EC2::Alpha"},
		{ResourceTypeName: "aws_ec2_backlog", CloudFormationTypeName: "AWS::EC2::Backlog"},
	}
}

func TestNormalize(t *testing.T) {
	t.Parallel()

	overlay := testOverlay()
	checkout := map[string]bool{} // Old::Gone intentionally NOT pinned.

	out, report, err := normalize(overlay, testBase(), testPrevious(), checkout)
	if err != nil {
		t.Fatalf("normalize: %v", err)
	}

	// Header count reflects the base size.
	if !strings.Contains(out, "# 6 CloudFormation resource types schemas are available") {
		t.Errorf("count header not updated to 6:\n%s", out)
	}

	// Complex block preserved byte-for-byte (comments and all).
	if !strings.Contains(out, complexBlock) {
		t.Errorf("complex block not preserved verbatim:\n%s", out)
	}

	// New row rendered canonically with aligned plural flag and its structural
	// suppression reason (canonicalBlock; contributing/docs/suppressed-and-frozen.md).
	wantNew := `resource_schema "aws_ec2_newthing" {
  cloudformation_type_name               = "AWS::EC2::Newthing"
  suppress_plural_data_source_generation = true
  suppression_reason                     = "structural: no list handler with zero required arguments"
}`
	if !strings.Contains(out, wantNew) {
		t.Errorf("new block not rendered as expected:\n%s", out)
	}

	// Backlog row added without a plural flag.
	wantBacklog := `resource_schema "aws_ec2_backlog" {
  cloudformation_type_name = "AWS::EC2::Backlog"
}`
	if !strings.Contains(out, wantBacklog) {
		t.Errorf("backlog block not rendered as expected:\n%s", out)
	}

	// Sorted by CloudFormation type name.
	order := []string{
		`"AWS::EC2::Alpha"`, `"AWS::EC2::Backlog"`, `"AWS::EC2::Beta"`,
		`"AWS::EC2::Complex"`, `"AWS::EC2::Newthing"`, `"AWS::EC2::Widget"`, `"AWS::Old::Gone"`,
	}
	prev := -1
	for _, marker := range order {
		idx := strings.Index(out, marker)
		if idx < 0 {
			t.Fatalf("missing %s in output", marker)
		}
		if idx < prev {
			t.Errorf("out of order: %s at %d precedes previous %d\n%s", marker, idx, prev, out)
		}
		prev = idx
	}

	// Report contents.
	if len(report.AddedNew) != 1 || report.AddedNew[0].cfn != "AWS::EC2::Newthing" {
		t.Errorf("AddedNew = %+v, want [AWS::EC2::Newthing]", report.AddedNew)
	}
	if len(report.AddedBacklog) != 1 || report.AddedBacklog[0].cfn != "AWS::EC2::Backlog" {
		t.Errorf("AddedBacklog = %+v, want [AWS::EC2::Backlog]", report.AddedBacklog)
	}
	if len(report.Retained) != 1 || report.Retained[0].cfn != "AWS::Old::Gone" {
		t.Errorf("Retained = %+v, want [AWS::Old::Gone]", report.Retained)
	}
	if len(report.UnexplainedRetained) != 1 || report.UnexplainedRetained[0].cfn != "AWS::Old::Gone" {
		t.Errorf("UnexplainedRetained = %+v, want [AWS::Old::Gone]", report.UnexplainedRetained)
	}
	if len(report.NamingViolate) != 1 || !strings.Contains(report.NamingViolate[0], "AWS::EC2::Widget") {
		t.Errorf("NamingViolate = %v, want one entry for AWS::EC2::Widget", report.NamingViolate)
	}
}

func TestNormalizeIdempotent(t *testing.T) {
	t.Parallel()

	overlay := testOverlay()
	checkout := map[string]bool{}

	first, _, err := normalize(overlay, testBase(), testPrevious(), checkout)
	if err != nil {
		t.Fatalf("first normalize: %v", err)
	}
	second, _, err := normalize(first, testBase(), testPrevious(), checkout)
	if err != nil {
		t.Fatalf("second normalize: %v", err)
	}
	if first != second {
		t.Errorf("normalize is not idempotent:\n--- first ---\n%s\n--- second ---\n%s", first, second)
	}
}

func TestExpectedLabel(t *testing.T) {
	t.Parallel()

	got, err := expectedLabel("AWS::EC2::FpgaImage")
	if err != nil {
		t.Fatalf("expectedLabel: %v", err)
	}
	if want := "aws_ec2_fpga_image"; got != want {
		t.Errorf("expectedLabel = %q, want %q", got, want)
	}
}

func TestCheckoutPinSuppressesAnomaly(t *testing.T) {
	t.Parallel()

	overlay := testOverlay()
	checkout := map[string]bool{"AWS::Old::Gone": true}

	_, report, err := normalize(overlay, testBase(), testPrevious(), checkout)
	if err != nil {
		t.Fatalf("normalize: %v", err)
	}
	if len(report.Retained) != 1 {
		t.Errorf("Retained = %+v, want 1", report.Retained)
	}
	if len(report.UnexplainedRetained) != 0 {
		t.Errorf("UnexplainedRetained = %+v, want 0 (Old::Gone is pinned)", report.UnexplainedRetained)
	}
}

func TestPreservesCommentsAndCommentedBlocks(t *testing.T) {
	t.Parallel()

	leadCommentBlock := `# note one
# note two
resource_schema "aws_svc_kept" {
  cloudformation_type_name = "AWS::Svc::Kept"
}`
	commentedOutBlock := `# This resource was not present in the 01/01/2026 refresh
#resource_schema "aws_svc_gone" {
#  cloudformation_type_name = "AWS::Svc::Gone"
#
#  suppress_resource_generation = true
#}`
	alpha := `resource_schema "aws_svc_alpha" {
  cloudformation_type_name = "AWS::Svc::Alpha"
}`

	overlay := testHead +
		"# 1 CloudFormation resource types schemas are available for use with the Cloud Control API.\n\n" +
		strings.Join([]string{leadCommentBlock, commentedOutBlock, alpha}, "\n\n") + "\n"

	base := []resourceRow{
		{ResourceTypeName: "aws_svc_alpha", CloudFormationTypeName: "AWS::Svc::Alpha"},
	}
	checkout := map[string]bool{"AWS::Svc::Kept": true}

	out, report, err := normalize(overlay, base, nil, checkout)
	if err != nil {
		t.Fatalf("normalize: %v", err)
	}

	// Leading comments stay attached to their live block.
	if !strings.Contains(out, leadCommentBlock) {
		t.Errorf("leading-comment block not preserved:\n%s", out)
	}
	// The fully commented-out block survives byte-for-byte.
	if !strings.Contains(out, commentedOutBlock) {
		t.Errorf("commented-out block not preserved:\n%s", out)
	}

	// Sorted by CFN: Alpha < Gone (commented) < Kept.
	order := []string{`"AWS::Svc::Alpha"`, `"AWS::Svc::Gone"`, `"AWS::Svc::Kept"`}
	prev := -1
	for _, m := range order {
		idx := strings.Index(out, m)
		if idx < 0 {
			t.Fatalf("missing %s", m)
		}
		if idx < prev {
			t.Errorf("out of order: %s\n%s", m, out)
		}
		prev = idx
	}

	// Kept is retained (not in base) but pinned, so no anomaly.
	if len(report.Retained) != 1 || report.Retained[0].cfn != "AWS::Svc::Kept" {
		t.Errorf("Retained = %+v, want [AWS::Svc::Kept]", report.Retained)
	}
	if len(report.UnexplainedRetained) != 0 {
		t.Errorf("UnexplainedRetained = %+v, want none", report.UnexplainedRetained)
	}

	// Idempotent.
	again, _, err := normalize(out, base, nil, checkout)
	if err != nil {
		t.Fatalf("second normalize: %v", err)
	}
	if again != out {
		t.Errorf("not idempotent:\n--- first ---\n%s\n--- second ---\n%s", out, again)
	}
}

func TestDetectsDuplicateLiveBlocks(t *testing.T) {
	t.Parallel()

	dup := `resource_schema "aws_svc_thing" {
  cloudformation_type_name = "AWS::Svc::Thing"
}`
	overlay := testHead +
		"# 1 CloudFormation resource types schemas are available for use with the Cloud Control API.\n\n" +
		strings.Join([]string{dup, dup}, "\n\n") + "\n"

	base := []resourceRow{
		{ResourceTypeName: "aws_svc_thing", CloudFormationTypeName: "AWS::Svc::Thing"},
	}

	_, report, err := normalize(overlay, base, nil, map[string]bool{})
	if err != nil {
		t.Fatalf("normalize: %v", err)
	}
	if len(report.Duplicates) != 1 || !strings.Contains(report.Duplicates[0], "AWS::Svc::Thing") {
		t.Errorf("Duplicates = %v, want one entry for AWS::Svc::Thing", report.Duplicates)
	}
}

func TestAnomalyProblemsGating(t *testing.T) {
	t.Parallel()

	// Unpinned-retained is advisory only and must NOT be a blocking problem.
	r := Report{UnexplainedRetained: []rowRef{{cfn: "AWS::Svc::Gone"}}}
	if got := r.anomalyProblems(); len(got) != 0 {
		t.Errorf("unpinned-retained should not block -check, got %v", got)
	}

	// Duplicates and naming violations are blocking.
	r = Report{
		Duplicates:    []string{"AWS::Svc::Thing (2 live blocks)"},
		NamingViolate: []string{"AWS::Svc::Other: label ..."},
	}
	if got := r.anomalyProblems(); len(got) != 2 {
		t.Errorf("anomalyProblems() = %v, want 2 problems", got)
	}
}

func TestFrozenAndNonProvisionableSuppressAnomaly(t *testing.T) {
	t.Parallel()

	// All three are retained (absent from base). Frozen and Nonprov are
	// explained by their attributes; Orphan is not and must be flagged.
	frozen := `resource_schema "aws_svc_frozen" {
  cloudformation_type_name = "AWS::Svc::Frozen"
  frozen_since             = "2025-01-02"
}`
	nonprov := `resource_schema "aws_svc_nonprov" {
  cloudformation_type_name = "AWS::Svc::Nonprov"
  non_provisionable        = true
}`
	orphan := `resource_schema "aws_svc_orphan" {
  cloudformation_type_name = "AWS::Svc::Orphan"
}`

	overlay := testHead +
		"# 0 CloudFormation resource types schemas are available for use with the Cloud Control API.\n\n" +
		strings.Join([]string{frozen, nonprov, orphan}, "\n\n") + "\n"

	// Empty base: every row is retained.
	out, report, err := normalize(overlay, nil, nil, map[string]bool{})
	if err != nil {
		t.Fatalf("normalize: %v", err)
	}

	if len(report.Retained) != 3 {
		t.Errorf("Retained = %+v, want 3", report.Retained)
	}
	if len(report.UnexplainedRetained) != 1 || report.UnexplainedRetained[0].cfn != "AWS::Svc::Orphan" {
		t.Errorf("UnexplainedRetained = %+v, want [AWS::Svc::Orphan]", report.UnexplainedRetained)
	}

	// Idempotent.
	again, _, err := normalize(out, nil, nil, map[string]bool{})
	if err != nil {
		t.Fatalf("second normalize: %v", err)
	}
	if again != out {
		t.Errorf("not idempotent:\n--- first ---\n%s\n--- second ---\n%s", out, again)
	}
}
