// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type defaultValueAttributePlanModifier struct{}

func (defaultValueAttributePlanModifier) Description(context.Context) string {
	return "If the value is omitted, a matching service default already stored in state is preserved."
}

func (m defaultValueAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

type stringDefaultValueAttributePlanModifier struct {
	defaultValueAttributePlanModifier
	value types.String
}

// StringDefaultValue preserves a matching service default from state without
// adding the default to a create plan.
func StringDefaultValue(value string) planmodifier.String {
	return stringDefaultValueAttributePlanModifier{value: types.StringValue(value)}
}

func (m stringDefaultValueAttributePlanModifier) PlanModifyString(_ context.Context, request planmodifier.StringRequest, response *planmodifier.StringResponse) {
	if request.PlanValue.IsNull() && !request.StateValue.IsNull() && request.StateValue.Equal(m.value) {
		response.PlanValue = request.StateValue
		return
	}
	response.PlanValue = request.PlanValue
}
