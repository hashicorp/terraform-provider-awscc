// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_subnet", subnetDataSource)
}

// subnetDataSource returns the Terraform awscc_ec2_subnet data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::Subnet resource.
func subnetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssignIpv6AddressOnCreation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"assign_ipv_6_address_on_creation": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZoneId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"availability_zone_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EnableDns64
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_dns_64": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv4NetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The netmask length of the IPv4 CIDR you want to allocate to this subnet from an Amazon VPC IP Address Manager (IPAM) pool",
		//	  "type": "integer"
		//	}
		"ipv_4_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The netmask length of the IPv4 CIDR you want to allocate to this subnet from an Amazon VPC IP Address Manager (IPAM) pool",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv6CidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"ipv_6_cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv6CidrBlocks
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"ipv_6_cidr_blocks": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv6Native
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"ipv_6_native": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv6NetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The netmask length of the IPv6 CIDR you want to allocate to this subnet from an Amazon VPC IP Address Manager (IPAM) pool",
		//	  "type": "integer"
		//	}
		"ipv_6_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The netmask length of the IPv6 CIDR you want to allocate to this subnet from an Amazon VPC IP Address Manager (IPAM) pool",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MapPublicIpOnLaunch
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"map_public_ip_on_launch": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkAclAssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"network_acl_association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: OutpostArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"outpost_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PrivateDnsNameOptionsOnLaunch
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "EnableResourceNameDnsAAAARecord": {
		//	      "type": "boolean"
		//	    },
		//	    "EnableResourceNameDnsARecord": {
		//	      "type": "boolean"
		//	    },
		//	    "HostnameType": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"private_dns_name_options_on_launch": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EnableResourceNameDnsAAAARecord
				"enable_resource_name_dns_aaaa_record": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: EnableResourceNameDnsARecord
				"enable_resource_name_dns_a_record": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: HostnameType
				"hostname_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::Subnet",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::Subnet").WithTerraformTypeName("awscc_ec2_subnet")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"assign_ipv_6_address_on_creation":     "AssignIpv6AddressOnCreation",
		"availability_zone":                    "AvailabilityZone",
		"availability_zone_id":                 "AvailabilityZoneId",
		"cidr_block":                           "CidrBlock",
		"enable_dns_64":                        "EnableDns64",
		"enable_resource_name_dns_a_record":    "EnableResourceNameDnsARecord",
		"enable_resource_name_dns_aaaa_record": "EnableResourceNameDnsAAAARecord",
		"hostname_type":                        "HostnameType",
		"ipv_4_netmask_length":                 "Ipv4NetmaskLength",
		"ipv_6_cidr_block":                     "Ipv6CidrBlock",
		"ipv_6_cidr_blocks":                    "Ipv6CidrBlocks",
		"ipv_6_native":                         "Ipv6Native",
		"ipv_6_netmask_length":                 "Ipv6NetmaskLength",
		"key":                                  "Key",
		"map_public_ip_on_launch":              "MapPublicIpOnLaunch",
		"network_acl_association_id":           "NetworkAclAssociationId",
		"outpost_arn":                          "OutpostArn",
		"private_dns_name_options_on_launch":   "PrivateDnsNameOptionsOnLaunch",
		"subnet_id":                            "SubnetId",
		"tags":                                 "Tags",
		"value":                                "Value",
		"vpc_id":                               "VpcId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
