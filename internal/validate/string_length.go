package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// stringLengthValidator validates that a string Attribute's length is valid.
type stringLengthValidator struct {
	tfsdk.AttributeValidator

	minLength, maxLength int
}

// Description describes the validation in plain text formatting.
func (v stringLengthValidator) Description(_ context.Context) string {
	return fmt.Sprintf("string length must be between %d and %d", v.minLength, v.maxLength)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v stringLengthValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v stringLengthValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	var l int
	s, ok := request.AttributeConfig.(types.String)

	if !ok {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value type",
			Detail:   fmt.Sprintf("received incorrect value type (%T) at path: %s", request.AttributeConfig, request.AttributePath),
		})
	}

	if s.Unknown || s.Null {
		return
	}

	l = len(s.Value)

	if l < v.minLength || l > v.maxLength {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid length",
			Detail:   fmt.Sprintf("expected length of %s to be in the range [%d, %d], got %d", request.AttributePath, v.minLength, v.maxLength, l),
		})

		return
	}
}

// StringLength returns a new string length validator.
func StringLength(minLength, maxLength int) tfsdk.AttributeValidator {
	if minLength < 0 || maxLength < 0 || minLength > maxLength {
		return nil
	}

	return stringLengthValidator{
		minLength: minLength,
		maxLength: maxLength,
	}
}
