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

func TestMultiset(t *testing.T) {
	t.Parallel()

	type testCase struct {
		plannedValue  attr.Value
		currentValue  attr.Value
		expectedValue attr.Value
		expectError   bool
	}
	tests := map[string]testCase{
		"not lists": {
			plannedValue: types.String{Value: "gamma"},
			currentValue: types.String{Value: "beta"},
			expectError:  true,
		},
		"null lists": {
			plannedValue:  types.List{ElemType: types.StringType, Null: true},
			currentValue:  types.List{ElemType: types.StringType, Null: true},
			expectedValue: types.List{ElemType: types.StringType, Null: true},
		},
		"single item": {
			plannedValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
				types.String{Value: "gamma"},
			}},
			currentValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
				types.String{Value: "beta"},
			}},
			expectedValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
				types.String{Value: "gamma"},
			}},
		},
		"different lengths": {
			plannedValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
				types.String{Value: "gamma"},
				types.String{Value: "gamma"},
			}},
			currentValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
				types.String{Value: "gamma"},
			}},
			expectedValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
				types.String{Value: "gamma"},
				types.String{Value: "gamma"},
			}},
		},
		"equivalent": {
			plannedValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
				types.String{Value: "gamma"},
				types.String{Value: "gamma"},
				types.String{Value: "beta"},
			}},
			currentValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
				types.String{Value: "gamma"},
				types.String{Value: "beta"},
				types.String{Value: "gamma"},
			}},
			expectedValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
				types.String{Value: "gamma"},
				types.String{Value: "beta"},
				types.String{Value: "gamma"},
			}},
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
			Multiset().Modify(ctx, request, &response)

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
