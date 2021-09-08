package validate

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func TestRequired(t *testing.T) {
	t.Parallel()

	type testCase struct {
		names       []string
		required    []string
		expectError bool
	}
	tests := map[string]testCase{
		"both empty": {},
		"none required": {
			names: []string{"alpha", "beta", "gamma"},
		},
		"some required": {
			names:    []string{"alpha", "beta", "gamma"},
			required: []string{"alpha", "gamma"},
		},
		"all required": {
			names:    []string{"alpha", "beta", "gamma"},
			required: []string{"alpha", "beta", "gamma"},
		},
		"missing one": {
			names:       []string{"alpha", "beta", "gamma"},
			required:    []string{"beta", "delta"},
			expectError: true,
		},
		"missing all": {
			names:       []string{"alpha", "beta", "gamma"},
			required:    []string{"sigma", "tau"},
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			diags := Required(test.required...)(test.names)

			if !diags.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if diags.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(diags))
			}
		})
	}
}

func TestAllOfRequired(t *testing.T) {
	t.Parallel()

	type testCase struct {
		names       []string
		fs          []RequiredAttributesFunc
		expectError bool
	}
	tests := map[string]testCase{
		"both empty": {},
		"none required": {
			names: []string{"alpha", "beta", "gamma"},
		},
		"some required": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("alpha"), Required("gamma")},
		},
		"all required": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("alpha"), AllOfRequired(Required("beta"), Required("gamma"))},
		},
		"missing one": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("beta"), Required("delta")},
			expectError: true,
		},
		"missing all": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("sigma"), AllOfRequired(Required("tau"))},
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			diags := AllOfRequired(test.fs...)(test.names)

			if !diags.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if diags.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(diags))
			}
		})
	}
}

func TestAnyOfRequired(t *testing.T) {
	t.Parallel()

	type testCase struct {
		names       []string
		fs          []RequiredAttributesFunc
		expectError bool
	}
	tests := map[string]testCase{
		"both empty": {},
		"none required": {
			names: []string{"alpha", "beta", "gamma"},
		},
		"some required": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("alpha"), Required("delta")},
		},
		"nested allOf OK": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("delta"), AllOfRequired(Required("beta"), Required("gamma"))},
		},
		"nested allOf error": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("delta"), AllOfRequired(Required("beta"), Required("epsilon"))},
			expectError: true,
		},
		"nested anyOf OK": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("delta"), AnyOfRequired(Required("beta"), Required("gamma"))},
		},
		"nested anyOf error": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("delta"), AnyOfRequired(Required("sigma"), Required("tau"))},
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			diags := AnyOfRequired(test.fs...)(test.names)

			if !diags.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if diags.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(diags))
			}
		})
	}
}

func TestOneOfRequired(t *testing.T) {
	t.Parallel()

	type testCase struct {
		names       []string
		fs          []RequiredAttributesFunc
		expectError bool
	}
	tests := map[string]testCase{
		"both empty": {},
		"none required": {
			names: []string{"alpha", "beta", "gamma"},
		},
		"some required OK": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("alpha"), Required("delta")},
		},
		"some required error": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("alpha"), Required("beta")},
			expectError: true,
		},
		"nested allOf OK": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("delta"), AllOfRequired(Required("beta"), Required("gamma"))},
		},
		"nested allOf error": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("delta"), AllOfRequired(Required("beta"), Required("epsilon"))},
			expectError: true,
		},
		"nested anyOf OK": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("delta"), AnyOfRequired(Required("beta"), Required("gamma"))},
		},
		"nested anyOf error": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("delta"), AnyOfRequired(Required("sigma"), Required("tau"))},
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			diags := OneOfRequired(test.fs...)(test.names)

			if !diags.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if diags.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(diags))
			}
		})
	}
}

