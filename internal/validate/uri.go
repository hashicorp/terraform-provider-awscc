package validate

import (
	"context"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/diag"
)

// uriValidator validates that a string is a URI.
type uriValidator struct {
	tfsdk.AttributeValidator
}

// Description describes the validation in plain text formatting.
func (validator uriValidator) Description(_ context.Context) string {
	return "string must be a URI"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator uriValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator uriValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s, ok := validateString(ctx, request, response)
	if !ok {
		return
	}

	if _, err := url.Parse(s); err != nil {
		response.Diagnostics.Append(ccdiag.NewInvalidFormatAttributeError(
			request.AttributePath,
			"expected value to be a URI",
		))

		return
	}
}

// ARN returns a new ARN validator.
func IsURI() tfsdk.AttributeValidator {
	return uriValidator{}
}
