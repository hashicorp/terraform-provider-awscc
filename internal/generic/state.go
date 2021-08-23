package generic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	tflog "github.com/hashicorp/terraform-plugin-log"
	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

// CopyValueAtPath copies the value at a specified path from source State to destination State.
func CopyValueAtPath(ctx context.Context, dst, src *tfsdk.State, path *tftypes.AttributePath) error {
	val, diags := src.GetAttribute(ctx, path)

	if tfresource.DiagsHasError(diags) {
		return tfresource.DiagsError(diags)
	}

	diags = dst.SetAttribute(ctx, path, val)

	if tfresource.DiagsHasError(diags) {
		return tfresource.DiagsError(diags)
	}

	return nil
}

// SetUnknownValuesFromCloudFormationResourceModel fills any unknown State values from a CloudFormation ResourceModel (string).
func SetUnknownValuesFromCloudFormationResourceModel(ctx context.Context, state *tfsdk.State, resourceModel string) error {
	var v interface{}

	if err := json.Unmarshal([]byte(resourceModel), &v); err != nil {
		return err
	}

	if v, ok := v.(map[string]interface{}); ok {
		return SetUnknownValuesFromCloudFormationResourceModelRaw(ctx, state, v)
	}

	return fmt.Errorf("CloudFormation ResourceModel value produced unexpected raw type: %T", v)
}

