package acctest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

type TestData struct {
	// ResourceType is the Terraform resource type, e.g. "aws_s3_bucket".
	ResourceType string
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
func NewTestData(t *testing.T, resourceType string) TestData {
	data := TestData{
		ResourceType: resourceType,
	}

	return data
}
