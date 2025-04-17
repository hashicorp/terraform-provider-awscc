// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_vpc_endpoint", vPCEndpointResource)
}

// vPCEndpointResource returns the Terraform awscc_ec2_vpc_endpoint resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::VPCEndpoint resource.
func vPCEndpointResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationTimestamp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"creation_timestamp": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DnsEntries
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"dns_entries": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DnsOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Describes the DNS options for an endpoint.",
		//	  "properties": {
		//	    "DnsRecordIpType": {
		//	      "description": "The DNS records created for the endpoint.",
		//	      "enum": [
		//	        "ipv4",
		//	        "ipv6",
		//	        "dualstack",
		//	        "service-defined",
		//	        "not-specified"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "PrivateDnsOnlyForInboundResolverEndpoint": {
		//	      "description": "Indicates whether to enable private DNS only for inbound endpoints. This option is available only for services that support both gateway and interface endpoints. It routes traffic that originates from the VPC to the gateway endpoint and traffic that originates from on-premises to the interface endpoint.",
		//	      "enum": [
		//	        "OnlyInboundResolver",
		//	        "AllResolvers",
		//	        "NotSpecified"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"dns_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DnsRecordIpType
				"dns_record_ip_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The DNS records created for the endpoint.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"ipv4",
							"ipv6",
							"dualstack",
							"service-defined",
							"not-specified",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: PrivateDnsOnlyForInboundResolverEndpoint
				"private_dns_only_for_inbound_resolver_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether to enable private DNS only for inbound endpoints. This option is available only for services that support both gateway and interface endpoints. It routes traffic that originates from the VPC to the gateway endpoint and traffic that originates from on-premises to the interface endpoint.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"OnlyInboundResolver",
							"AllResolvers",
							"NotSpecified",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Describes the DNS options for an endpoint.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"vpc_endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpAddressType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The supported IP address types.",
		//	  "enum": [
		//	    "ipv4",
		//	    "ipv6",
		//	    "dualstack",
		//	    "not-specified"
		//	  ],
		//	  "type": "string"
		//	}
		"ip_address_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The supported IP address types.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ipv4",
					"ipv6",
					"dualstack",
					"not-specified",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkInterfaceIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"network_interface_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.\n For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. For example, if you have a JSON policy, you can convert it to YAML before including it in the YAML template, and CFNlong converts the policy to JSON format before calling the API actions for privatelink. Alternatively, you can include the JSON directly in the YAML, as shown in the following ``Properties`` section:\n ``Properties: VpcEndpointType: 'Interface' ServiceName: !Sub 'com.amazonaws.${AWS::Region}.logs' PolicyDocument: '{ \"Version\":\"2012-10-17\", \"Statement\": [{ \"Effect\":\"Allow\", \"Principal\":\"*\", \"Action\":[\"logs:Describe*\",\"logs:Get*\",\"logs:List*\",\"logs:FilterLogEvents\"], \"Resource\":\"*\" }] }'``",
		//	  "type": "string"
		//	}
		"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.\n For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. For example, if you have a JSON policy, you can convert it to YAML before including it in the YAML template, and CFNlong converts the policy to JSON format before calling the API actions for privatelink. Alternatively, you can include the JSON directly in the YAML, as shown in the following ``Properties`` section:\n ``Properties: VpcEndpointType: 'Interface' ServiceName: !Sub 'com.amazonaws.${AWS::Region}.logs' PolicyDocument: '{ \"Version\":\"2012-10-17\", \"Statement\": [{ \"Effect\":\"Allow\", \"Principal\":\"*\", \"Action\":[\"logs:Describe*\",\"logs:Get*\",\"logs:List*\",\"logs:FilterLogEvents\"], \"Resource\":\"*\" }] }'``",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PrivateDnsEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.\n To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.\n This property is supported only for interface endpoints.\n Default: ``false``",
		//	  "type": "boolean"
		//	}
		"private_dns_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.\n To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.\n This property is supported only for interface endpoints.\n Default: ``false``",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceConfigurationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the resource configuration.",
		//	  "type": "string"
		//	}
		"resource_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the resource configuration.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RouteTableIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IDs of the route tables. Routing is supported only for gateway endpoints.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "relationshipRef": {
		//	      "propertyPath": "/properties/RouteTableId",
		//	      "typeName": "AWS::EC2::RouteTable"
		//	    },
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"route_table_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The IDs of the route tables. Routing is supported only for gateway endpoints.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "anyOf": [
		//	      {},
		//	      {},
		//	      {}
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"security_group_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the endpoint service.",
		//	  "type": "string"
		//	}
		"service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the endpoint service.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceNetworkArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the service network.",
		//	  "type": "string"
		//	}
		"service_network_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the service network.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceRegion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"service_region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "relationshipRef": {
		//	      "propertyPath": "/properties/SubnetId",
		//	      "typeName": "AWS::EC2::Subnet"
		//	    },
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"subnet_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags to associate with the endpoint.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Describes a tag.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key of the tag.\n Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with ``aws:``.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value of the tag.\n Constraints: Tag values are case-sensitive and accept a maximum of 256 Unicode characters.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key of the tag.\n Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with ``aws:``.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of the tag.\n Constraints: Tag values are case-sensitive and accept a maximum of 256 Unicode characters.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags to associate with the endpoint.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcEndpointType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of endpoint.\n Default: Gateway",
		//	  "enum": [
		//	    "Interface",
		//	    "Gateway",
		//	    "GatewayLoadBalancer",
		//	    "ServiceNetwork",
		//	    "Resource"
		//	  ],
		//	  "type": "string"
		//	}
		"vpc_endpoint_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of endpoint.\n Default: Gateway",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"Interface",
					"Gateway",
					"GatewayLoadBalancer",
					"ServiceNetwork",
					"Resource",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the VPC.",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the VPC.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS, an MKT Partner, or another AWS accounts in your organization. For more information, see the [User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/).\n An endpoint of type ``Interface`` establishes connections between the subnets in your VPC and an AWS-service, your own service, or a service hosted by another AWS-account. With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.\n An endpoint of type ``gateway`` serves as a target for a route in your route table for traffic destined for S3 or DDB. You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to S3, see [Why can't I connect to an S3 bucket using a gateway VPC endpoint?](https://docs.aws.amazon.com/premiumsupport/knowledge-center/connect-s3-vpc-endpoint) \n An endpoint of type ``GatewayLoadBalancer`` provides private connectivity between your VPC and virtual appliances from a service provider.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VPCEndpoint").WithTerraformTypeName("awscc_ec2_vpc_endpoint")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_timestamp":    "CreationTimestamp",
		"dns_entries":           "DnsEntries",
		"dns_options":           "DnsOptions",
		"dns_record_ip_type":    "DnsRecordIpType",
		"ip_address_type":       "IpAddressType",
		"key":                   "Key",
		"network_interface_ids": "NetworkInterfaceIds",
		"policy_document":       "PolicyDocument",
		"private_dns_enabled":   "PrivateDnsEnabled",
		"private_dns_only_for_inbound_resolver_endpoint": "PrivateDnsOnlyForInboundResolverEndpoint",
		"resource_configuration_arn":                     "ResourceConfigurationArn",
		"route_table_ids":                                "RouteTableIds",
		"security_group_ids":                             "SecurityGroupIds",
		"service_name":                                   "ServiceName",
		"service_network_arn":                            "ServiceNetworkArn",
		"service_region":                                 "ServiceRegion",
		"subnet_ids":                                     "SubnetIds",
		"tags":                                           "Tags",
		"value":                                          "Value",
		"vpc_endpoint_id":                                "Id",
		"vpc_endpoint_type":                              "VpcEndpointType",
		"vpc_id":                                         "VpcId",
	})

	opts = opts.WithCreateTimeoutInMinutes(210).WithDeleteTimeoutInMinutes(210)

	opts = opts.WithUpdateTimeoutInMinutes(210)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
