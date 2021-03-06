// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkmanager_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSNetworkManagerConnectPeer_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::NetworkManager::ConnectPeer", "awscc_networkmanager_connect_peer", "test")

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

func TestAccAWSNetworkManagerConnectPeer_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::NetworkManager::ConnectPeer", "awscc_networkmanager_connect_peer", "test")

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
