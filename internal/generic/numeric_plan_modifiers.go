// Package generic provides custom plan modifiers to address shadow drift issues.
// This file contains numeric type plan modifiers.
package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// CustomUseStateForUnknownInt64 returns a plan modifier for int64 attributes.
func CustomUseStateForUnknownInt64() planmodifier.Int64 {
	return customUseStateForUnknownInt64Modifier{}
}

type customUseStateForUnknownInt64Modifier struct{}

func (m customUseStateForUnknownInt64Modifier) Description(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownInt64Modifier) MarkdownDescription(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownInt64Modifier) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if !req.ConfigValue.IsNull() {
		return
	}
	// Use state value to prevent framework's "unknown" marking
	resp.PlanValue = req.StateValue
}

// CustomUseStateForUnknownFloat64 returns a plan modifier for float64 attributes.
func CustomUseStateForUnknownFloat64() planmodifier.Float64 {
	return customUseStateForUnknownFloat64Modifier{}
}

type customUseStateForUnknownFloat64Modifier struct{}

func (m customUseStateForUnknownFloat64Modifier) Description(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownFloat64Modifier) MarkdownDescription(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownFloat64Modifier) PlanModifyFloat64(ctx context.Context, req planmodifier.Float64Request, resp *planmodifier.Float64Response) {
	if !req.ConfigValue.IsNull() {
		return
	}
	// Use state value to prevent framework's "unknown" marking
	resp.PlanValue = req.StateValue
}

// CustomUseStateForUnknownNumber returns a plan modifier for number attributes.
func CustomUseStateForUnknownNumber() planmodifier.Number {
	return customUseStateForUnknownNumberModifier{}
}

type customUseStateForUnknownNumberModifier struct{}

func (m customUseStateForUnknownNumberModifier) Description(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownNumberModifier) MarkdownDescription(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownNumberModifier) PlanModifyNumber(ctx context.Context, req planmodifier.NumberRequest, resp *planmodifier.NumberResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}
	// Use state value to prevent framework's "unknown" marking
	resp.PlanValue = req.StateValue
}
