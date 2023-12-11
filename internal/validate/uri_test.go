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
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func TestURIValidator(t *testing.T) {
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
			val: basetypes.NewStringValue("http://mystack-mybucket-kdwwxmddtr2g.s3.dualstack.us-east-2.amazonaws.com/"),
		},
		"invalid string": {
			val:         basetypes.NewStringValue("://mystack-mybucket-kdwwxmddtr2g.s3-website-us-east-2.amazonaws.com/"),
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
			IsURI().ValidateString(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagnosticsError(response.Diagnostics))
			}
		})
	}
}
