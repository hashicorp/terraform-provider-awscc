// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaults

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// StaticPartialObject return an AttributePlanModifier that sets the specified value if the planned value is Null.
func StaticPartialObject(defaultVal map[string]interface{}) planmodifier.Object {
	return objectDefaultValueAttributePlanModifier{
		defaultVal: defaultVal,
	}
}

type objectDefaultValueAttributePlanModifier struct {
	defaultVal map[string]any
}

func (m objectDefaultValueAttributePlanModifier) Description(context.Context) string {
	return fmt.Sprintf("value defaults to %v", m.defaultVal)
}

func (m objectDefaultValueAttributePlanModifier) MarkdownDescription(context.Context) string {
	return fmt.Sprintf("value defaults to `%v`", m.defaultVal)
}

func (m objectDefaultValueAttributePlanModifier) PlanModifyObject(ctx context.Context, request planmodifier.ObjectRequest, response *planmodifier.ObjectResponse) {
	// NoOp.
	response.PlanValue = request.PlanValue
}
