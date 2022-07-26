package types

import (
	awsbase "github.com/hashicorp/aws-sdk-go-base/v2"
	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
)

type UserAgentProducts []userAgentProduct

type userAgentProduct struct {
	ProductName    tftypes.String `tfsdk:"product_name"`
	ProductVersion tftypes.String `tfsdk:"product_version"`
	Comment        tftypes.String `tfsdk:"comment"`
}

func (uap UserAgentProducts) UserAgentProducts() []awsbase.UserAgentProduct {
	results := make([]awsbase.UserAgentProduct, len(uap))
	for i, p := range uap {
		results[i] = awsbase.UserAgentProduct{
			Name:    p.ProductName.Value,
			Version: p.ProductVersion.Value,
			Comment: p.Comment.Value,
		}
	}
	return results
}
