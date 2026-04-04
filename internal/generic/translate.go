// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/errs/diag"
)

// Translates a Terraform Value to Cloud Control DesiredState.
type toCloudControl struct {
	tfToCfNameMap map[string]string
}

// AsRaw returns the raw map[string]interface{} representing Cloud Control DesiredState from a Terraform Value.
func (t toCloudControl) AsRaw(ctx context.Context, schema typeAtTerraformPather, val tftypes.Value) (map[string]any, error) {
	v, err := t.rawFromValue(ctx, schema, nil, val)

	if err != nil {
		return nil, err
	}

	if v == nil {
		return make(map[string]any), nil
	}

	if v, ok := v.(map[string]any); ok {
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
func (t toCloudControl) rawFromValue(ctx context.Context, schema typeAtTerraformPather, path *tftypes.AttributePath, val tftypes.Value) (any, error) {
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
			var v any
			diags := jsontypes.NewNormalizedValue(s).Unmarshal(&v)
			if diags.HasError() {
				return nil, ccdiag.DiagnosticsError(diags)
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
		vs := make([]any, 0)
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
		vs := make(map[string]any)
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
func (t toTerraform) FromRaw(ctx context.Context, schema typeAtTerraformPather, resourceModel map[string]any) (tftypes.Value, error) {
	return t.valueFromRaw(ctx, schema, nil, resourceModel)
}

// FromString returns the Terraform Value for the specified Cloud Control Properties (string).
// If priorStateMap is non-nil (e.g. from a previous Read), key-value list attributes
// (Tags, TargetGroupAttributes, LoadBalancerAttributes, etc.) in the resource model
// are reordered to match prior state so plan shows no diff regardless of user config order.
func (t toTerraform) FromString(ctx context.Context, schema typeAtTerraformPather, resourceModel string, priorStateMap map[string]any) (tftypes.Value, error) {
	var v any

	if err := json.Unmarshal([]byte(resourceModel), &v); err != nil {
		return tftypes.Value{}, err
	}

	if m, ok := v.(map[string]any); ok {
		if priorStateMap != nil {
			reorderKeyValueSlicesToMatchPrior(m, priorStateMap)
		} else {
			normalizeKeyValueSlices(m)
		}
		return t.FromRaw(ctx, schema, m)
	}

	return tftypes.Value{}, fmt.Errorf("unexpected raw type: %T", v)
}

// reorderKeyValueSlicesToMatchPrior reorders list values in m to match prior (key-value
// lists by key, primitive lists by value). Recurses into nested objects (e.g. VpcConfig).
// Preserves the user's order from the last apply so plan shows no diff.
func reorderKeyValueSlicesToMatchPrior(m, prior map[string]any) {
	for key, val := range m {
		switch v := val.(type) {
		case map[string]any:
			priorMap, _ := prior[key].(map[string]any)
			if priorMap != nil {
				reorderKeyValueSlicesToMatchPrior(v, priorMap)
			}
		case []any:
			if len(v) == 0 {
				continue
			}
			priorSlice, _ := prior[key].([]any)
			reordered := reorderKeyValueSliceToMatch(v, priorSlice)
			if reordered != nil {
				m[key] = reordered
			} else if reorderedPrim := reorderPrimitiveSliceToMatch(v, priorSlice); reorderedPrim != nil {
				m[key] = reorderedPrim
			} else {
				sortSliceByKey(v)
			}
		}
	}
}

// reorderPrimitiveSliceToMatch reorders current to match the element order in prior.
// Elements in current that are not in prior are appended at the end (sorted).
// Returns the reordered slice, or nil if current is not a primitive slice (all string or all number).
func reorderPrimitiveSliceToMatch(current, prior []any) []any {
	if len(current) == 0 {
		return current
	}
	// Check all elements are primitive (string or float64)
	prim := primitiveKind(current[0])
	if prim == primKindOther {
		return nil
	}
	for i := 1; i < len(current); i++ {
		if primitiveKind(current[i]) != prim {
			return nil
		}
	}
	// Build set of current elements for lookup
	currentSet := make(map[string]any)
	keyOf := func(a any) string {
		switch x := a.(type) {
		case string:
			return x
		case float64:
			return strconv.FormatFloat(x, 'g', -1, 64)
		default:
			return fmt.Sprint(a)
		}
	}
	for _, el := range current {
		currentSet[keyOf(el)] = el
	}
	// Build result: prior order first, then current-only (sorted)
	seen := make(map[string]bool)
	var result []any
	if len(prior) > 0 {
		for _, el := range prior {
			k := keyOf(el)
			if cur, exists := currentSet[k]; exists {
				result = append(result, cur)
				seen[k] = true
			}
		}
	}
	var extra []string
	for k := range currentSet {
		if !seen[k] {
			extra = append(extra, k)
		}
	}
	sort.Strings(extra)
	for _, k := range extra {
		result = append(result, currentSet[k])
	}
	return result
}

// primitive kind constants for primitiveKind return value.
const (
	primKindOther = iota
	primKindString
	primKindFloat64
)

// primitiveKind returns primKindString for string, primKindFloat64 for float64, primKindOther for other.
func primitiveKind(a any) int {
	switch a.(type) {
	case string:
		return primKindString
	case float64:
		return primKindFloat64
	default:
		return primKindOther
	}
}

// reorderKeyValueSliceToMatch reorders current to match the key order in prior.
// Keys in current that are not in prior are appended at the end in sorted order.
// Returns the reordered slice, or nil if current is not a key-value slice.
func reorderKeyValueSliceToMatch(current, prior []any) []any {
	if len(current) == 0 {
		return current
	}
	// Build current by key
	byKey := make(map[string]map[string]any)
	for _, el := range current {
		m, ok := el.(map[string]any)
		if !ok || (m["Key"] == nil && m["key"] == nil) {
			return nil
		}
		k := keyFromMap(m)
		byKey[k] = m
	}
	// Build result: first in prior order, then any keys only in current (sorted)
	seen := make(map[string]bool)
	var result []any
	if len(prior) > 0 {
		for _, el := range prior {
			p, ok := el.(map[string]any)
			if !ok {
				continue
			}
			k := keyFromMap(p)
			if k == "" {
				continue
			}
			if cur, exists := byKey[k]; exists {
				result = append(result, cur)
				seen[k] = true
			}
		}
	}
	var extra []string
	for k := range byKey {
		if !seen[k] {
			extra = append(extra, k)
		}
	}
	sort.Strings(extra)
	for _, k := range extra {
		result = append(result, byKey[k])
	}
	return result
}

// normalizeKeyValueSlices recursively walks the resource model and sorts any
// list of objects that have a "Key" field (Cloud Control API PascalCase) by
// that key. Used when there is no prior state (e.g. first Read after import).
func normalizeKeyValueSlices(v any) {
	switch x := v.(type) {
	case map[string]any:
		for _, val := range x {
			normalizeKeyValueSlices(val)
		}
	case []any:
		if len(x) == 0 {
			return
		}
		if sortSliceByKey(x) {
			return
		}
		for _, el := range x {
			normalizeKeyValueSlices(el)
		}
	}
}

// sortSliceByKey sorts slice in place by each element's "Key" or "key" field
// (Cloud Control API may use PascalCase or lowercase). Returns true if the
// slice was sorted (all elements were key-value maps), false otherwise.
func sortSliceByKey(slice []any) bool {
	for _, el := range slice {
		m, ok := el.(map[string]any)
		if !ok {
			return false
		}
		// Must have either "Key" or "key" so we can sort
		if _, hasKey := m["Key"]; hasKey {
			continue
		}
		if _, hasKey := m["key"]; hasKey {
			continue
		}
		return false
	}
	sort.Slice(slice, func(i, j int) bool {
		mi := slice[i].(map[string]any)
		mj := slice[j].(map[string]any)
		return keyFromMap(mi) < keyFromMap(mj)
	})
	return true
}

// keyFromMap returns the sort key from a key-value object (Cloud Control "Key" or "key").
func keyFromMap(m map[string]any) string {
	if k, ok := m["Key"].(string); ok {
		return k
	}
	if k, ok := m["key"].(string); ok {
		return k
	}
	return ""
}

func (t toTerraform) valueFromRaw(ctx context.Context, schema typeAtTerraformPather, path *tftypes.AttributePath, v any) (tftypes.Value, error) {
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
	case []any:
		if len(v) == 0 {
			return tftypes.NewValue(typ, nil), nil
		}

		if typ.Is(tftypes.String) {
			// Value is JSON string.
			val, err := json.Marshal(v)

			if err != nil {
				return tftypes.Value{}, err
			}
			return tftypes.NewValue(typ, string(val)), nil
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

	case map[string]any:
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
					tflog.Info(ctx, "attribute name mapping not found", map[string]any{
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
					tflog.Info(ctx, "not found in Terraform schema", map[string]any{
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
