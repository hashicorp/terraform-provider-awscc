package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLogsLogGroup_update(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Logs::LogGroup", "awscc_logs_log_group", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSLogsLogGroupRetentionConfig(&td, rName, 30),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttrSet(resourceName, "arn"),
				resource.TestCheckNoResourceAttr(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "log_group_name", rName),
				resource.TestCheckResourceAttr(resourceName, "retention_in_days", "30"),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
		{
			Config: testAccAWSLogsLogGroupRetentionConfig(&td, rName, 60),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttrSet(resourceName, "arn"),
				resource.TestCheckNoResourceAttr(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "log_group_name", rName),
				resource.TestCheckResourceAttr(resourceName, "retention_in_days", "60"),
			),
		},
	})
}

func testAccAWSLogsLogGroupRetentionConfig(td *acctest.TestData, rName string, retentionInDays int) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  log_group_name = %[3]q

  retention_in_days = %[4]d
}
`, td.TerraformResourceType, td.ResourceLabel, rName, retentionInDays)
}
