// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package eventschemas_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEventSchemasRegistryPolicyDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EventSchemas::RegistryPolicy", "awscc_eventschemas_registry_policy", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSEventSchemasRegistryPolicyDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EventSchemas::RegistryPolicy", "awscc_eventschemas_registry_policy", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
