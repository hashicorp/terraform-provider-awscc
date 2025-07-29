// Package generic provides custom plan modifiers to address shadow drift issues.
// These replace the framework's UseStateForUnknown to avoid false positive drift detection.
// Reference: https://github.com/hashicorp/terraform-provider-awscc/issues/2726
package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// CustomUseStateForUnknownString returns a plan modifier that prevents shadow drift
// by using the state value when configuration is null.
func CustomUseStateForUnknownString() planmodifier.String {
	return customUseStateForUnknownModifier{}
}

type customUseStateForUnknownModifier struct{}

func (m customUseStateForUnknownModifier) Description(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownModifier) MarkdownDescription(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}
	// Use state value to prevent framework's "unknown" marking
	resp.PlanValue = req.StateValue
}

// CustomUseStateForUnknownBool returns a plan modifier for boolean attributes.
func CustomUseStateForUnknownBool() planmodifier.Bool {
	return customUseStateForUnknownBoolModifier{}
}

type customUseStateForUnknownBoolModifier struct{}

func (m customUseStateForUnknownBoolModifier) Description(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownBoolModifier) MarkdownDescription(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownBoolModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}
	// Use state value to prevent framework's "unknown" marking
	resp.PlanValue = req.StateValue
}
