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

func TestUniqueItemsValidator(t *testing.T) {
	t.Parallel()

	objectElementType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"key": tftypes.String,
			"val": tftypes.String,
		},
	}
	objectElementAttrType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key": types.StringType,
			"val": types.StringType,
		},
	}

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		expectError bool
	}
	tests := map[string]testCase{
		"not a list": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			f:           types.BoolType.ValueFromTerraform,
			expectError: true,
		},
		"null list": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, nil),
			f:   types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
		},
		"empty list": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:   types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
		},
		"unique list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f: types.ListType{ElemType: types.StringType}.ValueFromTerraform,
		},
		"duplicates in list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
				tftypes.NewValue(tftypes.String, "beta"),
			}),
			f:           types.ListType{ElemType: types.StringType}.ValueFromTerraform,
			expectError: true,
		},
		"unique list of object": {
			val: tftypes.NewValue(
				tftypes.List{ElementType: objectElementType},
				[]tftypes.Value{
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key1"),
						"val": tftypes.NewValue(tftypes.String, "val1"),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key1"),
						"val": tftypes.NewValue(tftypes.String, "val2"),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key2"),
						"val": tftypes.NewValue(tftypes.String, "val2"),
					}),
				},
			),
			f: types.ListType{ElemType: objectElementAttrType}.ValueFromTerraform,
		},
		"duplicates in list of object": {
			val: tftypes.NewValue(
				tftypes.List{ElementType: objectElementType},
				[]tftypes.Value{
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key1"),
						"val": tftypes.NewValue(tftypes.String, "val1"),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key2"),
						"val": tftypes.NewValue(tftypes.String, "val2"),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key1"),
						"val": tftypes.NewValue(tftypes.String, "val1"),
					}),
				},
			),
			f:           types.ListType{ElemType: objectElementAttrType}.ValueFromTerraform,
			expectError: true,
		},
		"unique list of object with unknowns": {
			val: tftypes.NewValue(
				tftypes.List{ElementType: objectElementType},
				[]tftypes.Value{
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key1"),
						"val": tftypes.NewValue(tftypes.String, "val1"),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key1"),
						"val": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key1"),
						"val": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key2"),
						"val": tftypes.NewValue(tftypes.String, "val2"),
					}),
				},
			),
			f: types.ListType{ElemType: objectElementAttrType}.ValueFromTerraform,
		},
		"duplicates in list of object with unknowns": {
			val: tftypes.NewValue(
				tftypes.List{ElementType: objectElementType},
				[]tftypes.Value{
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key1"),
						"val": tftypes.NewValue(tftypes.String, "val1"),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key2"),
						"val": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key2"),
						"val": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
					}),
					tftypes.NewValue(objectElementType, map[string]tftypes.Value{
						"key": tftypes.NewValue(tftypes.String, "key1"),
						"val": tftypes.NewValue(tftypes.String, "val1"),
					}),
				},
			),
			f:           types.ListType{ElemType: objectElementAttrType}.ValueFromTerraform,
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
			UniqueItems().Validate(ctx, request, &response)

			if !tfresource.DiagsHasError(response.Diagnostics) && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if tfresource.DiagsHasError(response.Diagnostics) && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}
