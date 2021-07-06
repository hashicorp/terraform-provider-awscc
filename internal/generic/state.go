package generic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/iancoleman/strcase"
)

var (
	identifierAttributePath = tftypes.NewAttributePath().WithAttributeName("identifier")
)

// GetIdentifier gets the well-known "identifier" attribute from State.
func GetIdentifier(ctx context.Context, state *tfsdk.State) (string, error) {
	val, err := state.GetAttribute(ctx, identifierAttributePath)

	if err != nil {
		return "", err
	}

	if val, ok := val.(types.String); ok {
		return val.Value, nil
	}

	return "", fmt.Errorf("invalid identifier type %T", val)
}

// SetIdentifier sets the well-known "identifier" attribute in State.
func SetIdentifier(ctx context.Context, state *tfsdk.State, id string) error {
	return state.SetAttribute(ctx, identifierAttributePath, id)
}

// SetUnknownValuesFromCloudFormationResourceModel fills any unknown State values from a CloudFormation ResourceModel.
func SetUnknownValuesFromCloudFormationResourceModel(ctx context.Context, state *tfsdk.State, resourceModel string) error {
	var v interface{}

	if err := json.Unmarshal([]byte(resourceModel), &v); err != nil {
		return err
	}

	if v, ok := v.(map[string]interface{}); ok {
		// Get the paths to the state's unknown values.
		paths, err := GetUnknownValuePaths(ctx, state.Raw)

		if err != nil {
			return fmt.Errorf("error getting unknown values: %w", err)
		}

		for _, path := range paths {
			// Get the value from the CloudFormation ResourceModel.
			val, _, err := tftypes.WalkAttributePath(v, path.InCloudFormationResourceModel)

			if err != nil {
				return fmt.Errorf("error getting value at %s: %w", path.InCloudFormationResourceModel, err)
			}

			// Set it in the Terraform State.
			err = state.SetAttribute(ctx, path.InTerraformState, val)

			if err != nil {
				return fmt.Errorf("error setting value at %s: %w", path.InTerraformState, err)
			}
		}

		return nil
	}

	return fmt.Errorf("CloudFormation ResourceModel value produced unexpected raw type: %T", v)
}

// SetCloudFormationResourceModel sets the raw map[string]interface{} representing CloudFormation ResourceModel in State.
func SetCloudFormationResourceModelRaw(ctx context.Context, state *tfsdk.State, v map[string]interface{}) error {
	val, err := valueFromRaw(ctx, v)

	if err != nil {
		return err
	}

	state.Raw = val

	return nil
}

// valueFromRaw returns the Terraform value for the specified raw (from JSON unmarshaling) value.
// Attribute names are converted to snake case (Terraform standard).
func valueFromRaw(ctx context.Context, v interface{}) (tftypes.Value, error) {
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
		for _, v := range v {
			val, err := valueFromRaw(ctx, v)
			if err != nil {
				return tftypes.Value{}, err
			}
			vals = append(vals, val)
		}
		// TODO
		// TODO List vs. Set vs. Tuple???
		// TODO
		if len(vals) == 0 {
			return tftypes.Value{}, fmt.Errorf("unsupported raw empty array")
		}
		return tftypes.NewValue(tftypes.List{ElementType: vals[0].Type()}, vals), nil
	case map[string]interface{}:
		vals := make(map[string]tftypes.Value)
		typs := make(map[string]tftypes.Type)
		for name, v := range v {
			val, err := valueFromRaw(ctx, v)
			if err != nil {
				return tftypes.Value{}, err
			}
			name := strcase.ToSnake(name)
			vals[name] = val
			typs[name] = val.Type()
		}
		// TODO
		// TODO Map vs. Object???
		// TODO
		return tftypes.NewValue(tftypes.Object{AttributeTypes: typs}, vals), nil
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
func GetUnknownValuePaths(ctx context.Context, val tftypes.Value) ([]UnknownValuePath, error) {
	return getAttributePathsForUnknownValues(ctx, nil, nil, val)
}

func getAttributePathsForUnknownValues(ctx context.Context, inTerraformState, inCloudFormationResourceModel *tftypes.AttributePath, val tftypes.Value) ([]UnknownValuePath, error) {
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
			paths, err := getAttributePathsForUnknownValues(ctx, inTerraformState, inCloudFormationResourceModel, val)
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
				inCloudFormationResourceModel = inCloudFormationResourceModel.WithAttributeName(strcase.ToCamel(key))
			}
			paths, err := getAttributePathsForUnknownValues(ctx, inTerraformState, inCloudFormationResourceModel, val)
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
