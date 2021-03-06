// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIoTCACertificateDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoT::CACertificate", "awscc_iot_ca_certificate", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSIoTCACertificateDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoT::CACertificate", "awscc_iot_ca_certificate", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
