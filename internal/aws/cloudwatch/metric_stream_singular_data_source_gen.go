// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudwatch_metric_stream", metricStreamDataSource)
}

// metricStreamDataSource returns the Terraform awscc_cloudwatch_metric_stream data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudWatch::MetricStream resource.
func metricStreamDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name of the metric stream.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "Amazon Resource Name of the metric stream.",
			Type:        types.StringType,
			Computed:    true,
		},
		"creation_date": {
			// Property: CreationDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date of creation of the metric stream.",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "The date of creation of the metric stream.",
			Type:        types.StringType,
			Computed:    true,
		},
		"exclude_filters": {
			// Property: ExcludeFilters
			// CloudFormation resource type schema:
			// {
			//   "description": "Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "This structure defines the metrics that will be streamed.",
			//     "properties": {
			//       "Namespace": {
			//         "description": "Only metrics with Namespace matching this value will be streamed.",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Namespace"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 1000,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"namespace": {
						// Property: Namespace
						Description: "Only metrics with Namespace matching this value will be streamed.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"firehose_arn": {
			// Property: FirehoseArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the Kinesis Firehose where to stream the data.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "The ARN of the Kinesis Firehose where to stream the data.",
			Type:        types.StringType,
			Computed:    true,
		},
		"include_filters": {
			// Property: IncludeFilters
			// CloudFormation resource type schema:
			// {
			//   "description": "Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "This structure defines the metrics that will be streamed.",
			//     "properties": {
			//       "Namespace": {
			//         "description": "Only metrics with Namespace matching this value will be streamed.",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Namespace"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 1000,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"namespace": {
						// Property: Namespace
						Description: "Only metrics with Namespace matching this value will be streamed.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"last_update_date": {
			// Property: LastUpdateDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date of the last update of the metric stream.",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "The date of the last update of the metric stream.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the metric stream.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of the metric stream.",
			Type:        types.StringType,
			Computed:    true,
		},
		"output_format": {
			// Property: OutputFormat
			// CloudFormation resource type schema:
			// {
			//   "description": "The output format of the data streamed to the Kinesis Firehose.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The output format of the data streamed to the Kinesis Firehose.",
			Type:        types.StringType,
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the role that provides access to the Kinesis Firehose.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "The ARN of the role that provides access to the Kinesis Firehose.",
			Type:        types.StringType,
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "Displays the state of the Metric Stream.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Displays the state of the Metric Stream.",
			Type:        types.StringType,
			Computed:    true,
		},
		"statistics_configurations": {
			// Property: StatisticsConfigurations
			// CloudFormation resource type schema:
			// {
			//   "description": "By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed. You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "This structure specifies a list of additional statistics to stream, and the metrics to stream those additional statistics for. All metrics that match the combination of metric name and namespace will be streamed with the extended statistics, no matter their dimensions.",
			//     "properties": {
			//       "AdditionalStatistics": {
			//         "description": "The additional statistics to stream for the metrics listed in IncludeMetrics.",
			//         "items": {
			//           "type": "string"
			//         },
			//         "maxItems": 20,
			//         "type": "array",
			//         "uniqueItems": true
			//       },
			//       "IncludeMetrics": {
			//         "description": "An array that defines the metrics that are to have additional statistics streamed.",
			//         "items": {
			//           "additionalProperties": false,
			//           "description": "A structure that specifies the metric name and namespace for one metric that is going to have additional statistics included in the stream.",
			//           "properties": {
			//             "MetricName": {
			//               "description": "The name of the metric.",
			//               "maxLength": 255,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Namespace": {
			//               "description": "The namespace of the metric.",
			//               "maxLength": 255,
			//               "minLength": 1,
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "MetricName",
			//             "Namespace"
			//           ],
			//           "type": "object"
			//         },
			//         "maxItems": 100,
			//         "type": "array",
			//         "uniqueItems": true
			//       }
			//     },
			//     "required": [
			//       "AdditionalStatistics",
			//       "IncludeMetrics"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 100,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed. You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"additional_statistics": {
						// Property: AdditionalStatistics
						Description: "The additional statistics to stream for the metrics listed in IncludeMetrics.",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
					"include_metrics": {
						// Property: IncludeMetrics
						Description: "An array that defines the metrics that are to have additional statistics streamed.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"metric_name": {
									// Property: MetricName
									Description: "The name of the metric.",
									Type:        types.StringType,
									Computed:    true,
								},
								"namespace": {
									// Property: Namespace
									Description: "The namespace of the metric.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A set of tags to assign to the delivery stream.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Metadata that you can assign to a Metric Stream, consisting of a key-value pair.",
			//     "properties": {
			//       "Key": {
			//         "description": "A unique identifier for the tag.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "An optional string, which you can use to describe or define the tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A set of tags to assign to the delivery stream.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A unique identifier for the tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "An optional string, which you can use to describe or define the tag.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
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
		Description: "Data Source schema for AWS::CloudWatch::MetricStream",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudWatch::MetricStream").WithTerraformTypeName("awscc_cloudwatch_metric_stream")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"additional_statistics":     "AdditionalStatistics",
		"arn":                       "Arn",
		"creation_date":             "CreationDate",
		"exclude_filters":           "ExcludeFilters",
		"firehose_arn":              "FirehoseArn",
		"include_filters":           "IncludeFilters",
		"include_metrics":           "IncludeMetrics",
		"key":                       "Key",
		"last_update_date":          "LastUpdateDate",
		"metric_name":               "MetricName",
		"name":                      "Name",
		"namespace":                 "Namespace",
		"output_format":             "OutputFormat",
		"role_arn":                  "RoleArn",
		"state":                     "State",
		"statistics_configurations": "StatisticsConfigurations",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
