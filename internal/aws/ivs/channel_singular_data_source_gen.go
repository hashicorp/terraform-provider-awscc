// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ivs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ivs_channel", channelDataSource)
}

// channelDataSource returns the Terraform awscc_ivs_channel data source.
// This Terraform data source corresponds to the CloudFormation AWS::IVS::Channel resource.
func channelDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Channel ARN is automatically generated on creation and assigned as the unique identifier.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws:ivs:[a-z0-9-]+:[0-9]+:channel/[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Channel ARN is automatically generated on creation and assigned as the unique identifier.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Authorized
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "Whether the channel is authorized.",
		//	  "type": "boolean"
		//	}
		"authorized": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether the channel is authorized.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ContainerFormat
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "TS",
		//	  "description": "Indicates which content-packaging format is used (MPEG-TS or fMP4). If multitrackInputConfiguration is specified and enabled is true, then containerFormat is required and must be set to FRAGMENTED_MP4. Otherwise, containerFormat may be set to TS or FRAGMENTED_MP4. Default: TS.",
		//	  "enum": [
		//	    "TS",
		//	    "FRAGMENTED_MP4"
		//	  ],
		//	  "type": "string"
		//	}
		"container_format": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates which content-packaging format is used (MPEG-TS or fMP4). If multitrackInputConfiguration is specified and enabled is true, then containerFormat is required and must be set to FRAGMENTED_MP4. Otherwise, containerFormat may be set to TS or FRAGMENTED_MP4. Default: TS.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IngestEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Channel ingest endpoint, part of the definition of an ingest server, used when you set up streaming software.",
		//	  "type": "string"
		//	}
		"ingest_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Channel ingest endpoint, part of the definition of an ingest server, used when you set up streaming software.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InsecureIngest
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "Whether the channel allows insecure ingest.",
		//	  "type": "boolean"
		//	}
		"insecure_ingest": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether the channel allows insecure ingest.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LatencyMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "LOW",
		//	  "description": "Channel latency mode.",
		//	  "enum": [
		//	    "NORMAL",
		//	    "LOW"
		//	  ],
		//	  "type": "string"
		//	}
		"latency_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Channel latency mode.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MultitrackInputConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Enabled": {
		//	      "default": false,
		//	      "description": "Indicates whether multitrack input is enabled. Can be set to true only if channel type is STANDARD. Setting enabled to true with any other channel type will cause an exception. If true, then policy, maximumResolution, and containerFormat are required, and containerFormat must be set to FRAGMENTED_MP4. Default: false.",
		//	      "type": "boolean"
		//	    },
		//	    "MaximumResolution": {
		//	      "description": "Maximum resolution for multitrack input. Required if enabled is true.",
		//	      "enum": [
		//	        "SD",
		//	        "HD",
		//	        "FULL_HD"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Policy": {
		//	      "description": "Indicates whether multitrack input is allowed or required. Required if enabled is true.",
		//	      "enum": [
		//	        "ALLOW",
		//	        "REQUIRE"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"multitrack_input_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Enabled
				"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether multitrack input is enabled. Can be set to true only if channel type is STANDARD. Setting enabled to true with any other channel type will cause an exception. If true, then policy, maximumResolution, and containerFormat are required, and containerFormat must be set to FRAGMENTED_MP4. Default: false.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MaximumResolution
				"maximum_resolution": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Maximum resolution for multitrack input. Required if enabled is true.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Policy
				"policy": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether multitrack input is allowed or required. Required if enabled is true.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "-",
		//	  "description": "Channel",
		//	  "maxLength": 128,
		//	  "minLength": 0,
		//	  "pattern": "^[a-zA-Z0-9-_]*$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Channel",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PlaybackUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Channel Playback URL.",
		//	  "type": "string"
		//	}
		"playback_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Channel Playback URL.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Preset
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Optional transcode preset for the channel. This is selectable only for ADVANCED_HD and ADVANCED_SD channel types. For those channel types, the default preset is HIGHER_BANDWIDTH_DELIVERY. For other channel types (BASIC and STANDARD), preset is the empty string (\"\").",
		//	  "enum": [
		//	    "",
		//	    "HIGHER_BANDWIDTH_DELIVERY",
		//	    "CONSTRAINED_BANDWIDTH_DELIVERY"
		//	  ],
		//	  "type": "string"
		//	}
		"preset": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Optional transcode preset for the channel. This is selectable only for ADVANCED_HD and ADVANCED_SD channel types. For those channel types, the default preset is HIGHER_BANDWIDTH_DELIVERY. For other channel types (BASIC and STANDARD), preset is the empty string (\"\").",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RecordingConfigurationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "",
		//	  "description": "Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: \"\" (recording is disabled).",
		//	  "maxLength": 128,
		//	  "minLength": 0,
		//	  "pattern": "^$|arn:aws:ivs:[a-z0-9-]+:[0-9]+:recording-configuration/[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"recording_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: \"\" (recording is disabled).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the asset model.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
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
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of key-value pairs that contain metadata for the asset model.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "STANDARD",
		//	  "description": "Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.",
		//	  "enum": [
		//	    "STANDARD",
		//	    "BASIC",
		//	    "ADVANCED_SD",
		//	    "ADVANCED_HD"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IVS::Channel",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IVS::Channel").WithTerraformTypeName("awscc_ivs_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                            "Arn",
		"authorized":                     "Authorized",
		"container_format":               "ContainerFormat",
		"enabled":                        "Enabled",
		"ingest_endpoint":                "IngestEndpoint",
		"insecure_ingest":                "InsecureIngest",
		"key":                            "Key",
		"latency_mode":                   "LatencyMode",
		"maximum_resolution":             "MaximumResolution",
		"multitrack_input_configuration": "MultitrackInputConfiguration",
		"name":                           "Name",
		"playback_url":                   "PlaybackUrl",
		"policy":                         "Policy",
		"preset":                         "Preset",
		"recording_configuration_arn":    "RecordingConfigurationArn",
		"tags":                           "Tags",
		"type":                           "Type",
		"value":                          "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
