// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validators_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func TestNotNullFloat64(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val                 types.Float64
		expectedDiagnostics diag.Diagnostics
	}
	tests := map[string]testCase{
		"null": {
			val: types.Float64Null(),
			expectedDiagnostics: diag.Diagnostics{
				diag.NewAttributeErrorDiagnostic(
					path.Root("test"),
					"Missing Attribute Value",
					`test: value must be configured`,
				),
			},
		},
		"unknown": {
			val: types.Float64Unknown(),
		},
		"valid": {
			val: types.Float64Value(1.1),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			request := validator.Float64Request{
				Path:           path.Root("test"),
				PathExpression: path.MatchRoot("test"),
				ConfigValue:    test.val,
			}
			response := validator.Float64Response{}
			fwvalidators.NotNullFloat64().ValidateFloat64(ctx, request, &response)

			if diff := cmp.Diff(response.Diagnostics, test.expectedDiagnostics); diff != "" {
				t.Errorf("unexpected diagnostics difference: %s", diff)
			}
		})
	}
}
