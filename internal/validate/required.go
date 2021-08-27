package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

type RequiredAttributesFunc func(as, rs []string) error

func Required(required ...string) RequiredAttributesFunc {
	return func(as, rs []string) error {
		return nil
	}
}

func AllOfRequired(fs ...RequiredAttributesFunc) RequiredAttributesFunc {
	return func(as, rs []string) error {
		return nil
	}
}

func AnyOfRequired(fs ...RequiredAttributesFunc) RequiredAttributesFunc {
	return func(as, rs []string) error {
		return nil
	}
}

func OneOfRequired(fs ...RequiredAttributesFunc) RequiredAttributesFunc {
	return func(as, rs []string) error {
		return nil
	}
}

// requiredAttributesValidator validates that required Attributes are specified.
type requiredAttributesValidator struct {
	tfsdk.AttributeValidator
}

// Description describes the validation in plain text formatting.
func (v requiredAttributesValidator) Description(_ context.Context) string {
	return "required Attributes are specified"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v requiredAttributesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v requiredAttributesValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	object, ok := request.AttributeConfig.(types.Object)

	if !ok {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value type",
			Detail:   fmt.Sprintf("received incorrect value type (%T) at path: %s", request.AttributeConfig, request.AttributePath),
		})

		return
	}

	if object.Null || object.Unknown {
		return
	}
}

// AttributeRequired returns a new required Attributes validator.
func RequiredAttributes(fs ...RequiredAttributesFunc) tfsdk.AttributeValidator {
	return requiredAttributesValidator{}
}
