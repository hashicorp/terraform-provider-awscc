// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func NotNullBool() validator.Bool {
	return notNullBoolAttributeValidator{}
}

type notNullBoolAttributeValidator struct{}

func (m notNullBoolAttributeValidator) Description(context.Context) string {
	return "value must be configured"
}

func (m notNullBoolAttributeValidator) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m notNullBoolAttributeValidator) ValidateBool(ctx context.Context, request validator.BoolRequest, response *validator.BoolResponse) {
	if !request.ConfigValue.IsNull() {
		return
	}

	response.Diagnostics.AddAttributeError(
		request.Path,
		"Missing Attribute Value",
		request.Path.String()+": "+m.Description(ctx),
	)
}
