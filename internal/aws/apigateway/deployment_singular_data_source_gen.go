// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_deployment", deploymentDataSource)
}

// deploymentDataSource returns the Terraform awscc_apigateway_deployment data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::Deployment resource.
func deploymentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DeploymentCanarySettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies settings for the canary deployment.",
		//	  "properties": {
		//	    "PercentTraffic": {
		//	      "description": "The percentage (0-100) of traffic diverted to a canary deployment.",
		//	      "type": "number"
		//	    },
		//	    "StageVariableOverrides": {
		//	      "additionalProperties": false,
		//	      "description": "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values. Duplicates are not allowed.",
		//	      "patternProperties": {
		//	        "": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "UseStageCache": {
		//	      "description": "Whether the canary deployment uses the stage cache.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"deployment_canary_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PercentTraffic
				"percent_traffic": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The percentage (0-100) of traffic diverted to a canary deployment.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: StageVariableOverrides
				"stage_variable_overrides": // Pattern: ""
				schema.MapAttribute{        /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values. Duplicates are not allowed.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UseStageCache
				"use_stage_cache": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Whether the canary deployment uses the stage cache.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies settings for the canary deployment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeploymentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Primary Id for this resource",
		//	  "type": "string"
		//	}
		"deployment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Primary Id for this resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the purpose of the API Gateway deployment.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the purpose of the API Gateway deployment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the RestApi resource to deploy. ",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the RestApi resource to deploy. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StageDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configures the stage that API Gateway creates with this deployment.",
		//	  "properties": {
		//	    "AccessLogSetting": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies settings for logging access in this stage.",
		//	      "properties": {
		//	        "DestinationArn": {
		//	          "description": "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with amazon-apigateway-. ",
		//	          "type": "string"
		//	        },
		//	        "Format": {
		//	          "description": "A single line format of the access logs of data, as specified by selected $context variables. The format must include at least $context.requestId. ",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "CacheClusterEnabled": {
		//	      "description": "Indicates whether cache clustering is enabled for the stage.",
		//	      "type": "boolean"
		//	    },
		//	    "CacheClusterSize": {
		//	      "description": "The size of the stage's cache cluster.",
		//	      "type": "string"
		//	    },
		//	    "CacheDataEncrypted": {
		//	      "description": "The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses. ",
		//	      "type": "boolean"
		//	    },
		//	    "CacheTtlInSeconds": {
		//	      "description": "The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses. ",
		//	      "type": "integer"
		//	    },
		//	    "CachingEnabled": {
		//	      "description": "Indicates whether responses are cached and returned for requests. You must enable a cache cluster on the stage to cache responses.",
		//	      "type": "boolean"
		//	    },
		//	    "CanarySetting": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies settings for the canary deployment in this stage.",
		//	      "properties": {
		//	        "PercentTraffic": {
		//	          "description": "The percent (0-100) of traffic diverted to a canary deployment.",
		//	          "type": "number"
		//	        },
		//	        "StageVariableOverrides": {
		//	          "additionalProperties": false,
		//	          "description": "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values. ",
		//	          "patternProperties": {
		//	            "": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "UseStageCache": {
		//	          "description": "Whether the canary deployment uses the stage cache or not.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ClientCertificateId": {
		//	      "description": "The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage. ",
		//	      "type": "string"
		//	    },
		//	    "DataTraceEnabled": {
		//	      "description": "Indicates whether data trace logging is enabled for methods in the stage. API Gateway pushes these logs to Amazon CloudWatch Logs. ",
		//	      "type": "boolean"
		//	    },
		//	    "Description": {
		//	      "description": "A description of the purpose of the stage.",
		//	      "type": "string"
		//	    },
		//	    "DocumentationVersion": {
		//	      "description": "The version identifier of the API documentation snapshot.",
		//	      "type": "string"
		//	    },
		//	    "LoggingLevel": {
		//	      "description": "The logging level for this method. For valid values, see the loggingLevel property of the Stage resource in the Amazon API Gateway API Reference. ",
		//	      "type": "string"
		//	    },
		//	    "MethodSettings": {
		//	      "description": "Configures settings for all of the stage's methods.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "CacheDataEncrypted": {
		//	            "description": "Indicates whether the cached responses are encrypted",
		//	            "type": "boolean"
		//	          },
		//	          "CacheTtlInSeconds": {
		//	            "description": "The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses. ",
		//	            "type": "integer"
		//	          },
		//	          "CachingEnabled": {
		//	            "description": "Indicates whether responses are cached and returned for requests. You must enable a cache cluster on the stage to cache responses.",
		//	            "type": "boolean"
		//	          },
		//	          "DataTraceEnabled": {
		//	            "description": "Indicates whether data trace logging is enabled for methods in the stage. API Gateway pushes these logs to Amazon CloudWatch Logs. ",
		//	            "type": "boolean"
		//	          },
		//	          "HttpMethod": {
		//	            "description": "The HTTP method.",
		//	            "type": "string"
		//	          },
		//	          "LoggingLevel": {
		//	            "description": "The logging level for this method. For valid values, see the loggingLevel property of the Stage resource in the Amazon API Gateway API Reference. ",
		//	            "type": "string"
		//	          },
		//	          "MetricsEnabled": {
		//	            "description": "Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.",
		//	            "type": "boolean"
		//	          },
		//	          "ResourcePath": {
		//	            "description": "The resource path for this method. Forward slashes (/) are encoded as ~1 and the initial slash must include a forward slash. ",
		//	            "type": "string"
		//	          },
		//	          "ThrottlingBurstLimit": {
		//	            "description": "The number of burst requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account.",
		//	            "type": "integer"
		//	          },
		//	          "ThrottlingRateLimit": {
		//	            "description": "The number of steady-state requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account.",
		//	            "type": "number"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "MetricsEnabled": {
		//	      "description": "Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.",
		//	      "type": "boolean"
		//	    },
		//	    "Tags": {
		//	      "description": "An array of arbitrary tags (key-value pairs) to associate with the stage.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Key": {
		//	            "description": "The key name of the tag",
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "description": "The value for the tag",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Value",
		//	          "Key"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "ThrottlingBurstLimit": {
		//	      "description": "The number of burst requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account.",
		//	      "type": "integer"
		//	    },
		//	    "ThrottlingRateLimit": {
		//	      "description": "The number of steady-state requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account.",
		//	      "type": "number"
		//	    },
		//	    "TracingEnabled": {
		//	      "description": "Specifies whether active tracing with X-ray is enabled for this stage.",
		//	      "type": "boolean"
		//	    },
		//	    "Variables": {
		//	      "additionalProperties": false,
		//	      "description": "A map that defines the stage variables. Variable names must consist of alphanumeric characters, and the values must match the following regular expression: [A-Za-z0-9-._~:/?#\u0026=,]+. ",
		//	      "patternProperties": {
		//	        "": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"stage_description": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AccessLogSetting
				"access_log_setting": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DestinationArn
						"destination_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with amazon-apigateway-. ",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Format
						"format": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "A single line format of the access logs of data, as specified by selected $context variables. The format must include at least $context.requestId. ",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specifies settings for logging access in this stage.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: CacheClusterEnabled
				"cache_cluster_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether cache clustering is enabled for the stage.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: CacheClusterSize
				"cache_cluster_size": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The size of the stage's cache cluster.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: CacheDataEncrypted
				"cache_data_encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: CacheTtlInSeconds
				"cache_ttl_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: CachingEnabled
				"caching_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether responses are cached and returned for requests. You must enable a cache cluster on the stage to cache responses.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: CanarySetting
				"canary_setting": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: PercentTraffic
						"percent_traffic": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Description: "The percent (0-100) of traffic diverted to a canary deployment.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: StageVariableOverrides
						"stage_variable_overrides": // Pattern: ""
						schema.MapAttribute{        /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values. ",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: UseStageCache
						"use_stage_cache": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Whether the canary deployment uses the stage cache or not.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specifies settings for the canary deployment in this stage.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ClientCertificateId
				"client_certificate_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DataTraceEnabled
				"data_trace_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether data trace logging is enabled for methods in the stage. API Gateway pushes these logs to Amazon CloudWatch Logs. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Description
				"description": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A description of the purpose of the stage.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DocumentationVersion
				"documentation_version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The version identifier of the API documentation snapshot.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: LoggingLevel
				"logging_level": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The logging level for this method. For valid values, see the loggingLevel property of the Stage resource in the Amazon API Gateway API Reference. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MethodSettings
				"method_settings": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: CacheDataEncrypted
							"cache_data_encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Indicates whether the cached responses are encrypted",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: CacheTtlInSeconds
							"cache_ttl_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses. ",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: CachingEnabled
							"caching_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Indicates whether responses are cached and returned for requests. You must enable a cache cluster on the stage to cache responses.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: DataTraceEnabled
							"data_trace_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Indicates whether data trace logging is enabled for methods in the stage. API Gateway pushes these logs to Amazon CloudWatch Logs. ",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: HttpMethod
							"http_method": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The HTTP method.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: LoggingLevel
							"logging_level": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The logging level for this method. For valid values, see the loggingLevel property of the Stage resource in the Amazon API Gateway API Reference. ",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: MetricsEnabled
							"metrics_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: ResourcePath
							"resource_path": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The resource path for this method. Forward slashes (/) are encoded as ~1 and the initial slash must include a forward slash. ",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: ThrottlingBurstLimit
							"throttling_burst_limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The number of burst requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: ThrottlingRateLimit
							"throttling_rate_limit": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "The number of steady-state requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Configures settings for all of the stage's methods.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MetricsEnabled
				"metrics_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Tags
				"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Key
							"key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The key name of the tag",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The value for the tag",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "An array of arbitrary tags (key-value pairs) to associate with the stage.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ThrottlingBurstLimit
				"throttling_burst_limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The number of burst requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ThrottlingRateLimit
				"throttling_rate_limit": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The number of steady-state requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: TracingEnabled
				"tracing_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether active tracing with X-ray is enabled for this stage.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Variables
				"variables":         // Pattern: ""
				schema.MapAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A map that defines the stage variables. Variable names must consist of alphanumeric characters, and the values must match the following regular expression: [A-Za-z0-9-._~:/?#&=,]+. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configures the stage that API Gateway creates with this deployment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StageName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the stage that API Gateway creates with this deployment. Use only alphanumeric characters.",
		//	  "type": "string"
		//	}
		"stage_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the stage that API Gateway creates with this deployment. Use only alphanumeric characters.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::Deployment",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::Deployment").WithTerraformTypeName("awscc_apigateway_deployment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_log_setting":         "AccessLogSetting",
		"cache_cluster_enabled":      "CacheClusterEnabled",
		"cache_cluster_size":         "CacheClusterSize",
		"cache_data_encrypted":       "CacheDataEncrypted",
		"cache_ttl_in_seconds":       "CacheTtlInSeconds",
		"caching_enabled":            "CachingEnabled",
		"canary_setting":             "CanarySetting",
		"client_certificate_id":      "ClientCertificateId",
		"data_trace_enabled":         "DataTraceEnabled",
		"deployment_canary_settings": "DeploymentCanarySettings",
		"deployment_id":              "DeploymentId",
		"description":                "Description",
		"destination_arn":            "DestinationArn",
		"documentation_version":      "DocumentationVersion",
		"format":                     "Format",
		"http_method":                "HttpMethod",
		"key":                        "Key",
		"logging_level":              "LoggingLevel",
		"method_settings":            "MethodSettings",
		"metrics_enabled":            "MetricsEnabled",
		"percent_traffic":            "PercentTraffic",
		"resource_path":              "ResourcePath",
		"rest_api_id":                "RestApiId",
		"stage_description":          "StageDescription",
		"stage_name":                 "StageName",
		"stage_variable_overrides":   "StageVariableOverrides",
		"tags":                       "Tags",
		"throttling_burst_limit":     "ThrottlingBurstLimit",
		"throttling_rate_limit":      "ThrottlingRateLimit",
		"tracing_enabled":            "TracingEnabled",
		"use_stage_cache":            "UseStageCache",
		"value":                      "Value",
		"variables":                  "Variables",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}