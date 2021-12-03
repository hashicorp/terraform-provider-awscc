package validate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/diag"
)

// uniqueItemsValidator validates that an Attribute's list items have unique values.
type uniqueItemsValidator struct {
	tfsdk.AttributeValidator
}

// Description describes the validation in plain text formatting.
func (v uniqueItemsValidator) Description(_ context.Context) string {
	return "list items must have unique values"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v uniqueItemsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v uniqueItemsValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	list, ok := request.AttributeConfig.(types.List)

	if !ok {
		response.Diagnostics.Append(diag.NewIncorrectValueTypeAttributeError(
			request.AttributePath,
			request.AttributeConfig,
		))

		return
	}

	if list.Null || list.Unknown {
		return
	}

	val, err := list.ToTerraformValue(ctx)

	if err != nil {
		response.Diagnostics.Append(diag.NewUnableToObtainValueAttributeError(
			request.AttributePath,
			err,
		))

		return
	}

	vals := val.([]tftypes.Value)
	for i1, val1 := range vals {
		if !val1.IsFullyKnown() {
			continue
		}

		for i2, val2 := range vals {
			if i2 == i1 {
				continue
			}

			if !val2.IsFullyKnown() {
				continue
			}

			if val1.Equal(val2) {
				response.Diagnostics.AddAttributeError(
					request.AttributePath,
					"Duplicate value",
					"duplicate values",
				)

				return
			}
		}
	}
}

// UniqueItems returns a new unique items validator.
func UniqueItems() tfsdk.AttributeValidator {
	return uniqueItemsValidator{}
}
