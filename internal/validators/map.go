// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func NotNullMap() validator.Map {
	return notNullMapAttributeValidator{}
}

type notNullMapAttributeValidator struct{}

func (m notNullMapAttributeValidator) Description(context.Context) string {
	return "value must be configured"
}

func (m notNullMapAttributeValidator) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m notNullMapAttributeValidator) ValidateMap(ctx context.Context, request validator.MapRequest, response *validator.MapResponse) {
	if !request.ConfigValue.IsNull() {
		return
	}

	response.Diagnostics.AddAttributeError(
		request.Path,
		"Missing Attribute Value",
		request.Path.String()+": "+m.Description(ctx),
	)
}
