package generic

import (
	"context"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func TestDefaultValue(t *testing.T) {
	t.Parallel()

	type testCase struct {
		plannedValue  attr.Value
		currentValue  attr.Value
		defaultValue  attr.Value
		expectedValue attr.Value
		expectError   bool
	}
	tests := map[string]testCase{
		"non-default non-Null string": {
			plannedValue:  types.String{Value: "gamma"},
			currentValue:  types.String{Value: "beta"},
			defaultValue:  types.String{Value: "alpha"},
			expectedValue: types.String{Value: "gamma"},
		},
		"non-default non-Null string, current Null": {
			plannedValue:  types.String{Value: "gamma"},
			currentValue:  types.String{Null: true},
			defaultValue:  types.String{Value: "alpha"},
			expectedValue: types.String{Value: "gamma"},
		},
		"non-default Null string, current Null": {
			plannedValue:  types.String{Null: true},
			currentValue:  types.String{Value: "beta"},
			defaultValue:  types.String{Value: "alpha"},
			expectedValue: types.String{Null: true},
		},
		"default string": {
			plannedValue:  types.String{Null: true},
			currentValue:  types.String{Value: "alpha"},
			defaultValue:  types.String{Value: "alpha"},
			expectedValue: types.String{Value: "alpha"},
		},
		"non-default non-Null number": {
			plannedValue:  types.Number{Value: big.NewFloat(30)},
			currentValue:  types.Number{Value: big.NewFloat(10)},
			defaultValue:  types.Number{Value: big.NewFloat(-10)},
			expectedValue: types.Number{Value: big.NewFloat(30)},
		},
		"non-default non-Null number, current Null": {
			plannedValue:  types.Number{Value: big.NewFloat(30)},
			currentValue:  types.Number{Null: true},
			defaultValue:  types.Number{Value: big.NewFloat(-10)},
			expectedValue: types.Number{Value: big.NewFloat(30)},
		},
		"non-default Null number, current Null": {
			plannedValue:  types.Number{Null: true},
			currentValue:  types.Number{Value: big.NewFloat(10)},
			defaultValue:  types.Number{Value: big.NewFloat(-10)},
			expectedValue: types.Number{Null: true},
		},
		"default number": {
			plannedValue:  types.Number{Null: true},
			currentValue:  types.Number{Value: big.NewFloat(-10)},
			defaultValue:  types.Number{Value: big.NewFloat(-10)},
			expectedValue: types.Number{Value: big.NewFloat(-10)},
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
			DefaultValue(test.defaultValue).Modify(ctx, request, &response)

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
