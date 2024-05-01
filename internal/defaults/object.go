// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaults

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
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
	if request.PlanValue.IsNull() {
		// Create "empty" attributes and then copy over from the default value.
		attributeTypes := request.PlanValue.AttributeTypes(ctx)
		attributes, diags := cctypes.EmptyAttributes(ctx, attributeTypes)
		response.Diagnostics.Append(diags...)
		if response.Diagnostics.HasError() {
			return
		}

		response.Diagnostics.Append(copyAttributeValues(ctx, attributes, m.defaultVal)...)
		if response.Diagnostics.HasError() {
			return
		}

		objectValue, diags := types.ObjectValue(attributeTypes, attributes)
		response.Diagnostics.Append(diags...)
		if response.Diagnostics.HasError() {
			return
		}

		response.PlanValue = objectValue

		return
	}

	// NoOp.
	response.PlanValue = request.PlanValue
}

func copyAttributeValues(ctx context.Context, dst map[string]attr.Value, src map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	for k, v := range src {
		if old, ok := dst[k]; ok {
			switch v := v.(type) {
			case bool:
				dst[k] = types.BoolValue(v)
			case string:
				dst[k] = types.StringValue(v)
			case map[string]interface{}:
				if old, ok := old.(types.Object); ok {
					attributeTypes := old.AttributeTypes(ctx)
					attributes := map[string]attr.Value{}

					// If old is null, create an empty.
					if old.IsNull() {
						empty, d := cctypes.EmptyAttributes(ctx, attributeTypes)
						diags.Append(d...)
						if diags.HasError() {
							return diags
						}

						attributes = empty
					}

					diags.Append(copyAttributeValues(ctx, attributes, v)...)
					if diags.HasError() {
						return diags
					}

					objectValue, d := types.ObjectValue(attributeTypes, attributes)
					diags.Append(d...)
					if diags.HasError() {
						return diags
					}

					dst[k] = objectValue
				}
			}
		}
	}

	return diags
}
