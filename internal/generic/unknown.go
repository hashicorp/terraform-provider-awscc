package generic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

// unknownValuePath represents the path to an unknown (!val.IsKnown()) value.
// It holds paths in both the Terraform State and CloudFormation ResourceModel (raw map[string]interface{}).
type unknownValuePath struct {
	InTerraformState              *tftypes.AttributePath
	InCloudFormationResourceModel *tftypes.AttributePath
}

type unknowns []unknownValuePath

// Unknowns returns all the unknowns in the specified Terraform Value.
func Unknowns(ctx context.Context, val tftypes.Value, tfToCfNameMap map[string]string) (unknowns, error) {
	unknowns, err := unknownValuePaths(ctx, nil, nil, val, tfToCfNameMap)

	if err != nil {
		return nil, err
	}

	return unknowns, nil
}

// unknownValuePaths returns all the unknownValuePaths for the specified Terraform Value.
func unknownValuePaths(ctx context.Context, inTerraformState, inCloudFormationResourceModel *tftypes.AttributePath, val tftypes.Value, tfToCfNameMap map[string]string) ([]unknownValuePath, error) { //nolint:unparam
	if !val.IsKnown() {
		return []unknownValuePath{
			{
				InTerraformState:              inTerraformState,
				InCloudFormationResourceModel: inCloudFormationResourceModel,
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
				inCloudFormationResourceModel = inCloudFormationResourceModel.WithElementKeyValue(val)
			} else {
				inTerraformState = inTerraformState.WithElementKeyInt(int64(idx))
				inCloudFormationResourceModel = inCloudFormationResourceModel.WithElementKeyInt(int64(idx))
			}
			paths, err := unknownValuePaths(ctx, inTerraformState, inCloudFormationResourceModel, val, tfToCfNameMap)
			if err != nil {
				return nil, err
			}
			unknowns = append(unknowns, paths...)
			inTerraformState = inTerraformState.WithoutLastStep()
			inCloudFormationResourceModel = inCloudFormationResourceModel.WithoutLastStep()
		}

	case typ.Is(tftypes.Map{}), typ.Is(tftypes.Object{}):
		var vals map[string]tftypes.Value
		if err := val.As(&vals); err != nil {
			return nil, inTerraformState.NewError(err)
		}

		for key, val := range vals {
			if typ.Is(tftypes.Map{}) {
				inTerraformState = inTerraformState.WithElementKeyString(key)
				inCloudFormationResourceModel = inCloudFormationResourceModel.WithElementKeyString(key)
			} else if typ.Is(tftypes.Object{}) {
				inTerraformState = inTerraformState.WithAttributeName(key)
				propertyName, ok := tfToCfNameMap[key]
				if !ok {
					return nil, fmt.Errorf("attribute name mapping not found: %s", key)
				}
				inCloudFormationResourceModel = inCloudFormationResourceModel.WithAttributeName(propertyName)
			}
			paths, err := unknownValuePaths(ctx, inTerraformState, inCloudFormationResourceModel, val, tfToCfNameMap)
			if err != nil {
				return nil, err
			}
			unknowns = append(unknowns, paths...)
			inTerraformState = inTerraformState.WithoutLastStep()
			inCloudFormationResourceModel = inCloudFormationResourceModel.WithoutLastStep()
		}
	}

	return unknowns, nil
}

// SetValuesFromRaw fills any unknown State values from a CloudFormation ResourceModel (raw map[string]interface{}).
func (u unknowns) SetValuesFromRaw(ctx context.Context, state *tfsdk.State, resourceModel map[string]interface{}) error {
	for _, path := range u {
		// Get the value from the CloudFormation ResourceModel.
		val, _, err := tftypes.WalkAttributePath(resourceModel, path.InCloudFormationResourceModel)

		if errors.Is(err, tftypes.ErrInvalidStep) {
			// Value not found in CloudFormation ResourceModel. Set to Nil in State.

			// TODO
			// TODO State.SetAttribute does not support passing `nil` to set a Null value.
			// TODO https://github.com/hashicorp/terraform-plugin-framework/issues/66.
			// TODO

			attrType, err := state.Schema.AttributeTypeAtPath(path.InTerraformState)

			if err != nil {
				return fmt.Errorf("error getting attribute type at %s: %w", path.InTerraformState, err)
			}

			state.Raw, err = tftypes.Transform(state.Raw, func(p *tftypes.AttributePath, v tftypes.Value) (tftypes.Value, error) {
				if p.Equal(path.InTerraformState) {
					return tftypes.NewValue(attrType.TerraformType(ctx), nil), nil
				}
				return v, nil
			})

			if err != nil {
				return fmt.Errorf("error setting attribute in state: %w", err)
			}

			continue
		}

		if err != nil {
			return fmt.Errorf("error getting value at %s: %w", path.InCloudFormationResourceModel, err)
		}

		// Set it in the Terraform State.
		diags := state.SetAttribute(ctx, path.InTerraformState, val)

		if tfresource.DiagsHasError(diags) {
			return fmt.Errorf("error setting value at %s: %w", path.InTerraformState, tfresource.DiagsError(diags))
		}
	}

	return nil
}

// SetValuesFromRaw fills any unknown State values from a CloudFormation ResourceModel (string).
func (u unknowns) SetValuesFromString(ctx context.Context, state *tfsdk.State, resourceModel string) error {
	var v interface{}

	if err := json.Unmarshal([]byte(resourceModel), &v); err != nil {
		return err
	}

	if v, ok := v.(map[string]interface{}); ok {
		return u.SetValuesFromRaw(ctx, state, v)
	}

	return fmt.Errorf("unexpected raw type: %T", v)
}
