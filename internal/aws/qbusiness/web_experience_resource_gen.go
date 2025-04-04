// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package qbusiness

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	registry.AddResourceFactory("awscc_qbusiness_web_experience", webExperienceResource)
}

// webExperienceResource returns the Terraform awscc_qbusiness_web_experience resource.
// This Terraform resource corresponds to the CloudFormation AWS::QBusiness::WebExperience resource.
func webExperienceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "^[a-zA-Z0-9][a-zA-Z0-9-]{35}$",
		//	  "type": "string"
		//	}
		"application_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(36, 36),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9-]{35}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BrowserExtensionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "EnabledBrowserExtensions": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "enum": [
		//	          "FIREFOX",
		//	          "CHROME"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "maxItems": 2,
		//	      "minItems": 0,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "required": [
		//	    "EnabledBrowserExtensions"
		//	  ],
		//	  "type": "object"
		//	}
		"browser_extension_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EnabledBrowserExtensions
				"enabled_browser_extensions": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(0, 2),
						setvalidator.ValueStringsAre(
							stringvalidator.OneOf(
								"FIREFOX",
								"CHROME",
							),
						),
						fwvalidators.NotNullSet(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CustomizationConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "CustomCSSUrl": {
		//	      "maxLength": 1284,
		//	      "minLength": 0,
		//	      "pattern": "^(https?://[a-zA-Z0-9-_.+%/]+\\.css)?$",
		//	      "type": "string"
		//	    },
		//	    "FaviconUrl": {
		//	      "maxLength": 1284,
		//	      "minLength": 0,
		//	      "pattern": "^(https?://[a-zA-Z0-9-_.+%/]+\\.(svg|ico))?$",
		//	      "type": "string"
		//	    },
		//	    "FontUrl": {
		//	      "maxLength": 1284,
		//	      "minLength": 0,
		//	      "pattern": "^(https?://[a-zA-Z0-9-_.+%/]+\\.(ttf|woff|woff2|otf))?$",
		//	      "type": "string"
		//	    },
		//	    "LogoUrl": {
		//	      "maxLength": 1284,
		//	      "minLength": 0,
		//	      "pattern": "^(https?://[a-zA-Z0-9-_.+%/]+\\.(svg|png))?$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"customization_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CustomCSSUrl
				"custom_css_url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(0, 1284),
						stringvalidator.RegexMatches(regexp.MustCompile("^(https?://[a-zA-Z0-9-_.+%/]+\\.css)?$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: FaviconUrl
				"favicon_url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(0, 1284),
						stringvalidator.RegexMatches(regexp.MustCompile("^(https?://[a-zA-Z0-9-_.+%/]+\\.(svg|ico))?$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: FontUrl
				"font_url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(0, 1284),
						stringvalidator.RegexMatches(regexp.MustCompile("^(https?://[a-zA-Z0-9-_.+%/]+\\.(ttf|woff|woff2|otf))?$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: LogoUrl
				"logo_url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(0, 1284),
						stringvalidator.RegexMatches(regexp.MustCompile("^(https?://[a-zA-Z0-9-_.+%/]+\\.(svg|png))?$"), ""),
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
		// Property: DefaultEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "^(https?|ftp|file)://([^\\s]*)$",
		//	  "type": "string"
		//	}
		"default_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IdentityProviderConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "properties": {
		//	    "OpenIDConnectConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "SecretsArn": {
		//	          "maxLength": 1284,
		//	          "minLength": 0,
		//	          "pattern": "",
		//	          "type": "string"
		//	        },
		//	        "SecretsRole": {
		//	          "maxLength": 1284,
		//	          "minLength": 0,
		//	          "pattern": "",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "SecretsArn",
		//	        "SecretsRole"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "SamlConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AuthenticationUrl": {
		//	          "maxLength": 1284,
		//	          "minLength": 1,
		//	          "pattern": "^https://.*$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "AuthenticationUrl"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"identity_provider_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: OpenIDConnectConfiguration
				"open_id_connect_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: SecretsArn
						"secrets_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(0, 1284),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: SecretsRole
						"secrets_role": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(0, 1284),
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
				// Property: SamlConfiguration
				"saml_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AuthenticationUrl
						"authentication_url": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 1284),
								stringvalidator.RegexMatches(regexp.MustCompile("^https://.*$"), ""),
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
		// Property: Origins
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 64,
		//	    "minLength": 1,
		//	    "pattern": "^(http:\\/\\/|https:\\/\\/)[a-zA-Z0-9-_.]+(?::[0-9]{1,5})?$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 10,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"origins": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 10),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 64),
					stringvalidator.RegexMatches(regexp.MustCompile("^(http:\\/\\/|https:\\/\\/)[a-zA-Z0-9-_.]+(?::[0-9]{1,5})?$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1284,
		//	  "minLength": 0,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 1284),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SamplePromptsControlMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "ENABLED",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"sample_prompts_control_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ENABLED",
					"DISABLED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "CREATING",
		//	    "ACTIVE",
		//	    "DELETING",
		//	    "FAILED",
		//	    "PENDING_AUTH_CONFIG"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Subtitle
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 500,
		//	  "minLength": 0,
		//	  "pattern": "^[\\s\\S]*$",
		//	  "type": "string"
		//	}
		"subtitle": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 500),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\s\\S]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
							stringvalidator.LengthBetween(0, 256),
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
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Title
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 500,
		//	  "minLength": 0,
		//	  "pattern": "^[\\s\\S]*$",
		//	  "type": "string"
		//	}
		"title": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 500),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\s\\S]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WebExperienceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1284,
		//	  "minLength": 0,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"web_experience_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WebExperienceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "^[a-zA-Z0-9][a-zA-Z0-9-]*$",
		//	  "type": "string"
		//	}
		"web_experience_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WelcomeMessage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 300,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"welcome_message": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 300),
			}, /*END VALIDATORS*/
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
		Description: "Definition of AWS::QBusiness::WebExperience Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::QBusiness::WebExperience").WithTerraformTypeName("awscc_qbusiness_web_experience")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_id":                  "ApplicationId",
		"authentication_url":              "AuthenticationUrl",
		"browser_extension_configuration": "BrowserExtensionConfiguration",
		"created_at":                      "CreatedAt",
		"custom_css_url":                  "CustomCSSUrl",
		"customization_configuration":     "CustomizationConfiguration",
		"default_endpoint":                "DefaultEndpoint",
		"enabled_browser_extensions":      "EnabledBrowserExtensions",
		"favicon_url":                     "FaviconUrl",
		"font_url":                        "FontUrl",
		"identity_provider_configuration": "IdentityProviderConfiguration",
		"key":                             "Key",
		"logo_url":                        "LogoUrl",
		"open_id_connect_configuration":   "OpenIDConnectConfiguration",
		"origins":                         "Origins",
		"role_arn":                        "RoleArn",
		"saml_configuration":              "SamlConfiguration",
		"sample_prompts_control_mode":     "SamplePromptsControlMode",
		"secrets_arn":                     "SecretsArn",
		"secrets_role":                    "SecretsRole",
		"status":                          "Status",
		"subtitle":                        "Subtitle",
		"tags":                            "Tags",
		"title":                           "Title",
		"updated_at":                      "UpdatedAt",
		"value":                           "Value",
		"web_experience_arn":              "WebExperienceArn",
		"web_experience_id":               "WebExperienceId",
		"welcome_message":                 "WelcomeMessage",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
