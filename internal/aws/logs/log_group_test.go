package logs_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

type logGroupResource struct{}

func TestAccLogGroup_basic(t *testing.T) {
	data := acctest.NewTestData(t, "aws_logs_log_group")
	r := logGroupResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check:  resource.ComposeTestCheckFunc(),
		},
	})
}

func (r logGroupResource) basic(data acctest.TestData) string {
	return `
resource "aws_logs_log_group" "test" {
  provider = cloudapi
}
`
}
