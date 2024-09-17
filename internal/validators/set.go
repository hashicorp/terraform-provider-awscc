package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func NotNullSet() validator.List {
	return notNullSetNestedObjectAttributeValidator{}
}

type notNullSetNestedObjectAttributeValidator struct{}

func (notNullSetNestedObjectAttributeValidator) Description(context.Context) string {
	return "value must be configured"
}

func (m notNullSetNestedObjectAttributeValidator) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m notNullSetNestedObjectAttributeValidator) ValidateList(ctx context.Context, request validator.ListRequest, response *validator.ListResponse) {
	if !request.ConfigValue.IsNull() {
		return
	}

	response.Diagnostics.AddAttributeError(
		request.Path,
		"Missing Attribute Value",
		request.Path.String()+": "+m.Description(ctx),
	)
}
