// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func NotNullFloat64() validator.Float64 {
	return notNullFloat64AttributeValidator{}
}

type notNullFloat64AttributeValidator struct{}

func (m notNullFloat64AttributeValidator) Description(context.Context) string {
	return "value must be configured"
}

func (m notNullFloat64AttributeValidator) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m notNullFloat64AttributeValidator) ValidateFloat64(ctx context.Context, request validator.Float64Request, response *validator.Float64Response) {
	if !request.ConfigValue.IsNull() {
		return
	}

	response.Diagnostics.AddAttributeError(
		request.Path,
		"Missing Attribute Value",
		request.Path.String()+": "+m.Description(ctx),
	)
}
