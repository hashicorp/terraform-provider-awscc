package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// unknownValuePath represents the path to an unknown (!val.IsKnown()) value in Terraform State.
type unknownValuePath struct {
	InTerraformState *tftypes.AttributePath
}

type unknowns []unknownValuePath

// Unknowns returns all the unknowns in the specified Terraform Value.
func Unknowns(ctx context.Context, val tftypes.Value) (unknowns, error) {
	unknowns, err := unknownValuePaths(ctx, nil, val)

	if err != nil {
		return nil, err
	}

	return unknowns, nil
}

// unknownValuePaths returns all the unknownValuePaths for the specified Terraform Value.
func unknownValuePaths(ctx context.Context, inTerraformState *tftypes.AttributePath, val tftypes.Value) ([]unknownValuePath, error) { //nolint:unparam
	if !val.IsKnown() {
		return []unknownValuePath{
			{
				InTerraformState: inTerraformState,
			},
		}, nil
	}

	unknowns := make([]unknownValuePath, 0)

	typ := val.Type()
	switch {
	case typ.Is(tftypes.List{}), typ.Is(tftypes.Set{}), typ.Is(tftypes.Tuple{}):
		var vals []tftypes.Value
		if err := val.As(&vals); err != nil {
			return nil, inTerraformState.NewError(err)
		}

		for idx, val := range vals {
			if typ.Is(tftypes.Set{}) {
				inTerraformState = inTerraformState.WithElementKeyValue(val)
			} else {
				inTerraformState = inTerraformState.WithElementKeyInt(idx)
			}
			paths, err := unknownValuePaths(ctx, inTerraformState, val)
			if err != nil {
				return nil, err
			}
			unknowns = append(unknowns, paths...)
			inTerraformState = inTerraformState.WithoutLastStep()
		}

	case typ.Is(tftypes.Map{}), typ.Is(tftypes.Object{}):
		var vals map[string]tftypes.Value
		if err := val.As(&vals); err != nil {
			return nil, inTerraformState.NewError(err)
		}

		for key, val := range vals {
			if typ.Is(tftypes.Map{}) {
				inTerraformState = inTerraformState.WithElementKeyString(key)
			} else if typ.Is(tftypes.Object{}) {
				inTerraformState = inTerraformState.WithAttributeName(key)
			}
			paths, err := unknownValuePaths(ctx, inTerraformState, val)
			if err != nil {
				return nil, err
			}
			unknowns = append(unknowns, paths...)
			inTerraformState = inTerraformState.WithoutLastStep()
		}
	}

	return unknowns, nil
}

// SetValuesFromRaw fills any unknown State values from a Cloud Control ResourceModel (raw map[string]interface{}).
// func (u unknowns) SetValuesFromRaw(ctx context.Context, state *tfsdk.State, resourceModel map[string]interface{}) error {
// 	for _, path := range u {
// 		// Get the value from the Cloud Control ResourceModel.
// 		val, _, err := tftypes.WalkAttributePath(resourceModel, path.InCloudControlResourceModel)

// 		if errors.Is(err, tftypes.ErrInvalidStep) {
// 			// Value not found in Cloud Control ResourceModel. Set to Nil in State.

// 			// TODO
// 			// TODO State.SetAttribute does not support passing `nil` to set a Null value.
// 			// TODO https://github.com/hashicorp/terraform-plugin-framework/issues/66.
// 			// TODO

// 			attrType, err := state.Schema.AttributeTypeAtPath(path.InTerraformState)

// 			if err != nil {
// 				return fmt.Errorf("error getting attribute type at %s: %w", path.InTerraformState, err)
// 			}

// 			state.Raw, err = tftypes.Transform(state.Raw, func(p *tftypes.AttributePath, v tftypes.Value) (tftypes.Value, error) {
// 				if p.Equal(path.InTerraformState) {
// 					return tftypes.NewValue(attrType.TerraformType(ctx), nil), nil
// 				}
// 				return v, nil
// 			})

// 			if err != nil {
// 				return fmt.Errorf("error setting attribute in state: %w", err)
// 			}

// 			continue
// 		}

// 		if err != nil {
// 			return fmt.Errorf("error getting value at %s: %w", path.InCloudControlResourceModel, err)
// 		}

// 		// Set it in the Terraform State.
// 		diags := state.SetAttribute(ctx, path.InTerraformState, val)

// 		if diags.HasError() {
// 			return fmt.Errorf("error setting value at %s: %w", path.InTerraformState, tfresource.DiagsError(diags))
// 		}
// 	}

// 	return nil
// }

// SetValuesFromRaw fills any unknown State values from a Cloud Control ResourceModel (string).
func (u unknowns) SetValuesFromString(ctx context.Context, state *tfsdk.State, resourceModel string, cfToTfNameMap map[string]string) error {
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

	for _, path := range u {
		err = CopyValueAtPath(ctx, state, &src, path.InTerraformState)

		if err != nil {
			return err
		}
	}

	return nil

	// var v interface{}

	// if err := json.Unmarshal([]byte(resourceModel), &v); err != nil {
	// 	return err
	// }

	// if v, ok := v.(map[string]interface{}); ok {
	// 	return u.SetValuesFromRaw(ctx, state, v)
	// }

	// return fmt.Errorf("unexpected raw type: %T", v)
}
