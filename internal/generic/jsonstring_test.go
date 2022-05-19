package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func TestJSONString(t *testing.T) {
	t.Parallel()

	type testCase struct {
		plannedValue  attr.Value
		currentValue  attr.Value
		expectedValue attr.Value
		expectError   bool
	}
	tests := map[string]testCase{
		"planned not string": {
			plannedValue: types.Int64{Value: 1},
			currentValue: types.String{Value: `{}`},
			expectError:  true,
		},
		"current not string": {
			plannedValue: types.String{Value: `{}`},
			currentValue: types.Int64{Value: 1},
			expectError:  true,
		},
		"exactly equal": {
			plannedValue:  types.String{Value: `{}`},
			currentValue:  types.String{Value: `{}`},
			expectedValue: types.String{Value: `{}`},
		},
		"leading and trailing whitespace": {
			plannedValue:  types.String{Value: ` {}`},
			currentValue:  types.String{Value: `{}  `},
			expectedValue: types.String{Value: `{}  `},
		},
		"not equal": {
			plannedValue:  types.String{Value: `{"k1": 42}`},
			currentValue:  types.String{Value: `{"k1": -1}`},
			expectedValue: types.String{Value: `{"k1": 42}`},
		},
		"fields reordered": {
			plannedValue:  types.String{Value: `{"k2": ["v2",  {"k3": true}],  "k1": 42 }`},
			currentValue:  types.String{Value: `{"k1": 42, "k2": ["v2", {"k3": true}]}`},
			expectedValue: types.String{Value: `{"k1": 42, "k2": ["v2", {"k3": true}]}`},
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := tfsdk.ModifyAttributePlanRequest{
				AttributePath:  tftypes.NewAttributePath().WithAttributeName("test"),
				AttributePlan:  test.plannedValue,
				AttributeState: test.currentValue,
			}
			response := tfsdk.ModifyAttributePlanResponse{}
			JSONString().Modify(ctx, request, &response)

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
