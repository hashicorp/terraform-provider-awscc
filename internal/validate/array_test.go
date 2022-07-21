package validate

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func TestArrayLenBetweenValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		minItems    int
		maxItems    int
		expectError bool
	}
	tests := map[string]testCase{
		"not a list or set": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			f:           types.BoolType.ValueFromTerraform,
			expectError: true,
		},
		"unknown list": {
			val:      tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, tftypes.UnknownValue),
			f:        types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems: 1,
			maxItems: 3,
		},
		"null list": {
			val:      tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, nil),
			f:        types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems: 1,
			maxItems: 3,
		},
		"valid empty list": {
			val:      tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:        types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems: 0,
			maxItems: 3,
		},
		"invalid empty list": {
			val:         tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:           types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems:    1,
			maxItems:    3,
			expectError: true,
		},
		"valid list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:        types.ListType{ElemType: types.StringType}.ValueFromTerraform,
			minItems: 2,
			maxItems: 3,
		},
		"invalid list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
				tftypes.NewValue(tftypes.String, "delta"),
			}),
			f:           types.ListType{ElemType: types.StringType}.ValueFromTerraform,
			minItems:    2,
			maxItems:    3,
			expectError: true,
		},
		"unknown set": {
			val:      tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, tftypes.UnknownValue),
			f:        types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems: 0,
			maxItems: 3,
		},
		"null set": {
			val:      tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, nil),
			f:        types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems: 0,
			maxItems: 3,
		},
		"valid empty set": {
			val:      tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:        types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems: 0,
			maxItems: 3,
		},
		"invalid empty set": {
			val:         tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:           types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems:    1,
			maxItems:    3,
			expectError: true,
		},
		"valid set of string": {
			val: tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:        types.SetType{ElemType: types.StringType}.ValueFromTerraform,
			minItems: 2,
			maxItems: 3,
		},
		"invalid set of string": {
			val: tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
				tftypes.NewValue(tftypes.String, "delta"),
			}),
			f:           types.SetType{ElemType: types.StringType}.ValueFromTerraform,
			minItems:    2,
			maxItems:    3,
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
				AttributePath:   path.Root("test"),
				AttributeConfig: val,
			}
			response := tfsdk.ValidateAttributeResponse{}
			ArrayLenBetween(test.minItems, test.maxItems).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}

func TestArrayLenAtLeastValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		minItems    int
		expectError bool
	}
	tests := map[string]testCase{
		"not a list or set": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			f:           types.BoolType.ValueFromTerraform,
			expectError: true,
		},
		"unknown list": {
			val:      tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, tftypes.UnknownValue),
			f:        types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems: 1,
		},
		"null list": {
			val:      tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, nil),
			f:        types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems: 1,
		},
		"valid empty list": {
			val:      tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:        types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems: 0,
		},
		"invalid empty list": {
			val:         tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:           types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems:    1,
			expectError: true,
		},
		"valid list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:        types.ListType{ElemType: types.StringType}.ValueFromTerraform,
			minItems: 2,
		},
		"invalid list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
			}),
			f:           types.ListType{ElemType: types.StringType}.ValueFromTerraform,
			minItems:    2,
			expectError: true,
		},
		"unknown set": {
			val:      tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, tftypes.UnknownValue),
			f:        types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems: 1,
		},
		"null set": {
			val:      tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, nil),
			f:        types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems: 1,
		},
		"valid empty set": {
			val:      tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:        types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems: 0,
		},
		"invalid empty set": {
			val:         tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:           types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
			minItems:    1,
			expectError: true,
		},
		"valid set of string": {
			val: tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:        types.SetType{ElemType: types.StringType}.ValueFromTerraform,
			minItems: 2,
		},
		"invalid set of string": {
			val: tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
			}),
			f:           types.SetType{ElemType: types.StringType}.ValueFromTerraform,
			minItems:    2,
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
				AttributePath:   path.Root("test"),
				AttributeConfig: val,
			}
			response := tfsdk.ValidateAttributeResponse{}
			ArrayLenAtLeast(test.minItems).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}

func TestArrayLenAtMostValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		maxItems    int
		expectError bool
	}
	tests := map[string]testCase{
		"not a list or set": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			f:           types.BoolType.ValueFromTerraform,
			expectError: true,
		},
		"unknown list": {
			val:      tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, tftypes.UnknownValue),
			f:        types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			maxItems: 1,
		},
		"null list": {
			val:      tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, nil),
			f:        types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			maxItems: 1,
		},
		"valid empty list": {
			val:      tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:        types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			maxItems: 1,
		},
		"valid list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:        types.ListType{ElemType: types.StringType}.ValueFromTerraform,
			maxItems: 4,
		},
		"invalid list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:           types.ListType{ElemType: types.StringType}.ValueFromTerraform,
			maxItems:    2,
			expectError: true,
		},
		"unknown set": {
			val: tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, tftypes.UnknownValue),
			f:   types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
		},
		"null set": {
			val: tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, nil),
			f:   types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
		},
		"valid empty set": {
			val:      tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:        types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
			maxItems: 1,
		},
		"valid set of string": {
			val: tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:        types.SetType{ElemType: types.StringType}.ValueFromTerraform,
			maxItems: 4,
		},
		"invalid set of string": {
			val: tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:           types.SetType{ElemType: types.StringType}.ValueFromTerraform,
			maxItems:    2,
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
				AttributePath:   path.Root("test"),
				AttributeConfig: val,
			}
			response := tfsdk.ValidateAttributeResponse{}
			ArrayLenAtMost(test.maxItems).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}
