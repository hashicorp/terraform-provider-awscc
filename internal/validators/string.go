// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func NotNullString() validator.String {
	return notNullStringAttributeValidator{}
}

type notNullStringAttributeValidator struct{}

func (m notNullStringAttributeValidator) Description(context.Context) string {
	return "value must be configured"
}

func (m notNullStringAttributeValidator) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m notNullStringAttributeValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if !request.ConfigValue.IsNull() {
		return
	}

	response.Diagnostics.AddAttributeError(
		request.Path,
		"Missing Attribute Value",
		request.Path.String()+": "+m.Description(ctx),
	)
}
