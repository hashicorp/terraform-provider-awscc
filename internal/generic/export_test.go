// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/identity"
)

// PrimaryIdentifier exposes a generic resource's primary identifiers to external tests.
func PrimaryIdentifier(r resource.Resource) (identity.Identifiers, bool) {
	v, ok := r.(*genericResource)
	if !ok {
		return nil, false
	}

	return v.primaryIdentifier, true
}
