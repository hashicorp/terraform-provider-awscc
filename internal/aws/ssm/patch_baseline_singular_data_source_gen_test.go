// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ssm_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSSSMPatchBaselineDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SSM::PatchBaseline", "awscc_ssm_patch_baseline", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSSSMPatchBaselineDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SSM::PatchBaseline", "awscc_ssm_patch_baseline", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
