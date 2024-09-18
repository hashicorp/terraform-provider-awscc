// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func NotNullSet() validator.Set {
	return notNullSetNestedObjectAttributeValidator{}
}

type notNullSetNestedObjectAttributeValidator struct{}

func (notNullSetNestedObjectAttributeValidator) Description(context.Context) string {
	return "value must be configured"
}

func (m notNullSetNestedObjectAttributeValidator) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m notNullSetNestedObjectAttributeValidator) ValidateSet(ctx context.Context, request validator.SetRequest, response *validator.SetResponse) {
	if !request.ConfigValue.IsNull() {
		return
	}

	response.Diagnostics.AddAttributeError(
		request.Path,
		"Missing Attribute Value",
		request.Path.String()+": "+m.Description(ctx),
	)
}
