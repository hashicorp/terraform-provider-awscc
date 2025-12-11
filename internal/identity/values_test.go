// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package identity

import (
	"context"
	"math/big"
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
			name:     "valid number",
			input:    types.NumberValue(big.NewFloat(3)),
			expected: "3",
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

			result := ValueAsString(context.TODO(), tt.input)
			if result != tt.expected {
				t.Fatalf("expected %q but got %q", tt.expected, result)
			}
		})
	}
}
