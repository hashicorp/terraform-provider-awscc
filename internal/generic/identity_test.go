// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	_ "github.com/hashicorp/terraform-provider-awscc/internal/provider"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

// TestPrimaryIdentifierStatePaths checks that every generated resource reads each
// primary identifier from an attribute that exists in its schema. A path that does
// not resolve fails every Create and Read with a State Read Error.
func TestPrimaryIdentifierStatePaths(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	for name, factory := range registry.ResourceFactories() {
		// A factory error is reported by the provider itself and is out of scope here.
		r, err := factory(ctx)
		if err != nil {
			t.Logf("%s: skipped, creating resource: %s", name, err)
			continue
		}

		identifiers, ok := generic.PrimaryIdentifier(r)
		if !ok {
			continue
		}

		var response resource.SchemaResponse
		r.Schema(ctx, resource.SchemaRequest{}, &response)

		for _, v := range identifiers {
			if _, diags := response.Schema.AttributeAtPath(ctx, v.StatePath()); diags.HasError() {
				t.Errorf("%s: primary identifier %q is read from %s, which is not an attribute of the resource schema", name, v.Name, v.StatePath())
			}
		}
	}
}
