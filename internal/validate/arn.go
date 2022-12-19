package validate

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/diag"
)

// arnValidator validates that a string is an Amazon Resource Name (ARN).
type arnValidator struct{}

// Description describes the validation in plain text formatting.
func (validator arnValidator) Description(_ context.Context) string {
	return "string must be an ARN"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator arnValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator arnValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	s, ok := validateString(ctx, request, response)
	if !ok {
		return
	}

	if !arn.IsARN(s) {
		response.Diagnostics.Append(ccdiag.NewInvalidFormatAttributeError(
			request.Path,
			"expected value to be an ARN",
		))

		return
	}
}

// ARN returns a new ARN validator.
func ARN() validator.String {
	return arnValidator{}
}

// iamPolicyARNValidator validates that a string is a valid IAM Policy ARN.
type iamPolicyARNValidator struct{}

func (validator iamPolicyARNValidator) Description(_ context.Context) string {
	return "string must be an IAM Policy ARN"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator iamPolicyARNValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator iamPolicyARNValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	errDiag := ccdiag.NewInvalidFormatAttributeError(
		request.Path,
		"expected an IAM policy ARN",
	)

	arn, ok := validateARN(ctx, request, response, errDiag)
	if !ok {
		return
	}

	if arn.Service != "iam" || !strings.HasPrefix(arn.Resource, "policy/") {
		response.Diagnostics.Append(errDiag)
	}
}

// IAMPolicyARN returns a new string is IAM policy ARN validator.
func IAMPolicyARN() validator.String {
	return iamPolicyARNValidator{}
}

func validateARN(ctx context.Context, request validator.StringRequest, response *validator.StringResponse, errDiag diag.Diagnostic) (arn.ARN, bool) {
	s, ok := validateString(ctx, request, response)
	if !ok {
		return arn.ARN{}, false
	}

	arn, err := arn.Parse(s)
	if err != nil {
		response.Diagnostics.Append(errDiag)

		return arn, false
	}

	return arn, true
}