// SetUnknownValuesFromCloudFormationResourceModelRaw fills any unknown State values from a CloudFormation ResourceModel (raw map[string]interface{}).
func SetUnknownValuesFromCloudFormationResourceModelRaw(ctx context.Context, state *tfsdk.State, resourceModel map[string]interface{}) error {
	// Get the paths to the state's unknown values.
	paths, err := GetUnknownValuePaths(ctx, state.Raw)

	if err != nil {
		return fmt.Errorf("error getting unknown values: %w", err)
	}

	for _, path := range paths {
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

// GetCloudFormationResourceModelValue returns the Terraform Value for the specified CloudFormation ResourceModel (string).
func GetCloudFormationResourceModelValue(ctx context.Context, schema *tfsdk.Schema, resourceModel string) (tftypes.Value, error) {
	var v interface{}

	if err := json.Unmarshal([]byte(resourceModel), &v); err != nil {
		return tftypes.Value{}, err
	}

	if v, ok := v.(map[string]interface{}); ok {
		return GetCloudFormationResourceModelRawValue(ctx, schema, v)
	}

	return tftypes.Value{}, fmt.Errorf("CloudFormation ResourceModel value produced unexpected raw type: %T", v)
}

// GetCloudFormationResourceModelRawValue returns the Terraform Value for the specified CloudFormation ResourceModel (raw map[string]interface{}).
func GetCloudFormationResourceModelRawValue(ctx context.Context, schema *tfsdk.Schema, resourceModel map[string]interface{}) (tftypes.Value, error) {
	return getCloudFormationResourceModelValue(ctx, schema, nil, resourceModel)
}

func getCloudFormationResourceModelValue(ctx context.Context, schema *tfsdk.Schema, path *tftypes.AttributePath, v interface{}) (tftypes.Value, error) {
	var typ tftypes.Type

	if len(path.Steps()) == 0 {
		typ = schema.AttributeType().TerraformType(ctx)
	} else {
		attrType, err := schema.AttributeTypeAtPath(path)

		if err != nil {
			return tftypes.Value{}, fmt.Errorf("error getting attribute type at %s: %w", path, err)
		}

		typ = attrType.TerraformType(ctx)
	}

	switch v := v.(type) {
	//
	// Primitive types.
	//
	case bool:
		return tftypes.NewValue(tftypes.Bool, v), nil

	case float64:
		return tftypes.NewValue(tftypes.Number, v), nil

	case string:
		return tftypes.NewValue(tftypes.String, v), nil

	//
	// Complex types.
	//
	case []interface{}:
		var vals []tftypes.Value
		for idx, v := range v {
			if typ.Is(tftypes.Set{}) {
				// No need to worry about a specific value here.
				path = path.WithElementKeyValue(tftypes.NewValue(typ.(tftypes.Set).ElementType, nil))
			} else {
				path = path.WithElementKeyInt(int64(idx))
			}
			val, err := getCloudFormationResourceModelValue(ctx, schema, path, v)
			if err != nil {
				return tftypes.Value{}, err
			}
			vals = append(vals, val)
			path = path.WithoutLastStep()
		}
		return tftypes.NewValue(typ, vals), nil

	case map[string]interface{}:
		if typ.Is(tftypes.String) {
			// Value is JSON string.
			val, err := json.Marshal(v)

			if err != nil {
				return tftypes.Value{}, err
			}

			return tftypes.NewValue(typ, string(val)), nil
		}

		isObject := typ.Is(tftypes.Object{})
		vals := make(map[string]tftypes.Value)
		for key, v := range v {
			if isObject {
				// In the Terraform Value attribute names are snake cased.
				path = path.WithAttributeName(naming.CloudFormationPropertyToTerraformAttribute(key))
			} else {
				path = path.WithElementKeyString(key)
			}
			val, err := getCloudFormationResourceModelValue(ctx, schema, path, v)
			if err != nil {
				if isObject {
					tflog.Info(ctx, "not found in Terraform schema", "key", key, "path", path, "error", err.Error())
					path = path.WithoutLastStep()
					continue
				}
				return tftypes.Value{}, err
			}
			if isObject {
				// In the Terraform Value attribute names are snake cased.
				vals[naming.CloudFormationPropertyToTerraformAttribute(key)] = val
			} else {
				vals[key] = val
			}
			path = path.WithoutLastStep()
		}
		if isObject {
			// Set any missing attributes to Null.
			for k, t := range typ.(tftypes.Object).AttributeTypes {
				if _, ok := vals[k]; !ok {
					vals[k] = tftypes.NewValue(t, nil)
				}
			}
		}
		return tftypes.NewValue(typ, vals), nil

	default:
		return tftypes.Value{}, fmt.Errorf("unsupported raw type: %T", v)
	}
}

// UnknownValuePath represents the path to an unknown (!val.IsKnown()) value.
// It holds paths in both the Terraform State and CloudFormation ResourceModel (raw map[string]interface{}).
type UnknownValuePath struct {
	InTerraformState              *tftypes.AttributePath
	InCloudFormationResourceModel *tftypes.AttributePath
}

// GetUnknownValuePaths returns all the UnknownValuePaths for the specified value.
func GetUnknownValuePaths(_ context.Context, val tftypes.Value) ([]UnknownValuePath, error) {
	return getAttributePathsForUnknownValues(nil, nil, val)
}

func getAttributePathsForUnknownValues(inTerraformState, inCloudFormationResourceModel *tftypes.AttributePath, val tftypes.Value) ([]UnknownValuePath, error) {
	if !val.IsKnown() {
		return []UnknownValuePath{
			{
				InTerraformState:              inTerraformState,
				InCloudFormationResourceModel: inCloudFormationResourceModel,
			},
		}, nil
	}

	unknownValuePaths := make([]UnknownValuePath, 0)

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
			paths, err := getAttributePathsForUnknownValues(inTerraformState, inCloudFormationResourceModel, val)
			if err != nil {
				return nil, err
			}
			unknownValuePaths = append(unknownValuePaths, paths...)
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
				// In the CloudFormation ResourceModel attribute names are camel cased.
				inCloudFormationResourceModel = inCloudFormationResourceModel.WithAttributeName(naming.TerraformAttributeToCloudFormationProperty(key))
			}
			paths, err := getAttributePathsForUnknownValues(inTerraformState, inCloudFormationResourceModel, val)
			if err != nil {
				return nil, err
			}
			unknownValuePaths = append(unknownValuePaths, paths...)
			inTerraformState = inTerraformState.WithoutLastStep()
			inCloudFormationResourceModel = inCloudFormationResourceModel.WithoutLastStep()
		}
	}

	return unknownValuePaths, nil
}
