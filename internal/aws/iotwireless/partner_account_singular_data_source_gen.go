// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotwireless_partner_account", partnerAccountDataSource)
}

// partnerAccountDataSource returns the Terraform awscc_iotwireless_partner_account data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTWireless::PartnerAccount resource.
func partnerAccountDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"account_linked": {
			// Property: AccountLinked
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether the partner account is linked to the AWS account.",
			//   "type": "boolean"
			// }
			Description: "Whether the partner account is linked to the AWS account.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "PartnerAccount arn. Returned after successful create.",
			//   "type": "string"
			// }
			Description: "PartnerAccount arn. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"fingerprint": {
			// Property: Fingerprint
			// CloudFormation resource type schema:
			// {
			//   "description": "The fingerprint of the Sidewalk application server private key.",
			//   "pattern": "[a-fA-F0-9]{64}",
			//   "type": "string"
			// }
			Description: "The fingerprint of the Sidewalk application server private key.",
			Type:        types.StringType,
			Computed:    true,
		},
		"partner_account_id": {
			// Property: PartnerAccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "The partner account ID to disassociate from the AWS account",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "The partner account ID to disassociate from the AWS account",
			Type:        types.StringType,
			Computed:    true,
		},
		"partner_type": {
			// Property: PartnerType
			// CloudFormation resource type schema:
			// {
			//   "description": "The partner type",
			//   "enum": [
			//     "Sidewalk"
			//   ],
			//   "type": "string"
			// }
			Description: "The partner type",
			Type:        types.StringType,
			Computed:    true,
		},
		"sidewalk": {
			// Property: Sidewalk
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Sidewalk account credentials.",
			//   "properties": {
			//     "AppServerPrivateKey": {
			//       "maxLength": 4096,
			//       "minLength": 1,
			//       "pattern": "[a-fA-F0-9]{64}",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "AppServerPrivateKey"
			//   ],
			//   "type": "object"
			// }
			Description: "The Sidewalk account credentials.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"app_server_private_key": {
						// Property: AppServerPrivateKey
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"sidewalk_response": {
			// Property: SidewalkResponse
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Sidewalk account credentials.",
			//   "properties": {
			//     "AmazonId": {
			//       "maxLength": 2048,
			//       "type": "string"
			//     },
			//     "Arn": {
			//       "type": "string"
			//     },
			//     "Fingerprint": {
			//       "maxLength": 64,
			//       "minLength": 64,
			//       "pattern": "[a-fA-F0-9]{64}",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The Sidewalk account credentials.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"amazon_id": {
						// Property: AmazonId
						Type:     types.StringType,
						Computed: true,
					},
					"arn": {
						// Property: Arn
						Type:     types.StringType,
						Computed: true,
					},
					"fingerprint": {
						// Property: Fingerprint
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"sidewalk_update": {
			// Property: SidewalkUpdate
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Sidewalk account credentials.",
			//   "properties": {
			//     "AppServerPrivateKey": {
			//       "maxLength": 4096,
			//       "minLength": 1,
			//       "pattern": "[a-fA-F0-9]{64}",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The Sidewalk account credentials.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"app_server_private_key": {
						// Property: AppServerPrivateKey
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the destination.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the destination.",
			Attributes: tfsdk.SetNestedAttributes(
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
		Description: "Data Source schema for AWS::IoTWireless::PartnerAccount",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::PartnerAccount").WithTerraformTypeName("awscc_iotwireless_partner_account")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_linked":         "AccountLinked",
		"amazon_id":              "AmazonId",
		"app_server_private_key": "AppServerPrivateKey",
		"arn":                    "Arn",
		"fingerprint":            "Fingerprint",
		"key":                    "Key",
		"partner_account_id":     "PartnerAccountId",
		"partner_type":           "PartnerType",
		"sidewalk":               "Sidewalk",
		"sidewalk_response":      "SidewalkResponse",
		"sidewalk_update":        "SidewalkUpdate",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
