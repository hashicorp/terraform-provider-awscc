// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaults

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfslices "github.com/hashicorp/terraform-provider-awscc/internal/slices"
)

// StaticSetOfString returns a static set of string value default handler.
func StaticSetOfString(values ...string) defaults.Set {
	return setdefault.StaticValue(types.SetValueMust(types.StringType, tfslices.ApplyToAll(values, func(val string) attr.Value {
		return types.StringValue(val)
	})))
}

// EmptySetNestedObject return an AttributePlanModifier that returns an empty set if the planned value is Null.
func EmptySetNestedObject() planmodifier.Set {
	return emptySetNestedObjectAttributePlanModifier{}
}

type emptySetNestedObjectAttributePlanModifier struct{}

func (emptySetNestedObjectAttributePlanModifier) Description(context.Context) string {
	return "value defaults to empty set"
}

func (m emptySetNestedObjectAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m emptySetNestedObjectAttributePlanModifier) PlanModifySet(ctx context.Context, request planmodifier.SetRequest, response *planmodifier.SetResponse) {
	if request.PlanValue.IsUnknown() {
		response.PlanValue = types.SetValueMust(request.PlanValue.ElementType(ctx), []attr.Value{})
		return
	}

	// NoOp.
	response.PlanValue = request.PlanValue
}
