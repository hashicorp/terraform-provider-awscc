// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkfirewall

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_networkfirewall_logging_configuration", loggingConfigurationResource)
}

// loggingConfigurationResource returns the Terraform awscc_networkfirewall_logging_configuration resource.
// This Terraform resource corresponds to the CloudFormation AWS::NetworkFirewall::LoggingConfiguration resource.
func loggingConfigurationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: EnableMonitoringDashboard
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_monitoring_dashboard": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FirewallArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A resource ARN.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws.*$",
		//	  "type": "string"
		//	}
		"firewall_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A resource ARN.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws.*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FirewallName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"firewall_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LoggingConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "LogDestinationConfigs": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "LogDestination": {
		//	            "additionalProperties": false,
		//	            "description": "A key-value pair to configure the logDestinations.",
		//	            "patternProperties": {
		//	              "": {
		//	                "maxLength": 1024,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "LogDestinationType": {
		//	            "enum": [
		//	              "S3",
		//	              "CloudWatchLogs",
		//	              "KinesisDataFirehose"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "LogType": {
		//	            "enum": [
		//	              "ALERT",
		//	              "FLOW",
		//	              "TLS"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "LogType",
		//	          "LogDestinationType",
		//	          "LogDestination"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "LogDestinationConfigs"
		//	  ],
		//	  "type": "object"
		//	}
		"logging_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LogDestinationConfigs
				"log_destination_configs": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: LogDestination
							"log_destination":   // Pattern: ""
							schema.MapAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "A key-value pair to configure the logDestinations.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: LogDestinationType
							"log_destination_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"S3",
										"CloudWatchLogs",
										"KinesisDataFirehose",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: LogType
							"log_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"ALERT",
										"FLOW",
										"TLS",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Required: true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeAtLeast(1),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
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
		Description: "Resource type definition for AWS::NetworkFirewall::LoggingConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkFirewall::LoggingConfiguration").WithTerraformTypeName("awscc_networkfirewall_logging_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"enable_monitoring_dashboard": "EnableMonitoringDashboard",
		"firewall_arn":                "FirewallArn",
		"firewall_name":               "FirewallName",
		"log_destination":             "LogDestination",
		"log_destination_configs":     "LogDestinationConfigs",
		"log_destination_type":        "LogDestinationType",
		"log_type":                    "LogType",
		"logging_configuration":       "LoggingConfiguration",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
