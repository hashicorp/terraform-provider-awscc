// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotcoredeviceadvisor_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIoTCoreDeviceAdvisorSuiteDefinition_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoTCoreDeviceAdvisor::SuiteDefinition", "awscc_iotcoredeviceadvisor_suite_definition", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
