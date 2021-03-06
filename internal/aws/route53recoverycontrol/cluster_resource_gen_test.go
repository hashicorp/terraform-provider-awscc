// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53recoverycontrol_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRoute53RecoveryControlCluster_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53RecoveryControl::Cluster", "awscc_route53recoverycontrol_cluster", "test")

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

func TestAccAWSRoute53RecoveryControlCluster_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53RecoveryControl::Cluster", "awscc_route53recoverycontrol_cluster", "test")

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
