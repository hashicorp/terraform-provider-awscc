// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3outposts_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSS3OutpostsBucketPolicyDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::S3Outposts::BucketPolicy", "awscc_s3outposts_bucket_policy", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSS3OutpostsBucketPolicyDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::S3Outposts::BucketPolicy", "awscc_s3outposts_bucket_policy", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
