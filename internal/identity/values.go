// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identity

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ValueAsString(ctx context.Context, v attr.Value) string {
	if v.IsNull() || v.IsUnknown() {
		return ""
	}

	switch v.Type(ctx) {
	case types.StringType:
		return v.(types.String).ValueString()
	case types.Float64Type:
		return fmt.Sprintf("%v", v.(types.Float64).ValueFloat64())
	case types.Int64Type:
		return fmt.Sprintf("%d", v.(types.Int64).ValueInt64())
	case types.Int32Type:
		return fmt.Sprintf("%d", v.(types.Int32).ValueInt32())
	case types.NumberType:
		return fmt.Sprintf("%v", v.(types.Number).ValueBigFloat())
	case types.BoolType:
		return fmt.Sprintf("%t", v.(types.Bool).ValueBool())
	default:
		return ""
	}
}
