package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/provider"
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

	provider tfsdk.Provider
}

// EmptyConfig returns an empty (no attributes) Terraform configuration for the resource.
func (td *TestData) EmptyConfig() string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  provider = cloudapi
}
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

// NewTestData returns a new TestData structure.
func NewTestData(t *testing.T, cfResourceType, tfResourceType, resourceLabel string) TestData {
	data := TestData{
		CloudFormationResourceType: cfResourceType,
		ResourceLabel:              resourceLabel,
		ResourceName:               fmt.Sprintf("%[1]s.%[2]s", tfResourceType, resourceLabel),
		TerraformResourceType:      tfResourceType,

		provider: provider.New(),
	}

	return data
}
