package diags

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const (
	summaryInvalidValue = "Invalid value"
)

func NewInvalidValueError(path *tftypes.AttributePath, detail string) diag.Diagnostic {
	return diag.NewAttributeErrorDiagnostic(
		path,
		summaryInvalidValue,
		detail,
	)
}
