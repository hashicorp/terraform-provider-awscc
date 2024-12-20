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
		//	  "description": "The ``AccessLogSetting`` property type specifies settings for logging access in this stage.\n  ``AccessLogSetting`` is a property of the [AWS::ApiGateway::Stage](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html) resource.",
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
			Description: "The ``AccessLogSetting`` property type specifies settings for logging access in this stage.\n  ``AccessLogSetting`` is a property of the [AWS::ApiGateway::Stage](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html) resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CacheClusterEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "boolean"
		//	}
		"cache_cluster_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CacheClusterSize
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"cache_cluster_size": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CanarySetting
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "",
		//	  "properties": {
		//	    "DeploymentId": {
		//	      "description": "",
		//	      "type": "string"
		//	    },
		//	    "PercentTraffic": {
		//	      "description": "",
		//	      "maximum": 100,
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "StageVariableOverrides": {
		//	      "additionalProperties": false,
		//	      "description": "",
		//	      "patternProperties": {
		//	        "": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "UseStageCache": {
		//	      "description": "",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"canary_setting": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DeploymentId
				"deployment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PercentTraffic
				"percent_traffic": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: StageVariableOverrides
				"stage_variable_overrides": // Pattern: ""
				schema.MapAttribute{        /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UseStageCache
				"use_stage_cache": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ClientCertificateId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"client_certificate_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeploymentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"deployment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DocumentationVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"documentation_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MethodSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The ``MethodSetting`` property type configures settings for all methods in a stage.\n The ``MethodSettings`` property of the ``AWS::ApiGateway::Stage`` resource contains a list of ``MethodSetting`` property types.",
		//	    "properties": {
		//	      "CacheDataEncrypted": {
		//	        "description": "",
		//	        "type": "boolean"
		//	      },
		//	      "CacheTtlInSeconds": {
		//	        "description": "",
		//	        "type": "integer"
		//	      },
		//	      "CachingEnabled": {
		//	        "description": "",
		//	        "type": "boolean"
		//	      },
		//	      "DataTraceEnabled": {
		//	        "description": "",
		//	        "type": "boolean"
		//	      },
		//	      "HttpMethod": {
		//	        "description": "The HTTP method. To apply settings to multiple resources and methods, specify an asterisk (``*``) for the ``HttpMethod`` and ``/*`` for the ``ResourcePath``. This parameter is required when you specify a ``MethodSetting``.",
		//	        "type": "string"
		//	      },
		//	      "LoggingLevel": {
		//	        "description": "",
		//	        "type": "string"
		//	      },
		//	      "MetricsEnabled": {
		//	        "description": "",
		//	        "type": "boolean"
		//	      },
		//	      "ResourcePath": {
		//	        "description": "The resource path for this method. Forward slashes (``/``) are encoded as ``~1`` and the initial slash must include a forward slash. For example, the path value ``/resource/subresource`` must be encoded as ``/~1resource~1subresource``. To specify the root path, use only a slash (``/``). To apply settings to multiple resources and methods, specify an asterisk (``*``) for the ``HttpMethod`` and ``/*`` for the ``ResourcePath``. This parameter is required when you specify a ``MethodSetting``.",
		//	        "type": "string"
		//	      },
		//	      "ThrottlingBurstLimit": {
		//	        "description": "",
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      },
		//	      "ThrottlingRateLimit": {
		//	        "description": "",
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
						Description: "",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: CacheTtlInSeconds
					"cache_ttl_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: CachingEnabled
					"caching_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: DataTraceEnabled
					"data_trace_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: HttpMethod
					"http_method": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The HTTP method. To apply settings to multiple resources and methods, specify an asterisk (``*``) for the ``HttpMethod`` and ``/*`` for the ``ResourcePath``. This parameter is required when you specify a ``MethodSetting``.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: LoggingLevel
					"logging_level": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: MetricsEnabled
					"metrics_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ResourcePath
					"resource_path": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The resource path for this method. Forward slashes (``/``) are encoded as ``~1`` and the initial slash must include a forward slash. For example, the path value ``/resource/subresource`` must be encoded as ``/~1resource~1subresource``. To specify the root path, use only a slash (``/``). To apply settings to multiple resources and methods, specify an asterisk (``*``) for the ``HttpMethod`` and ``/*`` for the ``ResourcePath``. This parameter is required when you specify a ``MethodSetting``.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ThrottlingBurstLimit
					"throttling_burst_limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ThrottlingRateLimit
					"throttling_rate_limit": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Description: "",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StageName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"stage_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
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
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TracingEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "boolean"
		//	}
		"tracing_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "",
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
