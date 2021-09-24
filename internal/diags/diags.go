package diags

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const (
	summaryInvalidValue     = "Invalid value"
	summaryInvalidValueType = "Invalid value type"
)

func NewInvalidValueAttributeError(path *tftypes.AttributePath, detail string) diag.Diagnostic {
	return diag.NewAttributeErrorDiagnostic(
		path,
		summaryInvalidValue,
		detail,
	)
}

func NewIncorrectValueTypeAttributeError(path *tftypes.AttributePath, v attr.Value) diag.Diagnostic {
	return diag.NewAttributeErrorDiagnostic(
		path,
		summaryInvalidValueType,
		fmt.Sprintf("received incorrect value type (%T)", v),
	)
}

func NewIncorrectValueTypeResourceConfigError(t tftypes.Type) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		summaryInvalidValueType,
		fmt.Sprintf("received incorrect value type (%s)", t),
	)
}

func NewUnableToConvertValueTypeAttributeError(path *tftypes.AttributePath, err error) diag.Diagnostic {
	return diag.NewAttributeErrorDiagnostic(
		path,
		summaryInvalidValueType,
		fmt.Sprintf("unable to convert value type:\n\n%s", err),
	)
}

func NewUnableToConvertValueTypeResourceConfigError(err error) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		summaryInvalidValueType,
		fmt.Sprintf("unable to convert value type:\n\n%s", err),
	)
}
