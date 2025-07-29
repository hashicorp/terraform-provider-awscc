// Package generic provides custom plan modifiers to address shadow drift issues.
// This file contains complex object type plan modifiers.
package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// CustomUseStateForUnknownObject returns a plan modifier for object attributes.
func CustomUseStateForUnknownObject() planmodifier.Object {
	return customUseStateForUnknownObjectModifier{}
}

type customUseStateForUnknownObjectModifier struct{}

func (m customUseStateForUnknownObjectModifier) Description(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownObjectModifier) MarkdownDescription(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownObjectModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}
	// Use state value to prevent framework's "unknown" marking
	resp.PlanValue = req.StateValue
}
