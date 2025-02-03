// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudfront_key_group", keyGroupDataSource)
}

// keyGroupDataSource returns the Terraform awscc_cloudfront_key_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFront::KeyGroup resource.
func keyGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"key_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KeyGroupConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The key group configuration.",
		//	  "properties": {
		//	    "Comment": {
		//	      "description": "A comment to describe the key group. The comment cannot be longer than 128 characters.",
		//	      "type": "string"
		//	    },
		//	    "Items": {
		//	      "description": "A list of the identifiers of the public keys in the key group.",
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "Name": {
		//	      "description": "A name to identify the key group.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Name",
		//	    "Items"
		//	  ],
		//	  "type": "object"
		//	}
		"key_group_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Comment
				"comment": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A comment to describe the key group. The comment cannot be longer than 128 characters.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Items
				"items": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of the identifiers of the public keys in the key group.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A name to identify the key group.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The key group configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"last_modified_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFront::KeyGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::KeyGroup").WithTerraformTypeName("awscc_cloudfront_key_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"comment":            "Comment",
		"items":              "Items",
		"key_group_config":   "KeyGroupConfig",
		"key_group_id":       "Id",
		"last_modified_time": "LastModifiedTime",
		"name":               "Name",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
