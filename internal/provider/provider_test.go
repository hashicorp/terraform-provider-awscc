package provider

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	awsbase "github.com/hashicorp/aws-sdk-go-base/v2"
)

func TestProvider(t *testing.T) {}

func TestAppendProducts(t *testing.T) {
	t.Parallel()

	defaultProducts := []awsbase.APNProduct{
		{Name: "awscc-test", Version: "0.0.0", Comment: "unit test"},
		{Name: "terraform", Version: "0.0.0", Comment: "unit test"},
	}
	simpleProduct := awsbase.APNProduct{Name: "simple", Version: "t", Comment: "t"}
	simpleAddProduct := apnProduct{Name: &simpleProduct.Name, Version: &simpleProduct.Version, Comment: &simpleProduct.Comment}
	minimalProduct := awsbase.APNProduct{Name: "minimal", Version: defaultUserAgentVersion, Comment: defaultUserAgentComment}
	minimalAddProduct := apnProduct{Name: &minimalProduct.Name}

	testcases := map[string]struct {
		products    []awsbase.APNProduct
		addProducts []*apnProduct
		expected    []awsbase.APNProduct
	}{
		"none_added": {
			products:    defaultProducts,
			addProducts: []*apnProduct{},
			expected:    defaultProducts,
		},
		"simple_added": {
			products:    defaultProducts,
			addProducts: []*apnProduct{&simpleAddProduct},
			expected:    append(defaultProducts, simpleProduct),
		},
		"minimal_added": {
			products:    defaultProducts,
			addProducts: []*apnProduct{&minimalAddProduct},
			expected:    append(defaultProducts, minimalProduct),
		},
		"both_added": {
			products:    defaultProducts,
			addProducts: []*apnProduct{&simpleAddProduct, &minimalAddProduct},
			expected:    append(defaultProducts, []awsbase.APNProduct{simpleProduct, minimalProduct}...),
		},
	}

	for name, testcase := range testcases {
		name, testcase := name, testcase

		t.Run(name, func(t *testing.T) {
			actual := appendProducts(testcase.products, testcase.addProducts)
			if !cmp.Equal(testcase.expected, actual) {
				t.Errorf("expected %q, got %q", testcase.expected, actual)
			}
		})
	}
}
