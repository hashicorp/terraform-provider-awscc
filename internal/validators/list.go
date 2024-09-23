// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func NotNullList() validator.List {
	return notNullListNestedObjectAttributeValidator{}
}

type notNullListNestedObjectAttributeValidator struct{}

func (notNullListNestedObjectAttributeValidator) Description(context.Context) string {
	return "value must be configured"
}

func (m notNullListNestedObjectAttributeValidator) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m notNullListNestedObjectAttributeValidator) ValidateList(ctx context.Context, request validator.ListRequest, response *validator.ListResponse) {
	if !request.ConfigValue.IsNull() {
		return
	}

	response.Diagnostics.AddAttributeError(
		request.Path,
		"Missing Attribute Value",
		request.Path.String()+": "+m.Description(ctx),
	)
}
