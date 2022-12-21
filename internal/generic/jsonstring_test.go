package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func TestJSONStringTypeValueFromTerraform(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		val         tftypes.Value
		expected    attr.Value
		expectError bool
	}{
		"null value": {
			val:      tftypes.NewValue(tftypes.String, nil),
			expected: JSONStringNull(),
		},
		"unknown value": {
			val:      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			expected: JSONStringUnknown(),
		},
		"empty string": {
			val:      tftypes.NewValue(tftypes.String, ""),
			expected: JSONStringValue(""),
		},
		"valid string": {
			val:      tftypes.NewValue(tftypes.String, `{"k1": 42}`),
			expected: JSONStringValue(`{"k1": 42}`),
		},
		"invalid string": {
			val:         tftypes.NewValue(tftypes.String, "not ok"),
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			val, err := JSONStringType.ValueFromTerraform(ctx, test.val)

			if err == nil && test.expectError {
				t.Fatal("expected error, got no error")
			}
			if err != nil && !test.expectError {
				t.Fatalf("got unexpected error: %s", err)
			}

			if diff := cmp.Diff(val, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestJSONStringTypeValidate(t *testing.T) {
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
		"empty string": {
			val: tftypes.NewValue(tftypes.String, ""),
		},
		"valid string": {
			val: tftypes.NewValue(tftypes.String, `{"k1": 42, "k2": ["v2", {"k3": true}]}`),
		},
		"invalid string": {
			val:         tftypes.NewValue(tftypes.String, "not ok"),
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			attributePath := path.Root("test")
			diags := JSONStringType.Validate(ctx, test.val, attributePath)

			if !diags.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if diags.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(diags))
			}
		})
	}
}

func TestJSONStringTypeAttributePlanModifier(t *testing.T) {
	t.Parallel()

	type testCase struct {
		plannedValue  basetypes.StringValuable
		currentValue  basetypes.StringValuable
		expectedValue basetypes.StringValuable
		expectError   bool
	}
	tests := map[string]testCase{

		"current null": {
			plannedValue:  JSONStringValue(`{"k1": 42}`),
			currentValue:  JSONStringNull(),
			expectedValue: JSONStringValue(`{"k1": 42}`),
		},
		"exactly equal": {
			plannedValue:  JSONStringValue(`{}`),
			currentValue:  JSONStringValue(`{}`),
			expectedValue: JSONStringValue(`{}`),
		},
		"leading and trailing whitespace": {
			plannedValue:  JSONStringValue(` {}`),
			currentValue:  JSONStringValue(`{}  `),
			expectedValue: JSONStringValue(`{}  `),
		},
		"not equal": {
			plannedValue:  JSONStringValue(`{"k1": 42}`),
			currentValue:  JSONStringValue(`{"k1": -1}`),
			expectedValue: JSONStringValue(`{"k1": 42}`),
		},
		"fields reordered": {
			plannedValue:  JSONStringValue(`{"k2": ["v2",  {"k3": true}],  "k1": 42 }`),
			currentValue:  JSONStringValue(`{"k1": 42, "k2": ["v2", {"k3": true}]}`),
			expectedValue: JSONStringValue(`{"k1": 42, "k2": ["v2", {"k3": true}]}`),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			plannedValue, diags := test.plannedValue.ToStringValue(ctx)

			if diags.HasError() {
				t.Fatal(tfresource.DiagsError(diags))
			}

			currentValue, diags := test.currentValue.ToStringValue(ctx)

			if diags.HasError() {
				t.Fatal(tfresource.DiagsError(diags))
			}

			request := planmodifier.StringRequest{
				PlanValue:  plannedValue,
				Path:       path.Root("test"),
				StateValue: currentValue,
			}
			response := planmodifier.StringResponse{}
			JSONStringType.AttributePlanModifier().PlanModifyString(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}

			if diff := cmp.Diff(response.PlanValue, test.expectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
