package validate

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func TestStringLenBetweenValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		minLength   int
		maxLength   int
		expectError bool
	}
	tests := map[string]testCase{
		"not a string": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			f:           types.BoolType.ValueFromTerraform,
			expectError: true,
		},
		"unknown string": {
			val:       tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			f:         types.StringType.ValueFromTerraform,
			minLength: 1,
			maxLength: 3,
		},
		"null string": {
			val:       tftypes.NewValue(tftypes.String, nil),
			f:         types.StringType.ValueFromTerraform,
			minLength: 1,
			maxLength: 3,
		},
		"valid string": {
			val:       tftypes.NewValue(tftypes.String, "ok"),
			f:         types.StringType.ValueFromTerraform,
			minLength: 1,
			maxLength: 3,
		},
		"too long string": {
			val:         tftypes.NewValue(tftypes.String, "not ok"),
			f:           types.StringType.ValueFromTerraform,
			minLength:   1,
			maxLength:   3,
			expectError: true,
		},
		"too short string": {
			val:         tftypes.NewValue(tftypes.String, ""),
			f:           types.StringType.ValueFromTerraform,
			minLength:   1,
			maxLength:   3,
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			val, err := test.f(ctx, test.val)

			if err != nil {
				t.Fatalf("got unexpected error: %s", err)
			}

			request := tfsdk.ValidateAttributeRequest{
				AttributePath:   tftypes.NewAttributePath().WithAttributeName("test"),
				AttributeConfig: val,
			}
			response := tfsdk.ValidateAttributeResponse{}
			StringLenBetween(test.minLength, test.maxLength).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}

func TestStringLenAtLeastValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		minLength   int
		expectError bool
	}
	tests := map[string]testCase{
		"not a string": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			f:           types.BoolType.ValueFromTerraform,
			expectError: true,
		},
		"unknown string": {
			val:       tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			f:         types.StringType.ValueFromTerraform,
			minLength: 1,
		},
		"null string": {
			val:       tftypes.NewValue(tftypes.String, nil),
			f:         types.StringType.ValueFromTerraform,
			minLength: 1,
		},
		"valid string": {
			val:       tftypes.NewValue(tftypes.String, "ok"),
			f:         types.StringType.ValueFromTerraform,
			minLength: 1,
		},
		"too short string": {
			val:         tftypes.NewValue(tftypes.String, ""),
			f:           types.StringType.ValueFromTerraform,
			minLength:   1,
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			val, err := test.f(ctx, test.val)

			if err != nil {
				t.Fatalf("got unexpected error: %s", err)
			}

			request := tfsdk.ValidateAttributeRequest{
				AttributePath:   tftypes.NewAttributePath().WithAttributeName("test"),
				AttributeConfig: val,
			}
			response := tfsdk.ValidateAttributeResponse{}
			StringLenAtLeast(test.minLength).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}

func TestStringinSliceValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		valid       []string
		expectError bool
	}
	tests := map[string]testCase{
		"not a string": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			f:           types.BoolType.ValueFromTerraform,
			expectError: true,
		},
		"unknown string": {
			val: tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			f:   types.StringType.ValueFromTerraform,
		},
		"null string": {
			val: tftypes.NewValue(tftypes.String, nil),
			f:   types.StringType.ValueFromTerraform,
		},
		"valid string": {
			val:   tftypes.NewValue(tftypes.String, "ok"),
			f:     types.StringType.ValueFromTerraform,
			valid: []string{"alpha", "ok", "beta"},
		},
		"invalid string": {
			val:         tftypes.NewValue(tftypes.String, "not ok"),
			f:           types.StringType.ValueFromTerraform,
			valid:       []string{"alpha", "ok", "beta"},
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			val, err := test.f(ctx, test.val)

			if err != nil {
				t.Fatalf("got unexpected error: %s", err)
			}

			request := tfsdk.ValidateAttributeRequest{
				AttributePath:   tftypes.NewAttributePath().WithAttributeName("test"),
				AttributeConfig: val,
			}
			response := tfsdk.ValidateAttributeResponse{}
			StringInSlice(test.valid).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}
