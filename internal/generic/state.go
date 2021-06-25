package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/iancoleman/strcase"
)

// cloudFormationDesiredState returns the string representing CloudFormation DesiredState from a Terraform Plan.
func cloudFormationDesiredState(ctx context.Context, plan *tfsdk.Plan) (string, error) {
	m := make(map[string]interface{})

	err := tftypes.Walk(plan.Raw, func(ap *tftypes.AttributePath, v tftypes.Value) (bool, error) {
		steps := ap.Steps()

		var name string

		switch n := len(steps); n {
		case 0:
			// Top-level.
			return true, nil
		case 1:
			switch typ := steps[0].(type) {
			case tftypes.AttributeName:
				name = string(typ)
			default:
				return false, fmt.Errorf("AttributePath (%s) has unsupported step type: %T", ap, typ)
			}
		default:
			return false, fmt.Errorf("AttributePath (%s) has unsupported number of steps: %d", ap, n)
		}

		if v.IsKnown() && !v.IsNull() {
			typ := v.Type()
			switch {
			case typ.Is(tftypes.Bool):
				var b bool
				if err := v.As(&b); err != nil {
					return false, err
				}
				m[strcase.ToCamel(name)] = b
			case typ.Is(tftypes.Number):
				n := big.NewFloat(0)
				if err := v.As(&n); err != nil {
					return false, err
				}
				f, _ := n.Float64()
				m[strcase.ToCamel(name)] = f
			case typ.Is(tftypes.String):
				var s string
				if err := v.As(&s); err != nil {
					return false, err
				}
				m[strcase.ToCamel(name)] = s
			default:
				return false, fmt.Errorf("attribute %s has unsupported value type: %s", name, typ)
			}
		} else {
			log.Printf("[DEBUG] %s is Null or not Known", name)
		}

		return true, nil
	})

	if err != nil {
		return "", err
	}

	desiredState, err := json.Marshal(m)

	if err != nil {
		return "", err
	}

	return string(desiredState), nil
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
