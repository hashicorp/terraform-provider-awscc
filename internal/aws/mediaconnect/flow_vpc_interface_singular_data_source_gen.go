// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_mediaconnect_flow_vpc_interface", flowVpcInterfaceDataSource)
}

// flowVpcInterfaceDataSource returns the Terraform awscc_mediaconnect_flow_vpc_interface data source.
// This Terraform data source corresponds to the CloudFormation AWS::MediaConnect::FlowVpcInterface resource.
func flowVpcInterfaceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"flow_arn": {
			// Property: FlowArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Immutable and has to be a unique against other VpcInterfaces in this Flow.",
			//   "type": "string"
			// }
			Description: "Immutable and has to be a unique against other VpcInterfaces in this Flow.",
			Type:        types.StringType,
			Computed:    true,
		},
		"network_interface_ids": {
			// Property: NetworkInterfaceIds
			// CloudFormation resource type schema:
			// {
			//   "description": "IDs of the network interfaces created in customer's account by MediaConnect.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "IDs of the network interfaces created in customer's account by MediaConnect.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role Arn MediaConnect can assumes to create ENIs in customer's account.",
			//   "type": "string"
			// }
			Description: "Role Arn MediaConnect can assumes to create ENIs in customer's account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"security_group_ids": {
			// Property: SecurityGroupIds
			// CloudFormation resource type schema:
			// {
			//   "description": "Security Group IDs to be used on ENI.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "Security Group IDs to be used on ENI.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"subnet_id": {
			// Property: SubnetId
			// CloudFormation resource type schema:
			// {
			//   "description": "Subnet must be in the AZ of the Flow",
			//   "type": "string"
			// }
			Description: "Subnet must be in the AZ of the Flow",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::MediaConnect::FlowVpcInterface",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::FlowVpcInterface").WithTerraformTypeName("awscc_mediaconnect_flow_vpc_interface")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"flow_arn":              "FlowArn",
		"name":                  "Name",
		"network_interface_ids": "NetworkInterfaceIds",
		"role_arn":              "RoleArn",
		"security_group_ids":    "SecurityGroupIds",
		"subnet_id":             "SubnetId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
