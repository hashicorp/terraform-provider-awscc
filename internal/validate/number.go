package validate

import (
	"context"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// intBetweenValidator validates that an integer Attribute's value is in a range.
type intBetweenValidator struct {
	tfsdk.AttributeValidator

	min, max int
}

// Description describes the validation in plain text formatting.
func (v intBetweenValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be between %d and %d", v.min, v.max)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v intBetweenValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v intBetweenValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	n, ok := request.AttributeConfig.(types.Number)

	if !ok {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value type",
			Detail:   fmt.Sprintf("received incorrect value type (%T) at path: %s", request.AttributeConfig, request.AttributePath),
		})

		return
	}

	if n.Unknown || n.Null {
		return
	}

	val := n.Value

	if !val.IsInt() {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value",
			Detail:   fmt.Sprintf("invalid value (%s) at path: %s", request.AttributeConfig, request.AttributePath),
		})

		return
	}

	var i big.Int
	_, _ = val.Int(&i)

	if i := i.Int64(); i < int64(v.min) || i > int64(v.max) {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid length",
			Detail:   fmt.Sprintf("expected %s to be in the range [%d, %d], got %d", request.AttributePath, v.min, v.max, i),
		})

		return
	}
}

// IntBetween returns a new integer value between validator.
func IntBetween(min, max int) tfsdk.AttributeValidator {
	if min > max {
		return nil
	}

	return intBetweenValidator{
		min: min,
		max: max,
	}
}
