// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_connect_quick_connect", quickConnectDataSource)
}

// quickConnectDataSource returns the Terraform awscc_connect_quick_connect data source.
// This Terraform data source corresponds to the CloudFormation AWS::Connect::QuickConnect resource.
func quickConnectDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the quick connect.",
		//	  "maxLength": 250,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the quick connect.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the Amazon Connect instance.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the Amazon Connect instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the quick connect.",
		//	  "maxLength": 127,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the quick connect.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: QuickConnectArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the quick connect.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/transfer-destination/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"quick_connect_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the quick connect.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: QuickConnectConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration settings for the quick connect.",
		//	  "properties": {
		//	    "PhoneConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The phone configuration. This is required only if QuickConnectType is PHONE_NUMBER.",
		//	      "properties": {
		//	        "PhoneNumber": {
		//	          "description": "The phone number in E.164 format.",
		//	          "pattern": "^\\+[1-9]\\d{1,14}$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "PhoneNumber"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "QueueConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The queue configuration. This is required only if QuickConnectType is QUEUE.",
		//	      "properties": {
		//	        "ContactFlowArn": {
		//	          "description": "The identifier of the contact flow.",
		//	          "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$",
		//	          "type": "string"
		//	        },
		//	        "QueueArn": {
		//	          "description": "The identifier for the queue.",
		//	          "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/queue/[-a-zA-Z0-9]*$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ContactFlowArn",
		//	        "QueueArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "QuickConnectType": {
		//	      "description": "The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).",
		//	      "enum": [
		//	        "PHONE_NUMBER",
		//	        "QUEUE",
		//	        "USER"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "UserConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The user configuration. This is required only if QuickConnectType is USER.",
		//	      "properties": {
		//	        "ContactFlowArn": {
		//	          "description": "The identifier of the contact flow.",
		//	          "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$",
		//	          "type": "string"
		//	        },
		//	        "UserArn": {
		//	          "description": "The identifier of the user.",
		//	          "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/agent/[-a-zA-Z0-9]*$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ContactFlowArn",
		//	        "UserArn"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "QuickConnectType"
		//	  ],
		//	  "type": "object"
		//	}
		"quick_connect_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PhoneConfig
				"phone_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: PhoneNumber
						"phone_number": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The phone number in E.164 format.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The phone configuration. This is required only if QuickConnectType is PHONE_NUMBER.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: QueueConfig
				"queue_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ContactFlowArn
						"contact_flow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The identifier of the contact flow.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: QueueArn
						"queue_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The identifier for the queue.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The queue configuration. This is required only if QuickConnectType is QUEUE.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: QuickConnectType
				"quick_connect_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UserConfig
				"user_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ContactFlowArn
						"contact_flow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The identifier of the contact flow.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: UserArn
						"user_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The identifier of the user.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The user configuration. This is required only if QuickConnectType is USER.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration settings for the quick connect.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tags.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
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
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more tags.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Connect::QuickConnect",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::QuickConnect").WithTerraformTypeName("awscc_connect_quick_connect")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"contact_flow_arn":     "ContactFlowArn",
		"description":          "Description",
		"instance_arn":         "InstanceArn",
		"key":                  "Key",
		"name":                 "Name",
		"phone_config":         "PhoneConfig",
		"phone_number":         "PhoneNumber",
		"queue_arn":            "QueueArn",
		"queue_config":         "QueueConfig",
		"quick_connect_arn":    "QuickConnectArn",
		"quick_connect_config": "QuickConnectConfig",
		"quick_connect_type":   "QuickConnectType",
		"tags":                 "Tags",
		"user_arn":             "UserArn",
		"user_config":          "UserConfig",
		"value":                "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}