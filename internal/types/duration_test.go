// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func TestDurationTypeValueFromTerraform(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		val      tftypes.Value
		expected attr.Value
	}{
		"null value": {
			val:      tftypes.NewValue(tftypes.String, nil),
			expected: cctypes.DurationNull(),
		},
		"unknown value": {
			val:      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			expected: cctypes.DurationUnknown(),
		},
		"valid duration": {
			val:      tftypes.NewValue(tftypes.String, "2h"),
			expected: cctypes.DurationValue("2h"),
		},
		"invalid duration": {
			val:      tftypes.NewValue(tftypes.String, "not ok"),
			expected: cctypes.DurationUnknown(),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			val, err := cctypes.DurationType.ValueFromTerraform(ctx, test.val)

			if err != nil {
				t.Fatalf("got unexpected error: %s", err)
			}

			if diff := cmp.Diff(val, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestDurationValidateAttribute(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         cctypes.Duration
		expectError bool
	}
	tests := map[string]testCase{
		"unknown": {
			val: cctypes.DurationUnknown(),
		},
		"null": {
			val: cctypes.DurationNull(),
		},
		"valid": {
			val: cctypes.DurationValue("2h"),
		},
		"invalid": {
			val:         cctypes.DurationValue("not ok"),
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			req := xattr.ValidateAttributeRequest{}
			resp := xattr.ValidateAttributeResponse{}

			test.val.ValidateAttribute(ctx, req, &resp)
			if resp.Diagnostics.HasError() != test.expectError {
				t.Errorf("resp.Diagnostics.HasError() = %t, want = %t", resp.Diagnostics.HasError(), test.expectError)
			}
		})
	}
}

func TestDurationToStringValue(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		duration cctypes.Duration
		expected types.String
	}{
		"value": {
			duration: cctypes.DurationValue("2h"),
			expected: types.StringValue("2h"),
		},
		"null": {
			duration: cctypes.DurationNull(),
			expected: types.StringNull(),
		},
		"unknown": {
			duration: cctypes.DurationUnknown(),
			expected: types.StringUnknown(),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			s, _ := test.duration.ToStringValue(ctx)

			if !test.expected.Equal(s) {
				t.Fatalf("expected %#v to equal %#v", s, test.expected)
			}
		})
	}
}
