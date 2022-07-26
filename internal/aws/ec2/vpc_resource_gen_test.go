// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEC2VPC_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::VPC", "awscc_ec2_vpc", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func TestAccAWSEC2VPC_metadata(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::VPC", "awscc_ec2_vpc", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.MetadataConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}


func TestAccAWSEC2VPC_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::VPC", "awscc_ec2_vpc", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
