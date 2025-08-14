// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identity

const (
	NameAccountID string = "account_id"
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

func (a Identifiers) AddRegionID(optional bool) Identifiers {
	return append(a, Identifier{
		Name:              "region",
		Description:       "The AWS region where the resource is located",
		OptionalForImport: optional,
	})
}
