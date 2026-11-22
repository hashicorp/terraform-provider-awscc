// Copyright IBM Corp. 2021, 2026
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
)

// TestMultiset_fillsUnknownsFromMatchedState covers resolving Computed values inside unordered
// list elements from the state element with the same identity (#3110).
func TestMultiset_fillsUnknownsFromMatchedState(t *testing.T) {
	t.Parallel()

	patchAttrTypes := map[string]attr.Type{
		"mode":     types.StringType,
		"next_run": types.StringType,
	}
	groupAttrTypes := map[string]attr.Type{
		"instance_group_name":   types.StringType,
		"instance_count":        types.Int64Type,
		"image_release_version": types.StringType,
		"current_count":         types.Int64Type,
		"patch":                 types.ObjectType{AttrTypes: patchAttrTypes},
	}
	groupType := types.ObjectType{AttrTypes: groupAttrTypes}

	unknownString := types.StringUnknown()
	unknownInt := types.Int64Unknown()
	nullPatch := types.ObjectNull(patchAttrTypes)

	patch := func(mode string, nextRun attr.Value) attr.Value {
		return types.ObjectValueMust(patchAttrTypes, map[string]attr.Value{
			"mode":     types.StringValue(mode),
			"next_run": nextRun,
		})
	}
	group := func(name string, count int64, version, currentCount, p attr.Value) attr.Value {
		return types.ObjectValueMust(groupAttrTypes, map[string]attr.Value{
			"instance_group_name":   types.StringValue(name),
			"instance_count":        types.Int64Value(count),
			"image_release_version": version,
			"current_count":         currentCount,
			"patch":                 p,
		})
	}
	list := func(elems ...attr.Value) basetypes.ListValue {
		return types.ListValueMust(groupType, elems)
	}
	v := types.StringValue
	n := types.Int64Value

	state := list(
		group("worker-a", 1, v("1.1.0"), n(1), nullPatch),
		group("worker-b", 3, v("1.2.0"), n(3), patch("auto", v("mon"))),
		group("worker-c", 2, v("1.3.0"), n(2), nullPatch),
	)

	type testCase struct {
		plannedValue  basetypes.ListValue
		configValue   basetypes.ListValue
		currentValue  basetypes.ListValue
		expectedValue basetypes.ListValue
	}
	tests := map[string]testCase{
		"element removed: remaining elements keep their own computed values": {
			plannedValue: list(
				group("worker-b", 3, unknownString, unknownInt, patch("auto", unknownString)),
				group("worker-c", 2, unknownString, unknownInt, nullPatch),
			),
			currentValue: state,
			expectedValue: list(
				group("worker-b", 3, v("1.2.0"), n(3), patch("auto", v("mon"))),
				group("worker-c", 2, v("1.3.0"), n(2), nullPatch),
			),
		},
		"element added: new element stays unknown": {
			plannedValue: list(
				group("worker-a", 1, unknownString, unknownInt, nullPatch),
				group("worker-d", 1, unknownString, unknownInt, nullPatch),
			),
			currentValue: state,
			expectedValue: list(
				group("worker-a", 1, v("1.1.0"), n(1), nullPatch),
				group("worker-d", 1, unknownString, unknownInt, nullPatch),
			),
		},
		"reordered only: semantically equal, state value kept": {
			plannedValue: list(
				group("worker-c", 2, unknownString, unknownInt, nullPatch),
				group("worker-a", 1, unknownString, unknownInt, nullPatch),
				group("worker-b", 3, unknownString, unknownInt, patch("auto", unknownString)),
			),
			currentValue:  state,
			expectedValue: state,
		},
		"unknown in configuration: not filled": {
			plannedValue: list(
				group("worker-b", 3, unknownString, unknownInt, nullPatch),
			),
			configValue: list(
				group("worker-b", 3, unknownString, types.Int64Null(), nullPatch),
			),
			currentValue: state,
			expectedValue: list(
				group("worker-b", 3, unknownString, n(3), nullPatch),
			),
		},
		"no identity (duplicate names): unchanged": {
			plannedValue: list(
				group("worker-a", 1, unknownString, unknownInt, nullPatch),
				group("worker-a", 2, unknownString, unknownInt, nullPatch),
			),
			currentValue: list(
				group("worker-a", 2, v("1.2.0"), n(2), nullPatch),
				group("worker-a", 1, v("1.1.0"), n(1), nullPatch),
			),
			expectedValue: list(
				group("worker-a", 1, unknownString, unknownInt, nullPatch),
				group("worker-a", 2, unknownString, unknownInt, nullPatch),
			),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			configValue := test.configValue
			if configValue.IsNull() {
				configValue = types.ListNull(groupType)
			}
			request := planmodifier.ListRequest{
				Path:        path.Root("test"),
				ConfigValue: configValue,
				PlanValue:   test.plannedValue,
				StateValue:  test.currentValue,
			}
			response := planmodifier.ListResponse{PlanValue: request.PlanValue}

			Multiset().PlanModifyList(ctx, request, &response)

			if response.Diagnostics.HasError() {
				t.Fatalf("unexpected error: %v", response.Diagnostics)
			}
			if diff := cmp.Diff(test.expectedValue, response.PlanValue); diff != "" {
				t.Errorf("unexpected diff (-wanted, +got): %s", diff)
			}
		})
	}
}
