// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identity

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestValueAsString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    attr.Value
		expected string
	}{
		{
			name:     "null string",
			input:    types.StringNull(),
			expected: "",
		},
		{
			name:     "valid string",
			input:    types.StringValue("hello"),
			expected: "hello",
		},
		{
			name:     "valid int64",
			input:    types.Int64Value(1),
			expected: "1",
		},
		{
			name:     "valid float64",
			input:    types.Float64Value(3.14),
			expected: "3.14",
		},
		{
			name:     "valid bool",
			input:    types.BoolValue(true),
			expected: "true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ValueAsString(nil, tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}
