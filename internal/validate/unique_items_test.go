package validate

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
	schemaStringList := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"test": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
		},
	}
	schemaObjectList := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"test": schema.ListAttribute{
				Required:    true,
				ElementType: objectElementAttrType,
			},
		},
	}

	type testCase struct {
		val         tftypes.Value
		schema      schema.Schema
		expectError bool
	}
	tests := map[string]testCase{
		"not a list": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			schema:      schemaStringList,
			expectError: true,
		},
		"unknown list": {
			val:    tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, tftypes.UnknownValue),
			schema: schemaStringList,
		},
		"null list": {
			val:    tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, nil),
			schema: schemaStringList,
		},
		"empty list": {
			val:    tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, []tftypes.Value{}),
			schema: schemaStringList,
		},
		"unique list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			schema: schemaStringList,
		},
		"duplicates in list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
				tftypes.NewValue(tftypes.String, "beta"),
			}),
			schema:      schemaStringList,
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
			schema: schemaObjectList,
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
			schema:      schemaObjectList,
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
			schema: schemaObjectList,
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
			schema:      schemaObjectList,
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			request := validator.ListRequest{
				Config: tfsdk.Config{
					Raw:    test.val,
					Schema: test.schema,
				},
				Path: path.Root("test"),
			}
			response := validator.ListResponse{}
			UniqueItems().ValidateList(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}
