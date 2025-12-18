// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package types

import (
	awsbase "github.com/hashicorp/aws-sdk-go-base/v2"
	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
)

type userAgentProduct struct {
	Comment        tftypes.String `tfsdk:"comment"`
	ProductName    tftypes.String `tfsdk:"product_name"`
	ProductVersion tftypes.String `tfsdk:"product_version"`
}

type UserAgentProducts []userAgentProduct

func (uap UserAgentProducts) UserAgentProducts() []awsbase.UserAgentProduct {
	results := make([]awsbase.UserAgentProduct, len(uap))
	for i, p := range uap {
		results[i] = awsbase.UserAgentProduct{
			Comment: p.Comment.ValueString(),
			Name:    p.ProductName.ValueString(),
			Version: p.ProductVersion.ValueString(),
		}
	}
	return results
}
