// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iam_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIAMOIDCProviderDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IAM::OIDCProvider", "awscc_iam_oidc_provider", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSIAMOIDCProviderDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IAM::OIDCProvider", "awscc_iam_oidc_provider", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
