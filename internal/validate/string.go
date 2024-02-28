// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/errs/diag"
)

// stringIsJsonObjectValidator validates that a string Attribute's value is a valid JSON object.
type stringIsJsonObjectValidator struct{}

// Description describes the validation in plain text formatting.
func (validator stringIsJsonObjectValidator) Description(_ context.Context) string {
	return "value must be a valid JSON object"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator stringIsJsonObjectValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator stringIsJsonObjectValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	// A JSON object starts with a '{'
	if value[:1] != "{" {
		response.Diagnostics.Append(ccdiag.NewInvalidValueAttributeError(
			request.Path,
			"expected value to be a valid JSON object",
		))
		return
	}

	// Use json.Unmarshal() instead of just json.Valid() to get the parsing error
	var i interface{}
	err := json.Unmarshal([]byte(value), &i)
	if err != nil {
		response.Diagnostics.Append(ccdiag.NewInvalidValueAttributeError(
			request.Path,
			fmt.Sprintf("expected value to be valid JSON: %s", err),
		))
	}
}

// StringIsJsonObject returns a new string is JSON validator.
func StringIsJsonObject() validator.String {
	return stringIsJsonObjectValidator{}
}
