// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func TestARNTypeValueFromTerraform(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		val      tftypes.Value
		expected attr.Value
	}{
		"null value": {
			val:      tftypes.NewValue(tftypes.String, nil),
			expected: cctypes.ARNNull(),
		},
		"unknown value": {
			val:      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			expected: cctypes.ARNUnknown(),
		},
		"valid ARN": {
			val:      tftypes.NewValue(tftypes.String, "arn:aws:rds:us-east-1:123456789012:db:test"), // lintignore:AWSAT003,AWSAT005
			expected: cctypes.ARNValue("arn:aws:rds:us-east-1:123456789012:db:test"),                 // lintignore:AWSAT003,AWSAT005
		},
		"invalid ARN": {
			val:      tftypes.NewValue(tftypes.String, "not ok"),
			expected: cctypes.ARNUnknown(),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			val, err := cctypes.ARNType.ValueFromTerraform(ctx, test.val)

			if err != nil {
				t.Fatalf("got unexpected error: %s", err)
			}

			if diff := cmp.Diff(val, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestARNTypeValidate(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         tftypes.Value
		expectError bool
	}
	tests := map[string]testCase{
		"not a string": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			expectError: true,
		},
		"unknown string": {
			val: tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
		},
		"null string": {
			val: tftypes.NewValue(tftypes.String, nil),
		},
		"valid string": {
			val: tftypes.NewValue(tftypes.String, "arn:aws:rds:us-east-1:123456789012:db:test"), // lintignore:AWSAT003,AWSAT005
		},
		"invalid string": {
			val:         tftypes.NewValue(tftypes.String, "not ok"),
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			diags := cctypes.ARNType.Validate(ctx, test.val, path.Root("test"))

			if !diags.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if diags.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %#v", diags)
			}
		})
	}
}

func TestARNToStringValue(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		arn      cctypes.ARN
		expected types.String
	}{
		"value": {
			arn:      arnFromString(t, "arn:aws:rds:us-east-1:123456789012:db:test"),
			expected: types.StringValue("arn:aws:rds:us-east-1:123456789012:db:test"),
		},
		"null": {
			arn:      cctypes.ARNNull(),
			expected: types.StringNull(),
		},
		"unknown": {
			arn:      cctypes.ARNUnknown(),
			expected: types.StringUnknown(),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			s, _ := test.arn.ToStringValue(ctx)

			if !test.expected.Equal(s) {
				t.Fatalf("expected %#v to equal %#v", s, test.expected)
			}
		})
	}
}

func arnFromString(t *testing.T, s string) cctypes.ARN {
	ctx := context.Background()

	val := tftypes.NewValue(tftypes.String, s)

	attr, err := cctypes.ARNType.ValueFromTerraform(ctx, val)
	if err != nil {
		t.Fatalf("setting ARN: %s", err)
	}

	return attr.(cctypes.ARN)
}
