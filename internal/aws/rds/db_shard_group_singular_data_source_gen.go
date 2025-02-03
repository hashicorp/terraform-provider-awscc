// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package rds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_rds_db_shard_group", dBShardGroupDataSource)
}

// dBShardGroupDataSource returns the Terraform awscc_rds_db_shard_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::RDS::DBShardGroup resource.
func dBShardGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ComputeRedundancy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether to create standby DB shard groups for the DB shard group. Valid values are the following:\n  +  0 - Creates a DB shard group without a standby DB shard group. This is the default value.\n  +  1 - Creates a DB shard group with a standby DB shard group in a different Availability Zone (AZ).\n  +  2 - Creates a DB shard group with two standby DB shard groups in two different AZs.",
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"compute_redundancy": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether to create standby DB shard groups for the DB shard group. Valid values are the following:\n  +  0 - Creates a DB shard group without a standby DB shard group. This is the default value.\n  +  1 - Creates a DB shard group with a standby DB shard group in a different Availability Zone (AZ).\n  +  2 - Creates a DB shard group with two standby DB shard groups in two different AZs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DBClusterIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the primary DB cluster for the DB shard group.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"db_cluster_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the primary DB cluster for the DB shard group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DBShardGroupIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the DB shard group.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"db_shard_group_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the DB shard group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DBShardGroupResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"db_shard_group_resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaxACU
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum capacity of the DB shard group in Aurora capacity units (ACUs).",
		//	  "type": "number"
		//	}
		"max_acu": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum capacity of the DB shard group in Aurora capacity units (ACUs).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MinACU
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The minimum capacity of the DB shard group in Aurora capacity units (ACUs).",
		//	  "type": "number"
		//	}
		"min_acu": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The minimum capacity of the DB shard group in Aurora capacity units (ACUs).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PubliclyAccessible
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the DB shard group is publicly accessible.\n When the DB shard group is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB shard group's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB shard group's VPC. Access to the DB shard group is ultimately controlled by the security group it uses. That public access is not permitted if the security group assigned to the DB shard group doesn't permit it.\n When the DB shard group isn't publicly accessible, it is an internal DB shard group with a DNS name that resolves to a private IP address.\n Default: The default behavior varies depending on whether ``DBSubnetGroupName`` is specified.\n If ``DBSubnetGroupName`` isn't specified, and ``PubliclyAccessible`` isn't specified, the following applies:\n  +  If the default VPC in the target Region doesn?t have an internet gateway attached to it, the DB shard group is private.\n  +  If the default VPC in the target Region has an internet gateway attached to it, the DB shard group is public.\n  \n If ``DBSubnetGroupName`` is specified, and ``PubliclyAccessible`` isn't specified, the following applies:\n  +  If the subnets are part of a VPC that doesn?t have an internet gateway attached to it, the DB shard group is private.\n  +  If the subnets are part of a VPC that has an internet gateway attached to it, the DB shard group is public.",
		//	  "type": "boolean"
		//	}
		"publicly_accessible": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the DB shard group is publicly accessible.\n When the DB shard group is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB shard group's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB shard group's VPC. Access to the DB shard group is ultimately controlled by the security group it uses. That public access is not permitted if the security group assigned to the DB shard group doesn't permit it.\n When the DB shard group isn't publicly accessible, it is an internal DB shard group with a DNS name that resolves to a private IP address.\n Default: The default behavior varies depending on whether ``DBSubnetGroupName`` is specified.\n If ``DBSubnetGroupName`` isn't specified, and ``PubliclyAccessible`` isn't specified, the following applies:\n  +  If the default VPC in the target Region doesn?t have an internet gateway attached to it, the DB shard group is private.\n  +  If the default VPC in the target Region has an internet gateway attached to it, the DB shard group is public.\n  \n If ``DBSubnetGroupName`` is specified, and ``PubliclyAccessible`` isn't specified, the following applies:\n  +  If the subnets are part of a VPC that doesn?t have an internet gateway attached to it, the DB shard group is private.\n  +  If the subnets are part of a VPC that has an internet gateway attached to it, the DB shard group is public.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional set of key-value pairs to associate arbitrary data of your choosing with the DB shard group.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Metadata assigned to an Amazon RDS resource consisting of a key-value pair.\n For more information, see [Tagging Amazon RDS resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide* or [Tagging Amazon Aurora and Amazon RDS resources](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.html) in the *Amazon Aurora User Guide*.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An optional set of key-value pairs to associate arbitrary data of your choosing with the DB shard group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::RDS::DBShardGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::DBShardGroup").WithTerraformTypeName("awscc_rds_db_shard_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"compute_redundancy":         "ComputeRedundancy",
		"db_cluster_identifier":      "DBClusterIdentifier",
		"db_shard_group_identifier":  "DBShardGroupIdentifier",
		"db_shard_group_resource_id": "DBShardGroupResourceId",
		"endpoint":                   "Endpoint",
		"key":                        "Key",
		"max_acu":                    "MaxACU",
		"min_acu":                    "MinACU",
		"publicly_accessible":        "PubliclyAccessible",
		"tags":                       "Tags",
		"value":                      "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
