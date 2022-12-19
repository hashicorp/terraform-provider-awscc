package validate

import (
	"context"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/diag"
)

// uriValidator validates that a string is a URI.
type uriValidator struct{}

// Description describes the validation in plain text formatting.
func (validator uriValidator) Description(_ context.Context) string {
	return "string must be a URI"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator uriValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator uriValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	s, ok := validateString(ctx, request, response)
	if !ok {
		return
	}

	if _, err := url.Parse(s); err != nil {
		response.Diagnostics.Append(ccdiag.NewInvalidFormatAttributeError(
			request.Path,
			"expected value to be a URI",
		))

		return
	}
}

// ARN returns a new ARN validator.
func IsURI() validator.String {
	return uriValidator{}
}
