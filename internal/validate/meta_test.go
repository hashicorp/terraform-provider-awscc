package validate

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/diags"
)

func TestAllValidator(t *testing.T) {
	t.Parallel()

	rootPath := tftypes.NewAttributePath().WithAttributeName("test")

	type testCase struct {
		val           tftypes.Value
		f             func(context.Context, tftypes.Value) (attr.Value, error)
		validators    []tfsdk.AttributeValidator
		expectedDiags []diag.Diagnostic
	}
	tests := map[string]testCase{
		"unknown string": {
			val: tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			f:   types.StringType.ValueFromTerraform,
			validators: []tfsdk.AttributeValidator{
				StringInSlice([]string{"alpha", "beta", "gamma"}),
				StringLenAtMost(5),
			},
		},
		"null string": {
			val: tftypes.NewValue(tftypes.String, nil),
			f:   types.StringType.ValueFromTerraform,
			validators: []tfsdk.AttributeValidator{
				StringInSlice([]string{"alpha", "beta", "gamma"}),
				StringLenAtMost(5),
			},
		},
		"valid string": {
			val: tftypes.NewValue(tftypes.String, "gamma"),
			f:   types.StringType.ValueFromTerraform,
			validators: []tfsdk.AttributeValidator{
				StringInSlice([]string{"alpha", "beta", "gamma"}),
				StringLenAtMost(5),
			},
		},
		"invalid string single match": {
			val: tftypes.NewValue(tftypes.String, "alpha"),
			f:   types.StringType.ValueFromTerraform,
			validators: []tfsdk.AttributeValidator{
				StringInSlice([]string{"alpha", "beta", "gamma"}),
				StringLenAtMost(4),
			},
			expectedDiags: []diag.Diagnostic{
				diags.NewInvalidLengthAtMostAttributeError(rootPath, 4, 5),
			},
		},
		"invalid string multiple matches": {
			val: tftypes.NewValue(tftypes.String, "delta"),
			f:   types.StringType.ValueFromTerraform,
			validators: []tfsdk.AttributeValidator{
				StringInSlice([]string{"alpha", "beta", "gamma"}),
				StringLenAtMost(4),
			},
			expectedDiags: []diag.Diagnostic{
				diags.NewInvalidLengthAtMostAttributeError(rootPath, 4, 5),
				newStringNotInSliceError(
					rootPath,
					[]string{"alpha", "beta", "gamma"},
					"delta",
				),
			},
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
			All(test.validators...).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && len(test.expectedDiags) > 0 {
				t.Fatal("expected error diagnostics, got no error")
			}

			// if response.Diagnostics.HasError() && !response.Diagnostics.Contains(test.expectedDiag) {
			// 	t.Fatalf(`expected diagnostics to contain "%s", got "%s"`, printDiagnostic(test.expectedDiag), printDiagnostics(response.Diagnostics))
			// }

			// if response.Diagnostics.HasError() && len(test.expectedDiags) == 0 {
			// 	t.Fatalf(`got unexpected error diagnostics: "%s"`, printDiagnostics(response.Diagnostics))
			// }

			if response.Diagnostics.HasError() {
				if len(test.expectedDiags) == 0 {
					t.Fatalf(`got unexpected error diagnostics: %s`, printDiagnostics(response.Diagnostics))
				} else {
					for _, d := range test.expectedDiags {
						if !response.Diagnostics.Contains(d) {
							t.Errorf(`expected diagnostics to contain "%s", got %s`, printDiagnostic(d), printDiagnostics(response.Diagnostics))
						}
					}
				}
			}
		})
	}
}
