package types

import (
	awsbase "github.com/hashicorp/aws-sdk-go-base/v2"
	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
)

type userAgentProduct struct {
	Name    tftypes.String `tfsdk:"name"`
	Version tftypes.String `tfsdk:"version"`
	Comment tftypes.String `tfsdk:"comment"`
}

type UserAgentProducts []userAgentProduct

func (uap UserAgentProducts) UserAgentProducts() []awsbase.UserAgentProduct {
	results := make([]awsbase.UserAgentProduct, len(uap))
	for i, p := range uap {
		results[i] = awsbase.UserAgentProduct{
			Name:    p.Name.ValueString(),
			Version: p.Version.ValueString(),
			Comment: p.Comment.ValueString(),
		}
	}
	return results
}
