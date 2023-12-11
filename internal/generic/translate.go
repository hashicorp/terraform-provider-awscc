// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

// Translates a Terraform Value to Cloud Control DesiredState.
type toCloudControl struct {
	tfToCfNameMap map[string]string
}

// AsRaw returns the raw map[string]interface{} representing Cloud Control DesiredState from a Terraform Value.
func (t toCloudControl) AsRaw(ctx context.Context, schema typeAtTerraformPather, val tftypes.Value) (map[string]interface{}, error) {
	v, err := t.rawFromValue(ctx, schema, nil, val)

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
func (t toCloudControl) AsString(ctx context.Context, schema typeAtTerraformPather, val tftypes.Value) (string, error) {
	v, err := t.AsRaw(ctx, schema, val)

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
func (t toCloudControl) rawFromValue(ctx context.Context, schema typeAtTerraformPather, path *tftypes.AttributePath, val tftypes.Value) (interface{}, error) {
	if val.IsNull() || !val.IsKnown() {
		return nil, nil
	}

	attributeType, err := schema.TypeAtTerraformPath(ctx, path)

	if err != nil {
		return nil, fmt.Errorf("getting attribute type at %s: %w", path, err)
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
		if t := new(jsontypes.NormalizedType); t.Equal(attributeType) {
			var v interface{}
			diags := jsontypes.NewNormalizedValue(s).Unmarshal(&v)
			if diags.HasError() {
				return nil, tfresource.DiagnosticsError(diags)
			}
			return v, nil
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
		for idx, val := range vals {
			if typ.Is(tftypes.Set{}) {
				// No need to worry about a specific value here.
				path = path.WithElementKeyValue(tftypes.NewValue(typ.(tftypes.Set).ElementType, nil))
			} else {
				path = path.WithElementKeyInt(idx)
			}
			v, err := t.rawFromValue(ctx, schema, path, val)
			if err != nil {
				return nil, err
			}
			path = path.WithoutLastStep()
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
			if typ.Is(tftypes.Object{}) {
				path = path.WithAttributeName(name)
			} else {
				path = path.WithElementKeyString(name)
			}
			v, err := t.rawFromValue(ctx, schema, path, val)
			if err != nil {
				return nil, err
			}
			path = path.WithoutLastStep()
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
func (t toTerraform) FromRaw(ctx context.Context, schema typeAtTerraformPather, resourceModel map[string]interface{}) (tftypes.Value, error) {
	return t.valueFromRaw(ctx, schema, nil, resourceModel)
}

// FromString returns the Terraform Value for the specified Cloud Control Properties (string).
func (t toTerraform) FromString(ctx context.Context, schema typeAtTerraformPather, resourceModel string) (tftypes.Value, error) {
	var v interface{}

	if err := json.Unmarshal([]byte(resourceModel), &v); err != nil {
		return tftypes.Value{}, err
	}

	if v, ok := v.(map[string]interface{}); ok {
		return t.FromRaw(ctx, schema, v)
	}

	return tftypes.Value{}, fmt.Errorf("unexpected raw type: %T", v)
}

func (t toTerraform) valueFromRaw(ctx context.Context, schema typeAtTerraformPather, path *tftypes.AttributePath, v interface{}) (tftypes.Value, error) {
	attrType, err := schema.TypeAtTerraformPath(ctx, path)

	if err != nil {
		return tftypes.Value{}, fmt.Errorf("getting attribute type at %s: %w", path, err)
	}

	typ := attrType.TerraformType(ctx)

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
		if len(v) == 0 {
			return tftypes.NewValue(typ, nil), nil
		}
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
					tflog.Info(ctx, "attribute name mapping not found", map[string]interface{}{
						"key": key,
					})
					continue
				}
				path = path.WithAttributeName(attributeName)
			} else {
				path = path.WithElementKeyString(key)
			}
			val, err := t.valueFromRaw(ctx, schema, path, v)
			if err != nil {
				if isObject {
					tflog.Info(ctx, "not found in Terraform schema", map[string]interface{}{
						"key":   key,
						"path":  path,
						"error": err.Error(),
					})
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
