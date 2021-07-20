package types

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// SetType is an AttributeType representing a list of values.
// All values must be of the same type, which the provider must specify as the ElemType property.
// Implements attr.Type.
type SetType struct {
	ElemType attr.Type
}

// ElementType returns the attr.Type elements will be created from.
func (s SetType) ElementType() attr.Type {
	return s.ElemType
}

// TerraformType returns the tftypes.Type that should be used to
// represent this type. This constrains what user input will be
// accepted and what kind of data can be set in state. The framework
// will use this to translate the Type to something Terraform can
// understand.
func (s SetType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.Set{
		ElementType: s.ElemType.TerraformType(ctx),
	}
}

// ValueFromTerraform returns a Value given a tftypes.Value. This is
// meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (s SetType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if !in.Type().Is(s.TerraformType(ctx)) {
		return nil, fmt.Errorf("can't use %s as value of Set with ElementType %T, can only use %s values", in.String(), s.ElemType, s.ElemType.TerraformType(ctx).String())
	}
	set := Set{
		ElemType: s.ElemType,
	}
	if !in.IsKnown() {
		set.Unknown = true
		return set, nil
	}
	if in.IsNull() {
		set.Null = true
		return set, nil
	}
	val := []tftypes.Value{}
	err := in.As(&val)
	if err != nil {
		return nil, err
	}
	elems := make([]attr.Value, 0, len(val))
	for _, elem := range val {
		av, err := s.ElemType.ValueFromTerraform(ctx, elem)
		if err != nil {
			return nil, err
		}
		elems = append(elems, av)
	}
	set.Elems = elems
	return set, nil
}

// Equal must return true if the Type is considered semantically equal
// to the Type passed as an argument.
func (s SetType) Equal(o attr.Type) bool {
	if s.ElemType == nil {
		return false
	}
	other, ok := o.(SetType)
	if !ok {
		return false
	}
	return s.ElemType.Equal(other.ElemType)
}

// Return the attribute or element the AttributePathStep is referring
// to, or an error if the AttributePathStep is referring to an
// attribute or element that doesn't exist.
func (s SetType) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	if _, ok := step.(tftypes.ElementKeyValue); !ok {
		return nil, fmt.Errorf("cannot apply step %T to ListType", step)
	}

	return s.ElemType, nil
}

// Set represents a set of AttributeValues, all of the same type, indicated
// by ElemType.
type Set struct {
	// Unknown will be set to true if the entire set is an unknown value.
	// If only some of the members of the set are unknown, their known or
	// unknown status will be represented however that AttributeValue
	// surfaces that information. The Set's Unknown property only tracks
	// if the number of elements in a Set is known, not whether the
	// members of the set are known.
	Unknown bool

	// Null will be set to true if the set is null, either because it was
	// omitted from the configuration, state, or plan, or because it was
	// explicitly set to null.
	Null bool

	// Elems are the members of the set.
	Elems []attr.Value

	// ElemType is the tftypes.Type of the members of the set. All
	// members of the set must be of this type.
	ElemType attr.Type
}

// ToTerraformValue returns the data contained in the AttributeValue as
// a Go type that tftypes.NewValue will accept.
func (s Set) ToTerraformValue(ctx context.Context) (interface{}, error) {
	if s.Unknown {
		return tftypes.UnknownValue, nil
	}
	if s.Null {
		return nil, nil
	}
	vals := make([]tftypes.Value, 0, len(s.Elems))
	for _, elem := range s.Elems {
		val, err := elem.ToTerraformValue(ctx)
		if err != nil {
			return nil, err
		}
		err = tftypes.ValidateValue(s.ElemType.TerraformType(ctx), val)
		if err != nil {
			return nil, fmt.Errorf("error validating terraform type: %w", err)
		}
		vals = append(vals, tftypes.NewValue(s.ElemType.TerraformType(ctx), val))
	}
	return vals, nil
}

// Equal must return true if the AttributeValue is considered
// semantically equal to the AttributeValue passed as an argument.
func (s Set) Equal(o attr.Value) bool {
	other, ok := o.(Set)
	if !ok {
		return false
	}
	if s.Unknown != other.Unknown {
		return false
	}
	if s.Null != other.Null {
		return false
	}
	if !s.ElemType.Equal(other.ElemType) {
		return false
	}
	if len(s.Elems) != len(other.Elems) {
		return false
	}
	// Element order is not significant in Sets.
	for _, sElem := range s.Elems {
		found := false
		for _, oElem := range other.Elems {
			if sElem.Equal(oElem) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
