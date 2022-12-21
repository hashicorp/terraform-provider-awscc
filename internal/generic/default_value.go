package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type defaultValueAttributePlanModifier struct{}

func (attributePlanModifier defaultValueAttributePlanModifier) Description(_ context.Context) string {
	return "If the value of the attribute is missing, then the value is semantically the same as if the value was present with the default value."
}

func (attributePlanModifier defaultValueAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return attributePlanModifier.Description(ctx)
}

type stringDefaultValueAttributePlanModifier struct {
	defaultValueAttributePlanModifier
	val types.String
}

// StringDefaultValue return an AttributePlanModifier that sets the specified value if the planned value is Null and the current value is the default.
func StringDefaultValue(val types.String) planmodifier.String {
	return stringDefaultValueAttributePlanModifier{
		val: val,
	}
}

func (attributePlanModifier stringDefaultValueAttributePlanModifier) PlanModifyString(ctx context.Context, request planmodifier.StringRequest, response *planmodifier.StringResponse) {
	// If the planned value is Null and there is a current value and the current value is the default
	// then return the current value, else return the planned value.
	if request.PlanValue.IsNull() && !request.StateValue.IsNull() && request.StateValue.Equal(attributePlanModifier.val) {
		response.PlanValue = request.StateValue
	} else {
		response.PlanValue = request.PlanValue
	}
}
