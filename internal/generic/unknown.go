package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// UnknownValuePaths returns all the paths to all the unknown values in the specified Terraform Value.
func UnknownValuePaths(ctx context.Context, val tftypes.Value) ([]*tftypes.AttributePath, error) {
	return unknownValuePaths(ctx, nil, val)
}

// unknownValuePaths returns all the paths to all the unknown values for the specified Terraform Value.
func unknownValuePaths(ctx context.Context, path *tftypes.AttributePath, val tftypes.Value) ([]*tftypes.AttributePath, error) { //nolint:unparam
	if !val.IsKnown() {
		return []*tftypes.AttributePath{path}, nil
	}

	var unknowns []*tftypes.AttributePath

	typ := val.Type()
	switch {
	case typ.Is(tftypes.List{}), typ.Is(tftypes.Set{}), typ.Is(tftypes.Tuple{}):
		var vals []tftypes.Value
		if err := val.As(&vals); err != nil {
			return nil, path.NewError(err)
		}

		for idx, val := range vals {
			if typ.Is(tftypes.Set{}) {
				path = path.WithElementKeyValue(val)
			} else {
				path = path.WithElementKeyInt(idx)
			}
			paths, err := unknownValuePaths(ctx, path, val)
			if err != nil {
				return nil, err
			}
			unknowns = append(unknowns, paths...)
			path = path.WithoutLastStep()
		}

	case typ.Is(tftypes.Map{}), typ.Is(tftypes.Object{}):
		var vals map[string]tftypes.Value
		if err := val.As(&vals); err != nil {
			return nil, path.NewError(err)
		}

		for key, val := range vals {
			if typ.Is(tftypes.Map{}) {
				path = path.WithElementKeyString(key)
			} else if typ.Is(tftypes.Object{}) {
				path = path.WithAttributeName(key)
			}
			paths, err := unknownValuePaths(ctx, path, val)
			if err != nil {
				return nil, err
			}
			unknowns = append(unknowns, paths...)
			path = path.WithoutLastStep()
		}
	}

	return unknowns, nil
}

// SetUnknownValuesFromResourceModel fills any unknown State values from a Cloud Control ResourceModel (string).
// The unknown value paths are obtained from the State via a previous call to UnknownValuePaths.
// Functionality is split between these 2 functions, rather than calling UnknownValuePaths from within this function,
// so as to avoid unnecessary Cloud Control API calls to obtain the current ResourceModel.
func SetUnknownValuesFromResourceModel(ctx context.Context, state *tfsdk.State, unknowns []*tftypes.AttributePath, resourceModel string, cfToTfNameMap map[string]string) error {
	// Get the Terraform Value of the ResourceModel.
	translator := toTerraform{cfToTfNameMap: cfToTfNameMap}
	schema := &state.Schema
	val, err := translator.FromString(ctx, schema, resourceModel)

	if err != nil {
		return err
	}

	src := tfsdk.State{
		Schema: *schema,
		Raw:    val,
	}

	// Copy all unknown values from the ResourceModel to destination State.
	for _, path := range unknowns {
		err = CopyValueAtPath(ctx, state, &src, path)

		if err != nil {
			return err
		}
	}

	return nil
}
