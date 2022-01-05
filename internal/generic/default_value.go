package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-provider-awscc/internal/diag"
)

type defaultValueAttributePlanModifier struct {
	tfsdk.AttributePlanModifier
	val attr.Value
}

// DefaultValue return an AttributePlanModifier that sets the specified value if the planned value is Null and the current value is the default.
func DefaultValue(val attr.Value) tfsdk.AttributePlanModifier {
	return defaultValueAttributePlanModifier{
		val: val,
	}
}

func (attributePlanModifier defaultValueAttributePlanModifier) Description(_ context.Context) string {
	return "If the value of the attribute is missing, then the value is semantically the same as if the value was present with the default value."
}

func (attributePlanModifier defaultValueAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return attributePlanModifier.Description(ctx)
}

func (attributePlanModifier defaultValueAttributePlanModifier) Modify(ctx context.Context, request tfsdk.ModifyAttributePlanRequest, response *tfsdk.ModifyAttributePlanResponse) {
	// If the planned value is Null and there is a current value and the current value is the default
	// then return the current value, else return the planned value.
	if v, err := request.AttributePlan.ToTerraformValue(ctx); err != nil {
		response.Diagnostics.Append(diag.NewUnableToObtainValueAttributeError(
			request.AttributePath,
			err,
		))

		return
	} else if v.IsNull() && request.AttributeState != nil && request.AttributeState.Equal(attributePlanModifier.val) {
		response.AttributePlan = request.AttributeState
	} else {
		response.AttributePlan = request.AttributePlan
	}
}
