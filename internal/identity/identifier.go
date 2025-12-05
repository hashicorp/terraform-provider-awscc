// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package identity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/service/cloudcontrol"
)

const (
	NameAccountID = "account_id"
	NameRegion    = "region"
)

// Identifier represents a generic identifier with a name and an optional description.
type Identifier struct {
	Name              string
	Description       string
	RequiredForImport bool
}

type Identifiers []Identifier

// AppendDefaults appends default identifiers based on whether the resource is global or regional.
func (a Identifiers) AppendDefaults(isGlobal bool) Identifiers {
	idSpec := a.appendAccountID()

	if !isGlobal {
		idSpec = idSpec.appendRegionID()
	}

	return idSpec
}

func (a Identifiers) appendAccountID() Identifiers {
	return append(a, Identifier{
		Name:        NameAccountID,
		Description: "AWS Account where this resource is managed",
	})
}

func (a Identifiers) appendRegionID() Identifiers {
	return append(a, Identifier{
		Name:        NameRegion,
		Description: "Region where this resource is managed",
	})
}

// IdentitySetter is an interface for setting identity attributes.
type IdentitySetter interface {
	GetAttribute(context.Context, path.Path, any) diag.Diagnostics
	SetAttribute(context.Context, path.Path, any) diag.Diagnostics
}

// SetIdentity sets the identity in state using the primary identifiers.
func (a Identifiers) SetIdentity(ctx context.Context, provider cloudcontrol.Provider, state, identity IdentitySetter) diag.Diagnostics {
	var diags diag.Diagnostics
	for _, v := range a {
		if v.RequiredForImport {
			var out attr.Value
			diags.Append(state.GetAttribute(ctx, path.Root(v.Name), &out)...)
			if diags.HasError() {
				return diags
			}

			diags.Append(identity.SetAttribute(ctx, path.Root(v.Name), ValueAsString(ctx, out))...)
			if diags.HasError() {
				return diags
			}
		} else {
			switch v.Name {
			case NameAccountID:
				diags.Append(identity.SetAttribute(ctx, path.Root(NameAccountID), provider.AccountID(ctx))...)
				if diags.HasError() {
					return diags
				}
			case NameRegion:
				diags.Append(identity.SetAttribute(ctx, path.Root(NameRegion), provider.Region(ctx))...)
				if diags.HasError() {
					return diags
				}
			}
		}
	}

	return diags
}

func (a Identifiers) GetIdentity(ctx context.Context, state IdentitySetter) ([]string, types.String, types.String, diag.Diagnostics) {
	var diags diag.Diagnostics
	var identifier []string
	var accountID, region types.String
	for _, v := range a {
		if v.RequiredForImport {
			var out types.String
			diags.Append(state.GetAttribute(ctx, path.Root(v.Name), &out)...)
			if diags.HasError() {
				return nil, types.String{}, types.String{}, diags
			}

			identifier = append(identifier, out.ValueString())
		} else {
			switch v.Name {
			case NameAccountID:
				diags.Append(state.GetAttribute(ctx, path.Root(NameAccountID), &accountID)...)
				if diags.HasError() {
					return nil, types.String{}, types.String{}, diags
				}
			case NameRegion:
				diags.Append(state.GetAttribute(ctx, path.Root(NameRegion), &region)...)
				if diags.HasError() {
					return nil, types.String{}, types.String{}, diags
				}
			}
		}
	}

	return identifier, accountID, region, diags
}
