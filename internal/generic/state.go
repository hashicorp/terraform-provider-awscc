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

type State struct {
	inner *tfsdk.State
}

var (
	identifierAttributePath = tftypes.NewAttributePath().WithAttributeName("identifier")
)

// GetIdentifier gets the well-known "identifier" attribute from State.
func (s *State) GetIdentifier(ctx context.Context) (string, error) {
	val, err := s.inner.GetAttribute(ctx, identifierAttributePath)

	if err != nil {
		return "", err
	}

	if val, ok := val.(types.String); ok {
		return val.Value, nil
	}

	return "", fmt.Errorf("invalid identifier type %T", val)
}

// SetIdentifier sets the well-known "identifier" attribute in State.
func (s *State) SetIdentifier(ctx context.Context, id string) error {
	return s.inner.SetAttribute(ctx, identifierAttributePath, id)
}

// SetCloudFormationResourceModel sets the string representing CloudFormation ResourceModel in State.
func (s *State) SetCloudFormationResourceModel(ctx context.Context, resourceModel string) error {
	var v interface{}

	if err := json.Unmarshal([]byte(resourceModel), &v); err != nil {
		return err
	}

	if v, ok := v.(map[string]interface{}); ok {
		return s.SetCloudFormationResourceModelRaw(ctx, v)
	}

	return fmt.Errorf("CloudFormation ResourceModel value produced unexpected raw type: %T", v)
}

// SetCloudFormationResourceModel sets the raw map[string]interface{} representing CloudFormation ResourceModel in State.
func (s *State) SetCloudFormationResourceModelRaw(ctx context.Context, v map[string]interface{}) error {
	val, err := valueFromRaw(ctx, v)

	if err != nil {
		return err
	}

	s.inner.Raw = val

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
