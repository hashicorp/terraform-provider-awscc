package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// ComputedOptionalForceNew is used as a replacement for tfsdk.RequiresReplace for Computed,Optional,ForceNew attributes.
func ComputedOptionalForceNew() tfsdk.AttributePlanModifier {
	return tfsdk.RequiresReplaceIf(func(ctx context.Context, state, config attr.Value, path *tftypes.AttributePath) (bool, diag.Diagnostics) {
		if config.Equal(state) {
			return false, nil
		}

		// TODO: This can be replaced with a single config.IsNull() conditional when available
		// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/193
		switch config.(type) {
		case types.Bool:
			return !config.Equal(types.Bool{Null: true}), nil
		case types.Float64:
			return !config.Equal(types.Float64{Null: true}), nil
		case types.Int64:
			return !config.Equal(types.Int64{Null: true}), nil
		case types.List:
			return !config.Equal(types.List{Null: true}), nil
		case types.Map:
			return !config.Equal(types.Map{Null: true}), nil
		case types.Object:
			return !config.Equal(types.Object{Null: true}), nil
		case types.Number:
			return !config.Equal(types.Number{Null: true}), nil
		case types.Set:
			return !config.Equal(types.Set{Null: true}), nil
		case types.String:
			return !config.Equal(types.String{Null: true}), nil
		}

		return true, nil
	}, "require replacement if configuration value changes", "require replacement if configuration value changes")
}
