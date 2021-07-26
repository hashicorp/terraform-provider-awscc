package types

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type nestedAttributes map[string]schema.Attribute

func (n nestedAttributes) GetAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute(n)
}

func (n nestedAttributes) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	a, ok := step.(tftypes.AttributeName)
	if !ok {
		return nil, fmt.Errorf("can't apply %T to Attributes", step)
	}
	res, ok := n[string(a)]
	if !ok {
		return nil, fmt.Errorf("no attribute %q on Attributes", a)
	}
	return res, nil
}

// AttributeType returns an attr.Type corresponding to the nested attributes.
func (n nestedAttributes) AttributeType() attr.Type {
	attrTypes := map[string]attr.Type{}
	for name, attr := range n.GetAttributes() {
		if attr.Type != nil {
			attrTypes[name] = attr.Type
		}
		if attr.Attributes != nil {
			attrTypes[name] = attr.Attributes.AttributeType()
		}
	}
	return types.ObjectType{
		AttrTypes: attrTypes,
	}
}

// SetNestedAttributes nests `attributes` under another attribute, allowing
// multiple instances of that group of attributes to appear in the
// configuration, while requiring each group of values be unique. Minimum and
// maximum numbers of times the group can appear in the configuration can be
// set using `opts`.
func SetNestedAttributes(attributes map[string]schema.Attribute, opts SetNestedAttributesOptions) schema.NestedAttributes {
	return setNestedAttributes{
		nestedAttributes: nestedAttributes(attributes),
		min:              opts.MinItems,
		max:              opts.MaxItems,
	}
}

type setNestedAttributes struct {
	nestedAttributes

	min, max int
}

// SetNestedAttributesOptions captures additional, optional parameters for
// SetNestedAttributes.
type SetNestedAttributesOptions struct {
	MinItems int
	MaxItems int
}

func (s setNestedAttributes) GetNestingMode() schema.NestingMode {
	return schema.NestingModeSet
}

func (s setNestedAttributes) GetAttributes() map[string]schema.Attribute {
	return s.nestedAttributes
}

func (s setNestedAttributes) GetMinItems() int64 {
	return int64(s.min)
}

func (s setNestedAttributes) GetMaxItems() int64 {
	return int64(s.max)
}

// AttributeType returns an attr.Type corresponding to the nested attributes.
func (s setNestedAttributes) AttributeType() attr.Type {
	return SetType{
		ElemType: s.nestedAttributes.AttributeType(),
	}
}

func (s setNestedAttributes) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	_, ok := step.(tftypes.ElementKeyValue)
	if !ok {
		return nil, fmt.Errorf("can't apply %T to SetNestedAttributes", step)
	}
	return s.nestedAttributes, nil
}

func (s setNestedAttributes) Equal(o schema.NestedAttributes) bool {
	other, ok := o.(setNestedAttributes)
	if !ok {
		return false
	}
	if s.min != other.min {
		return false
	}
	if s.max != other.max {
		return false
	}
	if len(other.nestedAttributes) != len(s.nestedAttributes) {
		return false
	}
	for k, v := range s.nestedAttributes {
		otherV, ok := other.nestedAttributes[k]
		if !ok {
			return false
		}
		if !v.Equal(otherV) {
			return false
		}
	}
	return true
}
