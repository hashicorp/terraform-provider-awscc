// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodb_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

// https://github.com/hashicorp/terraform-provider-awscc/issues/951
// TODO: Plan not empty after create.
func TestAccAWSDynamoDBTableKeySchema_create(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DynamoDB::Table", "awscc_dynamodb_table", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSDynamoDBTableKeySchema(&td, rName),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttrSet(resourceName, "arn"),
				resource.TestCheckResourceAttr(resourceName, "table_name", rName),
			),
		},
	})
}

func testAccAWSDynamoDBTableKeySchema(td *acctest.TestData, rName string) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  table_name = %[3]q
  billing_mode = "PAY_PER_REQUEST"

  attribute_definitions = [{
    attribute_name = "UserId"
    attribute_type = "S"
  }]

  key_schema = jsonencode([{
    AttributeName = "UserId"
    KeyType = "HASH"
  }])
}
`, td.TerraformResourceType, td.ResourceLabel, rName)
}
