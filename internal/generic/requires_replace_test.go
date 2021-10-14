package generic

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func TestComputedOptionalForceNew(t *testing.T) {
	t.Parallel()

	type testCase struct {
		configValue             attr.Value
		plannedValue            attr.Value
		currentValue            attr.Value
		expectedRequiresReplace bool
	}
	tests := map[string]testCase{
		"string no diff": {
			configValue:  types.String{Value: "alpha"},
			plannedValue: types.String{Value: "alpha"},
			currentValue: types.String{Value: "alpha"},
		},
		"string diff": {
			configValue:             types.String{Value: "alpha"},
			plannedValue:            types.String{Value: "alpha"},
			currentValue:            types.String{Value: "beta"},
			expectedRequiresReplace: true,
		},
		"string no diff null config": {
			configValue:  types.String{Null: true},
			plannedValue: types.String{Value: "alpha"},
			currentValue: types.String{Value: "alpha"},
		},
		"string diff null config": {
			configValue:  types.String{Null: true},
			plannedValue: types.String{Value: "alpha"},
			currentValue: types.String{Value: "beta"},
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := tfsdk.ModifyAttributePlanRequest{
				AttributeConfig: test.configValue,
				AttributePath:   tftypes.NewAttributePath().WithAttributeName("test"),
				AttributePlan:   test.plannedValue,
				AttributeState:  test.currentValue,
			}
			response := tfsdk.ModifyAttributePlanResponse{}
			ComputedOptionalForceNew().Modify(ctx, request, &response)

			if response.Diagnostics.HasError() {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}

			if got, wanted := response.RequiresReplace, test.expectedRequiresReplace; got != wanted {
				t.Errorf("RequiresReplace wanted: %t, got: %t", wanted, got)
			}
		})
	}
}
