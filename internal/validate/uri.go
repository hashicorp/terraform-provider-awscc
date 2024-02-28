// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"context"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/errs/diag"
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
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	if _, err := url.Parse(value); err != nil {
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
