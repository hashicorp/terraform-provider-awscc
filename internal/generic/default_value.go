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

type boolDefaultValueAttributePlanModifier struct {
	defaultValueAttributePlanModifier
	val types.Bool
}

// BoolDefaultValue return an AttributePlanModifier that sets the specified value if the planned value is Null and the current value is the default.
func BoolDefaultValue(val types.Bool) planmodifier.Bool {
	return boolDefaultValueAttributePlanModifier{
		val: val,
	}
}

func (attributePlanModifier boolDefaultValueAttributePlanModifier) PlanModifyBool(ctx context.Context, request planmodifier.BoolRequest, response *planmodifier.BoolResponse) {
	// If the planned value is Null and there is a current value and the current value is the default
	// then return the current value, else return the planned value.
	if request.PlanValue.IsNull() && !request.StateValue.IsNull() && request.StateValue.Equal(attributePlanModifier.val) {
		response.PlanValue = request.StateValue
	} else {
		response.PlanValue = request.PlanValue
	}
}

type float64DefaultValueAttributePlanModifier struct {
	defaultValueAttributePlanModifier
	val types.Float64
}

// Float64DefaultValue return an AttributePlanModifier that sets the specified value if the planned value is Null and the current value is the default.
func Float64DefaultValue(val types.Float64) planmodifier.Float64 {
	return float64DefaultValueAttributePlanModifier{
		val: val,
	}
}

func (attributePlanModifier float64DefaultValueAttributePlanModifier) PlanModifyFloat64(ctx context.Context, request planmodifier.Float64Request, response *planmodifier.Float64Response) {
	// If the planned value is Null and there is a current value and the current value is the default
	// then return the current value, else return the planned value.
	if request.PlanValue.IsNull() && !request.StateValue.IsNull() && request.StateValue.Equal(attributePlanModifier.val) {
		response.PlanValue = request.StateValue
	} else {
		response.PlanValue = request.PlanValue
	}
}

type int64DefaultValueAttributePlanModifier struct {
	defaultValueAttributePlanModifier
	val types.Int64
}

// Int64DefaultValue return an AttributePlanModifier that sets the specified value if the planned value is Null and the current value is the default.
func Int64DefaultValue(val types.Int64) planmodifier.Int64 {
	return int64DefaultValueAttributePlanModifier{
		val: val,
	}
}

func (attributePlanModifier int64DefaultValueAttributePlanModifier) PlanModifyInt64(ctx context.Context, request planmodifier.Int64Request, response *planmodifier.Int64Response) {
	// If the planned value is Null and there is a current value and the current value is the default
	// then return the current value, else return the planned value.
	if request.PlanValue.IsNull() && !request.StateValue.IsNull() && request.StateValue.Equal(attributePlanModifier.val) {
		response.PlanValue = request.StateValue
	} else {
		response.PlanValue = request.PlanValue
	}
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
