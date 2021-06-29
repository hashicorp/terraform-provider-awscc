package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/iancoleman/strcase"
)

// cloudFormationDesiredState returns the string representing CloudFormation DesiredState from a Terraform Plan.
func cloudFormationDesiredStateString(ctx context.Context, plan *tfsdk.Plan) (string, error) {
	m, err := cloudFormationDesiredStateRaw(ctx, plan)

	if err != nil {
		return "", err
	}

	desiredState, err := json.Marshal(m)

	if err != nil {
		return "", err
	}

	return string(desiredState), nil
}

// cloudFormationDesiredState returns the raw map[string]interface{} representing CloudFormation DesiredState from a Terraform Plan.
func cloudFormationDesiredStateRaw(ctx context.Context, plan *tfsdk.Plan) (map[string]interface{}, error) {
	v, err := rawValue(ctx, plan.Raw)

	if err != nil {
		return nil, err
	}

	if v, ok := v.(map[string]interface{}); ok {
		return v, nil
	}

	return nil, fmt.Errorf("Plan.Raw value produced unexpected raw type: %T", v)
}

var (
	identifierAttributePath = tftypes.NewAttributePath().WithAttributeName("identifier")
)

// getIdentifier sets the well-known "identifier" attribute in State.
func getIdentifier(ctx context.Context, state *tfsdk.State) (string, error) {
	val, err := state.GetAttribute(ctx, identifierAttributePath)

	if err != nil {
		return "", err
	}

	if val, ok := val.(types.String); ok {
		return val.Value, nil
	}

	return "", fmt.Errorf("invalid identifier type %T", val)
}

// setIdentifier sets the well-known "identifier" attribute in State.
func setIdentifier(ctx context.Context, state *tfsdk.State, id string) error {
	return state.SetAttribute(ctx, identifierAttributePath, id)
}

// rawValue returns the raw value (suitable for JSON marshalling) of the specified Terraform value.
// Attribute names are converted to camel case (AWS standard).
func rawValue(ctx context.Context, val tftypes.Value) (interface{}, error) {
	if val.IsNull() || !val.IsKnown() {
		return nil, nil
	}

	typ := val.Type()
	switch {
	//
	// Primitive types.
	//
	case typ.Is(tftypes.Bool):
		var b bool
		if err := val.As(&b); err != nil {
			return nil, err
		}
		return b, nil

	case typ.Is(tftypes.Number):
		n := big.NewFloat(0)
		if err := val.As(&n); err != nil {
			return nil, err
		}
		f, _ := n.Float64()
		return f, nil

	case typ.Is(tftypes.String):
		var s string
		if err := val.As(&s); err != nil {
			return nil, err
		}
		return s, nil

	//
	// Complex types.
	//
	case typ.Is(tftypes.List{}), typ.Is(tftypes.Set{}), typ.Is(tftypes.Tuple{}):
		var vals []tftypes.Value
		if err := val.As(&vals); err != nil {
			return nil, err
		}
		vs := make([]interface{}, 0)
		for _, val := range vals {
			v, err := rawValue(ctx, val)
			if err != nil {
				return nil, err
			}
			if v == nil {
				continue
			}
			vs = append(vs, v)
		}
		if len(vs) == 0 {
			return nil, nil
		}
		return vs, nil

	case typ.Is(tftypes.Map{}), typ.Is(tftypes.Object{}):
		var vals map[string]tftypes.Value
		if err := val.As(&vals); err != nil {
			return nil, err
		}
		vs := make(map[string]interface{})
		for name, val := range vals {
			v, err := rawValue(ctx, val)
			if err != nil {
				return nil, err
			}
			if v == nil {
				continue
			}
			vs[strcase.ToCamel(name)] = v
		}
		if len(vs) == 0 {
			return nil, nil
		}
		return vs, nil
	}

	return nil, fmt.Errorf("unsupported value type: %s", typ)
}
