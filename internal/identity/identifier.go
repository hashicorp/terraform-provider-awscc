// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identity

const (
	NameAccountID = "account_id"
	NamesRegion   = "region"
)

// Identifier represents a generic identifier with a name and an optional description.
type Identifier struct {
	Name              string
	Description       string
	OptionalForImport bool
}

type Identifiers []Identifier

func (a Identifiers) AddAccountID() Identifiers {
	return append(a, Identifier{
		Name:              NameAccountID,
		Description:       "The ID of the AWS account",
		OptionalForImport: true,
	})
}

func (a Identifiers) AddRegionID() Identifiers {
	return append(a, Identifier{
		Name:              NamesRegion,
		Description:       "The AWS region where the resource is located",
		OptionalForImport: true,
	})
}
