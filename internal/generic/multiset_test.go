// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/errs/diag"
)

func TestMultiset(t *testing.T) {
	t.Parallel()

	tagAttrTypes := map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	}
	tagElemType := types.ObjectType{AttrTypes: tagAttrTypes}

	type testCase struct {
		plannedValue  basetypes.ListValue
		currentValue  basetypes.ListValue
		expectedValue basetypes.ListValue
		expectError   bool
	}
	tests := map[string]testCase{
		"null lists": {
			plannedValue:  types.ListNull(types.StringType),
			currentValue:  types.ListNull(types.StringType),
			expectedValue: types.ListNull(types.StringType),
		},
		"single item": {
			plannedValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
			}),
			currentValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("beta"),
			}),
			expectedValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
			}),
		},
		"different lengths": {
			plannedValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
				types.StringValue("gamma"),
			}),
			currentValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
			}),
			expectedValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
				types.StringValue("gamma"),
			}),
		},
		"equivalent": {
			plannedValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
				types.StringValue("gamma"),
				types.StringValue("beta"),
			}),
			currentValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
				types.StringValue("beta"),
				types.StringValue("gamma"),
			}),
			expectedValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
				types.StringValue("beta"),
				types.StringValue("gamma"),
			}),
		},
		"list of objects": {
			plannedValue: types.ListValueMust(tagElemType, []attr.Value{
				types.ObjectValueMust(tagAttrTypes, map[string]attr.Value{
					"key":   types.StringValue("k2"),
					"value": types.StringValue("v2"),
				}),
				types.ObjectValueMust(tagAttrTypes, map[string]attr.Value{
					"key":   types.StringValue("k1"),
					"value": types.StringValue("v1"),
				}),
			}),
			currentValue: types.ListValueMust(tagElemType, []attr.Value{
				types.ObjectValueMust(tagAttrTypes, map[string]attr.Value{
					"key":   types.StringValue("k1"),
					"value": types.StringValue("v1"),
				}),
				types.ObjectValueMust(tagAttrTypes, map[string]attr.Value{
					"key":   types.StringValue("k2"),
					"value": types.StringValue("v2"),
				}),
			}),
			expectedValue: types.ListValueMust(tagElemType, []attr.Value{
				types.ObjectValueMust(tagAttrTypes, map[string]attr.Value{
					"key":   types.StringValue("k1"),
					"value": types.StringValue("v1"),
				}),
				types.ObjectValueMust(tagAttrTypes, map[string]attr.Value{
					"key":   types.StringValue("k2"),
					"value": types.StringValue("v2"),
				}),
			}),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := planmodifier.ListRequest{
				Path:       path.Root("test"),
				PlanValue:  test.plannedValue,
				StateValue: test.currentValue,
			}
			response := planmodifier.ListResponse{}
			Multiset().PlanModifyList(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", ccdiag.DiagnosticsError(response.Diagnostics))
			}

			if diff := cmp.Diff(response.PlanValue, test.expectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
