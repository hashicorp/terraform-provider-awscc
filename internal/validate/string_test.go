// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/errs/diag"
)

func TestStringIsJsonValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.String
		expectError bool
	}
	tests := map[string]testCase{
		"unknown string": {
			val: basetypes.NewStringUnknown(),
		},
		"null string": {
			val: basetypes.NewStringNull(),
		},
		"valid JSON object": {
			val: basetypes.NewStringValue(`{
	"key": "value"
}`),
		},
		"parsing error": {
			val: basetypes.NewStringValue(`{
	"key": "value",
}`),
			expectError: true,
		},
		"not an object": {
			val:         basetypes.NewStringValue("[]"),
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			request := validator.StringRequest{
				ConfigValue: test.val,
				Path:        path.Root("test"),
			}
			response := validator.StringResponse{}
			StringIsJsonObject().ValidateString(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", ccdiag.DiagnosticsError(response.Diagnostics))
			}
		})
	}
}
