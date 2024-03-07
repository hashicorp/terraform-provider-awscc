// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func TestMultisetListSemanticEquals(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	type testCase struct {
		val1, val2 cctypes.MultisetValueOf[types.String]
		equals     bool
	}
	tests := map[string]testCase{
		"both empty": {
			val1:   cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{}),
			val2:   cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{}),
			equals: true,
		},
		"first empty, second single element": {
			val1: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{}),
			val2: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("GET"),
			}),
			equals: false,
		},
		"first single element, second empty": {
			val1: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("GET"),
			}),
			val2:   cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{}),
			equals: false,
		},
		"first single element, second single element, equal": {
			val1: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("GET"),
			}),
			val2: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("GET"),
			}),
			equals: true,
		},
		"first single element, second single element, not equal": {
			val1: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("GET"),
			}),
			val2: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("HEAD"),
			}),
			equals: false,
		},
		"first two elements, second three elements": {
			val1: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("GET"),
				types.StringValue("HEAD"),
			}),
			val2: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("HEAD"),
				types.StringValue("POST"),
				types.StringValue("GET"),
			}),
			equals: false,
		},
		"first three elements, second two elements": {
			val1: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("HEAD"),
				types.StringValue("POST"),
				types.StringValue("GET"),
			}),
			val2: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("GET"),
				types.StringValue("HEAD"),
			}),
			equals: false,
		},
		"first three elements, second three elements, not equal": {
			val1: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("HEAD"),
				types.StringValue("POST"),
				types.StringValue("GET"),
			}),
			val2: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("GET"),
				types.StringValue("HEAD"),
				types.StringValue("PUT"),
			}),
			equals: false,
		},
		"first three elements, second three elements, equal, same order": {
			val1: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("HEAD"),
				types.StringValue("POST"),
				types.StringValue("GET"),
			}),
			val2: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("HEAD"),
				types.StringValue("POST"),
				types.StringValue("GET"),
			}),
			equals: true,
		},
		"first three elements, second three elements, equal, different order": {
			val1: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("HEAD"),
				types.StringValue("POST"),
				types.StringValue("GET"),
			}),
			val2: cctypes.NewMultisetValueOfMust[types.String](ctx, []attr.Value{
				types.StringValue("HEAD"),
				types.StringValue("GET"),
				types.StringValue("POST"),
			}),
			equals: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			equals, _ := test.val1.ListSemanticEquals(ctx, test.val2)

			if got, want := equals, test.equals; got != want {
				t.Errorf("ListSemanticEquals(%q, %q) = %v, want %v", test.val1, test.val2, got, want)
			}
		})
	}
}
