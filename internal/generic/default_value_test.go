// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/errs/diag"
)

func TestDefaultInt64Value(t *testing.T) {
	t.Parallel()

	type testCase struct {
		plannedValue  types.Int64
		currentValue  types.Int64
		defaultValue  int64
		expectedValue types.Int64
		expectError   bool
	}
	tests := map[string]testCase{
		"non-default non-Null string": {
			plannedValue:  types.Int64Value(3),
			currentValue:  types.Int64Value(2),
			defaultValue:  1,
			expectedValue: types.Int64Value(3),
		},
		"non-default non-Null string, current Null": {
			plannedValue:  types.Int64Value(3),
			currentValue:  types.Int64Null(),
			defaultValue:  1,
			expectedValue: types.Int64Value(3),
		},
		"non-default Null string, current Null": {
			plannedValue:  types.Int64Null(),
			currentValue:  types.Int64Value(2),
			defaultValue:  1,
			expectedValue: types.Int64Null(),
		},
		"default string": {
			plannedValue:  types.Int64Null(),
			currentValue:  types.Int64Value(1),
			defaultValue:  1,
			expectedValue: types.Int64Value(1),
		},
		"default string on create": {
			plannedValue:  types.Int64Null(),
			currentValue:  types.Int64Null(),
			defaultValue:  1,
			expectedValue: types.Int64Null(),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := planmodifier.Int64Request{
				PlanValue:  test.plannedValue,
				Path:       path.Root("test"),
				StateValue: test.currentValue,
			}
			response := planmodifier.Int64Response{}
			Int64DefaultValue(test.defaultValue).PlanModifyInt64(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", ccdiag.DiagnosticsError(response.Diagnostics))
			}

			if diff := cmp.Diff(response.PlanValue, test.expectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestDefaultStringValue(t *testing.T) {
	t.Parallel()

	type testCase struct {
		plannedValue  types.String
		currentValue  types.String
		defaultValue  string
		expectedValue types.String
		expectError   bool
	}
	tests := map[string]testCase{
		"non-default non-Null string": {
			plannedValue:  types.StringValue("gamma"),
			currentValue:  types.StringValue("beta"),
			defaultValue:  "alpha",
			expectedValue: types.StringValue("gamma"),
		},
		"non-default non-Null string, current Null": {
			plannedValue:  types.StringValue("gamma"),
			currentValue:  types.StringNull(),
			defaultValue:  "alpha",
			expectedValue: types.StringValue("gamma"),
		},
		"non-default Null string, current Null": {
			plannedValue:  types.StringNull(),
			currentValue:  types.StringValue("beta"),
			defaultValue:  "alpha",
			expectedValue: types.StringNull(),
		},
		"default string": {
			plannedValue:  types.StringNull(),
			currentValue:  types.StringValue("alpha"),
			defaultValue:  "alpha",
			expectedValue: types.StringValue("alpha"),
		},
		"default string on create": {
			plannedValue:  types.StringNull(),
			currentValue:  types.StringNull(),
			defaultValue:  "alpha",
			expectedValue: types.StringNull(),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := planmodifier.StringRequest{
				PlanValue:  test.plannedValue,
				Path:       path.Root("test"),
				StateValue: test.currentValue,
			}
			response := planmodifier.StringResponse{}
			StringDefaultValue(test.defaultValue).PlanModifyString(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", ccdiag.DiagnosticsError(response.Diagnostics))
			}

			if diff := cmp.Diff(response.PlanValue, test.expectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
