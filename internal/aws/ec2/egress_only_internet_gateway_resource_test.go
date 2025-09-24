// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEC2EgressOnlyInternetGateway_success(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::EgressOnlyInternetGateway", "awscc_ec2_egress_only_internet_gateway", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSEC2EgressOnlyInternetGatewayConfig(&td, rName),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttrSet(resourceName, "egress_only_internet_gateway_id"),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func testAccAWSEC2EgressOnlyInternetGatewayConfig(td *acctest.TestData, rName string) string {
	return fmt.Sprintf(`
resource "awscc_ec2_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = [
    {
      key   = "Name"
      value = %[3]q
    }
  ]
}

resource %[1]q %[2]q {
  vpc_id = awscc_ec2_vpc.test.id
}
`, td.TerraformResourceType, td.ResourceLabel, rName)
}
