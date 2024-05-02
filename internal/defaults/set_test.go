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

func TestStaticSetOfString(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()
	request := defaults.SetRequest{}
	response := defaults.SetResponse{}
	StaticSetOfString("One", "bee").DefaultSet(ctx, request, &response)

	if response.Diagnostics.HasError() {
		t.Errorf("unexpected error: %v", response.Diagnostics)
	}
	if want, got := 2, len(response.PlanValue.Elements()); got != want {
		t.Errorf("StaticListOfString.PlanValue length %d, want %d", got, want)
	}
}

func TestEmptySetNestedObject(t *testing.T) {
	t.Parallel()

	attributeTypes := map[string]attr.Type{
		"name": types.StringType,
		"size": types.Int64Type,
	}

	type testCase struct {
		plannedValue  types.Set
		expectedValue types.Set
	}
	tests := map[string]testCase{
		"unknown set": {
			plannedValue:  types.SetUnknown(types.ObjectType{AttrTypes: attributeTypes}),
			expectedValue: types.SetValueMust(types.ObjectType{AttrTypes: attributeTypes}, []attr.Value{}),
		},
		"null set": {
			plannedValue:  types.SetNull(types.ObjectType{AttrTypes: attributeTypes}),
			expectedValue: types.SetNull(types.ObjectType{AttrTypes: attributeTypes}),
		},
		"empty set": {
			plannedValue:  types.SetValueMust(types.ObjectType{AttrTypes: attributeTypes}, []attr.Value{}),
			expectedValue: types.SetValueMust(types.ObjectType{AttrTypes: attributeTypes}, []attr.Value{}),
		},
		"non-empty set": {
			plannedValue: types.SetValueMust(types.ObjectType{AttrTypes: attributeTypes}, []attr.Value{
				types.ObjectValueMust(attributeTypes, map[string]attr.Value{
					"name": types.StringValue("n1"),
					"size": types.Int64Value(1),
				}),
				types.ObjectValueMust(attributeTypes, map[string]attr.Value{
					"name": types.StringValue("n2"),
					"size": types.Int64Value(2),
				}),
			}),
			expectedValue: types.SetValueMust(types.ObjectType{AttrTypes: attributeTypes}, []attr.Value{
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
			request := planmodifier.SetRequest{
				Path:      path.Root("test"),
				PlanValue: test.plannedValue,
			}
			response := planmodifier.SetResponse{}
			EmptySetNestedObject().PlanModifySet(ctx, request, &response)

			if diff := cmp.Diff(response.PlanValue, test.expectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
