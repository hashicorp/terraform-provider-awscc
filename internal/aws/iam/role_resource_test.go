package iam_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIAMRole_AssumeRolePolicyDocument(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IAM::Role", "awscc_iam_role", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSIAMRoleAssumeRolePolicyDocumentConfig(&td, rName),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttr(resourceName, "role_name", rName),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func testAccAWSIAMRoleAssumeRolePolicyDocumentConfig(td *acctest.TestData, rName string) string {
	return fmt.Sprintf(`
locals {
  assume_role_policy_document = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Principal": {
      "Service": "vpc-flow-logs.amazonaws.com"
    },
    "Action": "sts:AssumeRole"
  }]
}
EOF
}

resource %[1]q %[2]q {
  role_name                   = %[3]q
  assume_role_policy_document = jsonencode(jsondecode(local.assume_role_policy_document))
}
`, td.TerraformResourceType, td.ResourceLabel, rName)
}
