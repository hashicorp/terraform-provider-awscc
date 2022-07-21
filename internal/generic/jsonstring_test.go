package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
			expected: JSONString{Null: true},
		},
		"unknown value": {
			val:      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			expected: JSONString{Unknown: true},
		},
		"empty string": {
			val:      tftypes.NewValue(tftypes.String, ""),
			expected: JSONString{Value: ""},
		},
		"valid string": {
			val:      tftypes.NewValue(tftypes.String, `{"k1": 42}`),
			expected: JSONString{Value: `{"k1": 42}`},
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
		plannedValue  attr.Value
		currentValue  attr.Value
		expectedValue attr.Value
		expectError   bool
	}
	tests := map[string]testCase{
		"planned not JSONString": {
			plannedValue: types.Int64{Value: 1},
			currentValue: JSONString{Value: `{}`},
			expectError:  true,
		},
		"current not JSONString": {
			plannedValue: JSONString{Value: `{}`},
			currentValue: types.Int64{Value: 1},
			expectError:  true,
		},
		"exactly equal": {
			plannedValue:  JSONString{Value: `{}`},
			currentValue:  JSONString{Value: `{}`},
			expectedValue: JSONString{Value: `{}`},
		},
		"leading and trailing whitespace": {
			plannedValue:  JSONString{Value: ` {}`},
			currentValue:  JSONString{Value: `{}  `},
			expectedValue: JSONString{Value: `{}  `},
		},
		"not equal": {
			plannedValue:  JSONString{Value: `{"k1": 42}`},
			currentValue:  JSONString{Value: `{"k1": -1}`},
			expectedValue: JSONString{Value: `{"k1": 42}`},
		},
		"fields reordered": {
			plannedValue:  JSONString{Value: `{"k2": ["v2",  {"k3": true}],  "k1": 42 }`},
			currentValue:  JSONString{Value: `{"k1": 42, "k2": ["v2", {"k3": true}]}`},
			expectedValue: JSONString{Value: `{"k1": 42, "k2": ["v2", {"k3": true}]}`},
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := tfsdk.ModifyAttributePlanRequest{
				AttributePath:  path.Root("test"),
				AttributePlan:  test.plannedValue,
				AttributeState: test.currentValue,
			}
			response := tfsdk.ModifyAttributePlanResponse{}
			JSONStringType.AttributePlanModifier().Modify(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}

			if diff := cmp.Diff(response.AttributePlan, test.expectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
