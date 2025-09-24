// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEC2IPAMPool_update(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::IPAMPool", "awscc_ec2_ipam_pool", "test")
	resourceName := td.ResourceName

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSEC2IPAMPoolConfig(&td, "desc1"),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttr(resourceName, "description", "desc1"),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
		{
			Config: testAccAWSEC2IPAMPoolConfig(&td, "desc2"),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
				},
			},
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttr(resourceName, "description", "desc2"),
			),
		},
	})
}

func testAccAWSEC2IPAMPoolConfig(td *acctest.TestData, description string) string {
	return fmt.Sprintf(`
resource "awscc_ec2_ipam" "test" {
  operating_regions = [
    {
      region_name = %[4]q
    }
  ]
}

resource %[1]q %[2]q {
  address_family = "ipv4"
  description    = %[3]q
  ipam_scope_id  = awscc_ec2_ipam.test.private_default_scope_id
  locale         = %[4]q
}
`, td.TerraformResourceType, td.ResourceLabel, description, td.Region())
}
