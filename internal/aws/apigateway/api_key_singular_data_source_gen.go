// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_api_key", apiKeyDataSource)
}

// apiKeyDataSource returns the Terraform awscc_apigateway_api_key data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::ApiKey resource.
func apiKeyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: APIKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"api_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CustomerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An MKT customer identifier, when integrating with the AWS SaaS Marketplace.",
		//	  "type": "string"
		//	}
		"customer_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An MKT customer identifier, when integrating with the AWS SaaS Marketplace.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the ApiKey.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the ApiKey.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Enabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "Specifies whether the ApiKey can be used by callers.",
		//	  "type": "boolean"
		//	}
		"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the ApiKey can be used by callers.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GenerateDistinctId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether (``true``) or not (``false``) the key identifier is distinct from the created API key value. This parameter is deprecated and should not be used.",
		//	  "type": "boolean"
		//	}
		"generate_distinct_id": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether (``true``) or not (``false``) the key identifier is distinct from the created API key value. This parameter is deprecated and should not be used.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the API key. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the API key name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the API key. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the API key name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StageKeys
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "DEPRECATED FOR USAGE PLANS - Specifies stages associated with the API key.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "``StageKey`` is a property of the [AWS::ApiGateway::ApiKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html) resource that specifies the stage to associate with the API key. This association allows only clients with the key to make requests to methods in that stage.",
		//	    "properties": {
		//	      "RestApiId": {
		//	        "description": "The string identifier of the associated RestApi.",
		//	        "type": "string"
		//	      },
		//	      "StageName": {
		//	        "description": "The stage name associated with the stage key.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"stage_keys": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: RestApiId
					"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The string identifier of the associated RestApi.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: StageName
					"stage_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The stage name associated with the stage key.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "DEPRECATED FOR USAGE PLANS - Specifies stages associated with the API key.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The tag key can be up to 128 characters and must not start with ``aws:``. The tag value can be up to 256 characters.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The tag key can be up to 128 characters and must not start with ``aws:``. The tag value can be up to 256 characters.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Value
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies a value of the API key.",
		//	  "type": "string"
		//	}
		"value": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies a value of the API key.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::ApiKey",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::ApiKey").WithTerraformTypeName("awscc_apigateway_api_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_key_id":           "APIKeyId",
		"customer_id":          "CustomerId",
		"description":          "Description",
		"enabled":              "Enabled",
		"generate_distinct_id": "GenerateDistinctId",
		"key":                  "Key",
		"name":                 "Name",
		"rest_api_id":          "RestApiId",
		"stage_keys":           "StageKeys",
		"stage_name":           "StageName",
		"tags":                 "Tags",
		"value":                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
