// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ecs_cluster", clusterResource)
}

// clusterResource returns the Terraform awscc_ecs_cluster resource.
// This Terraform resource corresponds to the CloudFormation AWS::ECS::Cluster resource.
func clusterResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CapacityProviders
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The short name of one or more capacity providers to associate with the cluster. A capacity provider must be associated with a cluster before it can be included as part of the default capacity provider strategy of the cluster or used in a capacity provider strategy when calling the [CreateService](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html) or [RunTask](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) actions.\n If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must be created but not associated with another cluster. New Auto Scaling group capacity providers can be created with the [CreateCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateCapacityProvider.html) API operation.\n To use a FARGATElong capacity provider, specify either the ``FARGATE`` or ``FARGATE_SPOT`` capacity providers. The FARGATElong capacity providers are available to all accounts and only need to be associated with a cluster to be used.\n The [PutCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutCapacityProvider.html) API operation is used to update the list of available capacity providers for a cluster after the cluster is created.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"capacity_providers": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The short name of one or more capacity providers to associate with the cluster. A capacity provider must be associated with a cluster before it can be included as part of the default capacity provider strategy of the cluster or used in a capacity provider strategy when calling the [CreateService](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html) or [RunTask](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) actions.\n If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must be created but not associated with another cluster. New Auto Scaling group capacity providers can be created with the [CreateCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateCapacityProvider.html) API operation.\n To use a FARGATElong capacity provider, specify either the ``FARGATE`` or ``FARGATE_SPOT`` capacity providers. The FARGATElong capacity providers are available to all accounts and only need to be associated with a cluster to be used.\n The [PutCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutCapacityProvider.html) API operation is used to update the list of available capacity providers for a cluster after the cluster is created.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ClusterName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A user-generated string that you use to identify your cluster. If you don't specify a name, CFNlong generates a unique physical ID for the name.",
		//	  "type": "string"
		//	}
		"cluster_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A user-generated string that you use to identify your cluster. If you don't specify a name, CFNlong generates a unique physical ID for the name.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ClusterSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The settings to use when creating a cluster. This parameter is used to turn on CloudWatch Container Insights for a cluster.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The settings to use when creating a cluster. This parameter is used to turn on CloudWatch Container Insights for a cluster.",
		//	    "properties": {
		//	      "Name": {
		//	        "description": "The name of the cluster setting. The value is ``containerInsights`` .",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value to set for the cluster setting. The supported values are ``enabled`` and ``disabled``. \n If you set ``name`` to ``containerInsights`` and ``value`` to ``enabled``, CloudWatch Container Insights will be on for the cluster, otherwise it will be off unless the ``containerInsights`` account setting is turned on. If a cluster value is specified, it will override the ``containerInsights`` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html).",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"cluster_settings": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the cluster setting. The value is ``containerInsights`` .",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value to set for the cluster setting. The supported values are ``enabled`` and ``disabled``. \n If you set ``name`` to ``containerInsights`` and ``value`` to ``enabled``, CloudWatch Container Insights will be on for the cluster, otherwise it will be off unless the ``containerInsights`` account setting is turned on. If a cluster value is specified, it will override the ``containerInsights`` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html).",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The settings to use when creating a cluster. This parameter is used to turn on CloudWatch Container Insights for a cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Configuration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The execute command and managed storage configuration for the cluster.",
		//	  "properties": {
		//	    "ExecuteCommandConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "The details of the execute command configuration.",
		//	      "properties": {
		//	        "KmsKeyId": {
		//	          "description": "Specify an KMSlong key ID to encrypt the data between the local client and the container.",
		//	          "relationshipRef": {
		//	            "propertyPath": "/properties/Arn",
		//	            "typeName": "AWS::KMS::Key"
		//	          },
		//	          "type": "string"
		//	        },
		//	        "LogConfiguration": {
		//	          "additionalProperties": false,
		//	          "description": "The log configuration for the results of the execute command actions. The logs can be sent to CloudWatch Logs or an Amazon S3 bucket. When ``logging=OVERRIDE`` is specified, a ``logConfiguration`` must be provided.",
		//	          "properties": {
		//	            "CloudWatchEncryptionEnabled": {
		//	              "description": "Determines whether to use encryption on the CloudWatch logs. If not specified, encryption will be off.",
		//	              "type": "boolean"
		//	            },
		//	            "CloudWatchLogGroupName": {
		//	              "description": "The name of the CloudWatch log group to send logs to.\n  The CloudWatch log group must already be created.",
		//	              "relationshipRef": {
		//	                "propertyPath": "/properties/LogGroupName",
		//	                "typeName": "AWS::Logs::LogGroup"
		//	              },
		//	              "type": "string"
		//	            },
		//	            "S3BucketName": {
		//	              "description": "The name of the S3 bucket to send logs to.\n  The S3 bucket must already be created.",
		//	              "type": "string"
		//	            },
		//	            "S3EncryptionEnabled": {
		//	              "description": "Determines whether to use encryption on the S3 logs. If not specified, encryption is not used.",
		//	              "type": "boolean"
		//	            },
		//	            "S3KeyPrefix": {
		//	              "description": "An optional folder in the S3 bucket to place logs in.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "Logging": {
		//	          "description": "The log setting to use for redirecting logs for your execute command results. The following log settings are available.\n  +   ``NONE``: The execute command session is not logged.\n  +   ``DEFAULT``: The ``awslogs`` configuration in the task definition is used. If no logging parameter is specified, it defaults to this value. If no ``awslogs`` log driver is configured in the task definition, the output won't be logged.\n  +   ``OVERRIDE``: Specify the logging details as a part of ``logConfiguration``. If the ``OVERRIDE`` logging option is specified, the ``logConfiguration`` is required.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ManagedStorageConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "The details of the managed storage configuration.",
		//	      "properties": {
		//	        "FargateEphemeralStorageKmsKeyId": {
		//	          "description": "Specify the KMSlong key ID for the Fargate ephemeral storage.",
		//	          "type": "string"
		//	        },
		//	        "KmsKeyId": {
		//	          "description": "Specify a KMSlong key ID to encrypt the managed storage.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ExecuteCommandConfiguration
				"execute_command_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: KmsKeyId
						"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Specify an KMSlong key ID to encrypt the data between the local client and the container.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: LogConfiguration
						"log_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: CloudWatchEncryptionEnabled
								"cloudwatch_encryption_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Determines whether to use encryption on the CloudWatch logs. If not specified, encryption will be off.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
										boolplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: CloudWatchLogGroupName
								"cloudwatch_log_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the CloudWatch log group to send logs to.\n  The CloudWatch log group must already be created.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: S3BucketName
								"s3_bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the S3 bucket to send logs to.\n  The S3 bucket must already be created.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: S3EncryptionEnabled
								"s3_encryption_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Determines whether to use encryption on the S3 logs. If not specified, encryption is not used.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
										boolplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: S3KeyPrefix
								"s3_key_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "An optional folder in the S3 bucket to place logs in.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "The log configuration for the results of the execute command actions. The logs can be sent to CloudWatch Logs or an Amazon S3 bucket. When ``logging=OVERRIDE`` is specified, a ``logConfiguration`` must be provided.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Logging
						"logging": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The log setting to use for redirecting logs for your execute command results. The following log settings are available.\n  +   ``NONE``: The execute command session is not logged.\n  +   ``DEFAULT``: The ``awslogs`` configuration in the task definition is used. If no logging parameter is specified, it defaults to this value. If no ``awslogs`` log driver is configured in the task definition, the output won't be logged.\n  +   ``OVERRIDE``: Specify the logging details as a part of ``logConfiguration``. If the ``OVERRIDE`` logging option is specified, the ``logConfiguration`` is required.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The details of the execute command configuration.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ManagedStorageConfiguration
				"managed_storage_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: FargateEphemeralStorageKmsKeyId
						"fargate_ephemeral_storage_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Specify the KMSlong key ID for the Fargate ephemeral storage.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: KmsKeyId
						"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Specify a KMSlong key ID to encrypt the managed storage.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The details of the managed storage configuration.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The execute command and managed storage configuration for the cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultCapacityProviderStrategy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default capacity provider strategy for the cluster. When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The ``CapacityProviderStrategyItem`` property specifies the details of the default capacity provider strategy for the cluster. When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.",
		//	    "properties": {
		//	      "Base": {
		//	        "description": "The *base* value designates how many tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a *base* defined. If no value is specified, the default value of ``0`` is used.",
		//	        "type": "integer"
		//	      },
		//	      "CapacityProvider": {
		//	        "description": "The short name of the capacity provider.",
		//	        "relationshipRef": {
		//	          "propertyPath": "/properties/Name",
		//	          "typeName": "AWS::ECS::CapacityProvider"
		//	        },
		//	        "type": "string"
		//	      },
		//	      "Weight": {
		//	        "description": "The *weight* value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider. The ``weight`` value is taken into consideration after the ``base`` value, if defined, is satisfied.\n If no ``weight`` value is specified, the default value of ``0`` is used. When multiple capacity providers are specified within a capacity provider strategy, at least one of the capacity providers must have a weight value greater than zero and any capacity providers with a weight of ``0`` can't be used to place tasks. If you specify multiple capacity providers in a strategy that all have a weight of ``0``, any ``RunTask`` or ``CreateService`` actions using the capacity provider strategy will fail.\n An example scenario for using weights is defining a strategy that contains two capacity providers and both have a weight of ``1``, then when the ``base`` is satisfied, the tasks will be split evenly across the two capacity providers. Using that same logic, if you specify a weight of ``1`` for *capacityProviderA* and a weight of ``4`` for *capacityProviderB*, then for every one task that's run using *capacityProviderA*, four tasks would use *capacityProviderB*.",
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"default_capacity_provider_strategy": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Base
					"base": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "The *base* value designates how many tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a *base* defined. If no value is specified, the default value of ``0`` is used.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: CapacityProvider
					"capacity_provider": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The short name of the capacity provider.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Weight
					"weight": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "The *weight* value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider. The ``weight`` value is taken into consideration after the ``base`` value, if defined, is satisfied.\n If no ``weight`` value is specified, the default value of ``0`` is used. When multiple capacity providers are specified within a capacity provider strategy, at least one of the capacity providers must have a weight value greater than zero and any capacity providers with a weight of ``0`` can't be used to place tasks. If you specify multiple capacity providers in a strategy that all have a weight of ``0``, any ``RunTask`` or ``CreateService`` actions using the capacity provider strategy will fail.\n An example scenario for using weights is defining a strategy that contains two capacity providers and both have a weight of ``1``, then when the ``base`` is satisfied, the tasks will be split evenly across the two capacity providers. Using that same logic, if you specify a weight of ``1`` for *capacityProviderA* and a weight of ``4`` for *capacityProviderB*, then for every one task that's run using *capacityProviderA*, four tasks would use *capacityProviderB*.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The default capacity provider strategy for the cluster. When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceConnectDefaults
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Use this parameter to set a default Service Connect namespace. After you set a default Service Connect namespace, any new services with Service Connect turned on that are created in the cluster are added as client services in the namespace. This setting only applies to new services that set the ``enabled`` parameter to ``true`` in the ``ServiceConnectConfiguration``. You can set the namespace of each service individually in the ``ServiceConnectConfiguration`` to override this default parameter.\n Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.",
		//	  "properties": {
		//	    "Namespace": {
		//	      "description": "The namespace name or full Amazon Resource Name (ARN) of the CMAPlong namespace that's used when you create a service and don't specify a Service Connect configuration. The namespace name can include up to 1024 characters. The name is case-sensitive. The name can't include hyphens (-), tilde (~), greater than (\u003e), less than (\u003c), or slash (/).\n If you enter an existing namespace name or ARN, then that namespace will be used. Any namespace type is supported. The namespace must be in this account and this AWS Region.\n If you enter a new name, a CMAPlong namespace will be created. Amazon ECS creates a CMAP namespace with the \"API calls\" method of instance discovery only. This instance discovery method is the \"HTTP\" namespace type in the CLIlong. Other types of instance discovery aren't used by Service Connect.\n If you update the cluster with an empty string ``\"\"`` for the namespace name, the cluster configuration for Service Connect is removed. Note that the namespace will remain in CMAP and must be deleted separately.\n For more information about CMAPlong, see [Working with Services](https://docs.aws.amazon.com/cloud-map/latest/dg/working-with-services.html) in the *Developer Guide*.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"service_connect_defaults": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Namespace
				"namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The namespace name or full Amazon Resource Name (ARN) of the CMAPlong namespace that's used when you create a service and don't specify a Service Connect configuration. The namespace name can include up to 1024 characters. The name is case-sensitive. The name can't include hyphens (-), tilde (~), greater than (>), less than (<), or slash (/).\n If you enter an existing namespace name or ARN, then that namespace will be used. Any namespace type is supported. The namespace must be in this account and this AWS Region.\n If you enter a new name, a CMAPlong namespace will be created. Amazon ECS creates a CMAP namespace with the \"API calls\" method of instance discovery only. This instance discovery method is the \"HTTP\" namespace type in the CLIlong. Other types of instance discovery aren't used by Service Connect.\n If you update the cluster with an empty string ``\"\"`` for the namespace name, the cluster configuration for Service Connect is removed. Note that the namespace will remain in CMAP and must be deleted separately.\n For more information about CMAPlong, see [Working with Services](https://docs.aws.amazon.com/cloud-map/latest/dg/working-with-services.html) in the *Developer Guide*.",
					Optional:    true,
					// Namespace is a write-only property.
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Use this parameter to set a default Service Connect namespace. After you set a default Service Connect namespace, any new services with Service Connect turned on that are created in the cluster are added as client services in the namespace. This setting only applies to new services that set the ``enabled`` parameter to ``true`` in the ``ServiceConnectConfiguration``. You can set the namespace of each service individually in the ``ServiceConnectConfiguration`` to override this default parameter.\n Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.",
			Optional:    true,
			// ServiceConnectDefaults is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The metadata that you apply to the cluster to help you categorize and organize them. Each tag consists of a key and an optional value. You define both.\n The following basic restrictions apply to tags:\n  +  Maximum number of tags per resource - 50\n  +  For each resource, each tag key must be unique, and each tag key can have only one value.\n  +  Maximum key length - 128 Unicode characters in UTF-8\n  +  Maximum value length - 256 Unicode characters in UTF-8\n  +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.\n  +  Tag keys and values are case-sensitive.\n  +  Do not use ``aws:``, ``AWS:``, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The metadata that you apply to a resource to help you categorize and organize them. Each tag consists of a key and an optional value. You define them.\n The following basic restrictions apply to tags:\n  +  Maximum number of tags per resource - 50\n  +  For each resource, each tag key must be unique, and each tag key can have only one value.\n  +  Maximum key length - 128 Unicode characters in UTF-8\n  +  Maximum value length - 256 Unicode characters in UTF-8\n  +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.\n  +  Tag keys and values are case-sensitive.\n  +  Do not use ``aws:``, ``AWS:``, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "One part of a key-value pair that make up a tag. A ``key`` is a general label that acts like a category for more specific tag values.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The optional part of a key-value pair that make up a tag. A ``value`` acts as a descriptor within a tag category (key).",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "One part of a key-value pair that make up a tag. A ``key`` is a general label that acts like a category for more specific tag values.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The optional part of a key-value pair that make up a tag. A ``value`` acts as a descriptor within a tag category (key).",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The metadata that you apply to the cluster to help you categorize and organize them. Each tag consists of a key and an optional value. You define both.\n The following basic restrictions apply to tags:\n  +  Maximum number of tags per resource - 50\n  +  For each resource, each tag key must be unique, and each tag key can have only one value.\n  +  Maximum key length - 128 Unicode characters in UTF-8\n  +  Maximum value length - 256 Unicode characters in UTF-8\n  +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.\n  +  Tag keys and values are case-sensitive.\n  +  Do not use ``aws:``, ``AWS:``, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
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
		Description: "The ``AWS::ECS::Cluster`` resource creates an Amazon Elastic Container Service (Amazon ECS) cluster.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::Cluster").WithTerraformTypeName("awscc_ecs_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                  "Arn",
		"base":                                 "Base",
		"capacity_provider":                    "CapacityProvider",
		"capacity_providers":                   "CapacityProviders",
		"cloudwatch_encryption_enabled":        "CloudWatchEncryptionEnabled",
		"cloudwatch_log_group_name":            "CloudWatchLogGroupName",
		"cluster_name":                         "ClusterName",
		"cluster_settings":                     "ClusterSettings",
		"configuration":                        "Configuration",
		"default_capacity_provider_strategy":   "DefaultCapacityProviderStrategy",
		"execute_command_configuration":        "ExecuteCommandConfiguration",
		"fargate_ephemeral_storage_kms_key_id": "FargateEphemeralStorageKmsKeyId",
		"key":                                  "Key",
		"kms_key_id":                           "KmsKeyId",
		"log_configuration":                    "LogConfiguration",
		"logging":                              "Logging",
		"managed_storage_configuration":        "ManagedStorageConfiguration",
		"name":                                 "Name",
		"namespace":                            "Namespace",
		"s3_bucket_name":                       "S3BucketName",
		"s3_encryption_enabled":                "S3EncryptionEnabled",
		"s3_key_prefix":                        "S3KeyPrefix",
		"service_connect_defaults":             "ServiceConnectDefaults",
		"tags":                                 "Tags",
		"value":                                "Value",
		"weight":                               "Weight",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ServiceConnectDefaults",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
