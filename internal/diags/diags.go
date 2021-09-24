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
	summaryNoTerraformValue = "No Terraform value"
	summaryInvalidLength    = "Invalid length"
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

func NewUnableToObtainValueAttributeError(path *tftypes.AttributePath, err error) diag.Diagnostic {
	return diag.NewAttributeErrorDiagnostic(
		path,
		summaryNoTerraformValue,
		fmt.Sprintf("unable to obtain Terraform value:\n\n%s", err),
	)
}

func NewInvalidLengthBetweenAttributeError(path *tftypes.AttributePath, min, max, len int) diag.Diagnostic {
	return diag.NewAttributeErrorDiagnostic(
		path,
		summaryInvalidLength,
		fmt.Sprintf("expected length to be in the range [%d, %d], got %d", min, max, len),
	)
}

func NewInvalidLengthAtLeastAttributeError(path *tftypes.AttributePath, min, len int) diag.Diagnostic {
	return diag.NewAttributeErrorDiagnostic(
		path,
		summaryInvalidLength,
		fmt.Sprintf("expected length to be at least %d, got %d", min, len),
	)
}

func NewInvalidLengthAtMostAttributeError(path *tftypes.AttributePath, max, len int) diag.Diagnostic {
	return diag.NewAttributeErrorDiagnostic(
		path,
		summaryInvalidLength,
		fmt.Sprintf("expected length to be at most %d, got %d", max, len),
	)
}

func NewInvalidFormatAttributeError(path *tftypes.AttributePath, detail string) diag.Diagnostic {
	return diag.NewAttributeErrorDiagnostic(
		path,
		summaryInvalidLength,
		detail,
	)
}