func TestRequiredAttributesValidator_Object(t *testing.T) {
	t.Parallel()

	objectElementType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"alpha": tftypes.String,
			"beta":  tftypes.String,
			"gamma": tftypes.Bool,
			"delta": tftypes.Number,
		},
	}
	objectElementAttrType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alpha": types.StringType,
			"beta":  types.StringType,
			"gamma": types.BoolType,
			"delta": types.NumberType,
		},
	}

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		fs          []RequiredAttributesFunc
		expectError bool
	}
	tests := map[string]testCase{
		"not an object": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			f:           types.BoolType.ValueFromTerraform,
			expectError: true,
		},
		"nil object": {
			val: tftypes.NewValue(objectElementType, nil),
			f:   objectElementAttrType.ValueFromTerraform,
			fs:  []RequiredAttributesFunc{Required("alpha")},
		},
		"unknown object": {
			val: tftypes.NewValue(objectElementType, tftypes.UnknownValue),
			f:   objectElementAttrType.ValueFromTerraform,
			fs:  []RequiredAttributesFunc{Required("alpha")},
		},
		"not fully known object": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, nil),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, tftypes.UnknownValue),
				"delta": tftypes.NewValue(tftypes.Number, nil),
			}),
			f:  objectElementAttrType.ValueFromTerraform,
			fs: []RequiredAttributesFunc{Required("alpha")},
		},
		"none required": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, "a"),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, true),
				"delta": tftypes.NewValue(tftypes.Number, 42),
			}),
			f: objectElementAttrType.ValueFromTerraform,
		},
		"one required OK": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, "a"),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, true),
				"delta": tftypes.NewValue(tftypes.Number, 42),
			}),
			f:  objectElementAttrType.ValueFromTerraform,
			fs: []RequiredAttributesFunc{Required("alpha")},
		},
		"one required error": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, "a"),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, true),
				"delta": tftypes.NewValue(tftypes.Number, 42),
			}),
			f:           objectElementAttrType.ValueFromTerraform,
			fs:          []RequiredAttributesFunc{Required("beta")},
			expectError: true,
		},
		"two required OK anyOf": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, "a"),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, true),
				"delta": tftypes.NewValue(tftypes.Number, 42),
			}),
			f:  objectElementAttrType.ValueFromTerraform,
			fs: []RequiredAttributesFunc{AnyOfRequired(Required("beta"), AllOfRequired(Required("gamma"), Required("delta")))},
		},
		"two required OK oneOf": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, "a"),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, true),
				"delta": tftypes.NewValue(tftypes.Number, 42),
			}),
			f:  objectElementAttrType.ValueFromTerraform,
			fs: []RequiredAttributesFunc{OneOfRequired(Required("beta"), AllOfRequired(Required("gamma"), Required("delta")))},
		},
		"two required error oneOf": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, "a"),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, true),
				"delta": tftypes.NewValue(tftypes.Number, 42),
			}),
			f:           objectElementAttrType.ValueFromTerraform,
			fs:          []RequiredAttributesFunc{OneOfRequired(Required("alpha"), AllOfRequired(Required("gamma"), Required("delta")))},
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
			RequiredAttributes(test.fs...).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}

func TestRequiredAttributesValidator_List(t *testing.T) {
	t.Parallel()

	objectElementType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"alpha": tftypes.String,
			"beta":  tftypes.String,
			"gamma": tftypes.Bool,
			"delta": tftypes.Number,
		},
	}
	objectElementAttrType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alpha": types.StringType,
			"beta":  types.StringType,
			"gamma": types.BoolType,
			"delta": types.NumberType,
		},
	}

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		fs          []RequiredAttributesFunc
		expectError bool
	}
	tests := map[string]testCase{
		"unknown list": {
			val: tftypes.NewValue(
				tftypes.List{ElementType: objectElementType},
				tftypes.UnknownValue,
			),
			f: types.ListType{ElemType: objectElementAttrType}.ValueFromTerraform,
		},
		"null list": {
			val: tftypes.NewValue(
				tftypes.List{ElementType: objectElementType},
				nil,
			),
			f: types.ListType{ElemType: objectElementAttrType}.ValueFromTerraform,
		},
		"empty list": {
			val: tftypes.NewValue(
				tftypes.List{ElementType: objectElementType},
				[]tftypes.Value{},
			),
			f: types.ListType{ElemType: objectElementAttrType}.ValueFromTerraform,
		},
		"not fully known object": {
			val: tftypes.NewValue(
				tftypes.List{ElementType: objectElementType},
				[]tftypes.Value{
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a"),
						"beta":  tftypes.NewValue(tftypes.String, "b"),
						"gamma": tftypes.NewValue(tftypes.Bool, true),
						"delta": tftypes.NewValue(tftypes.Number, 42),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, nil),
						"beta":  tftypes.NewValue(tftypes.String, nil),
						"gamma": tftypes.NewValue(tftypes.Bool, tftypes.UnknownValue),
						"delta": tftypes.NewValue(tftypes.Number, nil),
					}),
				},
			),
			f:  types.ListType{ElemType: objectElementAttrType}.ValueFromTerraform,
			fs: []RequiredAttributesFunc{Required("alpha")},
		},
		"none required": {
			val: tftypes.NewValue(
				tftypes.List{ElementType: objectElementType},
				[]tftypes.Value{
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a1"),
						"beta":  tftypes.NewValue(tftypes.String, "b1"),
						"gamma": tftypes.NewValue(tftypes.Bool, true),
						"delta": tftypes.NewValue(tftypes.Number, 1),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a2"),
						"beta":  tftypes.NewValue(tftypes.String, nil),
						"gamma": tftypes.NewValue(tftypes.Bool, nil),
						"delta": tftypes.NewValue(tftypes.Number, 2),
					}),
				},
			),
			f: types.ListType{ElemType: objectElementAttrType}.ValueFromTerraform,
		},
		"one required OK": {
			val: tftypes.NewValue(
				tftypes.List{ElementType: objectElementType},
				[]tftypes.Value{
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a1"),
						"beta":  tftypes.NewValue(tftypes.String, "b1"),
						"gamma": tftypes.NewValue(tftypes.Bool, true),
						"delta": tftypes.NewValue(tftypes.Number, 1),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a2"),
						"beta":  tftypes.NewValue(tftypes.String, nil),
						"gamma": tftypes.NewValue(tftypes.Bool, nil),
						"delta": tftypes.NewValue(tftypes.Number, 2),
					}),
				},
			),
			f:  types.ListType{ElemType: objectElementAttrType}.ValueFromTerraform,
			fs: []RequiredAttributesFunc{Required("alpha")},
		},
		"one required error": {
			val: tftypes.NewValue(
				tftypes.List{ElementType: objectElementType},
				[]tftypes.Value{
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a1"),
						"beta":  tftypes.NewValue(tftypes.String, "b1"),
						"gamma": tftypes.NewValue(tftypes.Bool, true),
						"delta": tftypes.NewValue(tftypes.Number, 1),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a2"),
						"beta":  tftypes.NewValue(tftypes.String, nil),
						"gamma": tftypes.NewValue(tftypes.Bool, nil),
						"delta": tftypes.NewValue(tftypes.Number, 2),
					}),
				},
			),
			f:           types.ListType{ElemType: objectElementAttrType}.ValueFromTerraform,
			fs:          []RequiredAttributesFunc{Required("beta")},
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
			RequiredAttributes(test.fs...).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}

