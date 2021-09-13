package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

type defaultValueAttributePlanModifier struct {
	tfsdk.AttributePlanModifier
	val attr.Value
}

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
	// If the planned value is Null and the current value is the default then return the current value, else return the planned value.
	if v, err := request.AttributePlan.ToTerraformValue(ctx); err != nil {
		response.AddAttributeError(
			request.AttributePath,
			"No Terraform value",
			"unable to obtain Terraform value:\n\n"+err.Error(),
		)

		return
	} else if v == nil && request.AttributeState.Equal(attributePlanModifier.val) {
		response.AttributePlan = request.AttributeState
	} else {
		response.AttributePlan = request.AttributePlan
	}
}
