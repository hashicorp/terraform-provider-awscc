// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// EmptyObject returns "empty" attributes (all values set to null) for the given attribute types.
func EmptyAttributes(ctx context.Context, attributeTypes map[string]attr.Type) (map[string]attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	attributes := map[string]attr.Value{}

	for name, attrType := range attributeTypes {
		attrValue, err := NullValueForType(ctx, attrType)

		if err != nil {
			diags.AddError("NullValueForType", err.Error())

			return nil, diags
		}

		attributes[name] = attrValue
	}

	return attributes, diags
}

// EmptyObject returns an "empty" object (one with all attributes set to null) for the given attribute types.
func EmptyObject(ctx context.Context, attributeTypes map[string]attr.Type) (types.Object, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes, d := EmptyAttributes(ctx, attributeTypes)
	diags.Append(d...)
	if diags.HasError() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objectValue, d := types.ObjectValue(attributeTypes, attributes)
	diags.Append(d...)

	return objectValue, diags
}

// NullValueForType returns a null attr.Value for the specified attr.Type.
func NullValueForType(ctx context.Context, attrType attr.Type) (attr.Value, error) {
	return attrType.ValueFromTerraform(ctx, tftypes.NewValue(attrType.TerraformType(ctx), nil))
}
