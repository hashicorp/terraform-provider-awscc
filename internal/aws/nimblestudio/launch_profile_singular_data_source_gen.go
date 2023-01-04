// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package nimblestudio

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_nimblestudio_launch_profile", launchProfileDataSource)
}

// launchProfileDataSource returns the Terraform awscc_nimblestudio_launch_profile data source.
// This Terraform data source corresponds to the CloudFormation AWS::NimbleStudio::LaunchProfile resource.
func launchProfileDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe description.\u003c/p\u003e",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The description.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ec2SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eSpecifies the IDs of the EC2 subnets where streaming sessions will be accessible from.\n            These subnets must support the specified instance types. \u003c/p\u003e",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "maxItems": 6,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"ec_2_subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "<p>Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from.\n            These subnets must support the specified instance types. </p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LaunchProfileId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"launch_profile_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LaunchProfileProtocolVersions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe version number of the protocol that is used by the launch profile. The only valid\n            version is \"2021-03-31\".\u003c/p\u003e",
		//	  "items": {
		//	    "description": "\u003cp\u003eThe version number of the protocol that is used by the launch profile. The only valid\n            version is \"2021-03-31\".\u003c/p\u003e",
		//	    "maxLength": 10,
		//	    "minLength": 0,
		//	    "pattern": "^2021\\-03\\-31$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"launch_profile_protocol_versions": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "<p>The version number of the protocol that is used by the launch profile. The only valid\n            version is \"2021-03-31\".</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe name for the launch profile.\u003c/p\u003e",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The name for the launch profile.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StreamConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "\u003cp\u003eA configuration for a streaming session.\u003c/p\u003e",
		//	  "properties": {
		//	    "ClipboardMode": {
		//	      "enum": [
		//	        "ENABLED",
		//	        "DISABLED"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Ec2InstanceTypes": {
		//	      "description": "\u003cp\u003eThe EC2 instance types that users can select from when launching a streaming session\n            with this launch profile.\u003c/p\u003e",
		//	      "items": {
		//	        "enum": [
		//	          "g4dn.xlarge",
		//	          "g4dn.2xlarge",
		//	          "g4dn.4xlarge",
		//	          "g4dn.8xlarge",
		//	          "g4dn.12xlarge",
		//	          "g4dn.16xlarge",
		//	          "g3.4xlarge",
		//	          "g3s.xlarge",
		//	          "g5.xlarge",
		//	          "g5.2xlarge",
		//	          "g5.4xlarge",
		//	          "g5.8xlarge",
		//	          "g5.16xlarge"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "maxItems": 30,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    },
		//	    "MaxSessionLengthInMinutes": {
		//	      "description": "\u003cp\u003eThe length of time, in minutes, that a streaming session can be active before it is\n            stopped or terminated. After this point, Nimble Studio automatically terminates or\n            stops the session. The default length of time is 690 minutes, and the maximum length of\n            time is 30 days.\u003c/p\u003e",
		//	      "maximum": 43200,
		//	      "minimum": 1,
		//	      "type": "number"
		//	    },
		//	    "MaxStoppedSessionLengthInMinutes": {
		//	      "description": "\u003cp\u003eInteger that determines if you can start and stop your sessions and how long a session\n            can stay in the STOPPED state. The default value is 0. The maximum value is 5760.\u003c/p\u003e\n        \u003cp\u003eIf the value is missing or set to 0, your sessions can?t be stopped. If you then call\n                \u003ccode\u003eStopStreamingSession\u003c/code\u003e, the session fails. If the time that a session\n            stays in the READY state exceeds the \u003ccode\u003emaxSessionLengthInMinutes\u003c/code\u003e value, the\n            session will automatically be terminated (instead of stopped).\u003c/p\u003e\n        \u003cp\u003eIf the value is set to a positive number, the session can be stopped. You can call\n                \u003ccode\u003eStopStreamingSession\u003c/code\u003e to stop sessions in the READY state. If the time\n            that a session stays in the READY state exceeds the\n                \u003ccode\u003emaxSessionLengthInMinutes\u003c/code\u003e value, the session will automatically be\n            stopped (instead of terminated).\u003c/p\u003e",
		//	      "maximum": 5760,
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "SessionStorage": {
		//	      "additionalProperties": false,
		//	      "description": "\u003cp\u003eThe configuration for a streaming session?s upload storage.\u003c/p\u003e",
		//	      "properties": {
		//	        "Mode": {
		//	          "description": "\u003cp\u003eAllows artists to upload files to their workstations. The only valid option is\n                \u003ccode\u003eUPLOAD\u003c/code\u003e.\u003c/p\u003e",
		//	          "items": {
		//	            "enum": [
		//	              "UPLOAD"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "minItems": 1,
		//	          "type": "array"
		//	        },
		//	        "Root": {
		//	          "additionalProperties": false,
		//	          "description": "\u003cp\u003eThe upload storage root location (folder) on streaming workstations where files are\n            uploaded.\u003c/p\u003e",
		//	          "properties": {
		//	            "Linux": {
		//	              "description": "\u003cp\u003eThe folder path in Linux workstations where files are uploaded.\u003c/p\u003e",
		//	              "maxLength": 128,
		//	              "minLength": 1,
		//	              "pattern": "^(\\$HOME|/)[/]?([A-Za-z0-9-_]+/)*([A-Za-z0-9_-]+)$",
		//	              "type": "string"
		//	            },
		//	            "Windows": {
		//	              "description": "\u003cp\u003eThe folder path in Windows workstations where files are uploaded.\u003c/p\u003e",
		//	              "maxLength": 128,
		//	              "minLength": 1,
		//	              "pattern": "^((\\%HOMEPATH\\%)|[a-zA-Z]:)[\\\\/](?:[a-zA-Z0-9_-]+[\\\\/])*[a-zA-Z0-9_-]+$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "Mode"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "StreamingImageIds": {
		//	      "description": "\u003cp\u003eThe streaming images that users can select from when launching a streaming session\n            with this launch profile.\u003c/p\u003e",
		//	      "items": {
		//	        "maxLength": 22,
		//	        "minLength": 0,
		//	        "pattern": "^[a-zA-Z0-9-_]*$",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 20,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "ClipboardMode",
		//	    "Ec2InstanceTypes",
		//	    "StreamingImageIds"
		//	  ],
		//	  "type": "object"
		//	}
		"stream_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ClipboardMode
				"clipboard_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Ec2InstanceTypes
				"ec_2_instance_types": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "<p>The EC2 instance types that users can select from when launching a streaming session\n            with this launch profile.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MaxSessionLengthInMinutes
				"max_session_length_in_minutes": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "<p>The length of time, in minutes, that a streaming session can be active before it is\n            stopped or terminated. After this point, Nimble Studio automatically terminates or\n            stops the session. The default length of time is 690 minutes, and the maximum length of\n            time is 30 days.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MaxStoppedSessionLengthInMinutes
				"max_stopped_session_length_in_minutes": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "<p>Integer that determines if you can start and stop your sessions and how long a session\n            can stay in the STOPPED state. The default value is 0. The maximum value is 5760.</p>\n        <p>If the value is missing or set to 0, your sessions can?t be stopped. If you then call\n                <code>StopStreamingSession</code>, the session fails. If the time that a session\n            stays in the READY state exceeds the <code>maxSessionLengthInMinutes</code> value, the\n            session will automatically be terminated (instead of stopped).</p>\n        <p>If the value is set to a positive number, the session can be stopped. You can call\n                <code>StopStreamingSession</code> to stop sessions in the READY state. If the time\n            that a session stays in the READY state exceeds the\n                <code>maxSessionLengthInMinutes</code> value, the session will automatically be\n            stopped (instead of terminated).</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SessionStorage
				"session_storage": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Mode
						"mode": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "<p>Allows artists to upload files to their workstations. The only valid option is\n                <code>UPLOAD</code>.</p>",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Root
						"root": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Linux
								"linux": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "<p>The folder path in Linux workstations where files are uploaded.</p>",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Windows
								"windows": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "<p>The folder path in Windows workstations where files are uploaded.</p>",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "<p>The upload storage root location (folder) on streaming workstations where files are\n            uploaded.</p>",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "<p>The configuration for a streaming session?s upload storage.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: StreamingImageIds
				"streaming_image_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "<p>The streaming images that users can select from when launching a streaming session\n            with this launch profile.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "<p>A configuration for a streaming session.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StudioComponentIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eUnique identifiers for a collection of studio components that can be used with this\n            launch profile.\u003c/p\u003e",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "maxItems": 100,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"studio_component_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "<p>Unique identifiers for a collection of studio components that can be used with this\n            launch profile.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StudioId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe studio ID. \u003c/p\u003e",
		//	  "type": "string"
		//	}
		"studio_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The studio ID. </p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::NimbleStudio::LaunchProfile",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NimbleStudio::LaunchProfile").WithTerraformTypeName("awscc_nimblestudio_launch_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"clipboard_mode":                        "ClipboardMode",
		"description":                           "Description",
		"ec_2_instance_types":                   "Ec2InstanceTypes",
		"ec_2_subnet_ids":                       "Ec2SubnetIds",
		"launch_profile_id":                     "LaunchProfileId",
		"launch_profile_protocol_versions":      "LaunchProfileProtocolVersions",
		"linux":                                 "Linux",
		"max_session_length_in_minutes":         "MaxSessionLengthInMinutes",
		"max_stopped_session_length_in_minutes": "MaxStoppedSessionLengthInMinutes",
		"mode":                                  "Mode",
		"name":                                  "Name",
		"root":                                  "Root",
		"session_storage":                       "SessionStorage",
		"stream_configuration":                  "StreamConfiguration",
		"streaming_image_ids":                   "StreamingImageIds",
		"studio_component_ids":                  "StudioComponentIds",
		"studio_id":                             "StudioId",
		"tags":                                  "Tags",
		"windows":                               "Windows",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}