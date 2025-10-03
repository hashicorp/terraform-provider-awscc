package types

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	awsbase "github.com/hashicorp/aws-sdk-go-base/v2"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestUserAgentProducts(t *testing.T) {
	t.Parallel()

	simpleProduct := awsbase.UserAgentProduct{Name: "simple", Version: "0.0.1", Comment: "test comment"}
	simpleAddProduct := userAgentProduct{
		ProductName:    types.StringValue(simpleProduct.Name),
		ProductVersion: types.StringValue(simpleProduct.Version),
		Comment:        types.StringValue(simpleProduct.Comment),
	}
	minimalProduct := awsbase.UserAgentProduct{Name: "minimal"}
	minimalAddProduct := userAgentProduct{
		ProductName: types.StringValue(minimalProduct.Name),
	}

	testcases := map[string]struct {
		add      UserAgentProducts
		expected []awsbase.UserAgentProduct
	}{
		"none": {
			add:      []userAgentProduct{},
			expected: []awsbase.UserAgentProduct{},
		},
		"simple": {
			add:      []userAgentProduct{simpleAddProduct},
			expected: []awsbase.UserAgentProduct{simpleProduct},
		},
		"minimal": {
			add:      []userAgentProduct{minimalAddProduct},
			expected: []awsbase.UserAgentProduct{minimalProduct},
		},
		"both": {
			add:      []userAgentProduct{simpleAddProduct, minimalAddProduct},
			expected: []awsbase.UserAgentProduct{simpleProduct, minimalProduct},
		},
	}

	for name, testcase := range testcases {
		name, testcase := name, testcase

		t.Run(name, func(t *testing.T) {
			actual := testcase.add.UserAgentProducts()
			if !cmp.Equal(testcase.expected, actual) {
				t.Errorf("expected %q, got %q", testcase.expected, actual)
			}
		})
	}
}
