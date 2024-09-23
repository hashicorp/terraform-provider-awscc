// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validators_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func TestNotNullList(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val                 types.List
		expectedDiagnostics diag.Diagnostics
	}
	tests := map[string]testCase{
		"null": {
			val: types.ListNull(types.StringType),
			expectedDiagnostics: diag.Diagnostics{
				diag.NewAttributeErrorDiagnostic(
					path.Root("test"),
					"Missing Attribute Value",
					`test: value must be configured`,
				),
			},
		},
		"unknown": {
			val: types.ListUnknown(types.StringType),
		},
		"valid": {
			val: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("test"),
			}),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			request := validator.ListRequest{
				Path:           path.Root("test"),
				PathExpression: path.MatchRoot("test"),
				ConfigValue:    test.val,
			}
			response := validator.ListResponse{}
			fwvalidators.NotNullList().ValidateList(ctx, request, &response)

			if diff := cmp.Diff(response.Diagnostics, test.expectedDiagnostics); diff != "" {
				t.Errorf("unexpected diagnostics difference: %s", diff)
			}
		})
	}
}
