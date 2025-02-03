// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudfront_anycast_ip_list", anycastIpListDataSource)
}

// anycastIpListDataSource returns the Terraform awscc_cloudfront_anycast_ip_list data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFront::AnycastIpList resource.
func anycastIpListDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AnycastIpList
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AnycastIps": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "Arn": {
		//	      "type": "string"
		//	    },
		//	    "Id": {
		//	      "type": "string"
		//	    },
		//	    "IpCount": {
		//	      "type": "integer"
		//	    },
		//	    "LastModifiedTime": {
		//	      "format": "date-time",
		//	      "type": "string"
		//	    },
		//	    "Name": {
		//	      "maxLength": 64,
		//	      "minLength": 1,
		//	      "pattern": "^[a-zA-Z0-9-_]{1,64}$",
		//	      "type": "string"
		//	    },
		//	    "Status": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "AnycastIps",
		//	    "Arn",
		//	    "Id",
		//	    "IpCount",
		//	    "LastModifiedTime",
		//	    "Name",
		//	    "Status"
		//	  ],
		//	  "type": "object"
		//	}
		"anycast_ip_list": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AnycastIps
				"anycast_ips": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Arn
				"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Id
				"id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: IpCount
				"ip_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: LastModifiedTime
				"last_modified_time": schema.StringAttribute{ /*START ATTRIBUTE*/
					CustomType: timetypes.RFC3339Type{},
					Computed:   true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ETag
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"e_tag": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"anycast_ip_list_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IpCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"ip_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9-_]{1,64}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Items": {
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Key": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "maxLength": 256,
		//	            "minLength": 0,
		//	            "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Key"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Items
				"items": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFront::AnycastIpList",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::AnycastIpList").WithTerraformTypeName("awscc_cloudfront_anycast_ip_list")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"anycast_ip_list":    "AnycastIpList",
		"anycast_ip_list_id": "Id",
		"anycast_ips":        "AnycastIps",
		"arn":                "Arn",
		"e_tag":              "ETag",
		"id":                 "Id",
		"ip_count":           "IpCount",
		"items":              "Items",
		"key":                "Key",
		"last_modified_time": "LastModifiedTime",
		"name":               "Name",
		"status":             "Status",
		"tags":               "Tags",
		"value":              "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
