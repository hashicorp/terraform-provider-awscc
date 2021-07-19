package set

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Implements attr.Type.
type set struct {
}

// TerraformType returns the tftypes.Type that should be used to
// represent this type. This constrains what user input will be
// accepted and what kind of data can be set in state. The framework
// will use this to translate the Type to something Terraform can
// understand.
func (s *set) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.Set{}
}

// ValueFromTerraform returns a Value given a tftypes.Value. This is
// meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (s *set) ValueFromTerraform(ctx context.Context, val tftypes.Value) (attr.Value, error) {
	return nil, nil
}

// Equal must return true if the Type is considered semantically equal
// to the Type passed as an argument.
func (s *set) Equal(attr.Type) bool {
	return false
}

// Return the attribute or element the AttributePathStep is referring
// to, or an error if the AttributePathStep is referring to an
// attribute or element that doesn't exist.
func (s *set) ApplyTerraform5AttributePathStep(tftypes.AttributePathStep) (interface{}, error) {
	return nil, nil
}
