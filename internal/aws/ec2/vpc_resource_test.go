package ec2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEC2VPC_CidrBlock(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::VPC", "awscc_ec2_vpc", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()
	cidrBlock := "10.0.0.0/16"

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSEC2VPCCidrBlockConfig(&td, rName, cidrBlock),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", cidrBlock),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "Name"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", rName),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func testAccAWSEC2VPCCidrBlockConfig(td *acctest.TestData, rName, cidrBlock string) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  cidr_block = %[4]q

  tags = [
    {
      key   = "Name"
      value = %[3]q
    }
  ]
}
`, td.TerraformResourceType, td.ResourceLabel, rName, cidrBlock)
}
