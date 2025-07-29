// Package generic provides custom plan modifiers to address shadow drift issues.
// This file contains collection type plan modifiers.
package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// CustomUseStateForUnknownList returns a plan modifier for list attributes.
func CustomUseStateForUnknownList() planmodifier.List {
	return customUseStateForUnknownListModifier{}
}

type customUseStateForUnknownListModifier struct{}

func (m customUseStateForUnknownListModifier) Description(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownListModifier) MarkdownDescription(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownListModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}
	// Use state value to prevent framework's "unknown" marking
	resp.PlanValue = req.StateValue
}

// CustomUseStateForUnknownSet returns a plan modifier for set attributes.
func CustomUseStateForUnknownSet() planmodifier.Set {
	return customUseStateForUnknownSetModifier{}
}

type customUseStateForUnknownSetModifier struct{}

func (m customUseStateForUnknownSetModifier) Description(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownSetModifier) MarkdownDescription(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownSetModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}
	// Use state value to prevent framework's "unknown" marking
	resp.PlanValue = req.StateValue
}

// CustomUseStateForUnknownMap returns a plan modifier for map attributes.
func CustomUseStateForUnknownMap() planmodifier.Map {
	return customUseStateForUnknownMapModifier{}
}

type customUseStateForUnknownMapModifier struct{}

func (m customUseStateForUnknownMapModifier) Description(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownMapModifier) MarkdownDescription(ctx context.Context) string {
	return "If configuration is null, use the state value to avoid shadow drift."
}

func (m customUseStateForUnknownMapModifier) PlanModifyMap(ctx context.Context, req planmodifier.MapRequest, resp *planmodifier.MapResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}
	// Use state value to prevent framework's "unknown" marking
	resp.PlanValue = req.StateValue
}
