// Code generated by generators/resource/main.go; DO NOT EDIT.

package glue_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSGlueSchema_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Glue::Schema", "awscc_glue_schema", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
