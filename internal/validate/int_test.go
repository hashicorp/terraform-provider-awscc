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

func TestIntBetweenValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		min         int
		max         int
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
			min: 1,
			max: 3,
		},
		"null number": {
			val: tftypes.NewValue(tftypes.Number, nil),
			f:   types.NumberType.ValueFromTerraform,
			min: 1,
			max: 3,
		},
		"not an integer": {
			val:         tftypes.NewValue(tftypes.Number, 2.34),
			f:           types.NumberType.ValueFromTerraform,
			min:         1,
			max:         3,
			expectError: true,
		},
		"valid integer": {
			val: tftypes.NewValue(tftypes.Number, 2),
			f:   types.NumberType.ValueFromTerraform,
			min: 1,
			max: 3,
		},
		"valid integer min": {
			val: tftypes.NewValue(tftypes.Number, 1),
			f:   types.NumberType.ValueFromTerraform,
			min: 1,
			max: 3,
		},
		"valid integer max": {
			val: tftypes.NewValue(tftypes.Number, 3),
			f:   types.NumberType.ValueFromTerraform,
			min: 1,
			max: 3,
		},
		"too small integer": {
			val:         tftypes.NewValue(tftypes.Number, -1),
			f:           types.NumberType.ValueFromTerraform,
			min:         1,
			max:         3,
			expectError: true,
		},
		"too large integer": {
			val:         tftypes.NewValue(tftypes.Number, 42),
			f:           types.NumberType.ValueFromTerraform,
			min:         1,
			max:         3,
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
			IntBetween(test.min, test.max).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}

func TestIntAtLeastValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		min         int
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
			min: 1,
		},
		"null number": {
			val: tftypes.NewValue(tftypes.Number, nil),
			f:   types.NumberType.ValueFromTerraform,
			min: 1,
		},
		"not an integer": {
			val:         tftypes.NewValue(tftypes.Number, 2.34),
			f:           types.NumberType.ValueFromTerraform,
			min:         1,
			expectError: true,
		},
		"valid integer": {
			val: tftypes.NewValue(tftypes.Number, 2),
			f:   types.NumberType.ValueFromTerraform,
			min: 1,
		},
		"valid integer min": {
			val: tftypes.NewValue(tftypes.Number, 1),
			f:   types.NumberType.ValueFromTerraform,
			min: 1,
		},
		"too small integer": {
			val:         tftypes.NewValue(tftypes.Number, -1),
			f:           types.NumberType.ValueFromTerraform,
			min:         1,
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
			IntAtLeast(test.min).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}

func TestIntInSliceValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		valid       []int
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
		},
		"null number": {
			val: tftypes.NewValue(tftypes.Number, nil),
			f:   types.NumberType.ValueFromTerraform,
		},
		"not an integer": {
			val:         tftypes.NewValue(tftypes.Number, 2.34),
			f:           types.NumberType.ValueFromTerraform,
			expectError: true,
		},
		"valid integer": {
			val:   tftypes.NewValue(tftypes.Number, 2),
			f:     types.NumberType.ValueFromTerraform,
			valid: []int{-1, 2, 42},
		},
		"invalid integer": {
			val:         tftypes.NewValue(tftypes.Number, 1),
			f:           types.NumberType.ValueFromTerraform,
			valid:       []int{-1, 2, 42},
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
			IntInSlice(test.valid).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}
