package validate

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestArrayForEachValidator(t *testing.T) {
	t.Parallel()

	rootPath := tftypes.NewAttributePath().WithAttributeName("test")

	type testCase struct {
		val          tftypes.Value
		f            func(context.Context, tftypes.Value) (attr.Value, error)
		validator    tfsdk.AttributeValidator
		expectedDiag diag.Diagnostic
	}
	tests := map[string]testCase{
		"not a list": {
			val:       tftypes.NewValue(tftypes.Bool, true),
			f:         types.BoolType.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
			expectedDiag: diag.NewAttributeErrorDiagnostic(
				rootPath,
				"Invalid value type",
				"received incorrect value type (types.Bool)",
			),
		},

		"unknown list": {
			val:       tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, tftypes.UnknownValue),
			f:         types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
		"null list": {
			val:       tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, nil),
			f:         types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
		"empty list": {
			val:       tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:         types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
		"valid list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:         types.ListType{ElemType: types.StringType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
		"invalid list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
				tftypes.NewValue(tftypes.String, "delta"),
			}),
			f:         types.ListType{ElemType: types.StringType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
			expectedDiag: newStringNotInSliceError(
				rootPath.WithElementKeyInt(3),
				[]string{"alpha", "beta", "gamma"},
				"delta",
			),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			val, err := test.f(ctx, test.val)

			if err != nil {
				t.Fatalf("got unexpected error: %s", err)
			}

			request := tfsdk.ValidateAttributeRequest{
				AttributePath:   rootPath,
				AttributeConfig: val,
			}
			response := tfsdk.ValidateAttributeResponse{}
			ArrayForEach(test.validator).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectedDiag != nil {
				t.Fatal("expected error diagnostics, got no error")
			}

			if response.Diagnostics.HasError() && !response.Diagnostics.Contains(test.expectedDiag) {
				t.Fatalf(`expected diagnostics to contain "%s", got %s`, printDiagnostic(test.expectedDiag), printDiagnostics(response.Diagnostics))
			}

			if response.Diagnostics.HasError() && test.expectedDiag == nil {
				t.Fatalf(`got unexpected error diagnostics: %s`, printDiagnostics(response.Diagnostics))
			}
		})
	}
}

func printDiagnostics(ds diag.Diagnostics) string {
	s := make([]string, len(ds))
	for i, d := range ds {
		s[i] = printDiagnostic(d)
	}

	return "[" + strings.Join(s, ", ") + "]"
}

func printDiagnostic(d diag.Diagnostic) string {
	b := &strings.Builder{}
	fmt.Fprintf(b, `%s "%s": "%s"`, d.Severity().String(), d.Summary(), d.Detail())

	if v, ok := d.(diag.DiagnosticWithPath); ok {
		fmt.Fprintf(b, " at %s", v.Path().String())
	}

	return b.String()
}
