// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_awscc_service_thing

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func AwsccServiceThingResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"attr1": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

type AwsccServiceThingModel struct {
	Attr1 types.String `tfsdk:"attr1"`
}
