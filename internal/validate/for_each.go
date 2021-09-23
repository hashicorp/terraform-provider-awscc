package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// arrayForEachValidator validates that a List Attribute's contents all satisfy the included validator.
type arrayForEachValidator struct {
	tfsdk.AttributeValidator

	validator tfsdk.AttributeValidator
}

// Description describes the validation in plain text formatting.
func (validator arrayForEachValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("each element must match: %s", validator.validator.Description(ctx))
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator arrayForEachValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator arrayForEachValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	l, ok := request.AttributeConfig.(types.List)

	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value type",
			fmt.Sprintf("received incorrect value type (%T)", request.AttributeConfig),
		)

		return
	}

	for i, e := range l.Elems {
		elemPath := request.AttributePath.WithElementKeyInt(int64(i))

		elemRequest := tfsdk.ValidateAttributeRequest{
			AttributePath:   elemPath,
			AttributeConfig: e,
			Config:          request.Config,
		}

		var elemResponse tfsdk.ValidateAttributeResponse
		validator.validator.Validate(ctx, elemRequest, &elemResponse)
		response.Diagnostics.Append(elemResponse.Diagnostics...)
	}
}

// ArrayForEach returns a new array for each validator.
func ArrayForEach(validator tfsdk.AttributeValidator) tfsdk.AttributeValidator {
	if validator == nil {
		return nil
	}

	return arrayForEachValidator{
		validator: validator,
	}
}
