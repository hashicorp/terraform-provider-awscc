// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaults

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestStaticPartialObject_simple(t *testing.T) {
	t.Parallel()

	attributeTypes := map[string]attr.Type{
		"name":   types.StringType,
		"wicked": types.BoolType,
	}
	defaultVal := map[string]interface{}{
		"name": "defaultName",
	}

	type testCase struct {
		plannedValue  types.Object
		expectedValue types.Object
	}
	tests := map[string]testCase{
		"unknown object": {
			plannedValue: types.ObjectUnknown(attributeTypes),
			expectedValue: types.ObjectValueMust(attributeTypes, map[string]attr.Value{
				"name":   types.StringValue("defaultName"),
				"wicked": types.BoolNull(),
			}),
		},
		"null object": {
			plannedValue:  types.ObjectNull(attributeTypes),
			expectedValue: types.ObjectNull(attributeTypes),
		},
		"non-null object": {
			plannedValue: types.ObjectValueMust(attributeTypes, map[string]attr.Value{
				"name":   types.StringValue("n1"),
				"wicked": types.BoolValue(true),
			}),
			expectedValue: types.ObjectValueMust(attributeTypes, map[string]attr.Value{
				"name":   types.StringValue("n1"),
				"wicked": types.BoolValue(true),
			}),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := planmodifier.ObjectRequest{
				Path:      path.Root("test"),
				PlanValue: test.plannedValue,
			}
			response := planmodifier.ObjectResponse{}
			StaticPartialObject(defaultVal).PlanModifyObject(ctx, request, &response)

			if diff := cmp.Diff(response.PlanValue, test.expectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestStaticPartialObject_complex(t *testing.T) {
	t.Parallel()

	innerAttributeTypes := map[string]attr.Type{
		"name":   types.StringType,
		"wicked": types.BoolType,
	}
	outerAttributeTypes := map[string]attr.Type{
		"config": types.ObjectType{AttrTypes: innerAttributeTypes},
		"id":     types.StringType,
	}
	defaultVal := map[string]interface{}{
		"config": map[string]interface{}{
			"name": "defaultName",
		},
	}

	type testCase struct {
		plannedValue  types.Object
		expectedValue types.Object
	}
	tests := map[string]testCase{
		"unknown object": {
			plannedValue: types.ObjectUnknown(outerAttributeTypes),
			expectedValue: types.ObjectValueMust(outerAttributeTypes, map[string]attr.Value{
				"config": types.ObjectValueMust(innerAttributeTypes, map[string]attr.Value{
					"name":   types.StringValue("defaultName"),
					"wicked": types.BoolNull(),
				}),
				"id": types.StringNull(),
			}),
		},
		"null object": {
			plannedValue:  types.ObjectNull(outerAttributeTypes),
			expectedValue: types.ObjectNull(outerAttributeTypes),
		},
		"non-null object": {
			plannedValue: types.ObjectValueMust(outerAttributeTypes, map[string]attr.Value{
				"config": types.ObjectValueMust(innerAttributeTypes, map[string]attr.Value{
					"name":   types.StringValue("n1"),
					"wicked": types.BoolValue(true),
				}),
				"id": types.StringValue("id1"),
			}),
			expectedValue: types.ObjectValueMust(outerAttributeTypes, map[string]attr.Value{
				"config": types.ObjectValueMust(innerAttributeTypes, map[string]attr.Value{
					"name":   types.StringValue("n1"),
					"wicked": types.BoolValue(true),
				}),
				"id": types.StringValue("id1"),
			}),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := planmodifier.ObjectRequest{
				Path:      path.Root("test"),
				PlanValue: test.plannedValue,
			}
			response := planmodifier.ObjectResponse{}
			StaticPartialObject(defaultVal).PlanModifyObject(ctx, request, &response)

			if diff := cmp.Diff(response.PlanValue, test.expectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
