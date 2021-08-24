package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// stringLenBetweenValidator validates that a string Attribute's length is in a range.
type stringLenBetweenValidator struct {
	tfsdk.AttributeValidator

	minLength, maxLength int
}

// Description describes the validation in plain text formatting.
func (v stringLenBetweenValidator) Description(_ context.Context) string {
	return fmt.Sprintf("string length must be between %d and %d", v.minLength, v.maxLength)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v stringLenBetweenValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v stringLenBetweenValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	var l int
	s, ok := request.AttributeConfig.(types.String)

	if !ok {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value type",
			Detail:   fmt.Sprintf("received incorrect value type (%T) at path: %s", request.AttributeConfig, request.AttributePath),
		})

		return
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

// StringLenBetween returns a new string length between validator.
func StringLenBetween(minLength, maxLength int) tfsdk.AttributeValidator {
	if minLength < 0 || maxLength < 0 || minLength > maxLength {
		return nil
	}

	return stringLenBetweenValidator{
		minLength: minLength,
		maxLength: maxLength,
	}
}

// stringLenAtLeastValidator validates that a string Attribute's length is at least a certain value.
type stringLenAtLeastValidator struct {
	tfsdk.AttributeValidator

	minLength int
}

// Description describes the validation in plain text formatting.
func (v stringLenAtLeastValidator) Description(_ context.Context) string {
	return fmt.Sprintf("string length must be at least %d", v.minLength)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v stringLenAtLeastValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v stringLenAtLeastValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	var l int
	s, ok := request.AttributeConfig.(types.String)

	if !ok {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value type",
			Detail:   fmt.Sprintf("received incorrect value type (%T) at path: %s", request.AttributeConfig, request.AttributePath),
		})

		return
	}

	if s.Unknown || s.Null {
		return
	}

	l = len(s.Value)

	if l < v.minLength {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid length",
			Detail:   fmt.Sprintf("expected length of %s to be at least %d, got %d", request.AttributePath, v.minLength, l),
		})

		return
	}
}

// StringLenAtLeast returns a new string length at least validator.
func StringLenAtLeast(minLength int) tfsdk.AttributeValidator {
	if minLength < 0 {
		return nil
	}

	return stringLenAtLeastValidator{
		minLength: minLength,
	}
}
