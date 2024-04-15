// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaults

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfslices "github.com/hashicorp/terraform-provider-awscc/internal/slices"
)

// StaticSetOfString returns a static set of string value default handler.
func StaticSetOfString(values ...string) defaults.Set {
	return setdefault.StaticValue(types.SetValueMust(types.StringType, tfslices.ApplyToAll(values, func(val string) attr.Value {
		return types.StringValue(val)
	})))
}
