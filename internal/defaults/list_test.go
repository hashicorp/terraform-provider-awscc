// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaults

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestStaticListOfString(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()
	request := defaults.ListRequest{}
	response := defaults.ListResponse{}
	StaticListOfString("One", "bee").DefaultList(ctx, request, &response)

	if response.Diagnostics.HasError() {
		t.Errorf("unexpected error: %v", response.Diagnostics)
	}
	if want, got := 2, len(response.PlanValue.Elements()); got != want {
		t.Errorf("StaticListOfString.PlanValue length %d, want %d", got, want)
	}
}

func TestEmptyListNestedObject(t *testing.T) {
	t.Parallel()

	attributeTypes := map[string]attr.Type{
		"name": types.StringType,
		"size": types.Int64Type,
	}

	type testCase struct {
		plannedValue  types.List
		expectedValue types.List
	}
	tests := map[string]testCase{
		"unknown list": {
			plannedValue:  types.ListUnknown(types.ObjectType{AttrTypes: attributeTypes}),
			expectedValue: types.ListValueMust(types.ObjectType{AttrTypes: attributeTypes}, []attr.Value{}),
		},
		"null list": {
			plannedValue:  types.ListNull(types.ObjectType{AttrTypes: attributeTypes}),
			expectedValue: types.ListNull(types.ObjectType{AttrTypes: attributeTypes}),
		},
		"empty list": {
			plannedValue:  types.ListValueMust(types.ObjectType{AttrTypes: attributeTypes}, []attr.Value{}),
			expectedValue: types.ListValueMust(types.ObjectType{AttrTypes: attributeTypes}, []attr.Value{}),
		},
		"non-empty list": {
			plannedValue: types.ListValueMust(types.ObjectType{AttrTypes: attributeTypes}, []attr.Value{
				types.ObjectValueMust(attributeTypes, map[string]attr.Value{
					"name": types.StringValue("n1"),
					"size": types.Int64Value(1),
				}),
				types.ObjectValueMust(attributeTypes, map[string]attr.Value{
					"name": types.StringValue("n2"),
					"size": types.Int64Value(2),
				}),
			}),
			expectedValue: types.ListValueMust(types.ObjectType{AttrTypes: attributeTypes}, []attr.Value{
				types.ObjectValueMust(attributeTypes, map[string]attr.Value{
					"name": types.StringValue("n1"),
					"size": types.Int64Value(1),
				}),
				types.ObjectValueMust(attributeTypes, map[string]attr.Value{
					"name": types.StringValue("n2"),
					"size": types.Int64Value(2),
				}),
			}),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := planmodifier.ListRequest{
				Path:      path.Root("test"),
				PlanValue: test.plannedValue,
			}
			response := planmodifier.ListResponse{}
			EmptyListNestedObject().PlanModifyList(ctx, request, &response)

			if diff := cmp.Diff(response.PlanValue, test.expectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
