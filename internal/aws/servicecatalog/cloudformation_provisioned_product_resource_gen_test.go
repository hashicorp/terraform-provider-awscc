// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package servicecatalog_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSServiceCatalogCloudFormationProvisionedProduct_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ServiceCatalog::CloudFormationProvisionedProduct", "awscc_servicecatalog_cloudformation_provisioned_product", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func TestAccAWSServiceCatalogCloudFormationProvisionedProduct_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ServiceCatalog::CloudFormationProvisionedProduct", "awscc_servicecatalog_cloudformation_provisioned_product", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
