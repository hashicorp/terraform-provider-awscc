// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ivs_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIVSRecordingConfigurationDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IVS::RecordingConfiguration", "awscc_ivs_recording_configuration", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSIVSRecordingConfigurationDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IVS::RecordingConfiguration", "awscc_ivs_recording_configuration", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
