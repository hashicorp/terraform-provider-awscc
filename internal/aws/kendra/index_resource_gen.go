// Code generated by generators/resource/main.go; DO NOT EDIT.

package kendra

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_kendra_index", indexResource)
}

// indexResource returns the Terraform awscc_kendra_index resource.
// This Terraform resource corresponds to the CloudFormation AWS::Kendra::Index resource.
func indexResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"capacity_units": {
			// Property: CapacityUnits
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Capacity units",
			//   "properties": {
			//     "QueryCapacityUnits": {
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "StorageCapacityUnits": {
			//       "minimum": 0,
			//       "type": "integer"
			//     }
			//   },
			//   "required": [
			//     "StorageCapacityUnits",
			//     "QueryCapacityUnits"
			//   ],
			//   "type": "object"
			// }
			Description: "Capacity units",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"query_capacity_units": {
						// Property: QueryCapacityUnits
						Type:     types.Int64Type,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(0),
						},
					},
					"storage_capacity_units": {
						// Property: StorageCapacityUnits
						Type:     types.Int64Type,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(0),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description for the index",
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Description: "A description for the index",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(1000),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"document_metadata_configurations": {
			// Property: DocumentMetadataConfigurations
			// CloudFormation resource type schema:
			// {
			//   "description": "Document metadata configurations",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "maxLength": 30,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Relevance": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Duration": {
			//             "maxLength": 10,
			//             "minLength": 1,
			//             "pattern": "[0-9]+[s]",
			//             "type": "string"
			//           },
			//           "Freshness": {
			//             "type": "boolean"
			//           },
			//           "Importance": {
			//             "maximum": 10,
			//             "minimum": 1,
			//             "type": "integer"
			//           },
			//           "RankOrder": {
			//             "enum": [
			//               "ASCENDING",
			//               "DESCENDING"
			//             ],
			//             "type": "string"
			//           },
			//           "ValueImportanceItems": {
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Key": {
			//                   "maxLength": 50,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "maximum": 10,
			//                   "minimum": 1,
			//                   "type": "integer"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Search": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Displayable": {
			//             "type": "boolean"
			//           },
			//           "Facetable": {
			//             "type": "boolean"
			//           },
			//           "Searchable": {
			//             "type": "boolean"
			//           },
			//           "Sortable": {
			//             "type": "boolean"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Type": {
			//         "enum": [
			//           "STRING_VALUE",
			//           "STRING_LIST_VALUE",
			//           "LONG_VALUE",
			//           "DATE_VALUE"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Name",
			//       "Type"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 500,
			//   "type": "array"
			// }
			Description: "Document metadata configurations",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 30),
						},
					},
					"relevance": {
						// Property: Relevance
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"duration": {
									// Property: Duration
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 10),
										validate.StringMatch(regexp.MustCompile("[0-9]+[s]"), ""),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"freshness": {
									// Property: Freshness
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"importance": {
									// Property: Importance
									Type:     types.Int64Type,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(1, 10),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"rank_order": {
									// Property: RankOrder
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"ASCENDING",
											"DESCENDING",
										}),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"value_importance_items": {
									// Property: ValueImportanceItems
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Type:     types.StringType,
												Optional: true,
												Computed: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 50),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"value": {
												// Property: Value
												Type:     types.Int64Type,
												Optional: true,
												Computed: true,
												Validators: []tfsdk.AttributeValidator{
													validate.IntBetween(1, 10),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"search": {
						// Property: Search
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"displayable": {
									// Property: Displayable
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"facetable": {
									// Property: Facetable
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"searchable": {
									// Property: Searchable
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"sortable": {
									// Property: Sortable
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"STRING_VALUE",
								"STRING_LIST_VALUE",
								"LONG_VALUE",
								"DATE_VALUE",
							}),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(500),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"edition": {
			// Property: Edition
			// CloudFormation resource type schema:
			// {
			//   "description": "Edition of index",
			//   "enum": [
			//     "DEVELOPER_EDITION",
			//     "ENTERPRISE_EDITION"
			//   ],
			//   "type": "string"
			// }
			Description: "Edition of index",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"DEVELOPER_EDITION",
					"ENTERPRISE_EDITION",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique ID of index",
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "type": "string"
			// }
			Description: "Unique ID of index",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of index",
			//   "maxLength": 1000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of index",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1000),
			},
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role Arn",
			//   "maxLength": 1284,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Role Arn",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1284),
			},
		},
		"server_side_encryption_configuration": {
			// Property: ServerSideEncryptionConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Server side encryption configuration",
			//   "properties": {
			//     "KmsKeyId": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Server side encryption configuration",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"kms_key_id": {
						// Property: KmsKeyId
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags for labeling the index",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A label for tagging Kendra resources",
			//     "properties": {
			//       "Key": {
			//         "description": "A string used to identify this tag",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "A string containing the value for the tag",
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
			//   "type": "array"
			// }
			Description: "Tags for labeling the index",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A string used to identify this tag",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "A string containing the value for the tag",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(200),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"user_context_policy": {
			// Property: UserContextPolicy
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ATTRIBUTE_FILTER",
			//     "USER_TOKEN"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ATTRIBUTE_FILTER",
					"USER_TOKEN",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"user_token_configurations": {
			// Property: UserTokenConfigurations
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "JsonTokenTypeConfiguration": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "GroupAttributeField": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "UserNameAttributeField": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "UserNameAttributeField",
			//           "GroupAttributeField"
			//         ],
			//         "type": "object"
			//       },
			//       "JwtTokenTypeConfiguration": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ClaimRegex": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "GroupAttributeField": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Issuer": {
			//             "maxLength": 65,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "KeyLocation": {
			//             "enum": [
			//               "URL",
			//               "SECRET_MANAGER"
			//             ],
			//             "type": "string"
			//           },
			//           "SecretManagerArn": {
			//             "description": "Role Arn",
			//             "maxLength": 1284,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "URL": {
			//             "maxLength": 2048,
			//             "minLength": 1,
			//             "pattern": "^(https?|ftp|file):\\/\\/([^\\s]*)",
			//             "type": "string"
			//           },
			//           "UserNameAttributeField": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "KeyLocation"
			//         ],
			//         "type": "object"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 1,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"json_token_type_configuration": {
						// Property: JsonTokenTypeConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"group_attribute_field": {
									// Property: GroupAttributeField
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 100),
									},
								},
								"user_name_attribute_field": {
									// Property: UserNameAttributeField
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 100),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"jwt_token_type_configuration": {
						// Property: JwtTokenTypeConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"claim_regex": {
									// Property: ClaimRegex
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 100),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"group_attribute_field": {
									// Property: GroupAttributeField
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 100),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"issuer": {
									// Property: Issuer
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 65),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"key_location": {
									// Property: KeyLocation
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"URL",
											"SECRET_MANAGER",
										}),
									},
								},
								"secret_manager_arn": {
									// Property: SecretManagerArn
									Description: "Role Arn",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 1284),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"url": {
									// Property: URL
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 2048),
										validate.StringMatch(regexp.MustCompile("^(https?|ftp|file):\\/\\/([^\\s]*)"), ""),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"user_name_attribute_field": {
									// Property: UserNameAttributeField
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 100),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "A Kendra index",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Kendra::Index").WithTerraformTypeName("awscc_kendra_index")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                  "Arn",
		"capacity_units":                       "CapacityUnits",
		"claim_regex":                          "ClaimRegex",
		"description":                          "Description",
		"displayable":                          "Displayable",
		"document_metadata_configurations":     "DocumentMetadataConfigurations",
		"duration":                             "Duration",
		"edition":                              "Edition",
		"facetable":                            "Facetable",
		"freshness":                            "Freshness",
		"group_attribute_field":                "GroupAttributeField",
		"id":                                   "Id",
		"importance":                           "Importance",
		"issuer":                               "Issuer",
		"json_token_type_configuration":        "JsonTokenTypeConfiguration",
		"jwt_token_type_configuration":         "JwtTokenTypeConfiguration",
		"key":                                  "Key",
		"key_location":                         "KeyLocation",
		"kms_key_id":                           "KmsKeyId",
		"name":                                 "Name",
		"query_capacity_units":                 "QueryCapacityUnits",
		"rank_order":                           "RankOrder",
		"relevance":                            "Relevance",
		"role_arn":                             "RoleArn",
		"search":                               "Search",
		"searchable":                           "Searchable",
		"secret_manager_arn":                   "SecretManagerArn",
		"server_side_encryption_configuration": "ServerSideEncryptionConfiguration",
		"sortable":                             "Sortable",
		"storage_capacity_units":               "StorageCapacityUnits",
		"tags":                                 "Tags",
		"type":                                 "Type",
		"url":                                  "URL",
		"user_context_policy":                  "UserContextPolicy",
		"user_name_attribute_field":            "UserNameAttributeField",
		"user_token_configurations":            "UserTokenConfigurations",
		"value":                                "Value",
		"value_importance_items":               "ValueImportanceItems",
	})

	opts = opts.WithCreateTimeoutInMinutes(240).WithDeleteTimeoutInMinutes(720)

	opts = opts.WithUpdateTimeoutInMinutes(240)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
