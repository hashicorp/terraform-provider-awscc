// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func NotNullObject() validator.Object {
	return notNullObjectNestedAttributeValidator{}
}

type notNullObjectNestedAttributeValidator struct{}

func (notNullObjectNestedAttributeValidator) Description(context.Context) string {
	return "value must be configured"
}

func (m notNullObjectNestedAttributeValidator) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m notNullObjectNestedAttributeValidator) ValidateObject(ctx context.Context, request validator.ObjectRequest, response *validator.ObjectResponse) {
	if !request.ConfigValue.IsNull() {
		return
	}

	response.Diagnostics.AddAttributeError(
		request.Path,
		"Missing Attribute Value",
		request.Path.String()+": "+m.Description(ctx),
	)
}

// NestingDepthExceededObject returns a validator that rejects any configured value.
// It guards the placeholder attributes of a depth-limited (de-recursed) schema: the
// properties pruned at the depth boundary. Without it, configuration nested more
// deeply than the schema would be silently discarded, as Terraform drops unknown
// keys of an object-typed attribute during type conversion.
func NestingDepthExceededObject() validator.Object {
	return nestingDepthExceededObjectValidator{}
}

type nestingDepthExceededObjectValidator struct{}

func (nestingDepthExceededObjectValidator) Description(context.Context) string {
	return "value must not be configured: nesting here exceeds the maximum depth of this depth-limited schema"
}

func (m nestingDepthExceededObjectValidator) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m nestingDepthExceededObjectValidator) ValidateObject(ctx context.Context, request validator.ObjectRequest, response *validator.ObjectResponse) {
	if request.ConfigValue.IsNull() {
		return
	}

	response.Diagnostics.AddAttributeError(
		request.Path,
		"Maximum Nesting Depth Exceeded",
		request.Path.String()+": "+m.Description(ctx)+". "+
			"This resource type's schema models a recursive CloudFormation type to a fixed depth, and this configuration nests more deeply than that. "+
			"Restructure the configuration to fit within the supported depth.",
	)
}
