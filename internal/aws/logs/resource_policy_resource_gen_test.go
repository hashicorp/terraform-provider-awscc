// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLogsResourcePolicy_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Logs::ResourcePolicy", "awscc_logs_resource_policy", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
