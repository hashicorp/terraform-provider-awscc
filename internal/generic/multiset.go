package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type multisetAttributePlanModifier struct {
	tfsdk.AttributePlanModifier
}

// A multiset is an array allowing non-unique items with insertion order not significant.
// Multisets do not correspond directly with either Terraform Lists (insertion order is significant) or Sets (unique items).
// Multiset Attributes are declared as Lists with a plan modifier that suppresses semantically insignificant differences.
func Multiset() tfsdk.AttributePlanModifier {
	return multisetAttributePlanModifier{}
}

func (attributePlanModifier multisetAttributePlanModifier) Description(_ context.Context) string {
	return "Suppresses semantically insignificant differences."
}

func (attributePlanModifier multisetAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return attributePlanModifier.Description(ctx)
}

func (attributePlanModifier multisetAttributePlanModifier) Modify(ctx context.Context, request tfsdk.ModifyAttributePlanRequest, response *tfsdk.ModifyAttributePlanResponse) {
	if request.AttributeState == nil {
		response.AttributePlan = request.AttributePlan

		return
	}

	// If the current value is semantically equivalent to the planned value
	// then return the current value, else return the planned value.

	var planned types.List
	diags := tfsdk.ValueAs(ctx, request.AttributePlan, &planned)

	if diags.HasError() {
		response.Diagnostics = append(response.Diagnostics, diags...)

		return
	}

	var current types.List
	diags = tfsdk.ValueAs(ctx, request.AttributeState, &current)

	if diags.HasError() {
		response.Diagnostics = append(response.Diagnostics, diags...)

		return
	}

	if len(planned.Elems) != len(current.Elems) {
		response.AttributePlan = request.AttributePlan

		return
	}

	currentVals := make([]attr.Value, len(current.Elems))
	copy(currentVals, current.Elems)

	for _, plannedVal := range planned.Elems {
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
		response.AttributePlan = request.AttributeState
	} else {
		response.AttributePlan = request.AttributePlan
	}
}
