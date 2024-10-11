// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package finspace

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_finspace_environment", environmentResource)
}

// environmentResource returns the Terraform awscc_finspace_environment resource.
// This Terraform resource corresponds to the CloudFormation AWS::FinSpace::Environment resource.
func environmentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AwsAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AWS account ID associated with the Environment",
		//	  "pattern": "^[a-zA-Z0-9]{1,26}$",
		//	  "type": "string"
		//	}
		"aws_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AWS account ID associated with the Environment",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataBundles
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARNs of FinSpace Data Bundles to install",
		//	  "items": {
		//	    "pattern": "^arn:aws:finspace:[A-Za-z0-9_/.-]{0,63}:\\d*:data-bundle/[0-9A-Za-z_-]{1,128}$",
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"data_bundles": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "ARNs of FinSpace Data Bundles to install",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws:finspace:[A-Za-z0-9_/.-]{0,63}:\\d*:data-bundle/[0-9A-Za-z_-]{1,128}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DedicatedServiceAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ID for FinSpace created account used to store Environment artifacts",
		//	  "pattern": "^[a-zA-Z0-9]{1,26}$",
		//	  "type": "string"
		//	}
		"dedicated_service_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ID for FinSpace created account used to store Environment artifacts",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of the Environment",
		//	  "pattern": "^[a-zA-Z0-9. ]{1,1000}$",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of the Environment",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9. ]{1,1000}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of the Environment",
		//	  "pattern": "^arn:aws:finspace:[A-Za-z0-9_/.-]{0,63}:\\d+:environment/[0-9A-Za-z_-]{1,128}$",
		//	  "type": "string"
		//	}
		"environment_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of the Environment",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique identifier for representing FinSpace Environment",
		//	  "pattern": "^[a-zA-Z0-9]{1,26}$",
		//	  "type": "string"
		//	}
		"environment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique identifier for representing FinSpace Environment",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "URL used to login to the Environment",
		//	  "pattern": "^[-a-zA-Z0-9+\u0026amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\u0026amp;@#/%=~_|]{1,1000}",
		//	  "type": "string"
		//	}
		"environment_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "URL used to login to the Environment",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FederationMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Federation mode used with the Environment",
		//	  "enum": [
		//	    "LOCAL",
		//	    "FEDERATED"
		//	  ],
		//	  "type": "string"
		//	}
		"federation_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Federation mode used with the Environment",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"LOCAL",
					"FEDERATED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FederationParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Additional parameters to identify Federation mode",
		//	  "properties": {
		//	    "ApplicationCallBackURL": {
		//	      "description": "SAML metadata URL to link with the Environment",
		//	      "pattern": "^https?://[-a-zA-Z0-9+\u0026amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\u0026amp;@#/%=~_|]{1,1000}",
		//	      "type": "string"
		//	    },
		//	    "AttributeMap": {
		//	      "description": "Attribute map for SAML configuration",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Key": {
		//	            "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	            "maxLength": 256,
		//	            "minLength": 0,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "FederationProviderName": {
		//	      "description": "Federation provider name to link with the Environment",
		//	      "maxLength": 32,
		//	      "minLength": 1,
		//	      "pattern": "[^_\\p{Z}][\\p{L}\\p{M}\\p{S}\\p{N}\\p{P}][^_\\p{Z}]+",
		//	      "type": "string"
		//	    },
		//	    "FederationURN": {
		//	      "description": "SAML metadata URL to link with the Environment",
		//	      "type": "string"
		//	    },
		//	    "SamlMetadataDocument": {
		//	      "description": "SAML metadata document to link the federation provider to the Environment",
		//	      "maxLength": 10000000,
		//	      "minLength": 1000,
		//	      "pattern": ".*",
		//	      "type": "string"
		//	    },
		//	    "SamlMetadataURL": {
		//	      "description": "SAML metadata URL to link with the Environment",
		//	      "pattern": "^https?://[-a-zA-Z0-9+\u0026amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\u0026amp;@#/%=~_|]{1,1000}",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"federation_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ApplicationCallBackURL
				"application_call_back_url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "SAML metadata URL to link with the Environment",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("^https?://[-a-zA-Z0-9+&amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&amp;@#/%=~_|]{1,1000}"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: AttributeMap
				"attribute_map": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Key
							"key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 128),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
									stringplanmodifier.RequiresReplaceIfConfigured(),
								}, /*END PLAN MODIFIERS*/
								// Key is a write-only property.
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(0, 256),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
									stringplanmodifier.RequiresReplaceIfConfigured(),
								}, /*END PLAN MODIFIERS*/
								// Value is a write-only property.
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Attribute map for SAML configuration",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
						listplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
					// AttributeMap is a write-only property.
				}, /*END ATTRIBUTE*/
				// Property: FederationProviderName
				"federation_provider_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Federation provider name to link with the Environment",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 32),
						stringvalidator.RegexMatches(regexp.MustCompile("[^_\\p{Z}][\\p{L}\\p{M}\\p{S}\\p{N}\\p{P}][^_\\p{Z}]+"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: FederationURN
				"federation_urn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "SAML metadata URL to link with the Environment",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SamlMetadataDocument
				"saml_metadata_document": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "SAML metadata document to link the federation provider to the Environment",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1000, 10000000),
						stringvalidator.RegexMatches(regexp.MustCompile(".*"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SamlMetadataURL
				"saml_metadata_url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "SAML metadata URL to link with the Environment",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("^https?://[-a-zA-Z0-9+&amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&amp;@#/%=~_|]{1,1000}"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Additional parameters to identify Federation mode",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "KMS key used to encrypt customer data within FinSpace Environment infrastructure",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "KMS key used to encrypt customer data within FinSpace Environment infrastructure",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the Environment",
		//	  "pattern": "^[a-zA-Z0-9]+[a-zA-Z0-9-]*[a-zA-Z0-9]{1,255}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the Environment",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]+[a-zA-Z0-9-]*[a-zA-Z0-9]{1,255}$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: SageMakerStudioDomainUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "SageMaker Studio Domain URL associated with the Environment",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"sage_maker_studio_domain_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "SageMaker Studio Domain URL associated with the Environment",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "State of the Environment",
		//	  "enum": [
		//	    "CREATE_REQUESTED",
		//	    "CREATING",
		//	    "CREATED",
		//	    "DELETE_REQUESTED",
		//	    "DELETING",
		//	    "DELETED",
		//	    "FAILED_CREATION",
		//	    "FAILED_DELETION",
		//	    "RETRY_DELETION",
		//	    "SUSPENDED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "State of the Environment",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SuperuserParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Parameters of the first Superuser for the FinSpace Environment",
		//	  "properties": {
		//	    "EmailAddress": {
		//	      "description": "Email address",
		//	      "maxLength": 128,
		//	      "minLength": 1,
		//	      "pattern": "[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+[.]+[A-Za-z]+",
		//	      "type": "string"
		//	    },
		//	    "FirstName": {
		//	      "description": "First name",
		//	      "maxLength": 50,
		//	      "minLength": 1,
		//	      "pattern": "^[a-zA-Z0-9]{1,50}$",
		//	      "type": "string"
		//	    },
		//	    "LastName": {
		//	      "description": "Last name",
		//	      "maxLength": 50,
		//	      "minLength": 1,
		//	      "pattern": "^[a-zA-Z0-9]{1,50}$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"superuser_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EmailAddress
				"email_address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Email address",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 128),
						stringvalidator.RegexMatches(regexp.MustCompile("[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+[.]+[A-Za-z]+"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
					// EmailAddress is a write-only property.
				}, /*END ATTRIBUTE*/
				// Property: FirstName
				"first_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "First name",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 50),
						stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]{1,50}$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
					// FirstName is a write-only property.
				}, /*END ATTRIBUTE*/
				// Property: LastName
				"last_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Last name",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 50),
						stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]{1,50}$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
					// LastName is a write-only property.
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Parameters of the first Superuser for the FinSpace Environment",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// SuperuserParameters is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A list of all tags for a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
							stringplanmodifier.RequiresReplaceIfConfigured(),
						}, /*END PLAN MODIFIERS*/
						// Key is a write-only property.
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
							stringplanmodifier.RequiresReplaceIfConfigured(),
						}, /*END PLAN MODIFIERS*/
						// Value is a write-only property.
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// Tags is a write-only property.
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
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FinSpace::Environment").WithTerraformTypeName("awscc_finspace_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_call_back_url":    "ApplicationCallBackURL",
		"attribute_map":                "AttributeMap",
		"aws_account_id":               "AwsAccountId",
		"data_bundles":                 "DataBundles",
		"dedicated_service_account_id": "DedicatedServiceAccountId",
		"description":                  "Description",
		"email_address":                "EmailAddress",
		"environment_arn":              "EnvironmentArn",
		"environment_id":               "EnvironmentId",
		"environment_url":              "EnvironmentUrl",
		"federation_mode":              "FederationMode",
		"federation_parameters":        "FederationParameters",
		"federation_provider_name":     "FederationProviderName",
		"federation_urn":               "FederationURN",
		"first_name":                   "FirstName",
		"key":                          "Key",
		"kms_key_id":                   "KmsKeyId",
		"last_name":                    "LastName",
		"name":                         "Name",
		"sage_maker_studio_domain_url": "SageMakerStudioDomainUrl",
		"saml_metadata_document":       "SamlMetadataDocument",
		"saml_metadata_url":            "SamlMetadataURL",
		"status":                       "Status",
		"superuser_parameters":         "SuperuserParameters",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/SuperuserParameters",
		"/properties/FederationParameters/AttributeMap",
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
