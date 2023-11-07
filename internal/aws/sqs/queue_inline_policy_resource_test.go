// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

// https://github.com/hashicorp/terraform-provider-awscc/issues/1176
func TestAccAWSSQSQueueInlinePolicy_create(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SQS::QueueInlinePolicy", "awscc_sqs_queue_inline_policy", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSSQSQueueInlinePolicyConfig(&td, rName),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttrSet(resourceName, "queue"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_document"),
			),
		},
	})
}

func testAccAWSSQSQueueInlinePolicyConfig(td *acctest.TestData, rName string) string {
	return fmt.Sprintf(`
resource "awscc_sqs_queue" "this" {}

resource %[1]q %[2]q {
  queue = awscc_sqs_queue.this.id

  policy_document = jsonencode({
    Version : "2012-10-17",
    Statement : [{
      Effect : "Deny",
      Principal : {
        AWS : "*"
      },
      Action : "sqs:SendMessage",
      Resource : "*"
    }]
  })
}
`, td.TerraformResourceType, td.ResourceLabel, rName)
}
