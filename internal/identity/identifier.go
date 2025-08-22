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
		Description: "The ID of the AWS account",
	})
}

func (a Identifiers) AddRegionID() Identifiers {
	return append(a, Identifier{
		Name:        NameRegion,
		Description: "The AWS region where the resource is located",
	})
}
