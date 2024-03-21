// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceFactory("awscc_frauddetector_list", listDataSource)
}

// listDataSource returns the Terraform awscc_frauddetector_list data source.
// This Terraform data source corresponds to the CloudFormation AWS::FraudDetector::List resource.
func listDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list ARN.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The list ARN.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time when the list was created.",
		//	  "type": "string"
		//	}
		"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time when the list was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the list.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the list.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Elements
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The elements in this list.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "An element in a list.",
		//	    "maxLength": 64,
		//	    "minLength": 1,
		//	    "pattern": "^\\S+( +\\S+)*$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 100000,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"elements": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "The elements in this list.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time when the list was last updated.",
		//	  "type": "string"
		//	}
		"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time when the list was last updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the list.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9a-z_]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the list.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags associated with this list.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "Tags associated with this list.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VariableType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The variable type of the list.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[A-Z_]{1,64}$",
		//	  "type": "string"
		//	}
		"variable_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The variable type of the list.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::FraudDetector::List",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FraudDetector::List").WithTerraformTypeName("awscc_frauddetector_list")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"created_time":      "CreatedTime",
		"description":       "Description",
		"elements":          "Elements",
		"key":               "Key",
		"last_updated_time": "LastUpdatedTime",
		"name":              "Name",
		"tags":              "Tags",
		"value":             "Value",
		"variable_type":     "VariableType",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
