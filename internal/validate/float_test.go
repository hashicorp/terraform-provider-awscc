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

func TestFloatBetweenValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		min         float64
		max         float64
		expectError bool
	}
	tests := map[string]testCase{
		"not a number": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			f:           types.BoolType.ValueFromTerraform,
			expectError: true,
		},
		"unknown number": {
			val: tftypes.NewValue(tftypes.Number, tftypes.UnknownValue),
			f:   types.NumberType.ValueFromTerraform,
			min: 0.90,
			max: 3.10,
		},
		"null number": {
			val: tftypes.NewValue(tftypes.Number, nil),
			f:   types.NumberType.ValueFromTerraform,
			min: 0.90,
			max: 3.10,
		},
		"valid integer": {
			val: tftypes.NewValue(tftypes.Number, 2),
			f:   types.NumberType.ValueFromTerraform,
			min: 0.90,
			max: 3.10,
		},
		"valid float": {
			val: tftypes.NewValue(tftypes.Number, 2.2),
			f:   types.NumberType.ValueFromTerraform,
			min: 0.90,
			max: 3.10,
		},
		"valid float min": {
			val: tftypes.NewValue(tftypes.Number, 0.9),
			f:   types.NumberType.ValueFromTerraform,
			min: 0.90,
			max: 3.10,
		},
		"valid float max": {
			val: tftypes.NewValue(tftypes.Number, 3.10),
			f:   types.NumberType.ValueFromTerraform,
			min: 0.90,
			max: 3.10,
		},
		"too small float": {
			val:         tftypes.NewValue(tftypes.Number, -1.1111),
			f:           types.NumberType.ValueFromTerraform,
			min:         0.90,
			max:         3.10,
			expectError: true,
		},
		"too large float": {
			val:         tftypes.NewValue(tftypes.Number, 4.2),
			f:           types.NumberType.ValueFromTerraform,
			min:         0.90,
			max:         3.10,
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
			FloatBetween(test.min, test.max).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}

func TestFloatAtLeastValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		min         float64
		expectError bool
	}
	tests := map[string]testCase{
		"not a number": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			f:           types.BoolType.ValueFromTerraform,
			expectError: true,
		},
		"unknown number": {
			val: tftypes.NewValue(tftypes.Number, tftypes.UnknownValue),
			f:   types.NumberType.ValueFromTerraform,
			min: 0.90,
		},
		"null number": {
			val: tftypes.NewValue(tftypes.Number, nil),
			f:   types.NumberType.ValueFromTerraform,
			min: 0.90,
		},
		"valid integer": {
			val: tftypes.NewValue(tftypes.Number, 2),
			f:   types.NumberType.ValueFromTerraform,
			min: 0.90,
		},
		"valid float": {
			val: tftypes.NewValue(tftypes.Number, 2.2),
			f:   types.NumberType.ValueFromTerraform,
			min: 0.90,
		},
		"valid float min": {
			val: tftypes.NewValue(tftypes.Number, 0.9),
			f:   types.NumberType.ValueFromTerraform,
			min: 0.90,
		},
		"too small float": {
			val:         tftypes.NewValue(tftypes.Number, -1.1111),
			f:           types.NumberType.ValueFromTerraform,
			min:         0.90,
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
			FloatAtLeast(test.min).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}
