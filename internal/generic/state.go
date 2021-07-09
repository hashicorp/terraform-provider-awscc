package generic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema"
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

		if err == tftypes.ErrInvalidStep {
			// Value not found in CloudFormation ResourceModel. Set to Nil in State.

			// TODO
			// TODO State.SetAttribute does not support passing `nil` to set a Null value.
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
		err = state.SetAttribute(ctx, path.InTerraformState, val)

		if err != nil {
			return fmt.Errorf("error setting value at %s: %w", path.InTerraformState, err)
		}
	}

	return nil
}

// GetCloudFormationResourceModelValue returns the Terraform Value for the specified CloudFormation ResourceModel (string).
func GetCloudFormationResourceModelValue(ctx context.Context, schema *schema.Schema, resourceModel string) (tftypes.Value, error) {
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
func GetCloudFormationResourceModelRawValue(ctx context.Context, schema *schema.Schema, resourceModel map[string]interface{}) (tftypes.Value, error) {
	return getCloudFormationResourceModelValue(ctx, schema, nil, resourceModel)
}

func getCloudFormationResourceModelValue(ctx context.Context, schema *schema.Schema, path *tftypes.AttributePath, v interface{}) (tftypes.Value, error) {
	var typ tftypes.Type

	if len(path.Steps()) == 0 {
		typ = tftypes.Object{}
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
				// TODO
				// TODO How to express the path for the element without knowing its value???
				// TODO
				path = path.WithElementKeyValue(tftypes.NewValue(typ.(tftypes.Set).ElementType, v))
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
		vals := make(map[string]tftypes.Value)
		for key, v := range v {
			if typ.Is(tftypes.Map{}) {
				path = path.WithElementKeyString(key)
			} else {
				// In the Terraform Value attribute names are snake cased.
				path = path.WithAttributeName(strcase.ToSnake(key))
			}
			val, err := getCloudFormationResourceModelValue(ctx, schema, path, v)
			if err != nil {
				return tftypes.Value{}, err
			}
			if typ.Is(tftypes.Map{}) {
				vals[key] = val
			} else {
				// In the Terraform Value attribute names are snake cased.
				vals[strcase.ToSnake(key)] = val
			}
			path = path.WithoutLastStep()
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
