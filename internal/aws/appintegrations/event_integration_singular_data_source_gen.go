// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appintegrations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_appintegrations_event_integration", eventIntegrationDataSource)
}

// eventIntegrationDataSource returns the Terraform awscc_appintegrations_event_integration data source.
// This Terraform data source corresponds to the CloudFormation AWS::AppIntegrations::EventIntegration resource.
func eventIntegrationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"associations": {
			// Property: Associations
			// CloudFormation resource type schema:
			// {
			//   "description": "The associations with the event integration.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ClientAssociationMetadata": {
			//         "description": "The metadata associated with the client.",
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Key": {
			//               "description": "A key to identify the metadata.",
			//               "maxLength": 255,
			//               "minLength": 1,
			//               "pattern": ".*\\S.*",
			//               "type": "string"
			//             },
			//             "Value": {
			//               "description": "Corresponding metadata value for the key.",
			//               "maxLength": 255,
			//               "minLength": 1,
			//               "pattern": ".*\\S.*",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Key",
			//             "Value"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array"
			//       },
			//       "ClientId": {
			//         "description": "The identifier for the client that is associated with the event integration.",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "EventBridgeRuleName": {
			//         "description": "The name of the Eventbridge rule.",
			//         "maxLength": 2048,
			//         "minLength": 1,
			//         "pattern": "^[a-zA-Z0-9/\\._\\-]+$",
			//         "type": "string"
			//       },
			//       "EventIntegrationAssociationArn": {
			//         "description": "The Amazon Resource Name (ARN) for the event integration association.",
			//         "maxLength": 2048,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "EventIntegrationAssociationId": {
			//         "description": "The identifier for the event integration association.",
			//         "pattern": "[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The associations with the event integration.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"client_association_metadata": {
						// Property: ClientAssociationMetadata
						Description: "The metadata associated with the client.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Description: "A key to identify the metadata.",
									Type:        types.StringType,
									Computed:    true,
								},
								"value": {
									// Property: Value
									Description: "Corresponding metadata value for the key.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"client_id": {
						// Property: ClientId
						Description: "The identifier for the client that is associated with the event integration.",
						Type:        types.StringType,
						Computed:    true,
					},
					"event_bridge_rule_name": {
						// Property: EventBridgeRuleName
						Description: "The name of the Eventbridge rule.",
						Type:        types.StringType,
						Computed:    true,
					},
					"event_integration_association_arn": {
						// Property: EventIntegrationAssociationArn
						Description: "The Amazon Resource Name (ARN) for the event integration association.",
						Type:        types.StringType,
						Computed:    true,
					},
					"event_integration_association_id": {
						// Property: EventIntegrationAssociationId
						Description: "The identifier for the event integration association.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The event integration description.",
			//   "maxLength": 1000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The event integration description.",
			Type:        types.StringType,
			Computed:    true,
		},
		"event_bridge_bus": {
			// Property: EventBridgeBus
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Eventbridge bus for the event integration.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9/\\._\\-]+$",
			//   "type": "string"
			// }
			Description: "The Amazon Eventbridge bus for the event integration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"event_filter": {
			// Property: EventFilter
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The EventFilter (source) associated with the event integration.",
			//   "properties": {
			//     "Source": {
			//       "description": "The source of the events.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "pattern": "^aws\\.partner\\/.*$",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Source"
			//   ],
			//   "type": "object"
			// }
			Description: "The EventFilter (source) associated with the event integration.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"source": {
						// Property: Source
						Description: "The source of the events.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"event_integration_arn": {
			// Property: EventIntegrationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the event integration.",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the event integration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the event integration.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9/\\._\\-]+$",
			//   "type": "string"
			// }
			Description: "The name of the event integration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags (keys and values) associated with the event integration.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "A key to identify the tag.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Corresponding tag value for the key.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The tags (keys and values) associated with the event integration.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A key to identify the tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "Corresponding tag value for the key.",
						Type:        types.StringType,
						Computed:    true,
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
		Description: "Data Source schema for AWS::AppIntegrations::EventIntegration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppIntegrations::EventIntegration").WithTerraformTypeName("awscc_appintegrations_event_integration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"associations":                      "Associations",
		"client_association_metadata":       "ClientAssociationMetadata",
		"client_id":                         "ClientId",
		"description":                       "Description",
		"event_bridge_bus":                  "EventBridgeBus",
		"event_bridge_rule_name":            "EventBridgeRuleName",
		"event_filter":                      "EventFilter",
		"event_integration_arn":             "EventIntegrationArn",
		"event_integration_association_arn": "EventIntegrationAssociationArn",
		"event_integration_association_id":  "EventIntegrationAssociationId",
		"key":                               "Key",
		"name":                              "Name",
		"source":                            "Source",
		"tags":                              "Tags",
		"value":                             "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
