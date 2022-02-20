package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLogsQueryDefinition_queryString(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Logs::QueryDefinition", "awscc_logs_query_definition", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()
	expectedQueryString := `fields @timestamp, @message
| sort @timestamp desc
| limit 20
`

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSLogsQueryDefinitionConfig(&td, rName),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				resource.TestCheckResourceAttr(resourceName, "log_group_names.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "name", rName),
				resource.TestCheckResourceAttrSet(resourceName, "query_definition_id"),
				resource.TestCheckResourceAttr(resourceName, "query_string", expectedQueryString),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func testAccAWSLogsQueryDefinitionConfig(td *acctest.TestData, rName string) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  name = %[3]q

  query_string = <<EOF
fields @timestamp, @message
| sort @timestamp desc
| limit 20
EOF
}
`, td.TerraformResourceType, td.ResourceLabel, rName)
}
