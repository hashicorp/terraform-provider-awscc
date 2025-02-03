// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package wisdom

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_wisdom_knowledge_base", knowledgeBaseResource)
}

// knowledgeBaseResource returns the Terraform awscc_wisdom_knowledge_base resource.
// This Terraform resource corresponds to the CloudFormation AWS::Wisdom::KnowledgeBase resource.
func knowledgeBaseResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KnowledgeBaseArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^arn:[a-z-]*?:wisdom:[a-z0-9-]*?:[0-9]{12}:[a-z-]*?/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(?:/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})?$",
		//	  "type": "string"
		//	}
		"knowledge_base_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KnowledgeBaseId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$",
		//	  "type": "string"
		//	}
		"knowledge_base_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KnowledgeBaseType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "EXTERNAL",
		//	    "CUSTOM",
		//	    "MESSAGE_TEMPLATES",
		//	    "MANAGED"
		//	  ],
		//	  "type": "string"
		//	}
		"knowledge_base_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"EXTERNAL",
					"CUSTOM",
					"MESSAGE_TEMPLATES",
					"MANAGED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RenderingConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "TemplateUri": {
		//	      "maxLength": 4096,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"rendering_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: TemplateUri
				"template_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 4096),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServerSideEncryptionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "KmsKeyId": {
		//	      "maxLength": 4096,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"server_side_encryption_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KmsKeyId
				"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 4096),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "properties": {
		//	    "AppIntegrations": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AppIntegrationArn": {
		//	          "maxLength": 2048,
		//	          "minLength": 1,
		//	          "pattern": "^arn:[a-z-]+?:[a-z-]+?:[a-z0-9-]*?:([0-9]{12})?:[a-zA-Z0-9-:/]+$",
		//	          "type": "string"
		//	        },
		//	        "ObjectFields": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 4096,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "maxItems": 100,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "AppIntegrationArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ManagedSourceConfiguration": {
		//	      "properties": {
		//	        "WebCrawlerConfiguration": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "CrawlerLimits": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "RateLimit": {
		//	                  "maximum": 3000,
		//	                  "minimum": 1,
		//	                  "type": "number"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "ExclusionFilters": {
		//	              "items": {
		//	                "maxLength": 1000,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              },
		//	              "maxItems": 25,
		//	              "minItems": 1,
		//	              "type": "array"
		//	            },
		//	            "InclusionFilters": {
		//	              "items": {
		//	                "maxLength": 1000,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              },
		//	              "maxItems": 25,
		//	              "minItems": 1,
		//	              "type": "array"
		//	            },
		//	            "Scope": {
		//	              "enum": [
		//	                "HOST_ONLY",
		//	                "SUBDOMAINS"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "UrlConfiguration": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "SeedUrls": {
		//	                  "items": {
		//	                    "additionalProperties": false,
		//	                    "properties": {
		//	                      "Url": {
		//	                        "pattern": "^https?://[A-Za-z0-9][^\\s]*$",
		//	                        "type": "string"
		//	                      }
		//	                    },
		//	                    "type": "object"
		//	                  },
		//	                  "maxItems": 100,
		//	                  "minItems": 1,
		//	                  "type": "array"
		//	                }
		//	              },
		//	              "type": "object"
		//	            }
		//	          },
		//	          "required": [
		//	            "UrlConfiguration"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"source_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AppIntegrations
				"app_integrations": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AppIntegrationArn
						"app_integration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 2048),
								stringvalidator.RegexMatches(regexp.MustCompile("^arn:[a-z-]+?:[a-z-]+?:[a-z0-9-]*?:([0-9]{12})?:[a-zA-Z0-9-:/]+$"), ""),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ObjectFields
						"object_fields": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(1, 100),
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 4096),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ManagedSourceConfiguration
				"managed_source_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: WebCrawlerConfiguration
						"web_crawler_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: CrawlerLimits
								"crawler_limits": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: RateLimit
										"rate_limit": schema.Float64Attribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											Validators: []validator.Float64{ /*START VALIDATORS*/
												float64validator.Between(1.000000, 3000.000000),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
												float64planmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: ExclusionFilters
								"exclusion_filters": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []validator.List{ /*START VALIDATORS*/
										listvalidator.SizeBetween(1, 25),
										listvalidator.ValueStringsAre(
											stringvalidator.LengthBetween(1, 1000),
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: InclusionFilters
								"inclusion_filters": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []validator.List{ /*START VALIDATORS*/
										listvalidator.SizeBetween(1, 25),
										listvalidator.ValueStringsAre(
											stringvalidator.LengthBetween(1, 1000),
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Scope
								"scope": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.OneOf(
											"HOST_ONLY",
											"SUBDOMAINS",
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: UrlConfiguration
								"url_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: SeedUrls
										"seed_urls": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
											NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: Url
													"url": schema.StringAttribute{ /*START ATTRIBUTE*/
														Optional: true,
														Computed: true,
														Validators: []validator.String{ /*START VALIDATORS*/
															stringvalidator.RegexMatches(regexp.MustCompile("^https?://[A-Za-z0-9][^\\s]*$"), ""),
														}, /*END VALIDATORS*/
														PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
															stringplanmodifier.UseStateForUnknown(),
														}, /*END PLAN MODIFIERS*/
													}, /*END ATTRIBUTE*/
												}, /*END SCHEMA*/
											}, /*END NESTED OBJECT*/
											Optional: true,
											Computed: true,
											Validators: []validator.List{ /*START VALIDATORS*/
												listvalidator.SizeBetween(1, 100),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
												listplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Optional: true,
									Computed: true,
									Validators: []validator.Object{ /*START VALIDATORS*/
										fwvalidators.NotNullObject(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
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
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
				setplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VectorIngestionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ChunkingConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ChunkingStrategy": {
		//	          "enum": [
		//	            "FIXED_SIZE",
		//	            "NONE",
		//	            "HIERARCHICAL",
		//	            "SEMANTIC"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "FixedSizeChunkingConfiguration": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "MaxTokens": {
		//	              "minimum": 1,
		//	              "type": "number"
		//	            },
		//	            "OverlapPercentage": {
		//	              "maximum": 99,
		//	              "minimum": 1,
		//	              "type": "number"
		//	            }
		//	          },
		//	          "required": [
		//	            "MaxTokens",
		//	            "OverlapPercentage"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "HierarchicalChunkingConfiguration": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "LevelConfigurations": {
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "MaxTokens": {
		//	                    "maximum": 8192,
		//	                    "minimum": 1,
		//	                    "type": "number"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "MaxTokens"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "maxItems": 2,
		//	              "minItems": 2,
		//	              "type": "array"
		//	            },
		//	            "OverlapTokens": {
		//	              "minimum": 1,
		//	              "type": "number"
		//	            }
		//	          },
		//	          "required": [
		//	            "LevelConfigurations",
		//	            "OverlapTokens"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "SemanticChunkingConfiguration": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "BreakpointPercentileThreshold": {
		//	              "maximum": 99,
		//	              "minimum": 50,
		//	              "type": "number"
		//	            },
		//	            "BufferSize": {
		//	              "maximum": 1,
		//	              "minimum": 0,
		//	              "type": "number"
		//	            },
		//	            "MaxTokens": {
		//	              "minimum": 1,
		//	              "type": "number"
		//	            }
		//	          },
		//	          "required": [
		//	            "MaxTokens",
		//	            "BufferSize",
		//	            "BreakpointPercentileThreshold"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "ChunkingStrategy"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ParsingConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "BedrockFoundationModelConfiguration": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ModelArn": {
		//	              "maxLength": 2048,
		//	              "minLength": 1,
		//	              "pattern": "^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}::foundation-model\\/anthropic.claude-3-haiku-20240307-v1:0$",
		//	              "type": "string"
		//	            },
		//	            "ParsingPrompt": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "ParsingPromptText": {
		//	                  "maxLength": 10000,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "ParsingPromptText"
		//	              ],
		//	              "type": "object"
		//	            }
		//	          },
		//	          "required": [
		//	            "ModelArn"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "ParsingStrategy": {
		//	          "enum": [
		//	            "BEDROCK_FOUNDATION_MODEL"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ParsingStrategy"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"vector_ingestion_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ChunkingConfiguration
				"chunking_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ChunkingStrategy
						"chunking_strategy": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"FIXED_SIZE",
									"NONE",
									"HIERARCHICAL",
									"SEMANTIC",
								),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: FixedSizeChunkingConfiguration
						"fixed_size_chunking_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: MaxTokens
								"max_tokens": schema.Float64Attribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.Float64{ /*START VALIDATORS*/
										float64validator.AtLeast(1.000000),
										fwvalidators.NotNullFloat64(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
										float64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: OverlapPercentage
								"overlap_percentage": schema.Float64Attribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.Float64{ /*START VALIDATORS*/
										float64validator.Between(1.000000, 99.000000),
										fwvalidators.NotNullFloat64(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
										float64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: HierarchicalChunkingConfiguration
						"hierarchical_chunking_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: LevelConfigurations
								"level_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: MaxTokens
											"max_tokens": schema.Float64Attribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Validators: []validator.Float64{ /*START VALIDATORS*/
													float64validator.Between(1.000000, 8192.000000),
													fwvalidators.NotNullFloat64(),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
													float64planmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
									}, /*END NESTED OBJECT*/
									Optional: true,
									Computed: true,
									Validators: []validator.List{ /*START VALIDATORS*/
										listvalidator.SizeBetween(2, 2),
										fwvalidators.NotNullList(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: OverlapTokens
								"overlap_tokens": schema.Float64Attribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.Float64{ /*START VALIDATORS*/
										float64validator.AtLeast(1.000000),
										fwvalidators.NotNullFloat64(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
										float64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: SemanticChunkingConfiguration
						"semantic_chunking_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: BreakpointPercentileThreshold
								"breakpoint_percentile_threshold": schema.Float64Attribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.Float64{ /*START VALIDATORS*/
										float64validator.Between(50.000000, 99.000000),
										fwvalidators.NotNullFloat64(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
										float64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: BufferSize
								"buffer_size": schema.Float64Attribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.Float64{ /*START VALIDATORS*/
										float64validator.Between(0.000000, 1.000000),
										fwvalidators.NotNullFloat64(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
										float64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: MaxTokens
								"max_tokens": schema.Float64Attribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.Float64{ /*START VALIDATORS*/
										float64validator.AtLeast(1.000000),
										fwvalidators.NotNullFloat64(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
										float64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ParsingConfiguration
				"parsing_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BedrockFoundationModelConfiguration
						"bedrock_foundation_model_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ModelArn
								"model_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 2048),
										stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}::foundation-model\\/anthropic.claude-3-haiku-20240307-v1:0$"), ""),
										fwvalidators.NotNullString(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: ParsingPrompt
								"parsing_prompt": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ParsingPromptText
										"parsing_prompt_text": schema.StringAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.LengthBetween(1, 10000),
												fwvalidators.NotNullString(),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ParsingStrategy
						"parsing_strategy": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"BEDROCK_FOUNDATION_MODEL",
								),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Definition of AWS::Wisdom::KnowledgeBase Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Wisdom::KnowledgeBase").WithTerraformTypeName("awscc_wisdom_knowledge_base")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_integration_arn":                    "AppIntegrationArn",
		"app_integrations":                       "AppIntegrations",
		"bedrock_foundation_model_configuration": "BedrockFoundationModelConfiguration",
		"breakpoint_percentile_threshold":        "BreakpointPercentileThreshold",
		"buffer_size":                            "BufferSize",
		"chunking_configuration":                 "ChunkingConfiguration",
		"chunking_strategy":                      "ChunkingStrategy",
		"crawler_limits":                         "CrawlerLimits",
		"description":                            "Description",
		"exclusion_filters":                      "ExclusionFilters",
		"fixed_size_chunking_configuration":      "FixedSizeChunkingConfiguration",
		"hierarchical_chunking_configuration":    "HierarchicalChunkingConfiguration",
		"inclusion_filters":                      "InclusionFilters",
		"key":                                    "Key",
		"kms_key_id":                             "KmsKeyId",
		"knowledge_base_arn":                     "KnowledgeBaseArn",
		"knowledge_base_id":                      "KnowledgeBaseId",
		"knowledge_base_type":                    "KnowledgeBaseType",
		"level_configurations":                   "LevelConfigurations",
		"managed_source_configuration":           "ManagedSourceConfiguration",
		"max_tokens":                             "MaxTokens",
		"model_arn":                              "ModelArn",
		"name":                                   "Name",
		"object_fields":                          "ObjectFields",
		"overlap_percentage":                     "OverlapPercentage",
		"overlap_tokens":                         "OverlapTokens",
		"parsing_configuration":                  "ParsingConfiguration",
		"parsing_prompt":                         "ParsingPrompt",
		"parsing_prompt_text":                    "ParsingPromptText",
		"parsing_strategy":                       "ParsingStrategy",
		"rate_limit":                             "RateLimit",
		"rendering_configuration":                "RenderingConfiguration",
		"scope":                                  "Scope",
		"seed_urls":                              "SeedUrls",
		"semantic_chunking_configuration":        "SemanticChunkingConfiguration",
		"server_side_encryption_configuration":   "ServerSideEncryptionConfiguration",
		"source_configuration":                   "SourceConfiguration",
		"tags":                                   "Tags",
		"template_uri":                           "TemplateUri",
		"url":                                    "Url",
		"url_configuration":                      "UrlConfiguration",
		"value":                                  "Value",
		"vector_ingestion_configuration":         "VectorIngestionConfiguration",
		"web_crawler_configuration":              "WebCrawlerConfiguration",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
