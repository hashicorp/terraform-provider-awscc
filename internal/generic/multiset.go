// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

type multisetAttributePlanModifier struct{}

// A multiset is an array allowing non-unique items with insertion order not significant.
// Multisets do not correspond directly with either Terraform Lists (insertion order is significant) or Sets (unique items).
// Multiset Attributes are declared as Lists with a plan modifier that suppresses semantically insignificant differences.
func Multiset() planmodifier.List {
	return multisetAttributePlanModifier{}
}

func (attributePlanModifier multisetAttributePlanModifier) Description(_ context.Context) string {
	return "Suppresses semantically insignificant differences."
}

func (attributePlanModifier multisetAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return attributePlanModifier.Description(ctx)
}

func (attributePlanModifier multisetAttributePlanModifier) PlanModifyList(ctx context.Context, request planmodifier.ListRequest, response *planmodifier.ListResponse) {
	if request.StateValue.IsNull() {
		response.PlanValue = request.PlanValue

		return
	}

	// If the current value is semantically equivalent to the planned value
	// then return the current value, else return the planned value.

	planned, diags := request.PlanValue.ToListValue(ctx)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	current, diags := request.StateValue.ToListValue(ctx)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	if len(planned.Elements()) != len(current.Elements()) {
		response.PlanValue = request.PlanValue

		return
	}

	currentVals := make([]attr.Value, len(current.Elements()))
	copy(currentVals, current.Elements())

	for _, plannedVal := range planned.Elements() {
		for i, currentVal := range currentVals {
			if currentVal.Equal(plannedVal) {
				// Remove from the slice.
				currentVals = append(currentVals[:i], currentVals[i+1:]...)

				break
			}
		}
	}

	if len(currentVals) == 0 {
		// Every planned value is equal to a current value.
		response.PlanValue = request.StateValue
	} else {
		response.PlanValue = request.PlanValue
	}
}
