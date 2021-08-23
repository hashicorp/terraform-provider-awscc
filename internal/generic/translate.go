package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Translates a Terraform Value to CloudFormation DesiredState.
type toCloudFormation struct {
	tfToCfNameMap map[string]string
}

// AsRaw returns the raw map[string]interface{} representing CloudFormation DesiredState from a Terraform Value.
func (t toCloudFormation) AsRaw(ctx context.Context, val tftypes.Value) (map[string]interface{}, error) {
	v, err := t.rawFromValue(val)

	if err != nil {
		return nil, err
	}

	if v == nil {
		return make(map[string]interface{}), nil
	}

	if v, ok := v.(map[string]interface{}); ok {
		return v, nil
	}

	return nil, fmt.Errorf("unexpected raw type: %T", v)
}

// AsString returns the string representing CloudFormation DesiredState from a Terraform Value.
func (t toCloudFormation) AsString(ctx context.Context, val tftypes.Value) (string, error) {
	v, err := t.AsRaw(ctx, val)

	if err != nil {
		return "", err
	}

	desiredState, err := json.Marshal(v)

	if err != nil {
		return "", err
	}

	return string(desiredState), nil
}

// rawFromValue returns the raw value (suitable for JSON marshaling) of the specified Terraform value.
// Terraform attribute names are mapped to CloudFormation property names.
func (t toCloudFormation) rawFromValue(val tftypes.Value) (interface{}, error) {
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
			v, err := t.rawFromValue(val)
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
			v, err := t.rawFromValue(val)
			if err != nil {
				return nil, err
			}
			if v == nil {
				continue
			}
			propertyName, ok := t.tfToCfNameMap[name]
			if !ok {
				return nil, fmt.Errorf("attribute name mapping not found: %s", name)
			}
			vs[propertyName] = v
		}
		if len(vs) == 0 {
			return nil, nil
		}
		return vs, nil
	}

	return nil, fmt.Errorf("unsupported value type: %s", typ)
}

// Translates a CloudFormation ResourceModel to Terraform Value.
type toTerraform struct {
	cfToTfNameMap map[string]string
}
