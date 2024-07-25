// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package bedrock

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_bedrock_prompt_version", promptVersionResource)
}

// promptVersionResource returns the Terraform awscc_bedrock_prompt_version resource.
// This Terraform resource corresponds to the CloudFormation AWS::Bedrock::PromptVersion resource.
func promptVersionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of a prompt version resource",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "^(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:prompt/[0-9a-zA-Z]{10}:[0-9]{1,20})$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of a prompt version resource",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time Stamp.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Time Stamp.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultVariant
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name for a variant.",
		//	  "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	  "type": "string"
		//	}
		"default_variant": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name for a variant.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description for a prompt version resource.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description for a prompt version resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name for a prompt resource.",
		//	  "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name for a prompt resource.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PromptArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of a prompt resource possibly with a version",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "^(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:prompt/[0-9a-zA-Z]{10})$",
		//	  "type": "string"
		//	}
		"prompt_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of a prompt resource possibly with a version",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:prompt/[0-9a-zA-Z]{10})$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PromptId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Identifier for a Prompt",
		//	  "pattern": "^[0-9a-zA-Z]{10}$",
		//	  "type": "string"
		//	}
		"prompt_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Identifier for a Prompt",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time Stamp.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Time Stamp.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Variants
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of prompt variants",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Prompt variant",
		//	    "properties": {
		//	      "InferenceConfiguration": {
		//	        "description": "Model inference configuration",
		//	        "properties": {
		//	          "Text": {
		//	            "additionalProperties": false,
		//	            "description": "Prompt model inference configuration",
		//	            "properties": {
		//	              "MaxTokens": {
		//	                "description": "Maximum length of output",
		//	                "maximum": 4096,
		//	                "minimum": 0,
		//	                "type": "number"
		//	              },
		//	              "StopSequences": {
		//	                "description": "List of stop sequences",
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "maxItems": 4,
		//	                "minItems": 0,
		//	                "type": "array"
		//	              },
		//	              "Temperature": {
		//	                "description": "Controls randomness, higher values increase diversity",
		//	                "maximum": 1,
		//	                "minimum": 0,
		//	                "type": "number"
		//	              },
		//	              "TopK": {
		//	                "description": "Sample from the k most likely next tokens",
		//	                "maximum": 500,
		//	                "minimum": 0,
		//	                "type": "number"
		//	              },
		//	              "TopP": {
		//	                "description": "Cumulative probability cutoff for token selection",
		//	                "maximum": 1,
		//	                "minimum": 0,
		//	                "type": "number"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "ModelId": {
		//	        "description": "ARN or name of a Bedrock model.",
		//	        "maxLength": 2048,
		//	        "minLength": 1,
		//	        "pattern": "^(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:(([0-9]{12}:custom-model/[a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}/[a-z0-9]{12})|(:foundation-model/[a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}([.:]?[a-z0-9-]{1,63}))|([0-9]{12}:provisioned-model/[a-z0-9]{12})))|([a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}([.:]?[a-z0-9-]{1,63}))|(([0-9a-zA-Z][_-]?)+)$",
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "description": "Name for a variant.",
		//	        "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	        "type": "string"
		//	      },
		//	      "TemplateConfiguration": {
		//	        "description": "Prompt template configuration",
		//	        "properties": {
		//	          "Text": {
		//	            "additionalProperties": false,
		//	            "description": "Configuration for text prompt template",
		//	            "properties": {
		//	              "InputVariables": {
		//	                "description": "List of input variables",
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "additionalProperties": false,
		//	                  "description": "Input variable",
		//	                  "properties": {
		//	                    "Name": {
		//	                      "description": "Name for an input variable",
		//	                      "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "maxItems": 5,
		//	                "minItems": 1,
		//	                "type": "array"
		//	              },
		//	              "Text": {
		//	                "description": "Prompt content for String prompt template",
		//	                "maxLength": 200000,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Text"
		//	            ],
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "TemplateType": {
		//	        "description": "Prompt template type",
		//	        "enum": [
		//	          "TEXT"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name",
		//	      "TemplateType"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 3,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"variants": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: InferenceConfiguration
					"inference_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Text
							"text": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: MaxTokens
									"max_tokens": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Description: "Maximum length of output",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: StopSequences
									"stop_sequences": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "List of stop sequences",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Temperature
									"temperature": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Description: "Controls randomness, higher values increase diversity",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: TopK
									"top_k": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Description: "Sample from the k most likely next tokens",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: TopP
									"top_p": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Description: "Cumulative probability cutoff for token selection",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Prompt model inference configuration",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Model inference configuration",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ModelId
					"model_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "ARN or name of a Bedrock model.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Name for a variant.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: TemplateConfiguration
					"template_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Text
							"text": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: InputVariables
									"input_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
										NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: Name
												"name": schema.StringAttribute{ /*START ATTRIBUTE*/
													Description: "Name for an input variable",
													Computed:    true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
										}, /*END NESTED OBJECT*/
										Description: "List of input variables",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Text
									"text": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "Prompt content for String prompt template",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Configuration for text prompt template",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Prompt template configuration",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: TemplateType
					"template_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Prompt template type",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of prompt variants",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Version.",
		//	  "maxLength": 5,
		//	  "minLength": 1,
		//	  "pattern": "^(DRAFT|[0-9]{0,4}[1-9][0-9]{0,4})$",
		//	  "type": "string"
		//	}
		"version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Version.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Definition of AWS::Bedrock::PromptVersion Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Bedrock::PromptVersion").WithTerraformTypeName("awscc_bedrock_prompt_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                     "Arn",
		"created_at":              "CreatedAt",
		"default_variant":         "DefaultVariant",
		"description":             "Description",
		"inference_configuration": "InferenceConfiguration",
		"input_variables":         "InputVariables",
		"max_tokens":              "MaxTokens",
		"model_id":                "ModelId",
		"name":                    "Name",
		"prompt_arn":              "PromptArn",
		"prompt_id":               "PromptId",
		"stop_sequences":          "StopSequences",
		"temperature":             "Temperature",
		"template_configuration":  "TemplateConfiguration",
		"template_type":           "TemplateType",
		"text":                    "Text",
		"top_k":                   "TopK",
		"top_p":                   "TopP",
		"updated_at":              "UpdatedAt",
		"variants":                "Variants",
		"version":                 "Version",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
