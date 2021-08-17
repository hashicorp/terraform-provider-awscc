package types

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestOrderedSetTypeValidate(t *testing.T) {
	t.Parallel()

	type testCase struct {
		set         OrderedSetType
		in          tftypes.Value
		expectError bool
	}
	tests := map[string]testCase{
		"not a list": {
			set:         OrderedSetType{types.ListType{ElemType: types.BoolType}},
			in:          tftypes.NewValue(tftypes.Bool, true),
			expectError: true,
		},

		"empty list": {
			set: OrderedSetType{types.ListType{ElemType: types.BoolType}},
			in:  tftypes.NewValue(tftypes.List{ElementType: tftypes.Bool}, []tftypes.Value{}),
		},

		"unique list": {
			set: OrderedSetType{types.ListType{ElemType: types.StringType}},
			in: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
		},

		"duplicates in list": {
			set: OrderedSetType{types.ListType{ElemType: types.StringType}},
			in: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
				tftypes.NewValue(tftypes.String, "beta"),
			}),
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			set := OrderedSetType{}
			set.ElemType = types.BoolType

			test.set.ElemType = types.ListType{ElemType: types.BoolType}
			diags := test.set.Validate(context.TODO(), test.in)

			if !DiagsHasError(diags) && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if DiagsHasError(diags) && !test.expectError {
				t.Fatal("got unexpected error")
			}
		})
	}
}
