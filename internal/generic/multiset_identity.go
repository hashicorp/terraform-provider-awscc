// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// fillUnknownsFromMatchedState resolves unknown values inside the planned elements of an
// unordered list of objects from the prior state element with the same identity (see
// identityField), rather than from the state element at the same index.
//
// The code generator gives every Computed attribute a UseStateForUnknown plan modifier, which
// pairs plan and state list elements by index. For an "insertionOrder": false array that pairing
// is wrong whenever the configuration and state orders differ (e.g. after an element is removed),
// and would copy one element's computed values into another. Resolving the unknowns here first,
// with UseStateForUnknown's rules (plan unknown, config not unknown, state not null), leaves the
// nested modifiers nothing to do.
//
// Returns planned unchanged if the elements can't be identified.
func fillUnknownsFromMatchedState(ctx context.Context, planned, config, state basetypes.ListValue) basetypes.ListValue {
	if planned.IsNull() || planned.IsUnknown() || state.IsNull() || state.IsUnknown() {
		return planned
	}
	if v, err := planned.ToTerraformValue(ctx); err != nil || v.IsFullyKnown() {
		// Nothing to resolve.
		return planned
	}

	plannedObjs, ok := listObjects(ctx, planned)
	if !ok {
		return planned
	}
	stateObjs, ok := listObjects(ctx, state)
	if !ok {
		return planned
	}

	plannedKeys := objectsAsIdentityMaps(plannedObjs)
	stateKeys := objectsAsIdentityMaps(stateObjs)
	field := identityField(plannedKeys, stateKeys)
	if field == "" {
		return planned
	}

	keyOf := fieldKey(field)
	stateByKey := make(map[string]basetypes.ObjectValue, len(stateObjs))
	for i, m := range stateKeys {
		stateByKey[keyOf(m.(map[string]any))] = stateObjs[i]
	}

	var configElems []attr.Value
	if !config.IsNull() && !config.IsUnknown() {
		configElems = config.Elements()
	}

	elems := make([]attr.Value, len(plannedObjs))
	for i, p := range plannedObjs {
		elems[i] = p
		s, ok := stateByKey[keyOf(plannedKeys[i].(map[string]any))]
		if !ok {
			continue
		}
		var c attr.Value
		if i < len(configElems) {
			c = configElems[i]
		}
		elems[i] = fillUnknowns(ctx, p, c, s)
	}

	result, diags := types.ListValue(planned.ElementType(ctx), elems)
	if diags.HasError() {
		return planned
	}

	return result
}

// fillUnknowns returns plan with any unknown values replaced by the corresponding value in
// state, following UseStateForUnknown's rules. Objects are recursed into by attribute name
// and lists by index (state is already the matched element). config may be nil.
func fillUnknowns(ctx context.Context, plan, config, state attr.Value) attr.Value {
	if plan == nil || state == nil {
		return plan
	}

	if plan.IsUnknown() {
		if (config != nil && config.IsUnknown()) || state.IsNull() || state.IsUnknown() {
			return plan
		}
		return state
	}

	if plan.IsNull() || state.IsNull() || state.IsUnknown() {
		return plan
	}

	switch p := plan.(type) {
	case basetypes.ObjectValue:
		s, ok := state.(basetypes.ObjectValue)
		if !ok {
			return plan
		}
		var configAttrs map[string]attr.Value
		if c, ok := config.(basetypes.ObjectValue); ok && !c.IsNull() && !c.IsUnknown() {
			configAttrs = c.Attributes()
		}
		stateAttrs := s.Attributes()
		attrs := make(map[string]attr.Value, len(p.Attributes()))
		for name, v := range p.Attributes() {
			attrs[name] = fillUnknowns(ctx, v, configAttrs[name], stateAttrs[name])
		}
		result, diags := types.ObjectValue(p.AttributeTypes(ctx), attrs)
		if diags.HasError() {
			return plan
		}
		return result

	case basetypes.ListValue:
		s, ok := state.(basetypes.ListValue)
		if !ok {
			return plan
		}
		var configElems []attr.Value
		if c, ok := config.(basetypes.ListValue); ok && !c.IsNull() && !c.IsUnknown() {
			configElems = c.Elements()
		}
		stateElems := s.Elements()
		elems := make([]attr.Value, len(p.Elements()))
		for i, v := range p.Elements() {
			var c, st attr.Value
			if i < len(configElems) {
				c = configElems[i]
			}
			if i < len(stateElems) {
				st = stateElems[i]
			}
			elems[i] = fillUnknowns(ctx, v, c, st)
		}
		result, diags := types.ListValue(p.ElementType(ctx), elems)
		if diags.HasError() {
			return plan
		}
		return result

	default:
		return plan
	}
}

// listObjects returns the list's elements as objects, and false if any element is not an object.
func listObjects(ctx context.Context, list basetypes.ListValue) ([]basetypes.ObjectValue, bool) {
	elems := list.Elements()
	objs := make([]basetypes.ObjectValue, len(elems))
	for i, e := range elems {
		v, ok := e.(basetypes.ObjectValuable)
		if !ok {
			return nil, false
		}
		obj, diags := v.ToObjectValue(ctx)
		if diags.HasError() || obj.IsNull() || obj.IsUnknown() {
			return nil, false
		}
		objs[i] = obj
	}
	return objs, true
}

// objectsAsIdentityMaps returns, for each object, a map of its known string and integer
// attribute values, in the form identityField expects.
func objectsAsIdentityMaps(objs []basetypes.ObjectValue) []any {
	result := make([]any, len(objs))
	for i, obj := range objs {
		m := make(map[string]any)
		for name, v := range obj.Attributes() {
			if v.IsNull() || v.IsUnknown() {
				continue
			}
			switch v := v.(type) {
			case basetypes.StringValue:
				m[name] = v.ValueString()
			case basetypes.Int64Value:
				m[name] = strconv.FormatInt(v.ValueInt64(), 10)
			}
		}
		result[i] = m
	}
	return result
}
