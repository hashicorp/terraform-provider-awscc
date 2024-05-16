// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acctest

import (
	"fmt"
	"os"
	"testing"

	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-provider-awscc/internal/envvar"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider"
)

type TestData struct {
	// CloudFormationResourceType is the CloudFormation resource type, e.g. "AWS::S3::Bucket".
	CloudFormationResourceType string

	// ResourceName is the resource label (local name), e.g. "test".
	ResourceLabel string

	// ResourceName is the qualified resource name, e.g. "aws_s3_bucket.test".
	ResourceName string

	// TerraformResourceType is the Terraform resource type, e.g. "aws_s3_bucket".
	TerraformResourceType string

	provider fwprovider.Provider
}

// EmptyConfig returns an empty (no attributes) Terraform configuration for the resource.
func (td *TestData) EmptyConfig() string {
	config := fmt.Sprintf(`
resource %[1]q %[2]q {}
`, td.TerraformResourceType, td.ResourceLabel)

	if role := os.Getenv("TF_AWS_ASSUME_ROLE_ARN"); role != "" {
		config = fmt.Sprintf(`
provider "awscc" {
  assume_role={
    role_arn = %[1]q
  }
}
`, role) + config
	}
	return config
}

// DataSourceWithEmptyResourceConfig returns a Terraform configuration for the data source and its respective resource.
func (td *TestData) DataSourceWithEmptyResourceConfig() string {
	return td.EmptyConfig() + fmt.Sprintf(`
data %[1]q %[2]q {
  id = %[1]s.%[2]s.id
}
`, td.TerraformResourceType, td.ResourceLabel)
}

// DataSourceWithNonExistentIDConfig returns a Terraform configuration for the data source
// with the id attribute set to a non-existent resource.
func (td *TestData) DataSourceWithNonExistentIDConfig() string {
	return fmt.Sprintf(`
data %[1]q %[2]q {
  id = "non-existent"
}
`, td.TerraformResourceType, td.ResourceLabel)
}

// EmptyDataSourceConfig returns an empty (no attributes) Terraform configuration for the data source.
func (td *TestData) EmptyDataSourceConfig() string {
	return fmt.Sprintf(`
data %[1]q %[2]q {}
`, td.TerraformResourceType, td.ResourceLabel)
}

// RandomName returns a new random name with the standard prefix `tf-acc-test`.
func (td *TestData) RandomName() string {
	return acctest.RandomWithPrefix("tf-acc-test")
}

// RandomAlphaString returns a new alphabetic random string of length `n`.
func (td *TestData) RandomAlphaString(n int) string {
	return acctest.RandStringFromCharSet(n, acctest.CharSetAlpha)
}

// Region returns the AWS Region in effect.
func (td *TestData) Region() string {
	return envvar.GetWithDefault(envvar.DefaultRegion, "us-west-2")
}

// NewTestData returns a new TestData structure.
func NewTestData(_ *testing.T, cfResourceType, tfResourceType, resourceLabel string) TestData {
	data := TestData{
		CloudFormationResourceType: cfResourceType,
		ResourceLabel:              resourceLabel,
		ResourceName:               fmt.Sprintf("%[1]s.%[2]s", tfResourceType, resourceLabel),
		TerraformResourceType:      tfResourceType,

		provider: provider.New(),
	}

	return data
}
