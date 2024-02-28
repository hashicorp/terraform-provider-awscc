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

func TestARNValidator(t *testing.T) {
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
		"valid string": {
			val: basetypes.NewStringValue("arn:aws:kafka:us-west-2:123456789012:cluster/tf-acc-test-3972327032919409894/09494266-aaf8-48e7-b7ce-8548498bb813-11"),
		},
		"invalid string": {
			val:         basetypes.NewStringValue("not ok"),
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
			ARN().ValidateString(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", ccdiag.DiagnosticsError(response.Diagnostics))
			}
		})
	}
}

func TestIAMPolicyARNValidator(t *testing.T) {
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
		"valid IAM Policy ARN": {
			val: basetypes.NewStringValue("arn:aws:iam::123456789012:policy/policy_name"),
		},
		"invalid ARN": {
			val:         basetypes.NewStringValue("arn:aws:iam::123456789012:user/user_name"),
			expectError: true,
		},
		"not an ARN": {
			val:         basetypes.NewStringValue("not an ARN"),
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
			IAMPolicyARN().ValidateString(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", ccdiag.DiagnosticsError(response.Diagnostics))
			}
		})
	}
}
