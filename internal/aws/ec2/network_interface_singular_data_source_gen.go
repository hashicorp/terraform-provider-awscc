// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_network_interface", networkInterfaceDataSource)
}

// networkInterfaceDataSource returns the Terraform awscc_ec2_network_interface data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::NetworkInterface resource.
func networkInterfaceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the network interface.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the network interface.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GroupSet
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of security group IDs associated with this network interface.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"group_set": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of security group IDs associated with this network interface.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Network interface id.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Network interface id.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InterfaceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates the type of network interface.",
		//	  "type": "string"
		//	}
		"interface_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates the type of network interface.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv6AddressCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.",
		//	  "type": "integer"
		//	}
		"ipv_6_address_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv6Addresses
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Ipv6Address": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Ipv6Address"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"ipv_6_addresses": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Ipv6Address
					"ipv_6_address": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PrimaryPrivateIpAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Returns the primary private IP address of the network interface.",
		//	  "type": "string"
		//	}
		"primary_private_ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Returns the primary private IP address of the network interface.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PrivateIpAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the PrivateIpAddresses property. ",
		//	  "type": "string"
		//	}
		"private_ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the PrivateIpAddresses property. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PrivateIpAddresses
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Primary": {
		//	        "type": "boolean"
		//	      },
		//	      "PrivateIpAddress": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "PrivateIpAddress",
		//	      "Primary"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"private_ip_addresses": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Primary
					"primary": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: PrivateIpAddress
					"private_ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecondaryPrivateIpAddressCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses",
		//	  "type": "integer"
		//	}
		"secondary_private_ip_address_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecondaryPrivateIpAddresses
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Returns the secondary private IP addresses of the network interface.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"secondary_private_ip_addresses": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Returns the secondary private IP addresses of the network interface.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceDestCheck
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether traffic to or from the instance is validated.",
		//	  "type": "boolean"
		//	}
		"source_dest_check": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether traffic to or from the instance is validated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the subnet to associate with the network interface.",
		//	  "type": "string"
		//	}
		"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the subnet to associate with the network interface.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An arbitrary set of tags (key-value pairs) for this network interface.",
		//	  "insertionOrder": false,
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
			Description: "An arbitrary set of tags (key-value pairs) for this network interface.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::NetworkInterface",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::NetworkInterface").WithTerraformTypeName("awscc_ec2_network_interface")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":                        "Description",
		"group_set":                          "GroupSet",
		"id":                                 "Id",
		"interface_type":                     "InterfaceType",
		"ipv_6_address":                      "Ipv6Address",
		"ipv_6_address_count":                "Ipv6AddressCount",
		"ipv_6_addresses":                    "Ipv6Addresses",
		"key":                                "Key",
		"primary":                            "Primary",
		"primary_private_ip_address":         "PrimaryPrivateIpAddress",
		"private_ip_address":                 "PrivateIpAddress",
		"private_ip_addresses":               "PrivateIpAddresses",
		"secondary_private_ip_address_count": "SecondaryPrivateIpAddressCount",
		"secondary_private_ip_addresses":     "SecondaryPrivateIpAddresses",
		"source_dest_check":                  "SourceDestCheck",
		"subnet_id":                          "SubnetId",
		"tags":                               "Tags",
		"value":                              "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}