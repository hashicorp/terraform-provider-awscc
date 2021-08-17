package types

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ attr.TypeWithValidate = OrderedSetType{}
)

type OrderedSetType struct {
	types.ListType
}

func (t OrderedSetType) Validate(ctx context.Context, in tftypes.Value) []*tfprotov6.Diagnostic {
	var diags []*tfprotov6.Diagnostic

	if !in.Type().Is(tftypes.List{}) {
		err := fmt.Errorf("expected List value, received %T with value: %v", in, in)
		return append(diags, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "OrderedSet Type Validation Error",
			Detail:   "An unexpected error was encountered trying to validate an attribute value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
		})
	}

	var vals []tftypes.Value

	if err := in.As(&vals); err != nil {
		return append(diags, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "OrderedSet Type Validation Error",
			Detail:   "An unexpected error was encountered trying to validate an attribute value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
		})
	}

	for i1, val1 := range vals {
		for i2, val2 := range vals {
			if i2 == i1 {
				continue
			}

			if val1.Equal(val2) {
				return append(diags, &tfprotov6.Diagnostic{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "Duplicate Ordered Set Elements",
					Detail:   "This attribute contains duplicate elements",
				})
			}
		}
	}

	return nil
}

func DiagsHasError(diags []*tfprotov6.Diagnostic) bool {
	for _, diag := range diags {
		if diag == nil {
			continue
		}
		if diag.Severity == tfprotov6.DiagnosticSeverityError {
			return true
		}
	}

	return false
}
