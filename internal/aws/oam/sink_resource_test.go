// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oam_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

// https://github.com/hashicorp/terraform-provider-awscc/issues/836
// https://github.com/hashicorp/terraform-provider-awscc/issues/1059
func TestAccAWSOamSink_create(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Oam::Sink", "awscc_oam_sink", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSOamSinkPolicy(&td, rName),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttr(resourceName, "name", rName),
				resource.TestCheckResourceAttrSet(resourceName, "policy"),
			),
		},
	})
}

func testAccAWSOamSinkPolicy(td *acctest.TestData, rName string) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  name = %[3]q

  policy = jsonencode({
    Version : "2012-10-17",
    Statement : [
      {
        Action : [
          "oam:CreateLink",
          "oam:UpdateLink"
        ],
        Effect : "Allow",
        Resource : "*",
        Principal : "*",
        Condition : {
          StringEquals : {
            "aws:PrincipalOrgID" : "this-never-matches"
          }
        }
      }
    ]
  })
}
`, td.TerraformResourceType, td.ResourceLabel, rName)
}
