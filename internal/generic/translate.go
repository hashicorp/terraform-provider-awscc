package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Translates a Terraform Value to Cloud Control DesiredState.
type toCloudControl struct {
	tfToCfNameMap map[string]string
}

// AsRaw returns the raw map[string]interface{} representing Cloud Control DesiredState from a Terraform Value.
func (t toCloudControl) AsRaw(ctx context.Context, val tftypes.Value) (map[string]interface{}, error) {
	v, err := t.rawFromValue(ctx, val)

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

// AsString returns the string representing Cloud Control DesiredState from a Terraform Value.
func (t toCloudControl) AsString(ctx context.Context, val tftypes.Value) (string, error) {
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
// Terraform attribute names are mapped to Cloud Control property names.
func (t toCloudControl) rawFromValue(ctx context.Context, val tftypes.Value) (interface{}, error) { //nolint:unparam
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
			v, err := t.rawFromValue(ctx, val)
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
			v, err := t.rawFromValue(ctx, val)
			if err != nil {
				return nil, err
			}
			if v == nil {
				continue
			}
			if typ.Is(tftypes.Object{}) {
				propertyName, ok := t.tfToCfNameMap[name]
				if !ok {
					return nil, fmt.Errorf("attribute name mapping not found: %s", name)
				}
				vs[propertyName] = v
			} else {
				vs[name] = v
			}
		}
		if len(vs) == 0 {
			return nil, nil
		}
		return vs, nil
	}

	return nil, fmt.Errorf("unsupported value type: %s", typ)
}

// Translates Cloud Control Properties to Terraform Value.
type toTerraform struct {
	cfToTfNameMap map[string]string
}

// FromRaw returns the Terraform Value for the specified Cloud Control Properties (raw map[string]interface{}).
func (t toTerraform) FromRaw(ctx context.Context, schema *tfsdk.Schema, resourceModel map[string]interface{}) (tftypes.Value, error) {
	return t.valueFromRaw(ctx, schema, nil, resourceModel)
}

// FromString returns the Terraform Value for the specified Cloud Control Properties (string).
func (t toTerraform) FromString(ctx context.Context, schema *tfsdk.Schema, resourceModel string) (tftypes.Value, error) {
	var v interface{}

	if err := json.Unmarshal([]byte(resourceModel), &v); err != nil {
		return tftypes.Value{}, err
	}

	if v, ok := v.(map[string]interface{}); ok {
		return t.FromRaw(ctx, schema, v)
	}

	return tftypes.Value{}, fmt.Errorf("unexpected raw type: %T", v)
}

func (t toTerraform) valueFromRaw(ctx context.Context, schema *tfsdk.Schema, path *tftypes.AttributePath, v interface{}) (tftypes.Value, error) {
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
				path = path.WithElementKeyInt(idx)
			}
			val, err := t.valueFromRaw(ctx, schema, path, v)
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
				attributeName, ok := t.cfToTfNameMap[key]
				if !ok {
					log.Printf("[INFO] attribute name mapping not found. key: %s", key)
					continue
				}
				path = path.WithAttributeName(attributeName)
			} else {
				path = path.WithElementKeyString(key)
			}
			val, err := t.valueFromRaw(ctx, schema, path, v)
			if err != nil {
				if isObject {
					log.Printf("[INFO] not found in Terraform schema. key: %s, path: %s, error: %s", key, path, err.Error())
					path = path.WithoutLastStep()
					continue
				}
				return tftypes.Value{}, err
			}
			if isObject {
				// Attribute name mapping assured above.
				vals[t.cfToTfNameMap[key]] = val
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
