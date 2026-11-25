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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/errs/diag"
)

// Translates a Terraform Value to Cloud Control DesiredState.
type toCloudControl struct {
	tfToCfNameMap     map[string]string
	pathAwareNames    bool
	pathTfToCfNameMap map[string]string // "CfParentPath/tfAttr" → CF property name
}

// AsRaw returns the raw map[string]interface{} representing Cloud Control DesiredState from a Terraform Value.
func (t toCloudControl) AsRaw(ctx context.Context, schema typeAtTerraformPather, val tftypes.Value) (map[string]any, error) {
	v, err := t.rawFromValue(ctx, schema, nil, nil, val)

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
// cfPath tracks the CloudFormation property path for path-aware name resolution.
func (t toCloudControl) rawFromValue(ctx context.Context, schema typeAtTerraformPather, path *tftypes.AttributePath, cfPath []string, val tftypes.Value) (any, error) {
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
			v, err := t.rawFromValue(ctx, schema, path, cfPath, val)
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

			// Resolve CF property name — path-aware or flat.
			var propertyName string
			if typ.Is(tftypes.Object{}) {
				if t.pathAwareNames {
					pathKey := strings.Join(append(cfPath, name), "/")
					var ok bool
					propertyName, ok = t.pathTfToCfNameMap[pathKey]
					if !ok {
						return nil, fmt.Errorf("path-aware attribute name mapping not found: %s", pathKey)
					}
				} else {
					var ok bool
					propertyName, ok = t.tfToCfNameMap[name]
					if !ok {
						return nil, fmt.Errorf("attribute name mapping not found: %s", name)
					}
				}
			}

			// Build child cfPath for recursion.
			var childCfPath []string
			if typ.Is(tftypes.Object{}) {
				childCfPath = make([]string, len(cfPath)+1)
				copy(childCfPath, cfPath)
				childCfPath[len(cfPath)] = propertyName
			} else {
				childCfPath = cfPath
			}

			v, err := t.rawFromValue(ctx, schema, path, childCfPath, val)
			if err != nil {
				return nil, err
			}
			path = path.WithoutLastStep()
			if v == nil {
				continue
			}
			if typ.Is(tftypes.Object{}) {
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
	cfToTfNameMap     map[string]string
	pathAwareNames    bool
	pathCfToTfNameMap map[string]string // "CfParentPath/CfProp" → TF attribute name
	// unorderedPaths holds the CloudFormation JSON Pointer paths (array indices stripped)
	// of "insertionOrder": false arrays, as returned by unorderedArrayPaths. Object arrays
	// at these paths are matched to prior state by an identity field (see identityField).
	unorderedPaths map[string]bool
}

// FromRaw returns the Terraform Value for the specified Cloud Control Properties (raw map[string]interface{}).
func (t toTerraform) FromRaw(ctx context.Context, schema typeAtTerraformPather, resourceModel map[string]any) (tftypes.Value, error) {
	return t.valueFromRaw(ctx, schema, nil, nil, resourceModel)
}

// FromString returns the Terraform Value for the specified Cloud Control Properties (string).
// If priorStateMap is non-nil (e.g. from a previous Read), list attributes in the resource model
// are reordered to match prior state (see sliceMatcher) so plan shows no diff regardless of the
// order Cloud Control returns them in.
func (t toTerraform) FromString(ctx context.Context, schema typeAtTerraformPather, resourceModel string, priorStateMap map[string]any) (tftypes.Value, error) {
	var v any

	if err := json.Unmarshal([]byte(resourceModel), &v); err != nil {
		return tftypes.Value{}, err
	}

	if m, ok := v.(map[string]any); ok {
		if priorStateMap != nil {
			sliceMatcher{unorderedPaths: t.unorderedPaths}.reorder(m, priorStateMap, "")
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
	sliceMatcher{}.reorder(m, prior, "")
}

// sliceMatcher reorders the lists in a Cloud Control resource model to match prior state, so
// that each list element lines up with the prior element it corresponds to:
//   - key-value lists (elements with "Key"/"key", e.g. Tags) by key;
//   - object lists at an unordered path by an identity field chosen from the data (see
//     identityField), e.g. SageMaker Cluster InstanceGroups by InstanceGroupName;
//   - primitive lists by value.
//
// It recurses into nested objects and into list elements, pairing each element with its
// matched prior element.
type sliceMatcher struct {
	// unorderedPaths is the set of unordered array paths (see toTerraform.unorderedPaths).
	unorderedPaths map[string]bool
	// identityOnly restricts reordering to unordered object lists that have an identity field,
	// leaving all other lists in Cloud Control order.
	identityOnly bool
}

// reorder reorders the lists in m to match prior. prefix is the CloudFormation JSON Pointer
// path of m with array indices stripped ("" for the resource model).
func (sm sliceMatcher) reorder(m, prior map[string]any, prefix string) {
	for key, val := range m {
		path := prefix + "/" + key

		switch v := val.(type) {
		case map[string]any:
			priorMap, _ := prior[key].(map[string]any)
			if priorMap != nil {
				sm.reorder(v, priorMap, path)
			}
		case []any:
			if len(v) == 0 {
				continue
			}
			priorSlice, _ := prior[key].([]any)

			if !sm.identityOnly {
				if reordered := reorderKeyValueSliceToMatch(v, priorSlice); reordered != nil {
					m[key] = reordered
					sm.reorderMatchedElements(reordered, priorSlice, keyFromMap, path)
					continue
				}
			}

			if sm.unorderedPaths[path] {
				if field := identityField(v, priorSlice); field != "" {
					keyOf := fieldKey(field)
					reordered := reorderSliceByKeyToMatch(v, priorSlice, keyOf)
					m[key] = reordered
					sm.reorderMatchedElements(reordered, priorSlice, keyOf, path)
					continue
				}
			}

			if sm.identityOnly {
				continue
			}

			if reorderedPrim := reorderPrimitiveSliceToMatch(v, priorSlice); reorderedPrim != nil {
				m[key] = reorderedPrim
			} else {
				sortSliceByKey(v)
				for _, el := range v {
					if elMap, ok := el.(map[string]any); ok {
						sm.reorder(elMap, nil, path)
					}
				}
			}
		}
	}
}

// reorderMatchedElements recurses into each object element of current, pairing it with the
// prior element that has the same key (as returned by keyOf).
func (sm sliceMatcher) reorderMatchedElements(current, prior []any, keyOf func(map[string]any) string, path string) {
	priorByKey := make(map[string]map[string]any, len(prior))
	for _, el := range prior {
		if p, ok := el.(map[string]any); ok {
			if k := keyOf(p); k != "" {
				priorByKey[k] = p
			}
		}
	}
	for _, el := range current {
		if elMap, ok := el.(map[string]any); ok {
			sm.reorder(elMap, priorByKey[keyOf(elMap)], path)
		}
	}
}

// identityField picks the field that identifies each element of an object array, so
// that current elements (from Cloud Control) can be paired with prior state elements
// regardless of order.
//
// Candidate fields are named "Name"/"Id" or end in "Name"/"Id" (e.g. "InstanceGroupName"),
// tried in that preference order and then alphabetically. The first candidate whose
// value is a non-empty scalar on every element and unique within both current and prior,
// and that shares at least one value between them, is returned. Returns "" if there is
// no such field, in which case the array should be left as is.
//
// Any field that qualifies gives a valid pairing, because reordering only permutes the
// Cloud Control elements and never changes their values.
func identityField(current, prior []any) string {
	if len(current) == 0 || len(prior) == 0 {
		return ""
	}

	first, ok := current[0].(map[string]any)
	if !ok {
		return ""
	}

	candidates := make([]string, 0, len(first))
	for field := range first {
		if identityFieldRank(field) != identityRankNone {
			candidates = append(candidates, field)
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		ri, rj := identityFieldRank(candidates[i]), identityFieldRank(candidates[j])
		if ri != rj {
			return ri < rj
		}
		return candidates[i] < candidates[j]
	})

	for _, field := range candidates {
		keyOf := fieldKey(field)
		currentKeys, ok := uniqueKeys(current, keyOf)
		if !ok {
			continue
		}
		priorKeys, ok := uniqueKeys(prior, keyOf)
		if !ok {
			continue
		}
		for k := range currentKeys {
			if priorKeys[k] {
				return field
			}
		}
	}

	return ""
}

// identity field rank constants for identityFieldRank return value (lower is preferred).
const (
	identityRankNone = iota - 1
	identityRankName
	identityRankNameSuffix
	identityRankId
	identityRankIdSuffix
)

// identityFieldRank returns the preference rank of a candidate identity field name,
// or identityRankNone if the name is not a candidate.
// Both CloudFormation (e.g. "InstanceGroupName") and Terraform (e.g. "instance_group_name")
// naming is supported.
func identityFieldRank(field string) int {
	switch {
	case field == "Name" || field == "name":
		return identityRankName
	case len(field) > len("Name") && (strings.HasSuffix(field, "Name") || strings.HasSuffix(field, "_name")):
		return identityRankNameSuffix
	case field == "Id" || field == "id":
		return identityRankId
	case len(field) > len("Id") && (strings.HasSuffix(field, "Id") || strings.HasSuffix(field, "_id")):
		return identityRankIdSuffix
	default:
		return identityRankNone
	}
}

// fieldKey returns a function that extracts the scalar value of field from an object
// element as a string, or "" if the field is missing, empty, or not a scalar.
func fieldKey(field string) func(map[string]any) string {
	return func(m map[string]any) string {
		switch v := m[field].(type) {
		case string:
			return v
		case float64:
			return primitiveKey(v)
		default:
			return ""
		}
	}
}

// uniqueKeys returns the set of keys of slice's elements, and false if any element is
// not an object, has an empty key, or shares its key with another element.
func uniqueKeys(slice []any, keyOf func(map[string]any) string) (map[string]bool, bool) {
	keys := make(map[string]bool, len(slice))
	for _, el := range slice {
		m, ok := el.(map[string]any)
		if !ok {
			return nil, false
		}
		k := keyOf(m)
		if k == "" || keys[k] {
			return nil, false
		}
		keys[k] = true
	}

	return keys, true
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
	// Build set of current elements for lookup, using seen map to track both
	currentSet := make(map[string]any, len(current))
	for _, el := range current {
		k := primitiveKey(el)
		currentSet[k] = el
	}
	// Build result: prior order first, then current-only (sorted)
	result := make([]any, 0, len(current))
	if len(prior) > 0 {
		for _, el := range prior {
			k := primitiveKey(el)
			if cur, exists := currentSet[k]; exists {
				result = append(result, cur)
				delete(currentSet, k) // reuse map to track unseen
			}
		}
	}
	// Remaining items in currentSet are extras
	if len(currentSet) > 0 {
		extra := make([]string, 0, len(currentSet))
		for k := range currentSet {
			extra = append(extra, k)
		}
		sort.Strings(extra)
		for _, k := range extra {
			result = append(result, currentSet[k])
		}
	}
	return result
}

// primitiveKey returns the string key for a primitive value.
func primitiveKey(a any) string {
	switch x := a.(type) {
	case string:
		return x
	case float64:
		return strconv.FormatFloat(x, 'g', -1, 64)
	default:
		return fmt.Sprint(a)
	}
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
	for _, el := range current {
		m, ok := el.(map[string]any)
		if !ok || (m["Key"] == nil && m["key"] == nil) {
			return nil
		}
	}

	return reorderSliceByKeyToMatch(current, prior, keyFromMap)
}

// reorderSliceByKeyToMatch reorders current (a slice of objects) to match the key order
// in prior, where each element's key is returned by keyOf. Keys in current that are not
// in prior are appended at the end in sorted order.
func reorderSliceByKeyToMatch(current, prior []any, keyOf func(map[string]any) string) []any {
	if len(current) == 0 {
		return current
	}
	// Build current by key
	byKey := make(map[string]map[string]any, len(current))
	for _, el := range current {
		m := el.(map[string]any)
		byKey[keyOf(m)] = m
	}
	// Build result: first in prior order, then any keys only in current (sorted)
	result := make([]any, 0, len(current))
	if len(prior) > 0 {
		for _, el := range prior {
			p, ok := el.(map[string]any)
			if !ok {
				continue
			}
			k := keyOf(p)
			if k == "" {
				continue
			}
			if cur, exists := byKey[k]; exists {
				result = append(result, cur)
				delete(byKey, k) // reuse map to track unseen
			}
		}
	}
	// Remaining items in byKey are extras
	if len(byKey) > 0 {
		extra := make([]string, 0, len(byKey))
		for k := range byKey {
			extra = append(extra, k)
		}
		sort.Strings(extra)
		for _, k := range extra {
			result = append(result, byKey[k])
		}
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

func (t toTerraform) valueFromRaw(ctx context.Context, schema typeAtTerraformPather, path *tftypes.AttributePath, cfPath []string, v any) (tftypes.Value, error) {
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
			val, err := t.valueFromRaw(ctx, schema, path, cfPath, v)
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
			var attributeName string
			if isObject {
				if t.pathAwareNames {
					pathKey := strings.Join(append(cfPath, key), "/")
					var ok bool
					attributeName, ok = t.pathCfToTfNameMap[pathKey]
					if !ok {
						tflog.Info(ctx, "path-aware attribute name mapping not found", map[string]any{
							"key":     key,
							"pathKey": pathKey,
						})
						continue
					}
				} else {
					var ok bool
					attributeName, ok = t.cfToTfNameMap[key]
					if !ok {
						tflog.Info(ctx, "attribute name mapping not found", map[string]any{
							"key": key,
						})
						continue
					}
				}
				path = path.WithAttributeName(attributeName)
			} else {
				path = path.WithElementKeyString(key)
			}

			// Build child cfPath for recursion.
			var childCfPath []string
			if isObject {
				childCfPath = make([]string, len(cfPath)+1)
				copy(childCfPath, cfPath)
				childCfPath[len(cfPath)] = key
			} else {
				childCfPath = cfPath
			}

			val, err := t.valueFromRaw(ctx, schema, path, childCfPath, v)
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
				vals[attributeName] = val
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
