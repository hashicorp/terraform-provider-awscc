package generic

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type State struct {
	inner *tfsdk.State
}

var (
	identifierAttributePath = tftypes.NewAttributePath().WithAttributeName("identifier")
)

// GetIdentifier gets the well-known "identifier" attribute from State.
func (s *State) GetIdentifier(ctx context.Context) (string, error) {
	val, err := s.inner.GetAttribute(ctx, identifierAttributePath)

	if err != nil {
		return "", err
	}

	if val, ok := val.(types.String); ok {
		return val.Value, nil
	}

	return "", fmt.Errorf("invalid identifier type %T", val)
}

// SetIdentifier sets the well-known "identifier" attribute in State.
func (s *State) SetIdentifier(ctx context.Context, id string) error {
	return s.inner.SetAttribute(ctx, identifierAttributePath, id)
}
