// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSNetworkManagerCoreNetwork_create(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::NetworkManager::CoreNetwork", "awscc_networkmanager_core_network", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSNetworkManagerCoreNetworkConfig(&td, rName),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttrSet(resourceName, "core_network_arn"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_document"),
			),
		},
	})
}

func testAccAWSNetworkManagerCoreNetworkConfig(td *acctest.TestData, rName string) string {
	return fmt.Sprintf(`
resource "awscc_networkmanager_global_network" "this" {}

resource %[1]q %[2]q {
  global_network_id = awscc_networkmanager_global_network.this.id

  policy_document = jsonencode({
    core-network-configuration : {
      asn-ranges : ["64512-65534"],
      edge-locations : [
        { location : "us-east-1" }
      ]
    },
    version : "2021.12",
    segments : [
      {
        name : "potato",
      }
    ]
  })
}
`, td.TerraformResourceType, td.ResourceLabel, rName)
}
