// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	awsbase "github.com/hashicorp/aws-sdk-go-base/v2"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestProvider(t *testing.T) {}

func TestUserAgentProducts(t *testing.T) {
	t.Parallel()

	simpleProduct := awsbase.UserAgentProduct{Name: "simple", Version: "t", Comment: "t"}
	simpleAddProduct := userAgentProduct{ProductName: types.StringValue(simpleProduct.Name), ProductVersion: types.StringValue(simpleProduct.Version), Comment: types.StringValue(simpleProduct.Comment)}
	minimalProduct := awsbase.UserAgentProduct{Name: "minimal"}
	minimalAddProduct := userAgentProduct{ProductName: types.StringValue(minimalProduct.Name)}

	testcases := map[string]struct {
		addProducts []userAgentProduct
		expected    []awsbase.UserAgentProduct
	}{
		"none_added": {
			addProducts: []userAgentProduct{},
			expected:    []awsbase.UserAgentProduct{},
		},
		"simple_added": {
			addProducts: []userAgentProduct{simpleAddProduct},
			expected:    []awsbase.UserAgentProduct{simpleProduct},
		},
		"minimal_added": {
			addProducts: []userAgentProduct{minimalAddProduct},
			expected:    []awsbase.UserAgentProduct{minimalProduct},
		},
		"both_added": {
			addProducts: []userAgentProduct{simpleAddProduct, minimalAddProduct},
			expected:    []awsbase.UserAgentProduct{simpleProduct, minimalProduct},
		},
	}

	for name, testcase := range testcases {
		name, testcase := name, testcase

		t.Run(name, func(t *testing.T) {
			actual := userAgentProducts(testcase.addProducts)
			if !cmp.Equal(testcase.expected, actual) {
				t.Errorf("expected %q, got %q", testcase.expected, actual)
			}
		})
	}
}