func TestRequiredAttributesValidator_Set(t *testing.T) {
	t.Parallel()

	objectElementType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"alpha": tftypes.String,
			"beta":  tftypes.String,
			"gamma": tftypes.Bool,
			"delta": tftypes.Number,
		},
	}
	objectElementAttrType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alpha": types.StringType,
			"beta":  types.StringType,
			"gamma": types.BoolType,
			"delta": types.NumberType,
		},
	}

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		fs          []RequiredAttributesFunc
		expectError bool
	}
	tests := map[string]testCase{
		"unknown set": {
			val: tftypes.NewValue(
				tftypes.Set{ElementType: objectElementType},
				tftypes.UnknownValue,
			),
			f: providertypes.SetType{ElemType: objectElementAttrType}.ValueFromTerraform,
		},
		"null set": {
			val: tftypes.NewValue(
				tftypes.Set{ElementType: objectElementType},
				nil,
			),
			f: providertypes.SetType{ElemType: objectElementAttrType}.ValueFromTerraform,
		},
		"empty set": {
			val: tftypes.NewValue(
				tftypes.Set{ElementType: objectElementType},
				[]tftypes.Value{},
			),
			f: providertypes.SetType{ElemType: objectElementAttrType}.ValueFromTerraform,
		},
		"not fully known object": {
			val: tftypes.NewValue(
				tftypes.Set{ElementType: objectElementType},
				[]tftypes.Value{
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a"),
						"beta":  tftypes.NewValue(tftypes.String, "b"),
						"gamma": tftypes.NewValue(tftypes.Bool, true),
						"delta": tftypes.NewValue(tftypes.Number, 42),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, nil),
						"beta":  tftypes.NewValue(tftypes.String, nil),
						"gamma": tftypes.NewValue(tftypes.Bool, tftypes.UnknownValue),
						"delta": tftypes.NewValue(tftypes.Number, nil),
					}),
				},
			),
			f: providertypes.SetType{ElemType: objectElementAttrType}.ValueFromTerraform,
		},
		"none required": {
			val: tftypes.NewValue(
				tftypes.Set{ElementType: objectElementType},
				[]tftypes.Value{
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a1"),
						"beta":  tftypes.NewValue(tftypes.String, "b1"),
						"gamma": tftypes.NewValue(tftypes.Bool, true),
						"delta": tftypes.NewValue(tftypes.Number, 1),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a2"),
						"beta":  tftypes.NewValue(tftypes.String, nil),
						"gamma": tftypes.NewValue(tftypes.Bool, nil),
						"delta": tftypes.NewValue(tftypes.Number, 2),
					}),
				},
			),
			f: providertypes.SetType{ElemType: objectElementAttrType}.ValueFromTerraform,
		},
		"one required OK": {
			val: tftypes.NewValue(
				tftypes.Set{ElementType: objectElementType},
				[]tftypes.Value{
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a1"),
						"beta":  tftypes.NewValue(tftypes.String, "b1"),
						"gamma": tftypes.NewValue(tftypes.Bool, true),
						"delta": tftypes.NewValue(tftypes.Number, 1),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a2"),
						"beta":  tftypes.NewValue(tftypes.String, nil),
						"gamma": tftypes.NewValue(tftypes.Bool, nil),
						"delta": tftypes.NewValue(tftypes.Number, 2),
					}),
				},
			),
			f:  providertypes.SetType{ElemType: objectElementAttrType}.ValueFromTerraform,
			fs: []RequiredAttributesFunc{Required("alpha")},
		},
		"one required error": {
			val: tftypes.NewValue(
				tftypes.Set{ElementType: objectElementType},
				[]tftypes.Value{
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a1"),
						"beta":  tftypes.NewValue(tftypes.String, "b1"),
						"gamma": tftypes.NewValue(tftypes.Bool, true),
						"delta": tftypes.NewValue(tftypes.Number, 1),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"alpha": tftypes.NewValue(tftypes.String, "a2"),
						"beta":  tftypes.NewValue(tftypes.String, nil),
						"gamma": tftypes.NewValue(tftypes.Bool, nil),
						"delta": tftypes.NewValue(tftypes.Number, 2),
					}),
				},
			),
			f:           providertypes.SetType{ElemType: objectElementAttrType}.ValueFromTerraform,
			fs:          []RequiredAttributesFunc{Required("beta")},
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
			RequiredAttributes(test.fs...).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}

