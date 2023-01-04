// Code generated by generators/resource/main.go; DO NOT EDIT.

package lex

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_lex_bot_alias", botAliasResource)
}

// botAliasResource returns the Terraform awscc_lex_bot_alias resource.
// This Terraform resource corresponds to the CloudFormation AWS::Lex::BotAlias resource.
func botAliasResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BotAliasId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique ID of resource",
		//	  "maxLength": 10,
		//	  "minLength": 10,
		//	  "pattern": "^[0-9a-zA-Z]+$",
		//	  "type": "string"
		//	}
		"bot_alias_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique ID of resource",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BotAliasLocaleSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of bot alias locale settings to add to the bot alias.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A locale setting in alias",
		//	    "properties": {
		//	      "BotAliasLocaleSetting": {
		//	        "additionalProperties": false,
		//	        "description": "You can use this parameter to specify a specific Lambda function to run different functions in different locales.",
		//	        "properties": {
		//	          "CodeHookSpecification": {
		//	            "additionalProperties": false,
		//	            "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
		//	            "properties": {
		//	              "LambdaCodeHook": {
		//	                "additionalProperties": false,
		//	                "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
		//	                "properties": {
		//	                  "CodeHookInterfaceVersion": {
		//	                    "description": "The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.",
		//	                    "maxLength": 5,
		//	                    "minLength": 1,
		//	                    "type": "string"
		//	                  },
		//	                  "LambdaArn": {
		//	                    "description": "The Amazon Resource Name (ARN) of the Lambda function.",
		//	                    "maxLength": 2048,
		//	                    "minLength": 20,
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "CodeHookInterfaceVersion",
		//	                  "LambdaArn"
		//	                ],
		//	                "type": "object"
		//	              }
		//	            },
		//	            "required": [
		//	              "LambdaCodeHook"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "Enabled": {
		//	            "description": "Whether the Lambda code hook is enabled",
		//	            "type": "boolean"
		//	          }
		//	        },
		//	        "required": [
		//	          "Enabled"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "LocaleId": {
		//	        "description": "A string used to identify the locale",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "LocaleId",
		//	      "BotAliasLocaleSetting"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"bot_alias_locale_settings": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: BotAliasLocaleSetting
					"bot_alias_locale_setting": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: CodeHookSpecification
							"code_hook_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: LambdaCodeHook
									"lambda_code_hook": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: CodeHookInterfaceVersion
											"code_hook_interface_version": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.",
												Required:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(1, 5),
												}, /*END VALIDATORS*/
											}, /*END ATTRIBUTE*/
											// Property: LambdaArn
											"lambda_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The Amazon Resource Name (ARN) of the Lambda function.",
												Required:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(20, 2048),
												}, /*END VALIDATORS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "Contains information about code hooks that Amazon Lex calls during a conversation.",
										Required:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Contains information about code hooks that Amazon Lex calls during a conversation.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Enabled
							"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Whether the Lambda code hook is enabled",
								Required:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "You can use this parameter to specify a specific Lambda function to run different functions in different locales.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: LocaleId
					"locale_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string used to identify the locale",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of bot alias locale settings to add to the bot alias.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BotAliasName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for a resource.",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^([0-9a-zA-Z][_-]?)+$",
		//	  "type": "string"
		//	}
		"bot_alias_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for a resource.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 100),
				stringvalidator.RegexMatches(regexp.MustCompile("^([0-9a-zA-Z][_-]?)+$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: BotAliasStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "Creating",
		//	    "Available",
		//	    "Deleting",
		//	    "Failed"
		//	  ],
		//	  "type": "string"
		//	}
		"bot_alias_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BotAliasTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags to add to the bot alias.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A label for tagging Lex resources",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A string used to identify this tag",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "A string containing the value for the tag",
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
		//	  "uniqueItems": true
		//	}
		"bot_alias_tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string used to identify this tag",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string containing the value for the tag",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of tags to add to the bot alias.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// BotAliasTags is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: BotId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique ID of resource",
		//	  "maxLength": 10,
		//	  "minLength": 10,
		//	  "pattern": "^[0-9a-zA-Z]+$",
		//	  "type": "string"
		//	}
		"bot_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique ID of resource",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(10, 10),
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-zA-Z]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BotVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version of a bot.",
		//	  "maxLength": 5,
		//	  "minLength": 1,
		//	  "pattern": "^(DRAFT|[0-9]+)$",
		//	  "type": "string"
		//	}
		"bot_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version of a bot.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 5),
				stringvalidator.RegexMatches(regexp.MustCompile("^(DRAFT|[0-9]+)$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConversationLogSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
		//	  "properties": {
		//	    "AudioLogSettings": {
		//	      "description": "List of audio log settings",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Settings for logging audio of conversations between Amazon Lex and a user. You specify whether to log audio and the Amazon S3 bucket where the audio file is stored.",
		//	        "properties": {
		//	          "Destination": {
		//	            "additionalProperties": false,
		//	            "description": "The location of audio log files collected when conversation logging is enabled for a bot.",
		//	            "properties": {
		//	              "S3Bucket": {
		//	                "additionalProperties": false,
		//	                "description": "Specifies an Amazon S3 bucket for logging audio conversations",
		//	                "properties": {
		//	                  "KmsKeyArn": {
		//	                    "description": "The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for encrypting audio log files stored in an S3 bucket.",
		//	                    "maxLength": 2048,
		//	                    "minLength": 20,
		//	                    "pattern": "^arn:[\\w\\-]+:kms:[\\w\\-]+:[\\d]{12}:(?:key\\/[\\w\\-]+|alias\\/[a-zA-Z0-9:\\/_\\-]{1,256})$",
		//	                    "type": "string"
		//	                  },
		//	                  "LogPrefix": {
		//	                    "description": "The Amazon S3 key of the deployment package.",
		//	                    "maxLength": 1024,
		//	                    "minLength": 0,
		//	                    "type": "string"
		//	                  },
		//	                  "S3BucketArn": {
		//	                    "description": "The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.",
		//	                    "maxLength": 2048,
		//	                    "minLength": 1,
		//	                    "pattern": "^arn:[\\w\\-]+:s3:::[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$",
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "LogPrefix",
		//	                  "S3BucketArn"
		//	                ],
		//	                "type": "object"
		//	              }
		//	            },
		//	            "required": [
		//	              "S3Bucket"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "Enabled": {
		//	            "description": "",
		//	            "type": "boolean"
		//	          }
		//	        },
		//	        "required": [
		//	          "Destination",
		//	          "Enabled"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "TextLogSettings": {
		//	      "description": "List of text log settings",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
		//	        "properties": {
		//	          "Destination": {
		//	            "additionalProperties": false,
		//	            "description": "Defines the Amazon CloudWatch Logs destination log group for conversation text logs.",
		//	            "properties": {
		//	              "CloudWatch": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "CloudWatchLogGroupArn": {
		//	                    "description": "A string used to identify the groupArn for the Cloudwatch Log Group",
		//	                    "maxLength": 2048,
		//	                    "minLength": 1,
		//	                    "type": "string"
		//	                  },
		//	                  "LogPrefix": {
		//	                    "description": "A string containing the value for the Log Prefix",
		//	                    "maxLength": 1024,
		//	                    "minLength": 0,
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "CloudWatchLogGroupArn",
		//	                  "LogPrefix"
		//	                ],
		//	                "type": "object"
		//	              }
		//	            },
		//	            "required": [
		//	              "CloudWatch"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "Enabled": {
		//	            "description": "",
		//	            "type": "boolean"
		//	          }
		//	        },
		//	        "required": [
		//	          "Destination",
		//	          "Enabled"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"conversation_log_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AudioLogSettings
				"audio_log_settings": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Destination
							"destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: S3Bucket
									"s3_bucket": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: KmsKeyArn
											"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for encrypting audio log files stored in an S3 bucket.",
												Optional:    true,
												Computed:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(20, 2048),
													stringvalidator.RegexMatches(regexp.MustCompile("^arn:[\\w\\-]+:kms:[\\w\\-]+:[\\d]{12}:(?:key\\/[\\w\\-]+|alias\\/[a-zA-Z0-9:\\/_\\-]{1,256})$"), ""),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: LogPrefix
											"log_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The Amazon S3 key of the deployment package.",
												Required:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(0, 1024),
												}, /*END VALIDATORS*/
											}, /*END ATTRIBUTE*/
											// Property: S3BucketArn
											"s3_bucket_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.",
												Required:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(1, 2048),
													stringvalidator.RegexMatches(regexp.MustCompile("^arn:[\\w\\-]+:s3:::[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$"), ""),
												}, /*END VALIDATORS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "Specifies an Amazon S3 bucket for logging audio conversations",
										Required:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "The location of audio log files collected when conversation logging is enabled for a bot.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Enabled
							"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "",
								Required:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of audio log settings",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeAtMost(1),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TextLogSettings
				"text_log_settings": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Destination
							"destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: CloudWatch
									"cloudwatch": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: CloudWatchLogGroupArn
											"cloudwatch_log_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "A string used to identify the groupArn for the Cloudwatch Log Group",
												Required:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(1, 2048),
												}, /*END VALIDATORS*/
											}, /*END ATTRIBUTE*/
											// Property: LogPrefix
											"log_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "A string containing the value for the Log Prefix",
												Required:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(0, 1024),
												}, /*END VALIDATORS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Required: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Defines the Amazon CloudWatch Logs destination log group for conversation text logs.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Enabled
							"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "",
								Required:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of text log settings",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeAtMost(1),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Contains information about code hooks that Amazon Lex calls during a conversation.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the bot alias. Use the description to help identify the bot alias in lists.",
		//	  "maxLength": 200,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the bot alias. Use the description to help identify the bot alias in lists.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SentimentAnalysisSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.",
		//	  "properties": {
		//	    "DetectSentiment": {
		//	      "description": "Enable to call Amazon Comprehend for Sentiment natively within Lex",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "required": [
		//	    "DetectSentiment"
		//	  ],
		//	  "type": "object"
		//	}
		"sentiment_analysis_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DetectSentiment
				"detect_sentiment": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Enable to call Amazon Comprehend for Sentiment natively within Lex",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "A Bot Alias enables you to change the version of a bot without updating applications that use the bot",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lex::BotAlias").WithTerraformTypeName("awscc_lex_bot_alias")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"audio_log_settings":          "AudioLogSettings",
		"bot_alias_id":                "BotAliasId",
		"bot_alias_locale_setting":    "BotAliasLocaleSetting",
		"bot_alias_locale_settings":   "BotAliasLocaleSettings",
		"bot_alias_name":              "BotAliasName",
		"bot_alias_status":            "BotAliasStatus",
		"bot_alias_tags":              "BotAliasTags",
		"bot_id":                      "BotId",
		"bot_version":                 "BotVersion",
		"cloudwatch":                  "CloudWatch",
		"cloudwatch_log_group_arn":    "CloudWatchLogGroupArn",
		"code_hook_interface_version": "CodeHookInterfaceVersion",
		"code_hook_specification":     "CodeHookSpecification",
		"conversation_log_settings":   "ConversationLogSettings",
		"description":                 "Description",
		"destination":                 "Destination",
		"detect_sentiment":            "DetectSentiment",
		"enabled":                     "Enabled",
		"key":                         "Key",
		"kms_key_arn":                 "KmsKeyArn",
		"lambda_arn":                  "LambdaArn",
		"lambda_code_hook":            "LambdaCodeHook",
		"locale_id":                   "LocaleId",
		"log_prefix":                  "LogPrefix",
		"s3_bucket":                   "S3Bucket",
		"s3_bucket_arn":               "S3BucketArn",
		"sentiment_analysis_settings": "SentimentAnalysisSettings",
		"text_log_settings":           "TextLogSettings",
		"value":                       "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/BotAliasTags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}