// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cleanrooms

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
	registry.AddDataSourceFactory("awscc_cleanrooms_collaboration", collaborationDataSource)
}

// collaborationDataSource returns the Terraform awscc_cleanrooms_collaboration data source.
// This Terraform data source corresponds to the CloudFormation AWS::CleanRooms::Collaboration resource.
func collaborationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CollaborationIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
		//	  "type": "string"
		//	}
		"collaboration_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CreatorDisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"creator_display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CreatorMemberAbilities
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "enum": [
		//	      "CAN_QUERY",
		//	      "CAN_RECEIVE_RESULTS"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"creator_member_abilities": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatorPaymentConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "QueryCompute": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "IsResponsible": {
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "required": [
		//	        "IsResponsible"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "QueryCompute"
		//	  ],
		//	  "type": "object"
		//	}
		"creator_payment_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: QueryCompute
				"query_compute": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: IsResponsible
						"is_responsible": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DataEncryptionMetadata
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AllowCleartext": {
		//	      "type": "boolean"
		//	    },
		//	    "AllowDuplicates": {
		//	      "type": "boolean"
		//	    },
		//	    "AllowJoinsOnColumnsWithDifferentNames": {
		//	      "type": "boolean"
		//	    },
		//	    "PreserveNulls": {
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "required": [
		//	    "AllowCleartext",
		//	    "AllowDuplicates",
		//	    "AllowJoinsOnColumnsWithDifferentNames",
		//	    "PreserveNulls"
		//	  ],
		//	  "type": "object"
		//	}
		"data_encryption_metadata": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllowCleartext
				"allow_cleartext": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: AllowDuplicates
				"allow_duplicates": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: AllowJoinsOnColumnsWithDifferentNames
				"allow_joins_on_columns_with_different_names": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: PreserveNulls
				"preserve_nulls": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Members
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "AccountId": {
		//	        "maxLength": 12,
		//	        "minLength": 12,
		//	        "pattern": "^\\d+$",
		//	        "type": "string"
		//	      },
		//	      "DisplayName": {
		//	        "maxLength": 100,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "MemberAbilities": {
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "enum": [
		//	            "CAN_QUERY",
		//	            "CAN_RECEIVE_RESULTS"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "PaymentConfiguration": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "QueryCompute": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "IsResponsible": {
		//	                "type": "boolean"
		//	              }
		//	            },
		//	            "required": [
		//	              "IsResponsible"
		//	            ],
		//	            "type": "object"
		//	          }
		//	        },
		//	        "required": [
		//	          "QueryCompute"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "AccountId",
		//	      "DisplayName",
		//	      "MemberAbilities"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 9,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"members": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AccountId
					"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: DisplayName
					"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: MemberAbilities
					"member_abilities": schema.SetAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: PaymentConfiguration
					"payment_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: QueryCompute
							"query_compute": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: IsResponsible
									"is_responsible": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
			Computed:   true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: QueryLogStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "ENABLED",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"query_log_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.",
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
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
			Description: "An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CleanRooms::Collaboration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CleanRooms::Collaboration").WithTerraformTypeName("awscc_cleanrooms_collaboration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":       "AccountId",
		"allow_cleartext":  "AllowCleartext",
		"allow_duplicates": "AllowDuplicates",
		"allow_joins_on_columns_with_different_names": "AllowJoinsOnColumnsWithDifferentNames",
		"arn":                           "Arn",
		"collaboration_identifier":      "CollaborationIdentifier",
		"creator_display_name":          "CreatorDisplayName",
		"creator_member_abilities":      "CreatorMemberAbilities",
		"creator_payment_configuration": "CreatorPaymentConfiguration",
		"data_encryption_metadata":      "DataEncryptionMetadata",
		"description":                   "Description",
		"display_name":                  "DisplayName",
		"is_responsible":                "IsResponsible",
		"key":                           "Key",
		"member_abilities":              "MemberAbilities",
		"members":                       "Members",
		"name":                          "Name",
		"payment_configuration":         "PaymentConfiguration",
		"preserve_nulls":                "PreserveNulls",
		"query_compute":                 "QueryCompute",
		"query_log_status":              "QueryLogStatus",
		"tags":                          "Tags",
		"value":                         "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
