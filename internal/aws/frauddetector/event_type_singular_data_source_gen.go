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
	registry.AddDataSourceFactory("awscc_frauddetector_event_type", eventTypeDataSource)
}

// eventTypeDataSource returns the Terraform awscc_frauddetector_event_type data source.
// This Terraform data source corresponds to the CloudFormation AWS::FraudDetector::EventType resource.
func eventTypeDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the event type.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the event type.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time when the event type was created.",
		//	  "type": "string"
		//	}
		"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time when the event type was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the event type.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the event type.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EntityTypes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Arn": {
		//	        "type": "string"
		//	      },
		//	      "CreatedTime": {
		//	        "description": "The time when the event type was created.",
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "The description.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Inline": {
		//	        "type": "boolean"
		//	      },
		//	      "LastUpdatedTime": {
		//	        "description": "The time when the event type was last updated.",
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "type": "string"
		//	      },
		//	      "Tags": {
		//	        "description": "Tags associated with this event type.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "maxLength": 128,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "maxLength": 256,
		//	              "minLength": 0,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Key",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "maxItems": 200,
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"entity_types": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Arn
					"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: CreatedTime
					"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was created.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The description.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Inline
					"inline": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: LastUpdatedTime
					"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was last updated.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Tags
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
						CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
						Description: "Tags associated with this event type.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
			Computed:   true,
		}, /*END ATTRIBUTE*/
		// Property: EventVariables
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Arn": {
		//	        "type": "string"
		//	      },
		//	      "CreatedTime": {
		//	        "description": "The time when the event type was created.",
		//	        "type": "string"
		//	      },
		//	      "DataSource": {
		//	        "enum": [
		//	          "EVENT"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "DataType": {
		//	        "enum": [
		//	          "STRING",
		//	          "INTEGER",
		//	          "FLOAT",
		//	          "BOOLEAN"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "DefaultValue": {
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "The description.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Inline": {
		//	        "type": "boolean"
		//	      },
		//	      "LastUpdatedTime": {
		//	        "description": "The time when the event type was last updated.",
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "type": "string"
		//	      },
		//	      "Tags": {
		//	        "description": "Tags associated with this event type.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "maxLength": 128,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "maxLength": 256,
		//	              "minLength": 0,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Key",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "maxItems": 200,
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      },
		//	      "VariableType": {
		//	        "enum": [
		//	          "AUTH_CODE",
		//	          "AVS",
		//	          "BILLING_ADDRESS_L1",
		//	          "BILLING_ADDRESS_L2",
		//	          "BILLING_CITY",
		//	          "BILLING_COUNTRY",
		//	          "BILLING_NAME",
		//	          "BILLING_PHONE",
		//	          "BILLING_STATE",
		//	          "BILLING_ZIP",
		//	          "CARD_BIN",
		//	          "CATEGORICAL",
		//	          "CURRENCY_CODE",
		//	          "EMAIL_ADDRESS",
		//	          "FINGERPRINT",
		//	          "FRAUD_LABEL",
		//	          "FREE_FORM_TEXT",
		//	          "IP_ADDRESS",
		//	          "NUMERIC",
		//	          "ORDER_ID",
		//	          "PAYMENT_TYPE",
		//	          "PHONE_NUMBER",
		//	          "PRICE",
		//	          "PRODUCT_CATEGORY",
		//	          "SHIPPING_ADDRESS_L1",
		//	          "SHIPPING_ADDRESS_L2",
		//	          "SHIPPING_CITY",
		//	          "SHIPPING_COUNTRY",
		//	          "SHIPPING_NAME",
		//	          "SHIPPING_PHONE",
		//	          "SHIPPING_STATE",
		//	          "SHIPPING_ZIP",
		//	          "USERAGENT"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"event_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Arn
					"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: CreatedTime
					"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was created.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: DataSource
					"data_source": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: DataType
					"data_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: DefaultValue
					"default_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The description.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Inline
					"inline": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: LastUpdatedTime
					"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was last updated.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Tags
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
						CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
						Description: "Tags associated with this event type.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: VariableType
					"variable_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
			Computed:   true,
		}, /*END ATTRIBUTE*/
		// Property: Labels
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Arn": {
		//	        "type": "string"
		//	      },
		//	      "CreatedTime": {
		//	        "description": "The time when the event type was created.",
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "The description.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Inline": {
		//	        "type": "boolean"
		//	      },
		//	      "LastUpdatedTime": {
		//	        "description": "The time when the event type was last updated.",
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "type": "string"
		//	      },
		//	      "Tags": {
		//	        "description": "Tags associated with this event type.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "maxLength": 128,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "maxLength": 256,
		//	              "minLength": 0,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Key",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "maxItems": 200,
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 2,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"labels": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Arn
					"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: CreatedTime
					"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was created.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The description.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Inline
					"inline": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: LastUpdatedTime
					"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was last updated.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Tags
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
						CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
						Description: "Tags associated with this event type.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
			Computed:   true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time when the event type was last updated.",
		//	  "type": "string"
		//	}
		"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time when the event type was last updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the event type",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9a-z_-]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the event type",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags associated with this event type.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "Tags associated with this event type.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::FraudDetector::EventType",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FraudDetector::EventType").WithTerraformTypeName("awscc_frauddetector_event_type")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"created_time":      "CreatedTime",
		"data_source":       "DataSource",
		"data_type":         "DataType",
		"default_value":     "DefaultValue",
		"description":       "Description",
		"entity_types":      "EntityTypes",
		"event_variables":   "EventVariables",
		"inline":            "Inline",
		"key":               "Key",
		"labels":            "Labels",
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
