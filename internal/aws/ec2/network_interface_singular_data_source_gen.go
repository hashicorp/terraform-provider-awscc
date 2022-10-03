// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
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
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description for the network interface.",
			//   "type": "string"
			// }
			Description: "A description for the network interface.",
			Type:        types.StringType,
			Computed:    true,
		},
		"group_set": {
			// Property: GroupSet
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of security group IDs associated with this network interface.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of security group IDs associated with this network interface.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Network interface id.",
			//   "type": "string"
			// }
			Description: "Network interface id.",
			Type:        types.StringType,
			Computed:    true,
		},
		"interface_type": {
			// Property: InterfaceType
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates the type of network interface.",
			//   "type": "string"
			// }
			Description: "Indicates the type of network interface.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ipv_6_address_count": {
			// Property: Ipv6AddressCount
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.",
			//   "type": "integer"
			// }
			Description: "The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"ipv_6_addresses": {
			// Property: Ipv6Addresses
			// CloudFormation resource type schema:
			// {
			//   "description": "One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Ipv6Address": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Ipv6Address"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"ipv_6_address": {
						// Property: Ipv6Address
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"primary_private_ip_address": {
			// Property: PrimaryPrivateIpAddress
			// CloudFormation resource type schema:
			// {
			//   "description": "Returns the primary private IP address of the network interface.",
			//   "type": "string"
			// }
			Description: "Returns the primary private IP address of the network interface.",
			Type:        types.StringType,
			Computed:    true,
		},
		"private_ip_address": {
			// Property: PrivateIpAddress
			// CloudFormation resource type schema:
			// {
			//   "description": "Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the PrivateIpAddresses property. ",
			//   "type": "string"
			// }
			Description: "Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the PrivateIpAddresses property. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"private_ip_addresses": {
			// Property: PrivateIpAddresses
			// CloudFormation resource type schema:
			// {
			//   "description": "Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Primary": {
			//         "type": "boolean"
			//       },
			//       "PrivateIpAddress": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "PrivateIpAddress",
			//       "Primary"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"primary": {
						// Property: Primary
						Type:     types.BoolType,
						Computed: true,
					},
					"private_ip_address": {
						// Property: PrivateIpAddress
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"secondary_private_ip_address_count": {
			// Property: SecondaryPrivateIpAddressCount
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses",
			//   "type": "integer"
			// }
			Description: "The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"secondary_private_ip_addresses": {
			// Property: SecondaryPrivateIpAddresses
			// CloudFormation resource type schema:
			// {
			//   "description": "Returns the secondary private IP addresses of the network interface.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "Returns the secondary private IP addresses of the network interface.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"source_dest_check": {
			// Property: SourceDestCheck
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether traffic to or from the instance is validated.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether traffic to or from the instance is validated.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"subnet_id": {
			// Property: SubnetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the subnet to associate with the network interface.",
			//   "type": "string"
			// }
			Description: "The ID of the subnet to associate with the network interface.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An arbitrary set of tags (key-value pairs) for this network interface.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An arbitrary set of tags (key-value pairs) for this network interface.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::NetworkInterface",
		Version:     1,
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
