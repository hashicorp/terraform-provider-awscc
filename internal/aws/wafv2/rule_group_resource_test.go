// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

// The schema is de-recursed to 4 levels of Statement. Statements nested more deeply
// must fail validation rather than be silently discarded by Terraform's type
// conversion of the object-typed attributes.
func TestAccAWSWAFv2RuleGroup_nestingDepthExceeded(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::WAFv2::RuleGroup", "awscc_wafv2_rule_group", "test")
	rName := td.RandomName()

	td.ResourceTest(t, []resource.TestStep{
		{
			// Level-4 Statement offers no NotStatement.
			Config:      testAccAWSWAFv2RuleGroupStatementConfig(&td, rName, `{ and_statement = { statements = [{ or_statement = { statements = [{ not_statement = { statement = { not_statement = { statement = `+testAccAWSWAFv2ByteMatchStatement("/a")+` } } } }, `+testAccAWSWAFv2ByteMatchStatement("/b")+`] } }, `+testAccAWSWAFv2ByteMatchStatement("/c")+`] } }`),
			ExpectError: regexp.MustCompile(`Maximum Nesting Depth Exceeded`),
		},
		{
			// A rate-based statement's scope-down statement counts towards the depth.
			Config:      testAccAWSWAFv2RuleGroupStatementConfig(&td, rName, `{ rate_based_statement = { limit = 1000, aggregate_key_type = "IP", scope_down_statement = { and_statement = { statements = [{ not_statement = { statement = { or_statement = { statements = [`+testAccAWSWAFv2ByteMatchStatement("/a")+`, `+testAccAWSWAFv2ByteMatchStatement("/b")+`] } } } }, `+testAccAWSWAFv2ByteMatchStatement("/c")+`] } } } }`),
			ExpectError: regexp.MustCompile(`Maximum Nesting Depth Exceeded`),
		},
	})
}

func testAccAWSWAFv2RuleGroupStatementConfig(td *acctest.TestData, rName, statement string) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  name     = %[3]q
  scope    = "REGIONAL"
  capacity = 100

  visibility_config = {
    cloudwatch_metrics_enabled = false
    metric_name                = %[3]q
    sampled_requests_enabled   = false
  }

  rules = [{
    name     = "test"
    priority = 1
    action   = { block = {} }

    visibility_config = {
      cloudwatch_metrics_enabled = false
      metric_name                = "test"
      sampled_requests_enabled   = false
    }

    statement = %[4]s
  }]
}
`, td.TerraformResourceType, td.ResourceLabel, rName, statement)
}
