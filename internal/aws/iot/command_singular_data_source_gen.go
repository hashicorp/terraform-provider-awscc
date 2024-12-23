// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iot_command", commandDataSource)
}

// commandDataSource returns the Terraform awscc_iot_command data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoT::Command resource.
func commandDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CommandArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the command.",
		//	  "type": "string"
		//	}
		"command_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the command.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CommandId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier for the command.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"command_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier for the command.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time when the command was created.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time when the command was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Deprecated
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A flag indicating whether the command is deprecated.",
		//	  "type": "boolean"
		//	}
		"deprecated": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A flag indicating whether the command is deprecated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the command.",
		//	  "maxLength": 2028,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the command.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The display name for the command.",
		//	  "type": "string"
		//	}
		"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The display name for the command.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time when the command was last updated.",
		//	  "type": "string"
		//	}
		"last_updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time when the command was last updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MandatoryParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of mandatory parameters for the command.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "DefaultValue": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "B": {
		//	            "type": "boolean"
		//	          },
		//	          "BIN": {
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "D": {
		//	            "type": "number"
		//	          },
		//	          "I": {
		//	            "type": "integer"
		//	          },
		//	          "L": {
		//	            "maxLength": 19,
		//	            "pattern": "^-?\\d+$",
		//	            "type": "string"
		//	          },
		//	          "S": {
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "UL": {
		//	            "maxLength": 20,
		//	            "minLength": 1,
		//	            "pattern": "^[0-9]*$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Description": {
		//	        "maxLength": 2028,
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "maxLength": 192,
		//	        "minLength": 1,
		//	        "pattern": "^[.$a-zA-Z0-9_-]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "B": {
		//	            "type": "boolean"
		//	          },
		//	          "BIN": {
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "D": {
		//	            "type": "number"
		//	          },
		//	          "I": {
		//	            "type": "integer"
		//	          },
		//	          "L": {
		//	            "maxLength": 19,
		//	            "pattern": "^-?\\d+$",
		//	            "type": "string"
		//	          },
		//	          "S": {
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "UL": {
		//	            "maxLength": 20,
		//	            "minLength": 1,
		//	            "pattern": "^[0-9]*$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"mandatory_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DefaultValue
					"default_value": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: B
							"b": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: BIN
							"bin": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: D
							"d": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: I
							"i": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: L
							"l": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: S
							"s": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: UL
							"ul": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: B
							"b": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: BIN
							"bin": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: D
							"d": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: I
							"i": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: L
							"l": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: S
							"s": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: UL
							"ul": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of mandatory parameters for the command.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Namespace
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The namespace to which the command belongs.",
		//	  "enum": [
		//	    "AWS-IoT",
		//	    "AWS-IoT-FleetWise"
		//	  ],
		//	  "type": "string"
		//	}
		"namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The namespace to which the command belongs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Payload
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The payload associated with the command.",
		//	  "properties": {
		//	    "Content": {
		//	      "type": "string"
		//	    },
		//	    "ContentType": {
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"payload": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Content
				"content": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ContentType
				"content_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The payload associated with the command.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PendingDeletion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A flag indicating whether the command is pending deletion.",
		//	  "type": "boolean"
		//	}
		"pending_deletion": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A flag indicating whether the command is pending deletion.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The customer role associated with the command.",
		//	  "maxLength": 2028,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The customer role associated with the command.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags to be associated with the command.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag's key.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag's value.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's key.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's value.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags to be associated with the command.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoT::Command",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::Command").WithTerraformTypeName("awscc_iot_command")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"b":                    "B",
		"bin":                  "BIN",
		"command_arn":          "CommandArn",
		"command_id":           "CommandId",
		"content":              "Content",
		"content_type":         "ContentType",
		"created_at":           "CreatedAt",
		"d":                    "D",
		"default_value":        "DefaultValue",
		"deprecated":           "Deprecated",
		"description":          "Description",
		"display_name":         "DisplayName",
		"i":                    "I",
		"key":                  "Key",
		"l":                    "L",
		"last_updated_at":      "LastUpdatedAt",
		"mandatory_parameters": "MandatoryParameters",
		"name":                 "Name",
		"namespace":            "Namespace",
		"payload":              "Payload",
		"pending_deletion":     "PendingDeletion",
		"role_arn":             "RoleArn",
		"s":                    "S",
		"tags":                 "Tags",
		"ul":                   "UL",
		"value":                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
