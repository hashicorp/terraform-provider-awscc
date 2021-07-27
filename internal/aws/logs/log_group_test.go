package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

type logGroupResource struct{}

func TestAccLogGroup_basic(t *testing.T) {
	data := acctest.NewTestData(t, "AWS::Logs::LogGroup", "aws_logs_log_group", "test")
	r := logGroupResource{}

	data.ResourceTest(t, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				acctest.CheckThat(data.ResourceName).ExistsInAWS(),
			),
		},
	})
}

func (r logGroupResource) basic(data acctest.TestData) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  provider = cloudapi
}
`, data.TerraformResourceType, data.ResourceLabel)
}
