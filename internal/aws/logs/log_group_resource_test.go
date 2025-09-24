// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLogsLogGroup_update(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Logs::LogGroup", "awscc_logs_log_group", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSLogsLogGroupRetentionConfig(&td, rName, 30),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttrSet(resourceName, "arn"),
				resource.TestCheckNoResourceAttr(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "log_group_name", rName),
				resource.TestCheckResourceAttr(resourceName, "retention_in_days", "30"),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
		{
			Config: testAccAWSLogsLogGroupRetentionConfig(&td, rName, 60),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttrSet(resourceName, "arn"),
				resource.TestCheckNoResourceAttr(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "log_group_name", rName),
				resource.TestCheckResourceAttr(resourceName, "retention_in_days", "60"),
			),
		},
	})
}

// https://github.com/hashicorp/terraform-provider-awscc/issues/1020
func TestAccAWSLogsLogGroupWithDataProtectionPolicy_create(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Logs::LogGroup", "awscc_logs_log_group", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSLogsLogGroupDataProtectionPolicy(&td, rName),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttrSet(resourceName, "arn"),
				resource.TestCheckNoResourceAttr(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "log_group_name", rName),
			),
		},
	})
}

func testAccAWSLogsLogGroupRetentionConfig(td *acctest.TestData, rName string, retentionInDays int) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  log_group_name = %[3]q

  retention_in_days = %[4]d
}
`, td.TerraformResourceType, td.ResourceLabel, rName, retentionInDays)
}

func testAccAWSLogsLogGroupDataProtectionPolicy(td *acctest.TestData, rName string) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  log_group_name = %[3]q

  // Minimal policy
  data_protection_policy = jsonencode({
    "Name" : "data-protection-policy",
    "Version" : "2021-06-01",
    "Statement" : [{
      "Sid" : "audit-policy",
      "DataIdentifier" : [
        "arn:aws:dataprotection::aws:data-identifier/Address"
      ],
      "Operation" : {
        "Audit" : {
          "FindingsDestination" : {}
        }
      }
    },
    {
      "Sid" : "redact-policy",
      "DataIdentifier" : [
        "arn:aws:dataprotection::aws:data-identifier/Address"
      ],
      "Operation" : {
        "Deidentify" : {
          "MaskConfig" : {}
        }
      }
    }]
  })
}
`, td.TerraformResourceType, td.ResourceLabel, rName)
}
