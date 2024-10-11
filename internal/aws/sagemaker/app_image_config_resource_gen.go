// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_sagemaker_app_image_config", appImageConfigResource)
}

// appImageConfigResource returns the Terraform awscc_sagemaker_app_image_config resource.
// This Terraform resource corresponds to the CloudFormation AWS::SageMaker::AppImageConfig resource.
func appImageConfigResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AppImageConfigArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the AppImageConfig.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "arn:aws[a-z\\-]*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:app-image-config/.*",
		//	  "type": "string"
		//	}
		"app_image_config_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the AppImageConfig.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AppImageConfigName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Name of the AppImageConfig.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}",
		//	  "type": "string"
		//	}
		"app_image_config_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Name of the AppImageConfig.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CodeEditorAppImageConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The CodeEditorAppImageConfig.",
		//	  "properties": {
		//	    "ContainerConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The container configuration for a SageMaker image.",
		//	      "properties": {
		//	        "ContainerArguments": {
		//	          "description": "A list of arguments to apply to the container.",
		//	          "items": {
		//	            "description": "The container image arguments",
		//	            "maxLength": 64,
		//	            "minLength": 1,
		//	            "pattern": "",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 50,
		//	          "minItems": 0,
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        },
		//	        "ContainerEntrypoint": {
		//	          "description": "The custom entry point to use on container.",
		//	          "items": {
		//	            "description": "The container entry point",
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "pattern": "",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 1,
		//	          "minItems": 0,
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        },
		//	        "ContainerEnvironmentVariables": {
		//	          "description": "A list of variables to apply to the custom container.",
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Key": {
		//	                "maxLength": 256,
		//	                "minLength": 1,
		//	                "pattern": "",
		//	                "type": "string"
		//	              },
		//	              "Value": {
		//	                "maxLength": 256,
		//	                "minLength": 1,
		//	                "pattern": "",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Key",
		//	              "Value"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "maxItems": 25,
		//	          "minItems": 0,
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"code_editor_app_image_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ContainerConfig
				"container_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ContainerArguments
						"container_arguments": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "A list of arguments to apply to the container.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(0, 50),
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 64),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ContainerEntrypoint
						"container_entrypoint": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The custom entry point to use on container.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(0, 1),
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 256),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ContainerEnvironmentVariables
						"container_environment_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Key
									"key": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(1, 256),
											fwvalidators.NotNullString(),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Value
									"value": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(1, 256),
											fwvalidators.NotNullString(),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Description: "A list of variables to apply to the custom container.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(0, 25),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The container configuration for a SageMaker image.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The CodeEditorAppImageConfig.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: JupyterLabAppImageConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The JupyterLabAppImageConfig.",
		//	  "properties": {
		//	    "ContainerConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The container configuration for a SageMaker image.",
		//	      "properties": {
		//	        "ContainerArguments": {
		//	          "description": "A list of arguments to apply to the container.",
		//	          "items": {
		//	            "description": "The container image arguments",
		//	            "maxLength": 64,
		//	            "minLength": 1,
		//	            "pattern": "",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 50,
		//	          "minItems": 0,
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        },
		//	        "ContainerEntrypoint": {
		//	          "description": "The custom entry point to use on container.",
		//	          "items": {
		//	            "description": "The container entry point",
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "pattern": "",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 1,
		//	          "minItems": 0,
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        },
		//	        "ContainerEnvironmentVariables": {
		//	          "description": "A list of variables to apply to the custom container.",
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Key": {
		//	                "maxLength": 256,
		//	                "minLength": 1,
		//	                "pattern": "",
		//	                "type": "string"
		//	              },
		//	              "Value": {
		//	                "maxLength": 256,
		//	                "minLength": 1,
		//	                "pattern": "",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Key",
		//	              "Value"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "maxItems": 25,
		//	          "minItems": 0,
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"jupyter_lab_app_image_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ContainerConfig
				"container_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ContainerArguments
						"container_arguments": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "A list of arguments to apply to the container.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(0, 50),
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 64),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ContainerEntrypoint
						"container_entrypoint": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The custom entry point to use on container.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(0, 1),
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 256),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ContainerEnvironmentVariables
						"container_environment_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Key
									"key": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(1, 256),
											fwvalidators.NotNullString(),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Value
									"value": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(1, 256),
											fwvalidators.NotNullString(),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Description: "A list of variables to apply to the custom container.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(0, 25),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The container configuration for a SageMaker image.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The JupyterLabAppImageConfig.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KernelGatewayImageConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The KernelGatewayImageConfig.",
		//	  "properties": {
		//	    "FileSystemConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.",
		//	      "properties": {
		//	        "DefaultGid": {
		//	          "description": "The default POSIX group ID (GID). If not specified, defaults to 100.",
		//	          "maximum": 65535,
		//	          "minimum": 0,
		//	          "type": "integer"
		//	        },
		//	        "DefaultUid": {
		//	          "description": "The default POSIX user ID (UID). If not specified, defaults to 1000.",
		//	          "maximum": 65535,
		//	          "minimum": 0,
		//	          "type": "integer"
		//	        },
		//	        "MountPath": {
		//	          "description": "The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.",
		//	          "maxLength": 1024,
		//	          "minLength": 1,
		//	          "pattern": "^/.*",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "KernelSpecs": {
		//	      "description": "The specification of the Jupyter kernels in the image.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "DisplayName": {
		//	            "description": "The display name of the kernel.",
		//	            "maxLength": 1024,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "description": "The name of the kernel.",
		//	            "maxLength": 1024,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 1,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "KernelSpecs"
		//	  ],
		//	  "type": "object"
		//	}
		"kernel_gateway_image_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FileSystemConfig
				"file_system_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DefaultGid
						"default_gid": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The default POSIX group ID (GID). If not specified, defaults to 100.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.Between(0, 65535),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: DefaultUid
						"default_uid": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The default POSIX user ID (UID). If not specified, defaults to 1000.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.Between(0, 65535),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: MountPath
						"mount_path": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 1024),
								stringvalidator.RegexMatches(regexp.MustCompile("^/.*"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: KernelSpecs
				"kernel_specs": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: DisplayName
							"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The display name of the kernel.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 1024),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the kernel.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 1024),
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The specification of the Jupyter kernels in the image.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(1, 1),
						fwvalidators.NotNullList(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The KernelGatewayImageConfig.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags to apply to the AppImageConfig.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
							stringplanmodifier.RequiresReplaceIfConfigured(),
						}, /*END PLAN MODIFIERS*/
						// Key is a write-only property.
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
							stringplanmodifier.RequiresReplaceIfConfigured(),
						}, /*END PLAN MODIFIERS*/
						// Value is a write-only property.
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of tags to apply to the AppImageConfig.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// Tags is a write-only property.
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
		Description: "Resource Type definition for AWS::SageMaker::AppImageConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::AppImageConfig").WithTerraformTypeName("awscc_sagemaker_app_image_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_image_config_arn":            "AppImageConfigArn",
		"app_image_config_name":           "AppImageConfigName",
		"code_editor_app_image_config":    "CodeEditorAppImageConfig",
		"container_arguments":             "ContainerArguments",
		"container_config":                "ContainerConfig",
		"container_entrypoint":            "ContainerEntrypoint",
		"container_environment_variables": "ContainerEnvironmentVariables",
		"default_gid":                     "DefaultGid",
		"default_uid":                     "DefaultUid",
		"display_name":                    "DisplayName",
		"file_system_config":              "FileSystemConfig",
		"jupyter_lab_app_image_config":    "JupyterLabAppImageConfig",
		"kernel_gateway_image_config":     "KernelGatewayImageConfig",
		"kernel_specs":                    "KernelSpecs",
		"key":                             "Key",
		"mount_path":                      "MountPath",
		"name":                            "Name",
		"tags":                            "Tags",
		"value":                           "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
