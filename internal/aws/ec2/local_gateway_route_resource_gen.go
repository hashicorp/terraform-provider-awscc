// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_local_gateway_route", localGatewayRouteResource)
}

// localGatewayRouteResource returns the Terraform awscc_ec2_local_gateway_route resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::LocalGatewayRoute resource.
func localGatewayRouteResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DestinationCidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The CIDR block used for destination matches.",
		//	  "type": "string"
		//	}
		"destination_cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The CIDR block used for destination matches.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocalGatewayRouteTableId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the local gateway route table.",
		//	  "type": "string"
		//	}
		"local_gateway_route_table_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the local gateway route table.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocalGatewayVirtualInterfaceGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the virtual interface group.",
		//	  "type": "string"
		//	}
		"local_gateway_virtual_interface_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the virtual interface group.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkInterfaceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the network interface.",
		//	  "type": "string"
		//	}
		"network_interface_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the network interface.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of the route.",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of the route.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The route type.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The route type.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type definition for Local Gateway Route which describes a route for a local gateway route table.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::LocalGatewayRoute").WithTerraformTypeName("awscc_ec2_local_gateway_route")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"destination_cidr_block":                   "DestinationCidrBlock",
		"local_gateway_route_table_id":             "LocalGatewayRouteTableId",
		"local_gateway_virtual_interface_group_id": "LocalGatewayVirtualInterfaceGroupId",
		"network_interface_id":                     "NetworkInterfaceId",
		"state":                                    "State",
		"type":                                     "Type",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
