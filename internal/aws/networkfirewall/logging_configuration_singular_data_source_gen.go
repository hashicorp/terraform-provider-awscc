// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkfirewall

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_networkfirewall_logging_configuration", loggingConfigurationDataSource)
}

// loggingConfigurationDataSource returns the Terraform awscc_networkfirewall_logging_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::NetworkFirewall::LoggingConfiguration resource.
func loggingConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: EnableMonitoringDashboard
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_monitoring_dashboard": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
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
			Computed:    true,
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
			Computed: true,
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
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: LogDestinationType
							"log_destination_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: LogType
							"log_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::NetworkFirewall::LoggingConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

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

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
