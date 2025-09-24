// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEC2LaunchTemplate_update(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::LaunchTemplate", "awscc_ec2_launch_template", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()
	testExternalProviders := map[string]resource.ExternalProvider{
		"aws": {
			Source:            "hashicorp/aws",
			VersionConstraint: "~> 5.50",
		},
	}

	td.ResourceTestNoProviderFactories(t, []resource.TestStep{
		{
			ProtoV6ProviderFactories: td.ProviderFactories(),
			ExternalProviders:        testExternalProviders,
			Config:                   testAccAWSEC2LaunchTemplateInstanceTypeConfig(&td, rName, "t2.large"),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttr(resourceName, "launch_template_data.instance_type", "t2.large"),
			),
		},
		{
			ProtoV6ProviderFactories: td.ProviderFactories(),
			ExternalProviders:        testExternalProviders,
			ResourceName:             td.ResourceName,
			ImportState:              true,
			ImportStateVerify:        true,
			ImportStateVerifyIgnore:  []string{"launch_template_data"},
		},
		{
			ProtoV6ProviderFactories: td.ProviderFactories(),
			ExternalProviders:        testExternalProviders,
			Config:                   testAccAWSEC2LaunchTemplateInstanceTypeConfig(&td, rName, "t3.large"),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttr(resourceName, "launch_template_data.instance_type", "t3.large"),
			),
		},
	})
}

func testAccAWSEC2LaunchTemplateInstanceTypeConfig(td *acctest.TestData, rName string, instanceType string) string {
	return fmt.Sprintf(`
data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]
  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-gp2"]
  }

  filter {
    name   = "root-device-type"
    values = ["ebs"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  filter {
    name   = "architecture"
    values = ["x86_64"]
  }
}

resource %[1]q %[2]q {
  launch_template_data = {
    image_id      = data.aws_ami.amazon_linux.id
    instance_type = %[4]q
  }
  launch_template_name = %[3]q
}
`, td.TerraformResourceType, td.ResourceLabel, rName, instanceType)
}
