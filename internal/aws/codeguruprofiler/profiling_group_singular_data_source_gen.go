// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package codeguruprofiler

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_codeguruprofiler_profiling_group", profilingGroupDataSourceType)
}

// profilingGroupDataSourceType returns the Terraform awscc_codeguruprofiler_profiling_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CodeGuruProfiler::ProfilingGroup resource type.
func profilingGroupDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"agent_permissions": {
			// Property: AgentPermissions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The agent permissions attached to this profiling group.",
			//   "properties": {
			//     "Principals": {
			//       "description": "The principals for the agent permissions.",
			//       "items": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "Principals"
			//   ],
			//   "type": "object"
			// }
			Description: "The agent permissions attached to this profiling group.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"principals": {
						// Property: Principals
						Description: "The principals for the agent permissions.",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"anomaly_detection_notification_configuration": {
			// Property: AnomalyDetectionNotificationConfiguration
			// CloudFormation resource type schema:
			// {
			//   "description": "Configuration for Notification Channels for Anomaly Detection feature in CodeGuru Profiler which enables customers to detect anomalies in the application profile for those methods that represent the highest proportion of CPU time or latency",
			//   "items": {
			//     "description": "Notification medium for users to get alerted for events that occur in application profile. We support SNS topic as a notification channel.",
			//     "properties": {
			//       "channelId": {
			//         "description": "Unique identifier for each Channel in the notification configuration of a Profiling Group",
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "channelUri": {
			//         "description": "Unique arn of the resource to be used for notifications. We support a valid SNS topic arn as a channel uri.",
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "channelUri"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Configuration for Notification Channels for Anomaly Detection feature in CodeGuru Profiler which enables customers to detect anomalies in the application profile for those methods that represent the highest proportion of CPU time or latency",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"channel_id": {
						// Property: channelId
						Description: "Unique identifier for each Channel in the notification configuration of a Profiling Group",
						Type:        types.StringType,
						Computed:    true,
					},
					"channel_uri": {
						// Property: channelUri
						Description: "Unique arn of the resource to be used for notifications. We support a valid SNS topic arn as a channel uri.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the specified profiling group.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the specified profiling group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"compute_platform": {
			// Property: ComputePlatform
			// CloudFormation resource type schema:
			// {
			//   "description": "The compute platform of the profiling group.",
			//   "enum": [
			//     "Default",
			//     "AWSLambda"
			//   ],
			//   "type": "string"
			// }
			Description: "The compute platform of the profiling group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"profiling_group_name": {
			// Property: ProfilingGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the profiling group.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the profiling group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags associated with a profiling group.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The tags associated with a profiling group.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::CodeGuruProfiler::ProfilingGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeGuruProfiler::ProfilingGroup").WithTerraformTypeName("awscc_codeguruprofiler_profiling_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"agent_permissions":                            "AgentPermissions",
		"anomaly_detection_notification_configuration": "AnomalyDetectionNotificationConfiguration",
		"arn":                  "Arn",
		"channel_id":           "channelId",
		"channel_uri":          "channelUri",
		"compute_platform":     "ComputePlatform",
		"key":                  "Key",
		"principals":           "Principals",
		"profiling_group_name": "ProfilingGroupName",
		"tags":                 "Tags",
		"value":                "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
