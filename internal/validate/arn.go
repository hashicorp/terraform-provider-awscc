package validate

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// arnValidator validates that a string is an Amazon Resource Name (ARN).
type arnValidator struct {
	tfsdk.AttributeValidator
}

// Description describes the validation in plain text formatting.
func (validator arnValidator) Description(_ context.Context) string {
	return "string must be an ARN"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator arnValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator arnValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s, ok := validateString(request, response)
	if !ok {
		return
	}

	if !arn.IsARN(s) {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid format",
			"expected value to be an ARN",
		)

		return
	}
}

// ARN returns a new ARN validator.
func ARN() tfsdk.AttributeValidator {
	return arnValidator{}
}

// iamPolicyARNValidator validates that a string is a valid IAM Policy ARN.
type iamPolicyARNValidator struct {
	tfsdk.AttributeValidator
}

func (validator iamPolicyARNValidator) Description(_ context.Context) string {
	return "string must be an IAM Policy ARN"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator iamPolicyARNValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator iamPolicyARNValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	arn, ok := validateARN(request, response)
	if !ok {
		return
	}

	if arn.Service != "iam" || !strings.HasPrefix(arn.Resource, "policy/") {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid format",
			"expected an IAM policy ARN",
		)
	}
}

// IAMPolicyARN returns a new string is IAM policy ARN validator.
func IAMPolicyARN() tfsdk.AttributeValidator {
	return iamPolicyARNValidator{}
}

func validateARN(request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) (arn.ARN, bool) {
	s, ok := validateString(request, response)
	if !ok {
		return arn.ARN{}, false
	}

	arn, err := arn.Parse(s)
	if err != nil {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid format",
			"expected an IAM policy ARN",
		)

		return arn, false
	}

	return arn, true
}
