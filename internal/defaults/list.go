// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaults

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfslices "github.com/hashicorp/terraform-provider-awscc/internal/slices"
)

// StaticListOfString returns a static list of string value default handler.
func StaticListOfString(values ...string) defaults.List {
	return listdefault.StaticValue(types.ListValueMust(types.StringType, tfslices.ApplyToAll(values, func(val string) attr.Value {
		return types.StringValue(val)
	})))
}

// EmptyListNestedObject return an AttributePlanModifier that returns an empty list if the planned value is Null.
func EmptyListNestedObject() planmodifier.List {
	return emptyListNestedObjectAttributePlanModifier{}
}

type emptyListNestedObjectAttributePlanModifier struct{}

func (emptyListNestedObjectAttributePlanModifier) Description(context.Context) string {
	return "value defaults to empty list"
}

func (m emptyListNestedObjectAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m emptyListNestedObjectAttributePlanModifier) PlanModifyList(ctx context.Context, request planmodifier.ListRequest, response *planmodifier.ListResponse) {
	if request.PlanValue.IsNull() {
		response.PlanValue = types.ListValueMust(request.PlanValue.ElementType(ctx), []attr.Value{})
		return
	}

	// NoOp.
	response.PlanValue = request.PlanValue
}
