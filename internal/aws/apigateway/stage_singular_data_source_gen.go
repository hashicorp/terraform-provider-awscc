// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_stage", stageDataSource)
}

// stageDataSource returns the Terraform awscc_apigateway_stage data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::Stage resource.
func stageDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessLogSetting
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Access log settings, including the access log format and access log destination ARN.",
		//	  "properties": {
		//	    "DestinationArn": {
		//	      "description": "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with ``amazon-apigateway-``. This parameter is required to enable access logging.",
		//	      "type": "string"
		//	    },
		//	    "Format": {
		//	      "description": "A single line format of the access logs of data, as specified by selected [$context variables](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference). The format must include at least ``$context.requestId``. This parameter is required to enable access logging.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"access_log_setting": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DestinationArn
				"destination_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with ``amazon-apigateway-``. This parameter is required to enable access logging.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Format
				"format": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A single line format of the access logs of data, as specified by selected [$context variables](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference). The format must include at least ``$context.requestId``. This parameter is required to enable access logging.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Access log settings, including the access log format and access log destination ARN.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CacheClusterEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether a cache cluster is enabled for the stage.",
		//	  "type": "boolean"
		//	}
		"cache_cluster_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether a cache cluster is enabled for the stage.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CacheClusterSize
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The stage's cache capacity in GB. For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html).",
		//	  "type": "string"
		//	}
		"cache_cluster_size": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The stage's cache capacity in GB. For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CanarySetting
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Settings for the canary deployment in this stage.",
		//	  "properties": {
		//	    "DeploymentId": {
		//	      "description": "The ID of the canary deployment.",
		//	      "type": "string"
		//	    },
		//	    "PercentTraffic": {
		//	      "description": "The percent (0-100) of traffic diverted to a canary deployment.",
		//	      "maximum": 100,
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "StageVariableOverrides": {
		//	      "additionalProperties": false,
		//	      "description": "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.",
		//	      "patternProperties": {
		//	        "": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "UseStageCache": {
		//	      "description": "A Boolean flag to indicate whether the canary deployment uses the stage cache or not.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"canary_setting": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DeploymentId
				"deployment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ID of the canary deployment.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PercentTraffic
				"percent_traffic": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The percent (0-100) of traffic diverted to a canary deployment.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: StageVariableOverrides
				"stage_variable_overrides": // Pattern: ""
				schema.MapAttribute{        /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UseStageCache
				"use_stage_cache": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "A Boolean flag to indicate whether the canary deployment uses the stage cache or not.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Settings for the canary deployment in this stage.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ClientCertificateId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of a client certificate for an API stage.",
		//	  "type": "string"
		//	}
		"client_certificate_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of a client certificate for an API stage.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeploymentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the Deployment that the stage points to.",
		//	  "type": "string"
		//	}
		"deployment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the Deployment that the stage points to.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The stage's description.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The stage's description.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DocumentationVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version of the associated API documentation.",
		//	  "type": "string"
		//	}
		"documentation_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version of the associated API documentation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MethodSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A map that defines the method settings for a Stage resource. Keys (designated as ``/{method_setting_key`` below) are method paths defined as ``{resource_path}/{http_method}`` for an individual method override, or ``/\\*/\\*`` for overriding all methods in the stage.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The ``MethodSetting`` property type configures settings for all methods in a stage.\n The ``MethodSettings`` property of the ``AWS::ApiGateway::Stage`` resource contains a list of ``MethodSetting`` property types.",
		//	    "properties": {
		//	      "CacheDataEncrypted": {
		//	        "description": "Specifies whether the cached responses are encrypted.",
		//	        "type": "boolean"
		//	      },
		//	      "CacheTtlInSeconds": {
		//	        "description": "Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.",
		//	        "type": "integer"
		//	      },
		//	      "CachingEnabled": {
		//	        "description": "Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.",
		//	        "type": "boolean"
		//	      },
		//	      "DataTraceEnabled": {
		//	        "description": "Specifies whether data trace logging is enabled for this method, which affects the log entries pushed to Amazon CloudWatch Logs. This can be useful to troubleshoot APIs, but can result in logging sensitive data. We recommend that you don't enable this option for production APIs.",
		//	        "type": "boolean"
		//	      },
		//	      "HttpMethod": {
		//	        "description": "The HTTP method. To apply settings to multiple resources and methods, specify an asterisk (``*``) for the ``HttpMethod`` and ``/*`` for the ``ResourcePath``. This parameter is required when you specify a ``MethodSetting``.",
		//	        "type": "string"
		//	      },
		//	      "LoggingLevel": {
		//	        "description": "Specifies the logging level for this method, which affects the log entries pushed to Amazon CloudWatch Logs. Valid values are ``OFF``, ``ERROR``, and ``INFO``. Choose ``ERROR`` to write only error-level entries to CloudWatch Logs, or choose ``INFO`` to include all ``ERROR`` events as well as extra informational events.",
		//	        "type": "string"
		//	      },
		//	      "MetricsEnabled": {
		//	        "description": "Specifies whether Amazon CloudWatch metrics are enabled for this method.",
		//	        "type": "boolean"
		//	      },
		//	      "ResourcePath": {
		//	        "description": "The resource path for this method. Forward slashes (``/``) are encoded as ``~1`` and the initial slash must include a forward slash. For example, the path value ``/resource/subresource`` must be encoded as ``/~1resource~1subresource``. To specify the root path, use only a slash (``/``). To apply settings to multiple resources and methods, specify an asterisk (``*``) for the ``HttpMethod`` and ``/*`` for the ``ResourcePath``. This parameter is required when you specify a ``MethodSetting``.",
		//	        "type": "string"
		//	      },
		//	      "ThrottlingBurstLimit": {
		//	        "description": "Specifies the throttling burst limit.",
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      },
		//	      "ThrottlingRateLimit": {
		//	        "description": "Specifies the throttling rate limit.",
		//	        "minimum": 0,
		//	        "type": "number"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"method_settings": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CacheDataEncrypted
					"cache_data_encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "Specifies whether the cached responses are encrypted.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: CacheTtlInSeconds
					"cache_ttl_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: CachingEnabled
					"caching_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: DataTraceEnabled
					"data_trace_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "Specifies whether data trace logging is enabled for this method, which affects the log entries pushed to Amazon CloudWatch Logs. This can be useful to troubleshoot APIs, but can result in logging sensitive data. We recommend that you don't enable this option for production APIs.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: HttpMethod
					"http_method": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The HTTP method. To apply settings to multiple resources and methods, specify an asterisk (``*``) for the ``HttpMethod`` and ``/*`` for the ``ResourcePath``. This parameter is required when you specify a ``MethodSetting``.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: LoggingLevel
					"logging_level": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Specifies the logging level for this method, which affects the log entries pushed to Amazon CloudWatch Logs. Valid values are ``OFF``, ``ERROR``, and ``INFO``. Choose ``ERROR`` to write only error-level entries to CloudWatch Logs, or choose ``INFO`` to include all ``ERROR`` events as well as extra informational events.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: MetricsEnabled
					"metrics_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "Specifies whether Amazon CloudWatch metrics are enabled for this method.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ResourcePath
					"resource_path": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The resource path for this method. Forward slashes (``/``) are encoded as ``~1`` and the initial slash must include a forward slash. For example, the path value ``/resource/subresource`` must be encoded as ``/~1resource~1subresource``. To specify the root path, use only a slash (``/``). To apply settings to multiple resources and methods, specify an asterisk (``*``) for the ``HttpMethod`` and ``/*`` for the ``ResourcePath``. This parameter is required when you specify a ``MethodSetting``.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ThrottlingBurstLimit
					"throttling_burst_limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "Specifies the throttling burst limit.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ThrottlingRateLimit
					"throttling_rate_limit": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Description: "Specifies the throttling rate limit.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A map that defines the method settings for a Stage resource. Keys (designated as ``/{method_setting_key`` below) are method paths defined as ``{resource_path}/{http_method}`` for an individual method override, or ``/\\*/\\*`` for overriding all methods in the stage.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The string identifier of the associated RestApi.",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The string identifier of the associated RestApi.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StageName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.",
		//	  "type": "string"
		//	}
		"stage_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The collection of tags. Each tag element is associated with a given resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:.",
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The collection of tags. Each tag element is associated with a given resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TracingEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether active tracing with X-ray is enabled for the Stage.",
		//	  "type": "boolean"
		//	}
		"tracing_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether active tracing with X-ray is enabled for the Stage.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Variables
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value. Variable names are limited to alphanumeric characters. Values must match the following regular expression: ``[A-Za-z0-9-._~:/?#\u0026=,]+``.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"variables":         // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value. Variable names are limited to alphanumeric characters. Values must match the following regular expression: ``[A-Za-z0-9-._~:/?#&=,]+``.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::Stage",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::Stage").WithTerraformTypeName("awscc_apigateway_stage")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_log_setting":       "AccessLogSetting",
		"cache_cluster_enabled":    "CacheClusterEnabled",
		"cache_cluster_size":       "CacheClusterSize",
		"cache_data_encrypted":     "CacheDataEncrypted",
		"cache_ttl_in_seconds":     "CacheTtlInSeconds",
		"caching_enabled":          "CachingEnabled",
		"canary_setting":           "CanarySetting",
		"client_certificate_id":    "ClientCertificateId",
		"data_trace_enabled":       "DataTraceEnabled",
		"deployment_id":            "DeploymentId",
		"description":              "Description",
		"destination_arn":          "DestinationArn",
		"documentation_version":    "DocumentationVersion",
		"format":                   "Format",
		"http_method":              "HttpMethod",
		"key":                      "Key",
		"logging_level":            "LoggingLevel",
		"method_settings":          "MethodSettings",
		"metrics_enabled":          "MetricsEnabled",
		"percent_traffic":          "PercentTraffic",
		"resource_path":            "ResourcePath",
		"rest_api_id":              "RestApiId",
		"stage_name":               "StageName",
		"stage_variable_overrides": "StageVariableOverrides",
		"tags":                     "Tags",
		"throttling_burst_limit":   "ThrottlingBurstLimit",
		"throttling_rate_limit":    "ThrottlingRateLimit",
		"tracing_enabled":          "TracingEnabled",
		"use_stage_cache":          "UseStageCache",
		"value":                    "Value",
		"variables":                "Variables",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