func TestResourceConfigRequiredAttributesValidator(t *testing.T) {
	t.Parallel()

	objectElementType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"alpha": tftypes.String,
			"beta":  tftypes.String,
			"gamma": tftypes.Bool,
			"delta": tftypes.Number,
		},
	}

	type testCase struct {
		val         tftypes.Value
		fs          []RequiredAttributesFunc
		expectError bool
	}
	tests := map[string]testCase{
		"not an object": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			expectError: true,
		},
		"nil object": {
			val: tftypes.NewValue(objectElementType, nil),
			fs:  []RequiredAttributesFunc{Required("alpha")},
		},
		"unknown object": {
			val: tftypes.NewValue(objectElementType, tftypes.UnknownValue),
			fs:  []RequiredAttributesFunc{Required("alpha")},
		},
		"not fully known object": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, nil),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, tftypes.UnknownValue),
				"delta": tftypes.NewValue(tftypes.Number, nil),
			}),
			fs: []RequiredAttributesFunc{Required("alpha")},
		},
		"none required": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, "a"),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, true),
				"delta": tftypes.NewValue(tftypes.Number, 42),
			}),
		},
		"one required OK": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, "a"),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, true),
				"delta": tftypes.NewValue(tftypes.Number, 42),
			}),
			fs: []RequiredAttributesFunc{Required("alpha")},
		},
		"one required error": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, "a"),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, true),
				"delta": tftypes.NewValue(tftypes.Number, 42),
			}),
			fs:          []RequiredAttributesFunc{Required("beta")},
			expectError: true,
		},
		"two required OK anyOf": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, "a"),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, true),
				"delta": tftypes.NewValue(tftypes.Number, 42),
			}),
			fs: []RequiredAttributesFunc{AnyOfRequired(Required("beta"), AllOfRequired(Required("gamma"), Required("delta")))},
		},
		"two required OK oneOf": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, "a"),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, true),
				"delta": tftypes.NewValue(tftypes.Number, 42),
			}),
			fs: []RequiredAttributesFunc{OneOfRequired(Required("beta"), AllOfRequired(Required("gamma"), Required("delta")))},
		},
		"two required error oneOf": {
			val: tftypes.NewValue(objectElementType, map[string]tftypes.Value{
				"alpha": tftypes.NewValue(tftypes.String, "a"),
				"beta":  tftypes.NewValue(tftypes.String, nil),
				"gamma": tftypes.NewValue(tftypes.Bool, true),
				"delta": tftypes.NewValue(tftypes.Number, 42),
			}),
			fs:          []RequiredAttributesFunc{OneOfRequired(Required("alpha"), AllOfRequired(Required("gamma"), Required("delta")))},
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := tfsdk.ValidateResourceConfigRequest{
				Config: tfsdk.Config{
					Raw: test.val,
				},
			}
			response := tfsdk.ValidateResourceConfigResponse{}
			ResourceConfigRequiredAttributes(test.fs...).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}
