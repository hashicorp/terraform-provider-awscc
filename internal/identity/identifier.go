// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identity

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

func (a Identifiers) AddAccountID() Identifiers {
	return append(a, Identifier{
		Name:        NameAccountID,
		Description: "AWS Account where this resource is managed",
	})
}

func (a Identifiers) AddRegionID() Identifiers {
	return append(a, Identifier{
		Name:        NameRegion,
		Description: "Region where this resource is managed",
	})
}
