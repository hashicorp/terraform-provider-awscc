// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.ListTypable                    = (*multisetType)(nil)
	_ basetypes.ListValuable                   = (*Multiset)(nil)
	_ basetypes.ListValuableWithSemanticEquals = (*Multiset)(nil)
)

// A multiset is an array allowing non-unique items with insertion order not significant.
// Multisets do not correspond directly with either Terraform Lists (insertion order is significant) or Sets (unique items).
// Multiset Attributes are declared as Lists with a custom type.

type multisetType struct {
	basetypes.ListType
}

var (
	MultisetType = multisetType{}
)

func (t multisetType) Equal(o attr.Type) bool {
	other, ok := o.(multisetType)

	if !ok {
		return false
	}

	return t.ListType.Equal(other.ListType)
}

func (multisetType) String() string {
	return "MultisetType"
}

func (t multisetType) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	var diags diag.Diagnostics
	elementType := t.ListType.ElemType

	if in.IsNull() {
		return MultisetNull(elementType), diags
	}

	if in.IsUnknown() {
		return MultisetUnknown(elementType), diags
	}

	return MultisetValue(elementType, in.Elements())
}

func (t multisetType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	listValue, ok := attrValue.(basetypes.ListValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	listValuable, diags := t.ValueFromList(ctx, listValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return listValuable, nil
}

func (multisetType) ValueType(context.Context) attr.Value {
	return Multiset{}
}

type Multiset struct {
	basetypes.ListValue
}

func (v Multiset) Equal(o attr.Value) bool {
	other, ok := o.(Multiset)

	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

func (Multiset) Type(context.Context) attr.Type {
	return MultisetType
}

func (v Multiset) ListSemanticEquals(ctx context.Context, newValuable basetypes.ListValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(Multiset)
	if !ok {
		return false, diags
	}

	old, d := v.ToListValue(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return false, diags
	}

	new, d := newValue.ToListValue(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return false, diags
	}

	oldElems, newElems := old.Elements(), new.Elements()

	if len(oldElems) != len(newElems) {
		return false, diags
	}

	for _, newElem := range newElems {
		found := false
		for i, oldElem := range oldElems {
			if oldElem.Equal(newElem) {
				oldElems = slices.Delete(oldElems, i, i)
				found = true
				break
			}
		}
		if !found {
			return false, diags
		}
	}

	return len(oldElems) == 0, diags
}

func MultisetNull(elementType attr.Type) Multiset {
	return Multiset{ListValue: basetypes.NewListNull(elementType)}
}

func MultisetUnknown(elementType attr.Type) Multiset {
	return Multiset{ListValue: basetypes.NewListUnknown(elementType)}
}

func MultisetValue(elementType attr.Type, elements []attr.Value) (Multiset, diag.Diagnostics) {
	var diags diag.Diagnostics

	v, d := basetypes.NewListValue(elementType, elements)
	diags.Append(d...)
	if diags.HasError() {
		return MultisetUnknown(elementType), diags
	}

	return Multiset{ListValue: v}, diags
}
