package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func NotNullObjectNestedObject() planmodifier.Object {
	return notNullObjectNestedAttributeValidator{}
}

type notNullObjectNestedAttributeValidator struct{}

func (notNullObjectNestedAttributeValidator) Description(context.Context) string {
	return "value defaults to state value, if set"
}

func (m notNullObjectNestedAttributeValidator) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m notNullObjectNestedAttributeValidator) PlanModifyObject(ctx context.Context, request planmodifier.ObjectRequest, response *planmodifier.ObjectResponse) {
	if (request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown()) && !request.StateValue.IsNull() {
		response.PlanValue = request.StateValue
		return
	}

	// NoOp.
	response.PlanValue = request.PlanValue
}
