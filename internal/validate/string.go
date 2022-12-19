package validate

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/diag"
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
	s, ok := validateString(ctx, request, response)
	if !ok {
		return
	}

	// A JSON object starts with a '{'
	if s[:1] != "{" {
		response.Diagnostics.Append(ccdiag.NewInvalidValueAttributeError(
			request.Path,
			"expected value to be a valid JSON object",
		))
		return
	}

	// Use json.Unmarshal() instead of just json.Valid() to get the parsing error
	var i interface{}
	err := json.Unmarshal([]byte(s), &i)
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

func validateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) (string, bool) {
	var s types.String
	diags := tfsdk.ValueAs(ctx, request.ConfigValue, &s)

	if diags.HasError() {
		response.Diagnostics = append(response.Diagnostics, diags...)

		return "", false
	}

	if s.IsNull() || s.IsUnknown() {
		return "", false
	}

	return s.ValueString(), true
}
