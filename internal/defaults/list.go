// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaults

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfslices "github.com/hashicorp/terraform-provider-awscc/internal/slices"
)

// StaticListOfString returns a static list of string value default handler.
func StaticListOfString(values ...string) defaults.List {
	return listdefault.StaticValue(types.ListValueMust(types.StringType, tfslices.ApplyToAll(values, func(val string) attr.Value {
		return types.StringValue(val)
	})))
}
