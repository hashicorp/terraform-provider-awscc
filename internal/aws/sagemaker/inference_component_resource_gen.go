// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
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
	registry.AddResourceFactory("awscc_sagemaker_inference_component", inferenceComponentResource)
}

// inferenceComponentResource returns the Terraform awscc_sagemaker_inference_component resource.
// This Terraform resource corresponds to the CloudFormation AWS::SageMaker::InferenceComponent resource.
func inferenceComponentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EndpointArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the endpoint the inference component is associated with",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"endpoint_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the endpoint the inference component is associated with",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EndpointName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the endpoint the inference component is associated with",
		//	  "maxLength": 63,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"endpoint_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the endpoint the inference component is associated with",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: FailureReason
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The failure reason if the inference component is in a failed state",
		//	  "maxLength": 63,
		//	  "type": "string"
		//	}
		"failure_reason": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The failure reason if the inference component is in a failed state",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InferenceComponentArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the inference component",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"inference_component_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the inference component",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InferenceComponentName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the inference component",
		//	  "maxLength": 63,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"inference_component_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the inference component",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InferenceComponentStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "InService",
		//	    "Creating",
		//	    "Updating",
		//	    "Failed",
		//	    "Deleting"
		//	  ],
		//	  "type": "string"
		//	}
		"inference_component_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"last_modified_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RuntimeConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The runtime config for the inference component",
		//	  "properties": {
		//	    "CopyCount": {
		//	      "description": "The number of copies for the inference component",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "CurrentCopyCount": {
		//	      "description": "The number of copies for the inference component",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "DesiredCopyCount": {
		//	      "description": "The number of copies for the inference component",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"runtime_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CopyCount
				"copy_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The number of copies for the inference component",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(0),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
					// CopyCount is a write-only property.
				}, /*END ATTRIBUTE*/
				// Property: CurrentCopyCount
				"current_copy_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The number of copies for the inference component",
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DesiredCopyCount
				"desired_copy_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The number of copies for the inference component",
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The runtime config for the inference component",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Specification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The specification for the inference component",
		//	  "properties": {
		//	    "BaseInferenceComponentName": {
		//	      "description": "The name of the base inference component",
		//	      "maxLength": 63,
		//	      "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	      "type": "string"
		//	    },
		//	    "ComputeResourceRequirements": {
		//	      "additionalProperties": false,
		//	      "description": "",
		//	      "properties": {
		//	        "MaxMemoryRequiredInMb": {
		//	          "minimum": 128,
		//	          "type": "integer"
		//	        },
		//	        "MinMemoryRequiredInMb": {
		//	          "minimum": 128,
		//	          "type": "integer"
		//	        },
		//	        "NumberOfAcceleratorDevicesRequired": {
		//	          "minimum": 1,
		//	          "type": "number"
		//	        },
		//	        "NumberOfCpuCoresRequired": {
		//	          "minimum": 0.25,
		//	          "type": "number"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Container": {
		//	      "additionalProperties": false,
		//	      "description": "",
		//	      "properties": {
		//	        "ArtifactUrl": {
		//	          "maxLength": 1024,
		//	          "pattern": "^(https|s3)://([^/]+)/?(.*)$",
		//	          "type": "string"
		//	        },
		//	        "DeployedImage": {
		//	          "additionalProperties": false,
		//	          "description": "",
		//	          "properties": {
		//	            "ResolutionTime": {
		//	              "type": "string"
		//	            },
		//	            "ResolvedImage": {
		//	              "description": "The image to use for the container that will be materialized for the inference component",
		//	              "maxLength": 255,
		//	              "pattern": "[\\S]+",
		//	              "type": "string"
		//	            },
		//	            "SpecifiedImage": {
		//	              "description": "The image to use for the container that will be materialized for the inference component",
		//	              "maxLength": 255,
		//	              "pattern": "[\\S]+",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "Environment": {
		//	          "additionalProperties": false,
		//	          "description": "Environment variables to specify on the container",
		//	          "patternProperties": {
		//	            "": {
		//	              "maxLength": 1024,
		//	              "pattern": "^[\\S\\s]*$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "Image": {
		//	          "description": "The image to use for the container that will be materialized for the inference component",
		//	          "maxLength": 255,
		//	          "pattern": "[\\S]+",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ModelName": {
		//	      "description": "The name of the model to use with the inference component",
		//	      "maxLength": 63,
		//	      "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	      "type": "string"
		//	    },
		//	    "StartupParameters": {
		//	      "additionalProperties": false,
		//	      "description": "",
		//	      "properties": {
		//	        "ContainerStartupHealthCheckTimeoutInSeconds": {
		//	          "maximum": 3600,
		//	          "minimum": 60,
		//	          "type": "integer"
		//	        },
		//	        "ModelDataDownloadTimeoutInSeconds": {
		//	          "maximum": 3600,
		//	          "minimum": 60,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BaseInferenceComponentName
				"base_inference_component_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the base inference component",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(63),
						stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ComputeResourceRequirements
				"compute_resource_requirements": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: MaxMemoryRequiredInMb
						"max_memory_required_in_mb": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.AtLeast(128),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: MinMemoryRequiredInMb
						"min_memory_required_in_mb": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.AtLeast(128),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: NumberOfAcceleratorDevicesRequired
						"number_of_accelerator_devices_required": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Float64{ /*START VALIDATORS*/
								float64validator.AtLeast(1.000000),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
								float64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: NumberOfCpuCoresRequired
						"number_of_cpu_cores_required": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Float64{ /*START VALIDATORS*/
								float64validator.AtLeast(0.250000),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
								float64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Container
				"container": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ArtifactUrl
						"artifact_url": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(1024),
								stringvalidator.RegexMatches(regexp.MustCompile("^(https|s3)://([^/]+)/?(.*)$"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: DeployedImage
						"deployed_image": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ResolutionTime
								"resolution_time": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: ResolvedImage
								"resolved_image": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The image to use for the container that will be materialized for the inference component",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: SpecifiedImage
								"specified_image": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The image to use for the container that will be materialized for the inference component",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "",
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Environment
						"environment":       // Pattern: ""
						schema.MapAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Environment variables to specify on the container",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
								mapplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Image
						"image": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The image to use for the container that will be materialized for the inference component",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(255),
								stringvalidator.RegexMatches(regexp.MustCompile("[\\S]+"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
							// Image is a write-only property.
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ModelName
				"model_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the model to use with the inference component",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(63),
						stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: StartupParameters
				"startup_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ContainerStartupHealthCheckTimeoutInSeconds
						"container_startup_health_check_timeout_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.Between(60, 3600),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ModelDataDownloadTimeoutInSeconds
						"model_data_download_timeout_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.Between(60, 3600),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The specification for the inference component",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of tags to apply to the resource",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A tag in the form of a key-value pair to associate with the resource",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -",
		//	        "maxLength": 256,
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
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -",
						Optional:    true,
						Computed:    true,
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
			Description: "An array of tags to apply to the resource",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VariantName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the endpoint variant the inference component is associated with",
		//	  "maxLength": 63,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"variant_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the endpoint variant the inference component is associated with",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::SageMaker::InferenceComponent",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::InferenceComponent").WithTerraformTypeName("awscc_sagemaker_inference_component")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"artifact_url":                  "ArtifactUrl",
		"base_inference_component_name": "BaseInferenceComponentName",
		"compute_resource_requirements": "ComputeResourceRequirements",
		"container":                     "Container",
		"container_startup_health_check_timeout_in_seconds": "ContainerStartupHealthCheckTimeoutInSeconds",
		"copy_count":                             "CopyCount",
		"creation_time":                          "CreationTime",
		"current_copy_count":                     "CurrentCopyCount",
		"deployed_image":                         "DeployedImage",
		"desired_copy_count":                     "DesiredCopyCount",
		"endpoint_arn":                           "EndpointArn",
		"endpoint_name":                          "EndpointName",
		"environment":                            "Environment",
		"failure_reason":                         "FailureReason",
		"image":                                  "Image",
		"inference_component_arn":                "InferenceComponentArn",
		"inference_component_name":               "InferenceComponentName",
		"inference_component_status":             "InferenceComponentStatus",
		"key":                                    "Key",
		"last_modified_time":                     "LastModifiedTime",
		"max_memory_required_in_mb":              "MaxMemoryRequiredInMb",
		"min_memory_required_in_mb":              "MinMemoryRequiredInMb",
		"model_data_download_timeout_in_seconds": "ModelDataDownloadTimeoutInSeconds",
		"model_name":                             "ModelName",
		"number_of_accelerator_devices_required": "NumberOfAcceleratorDevicesRequired",
		"number_of_cpu_cores_required":           "NumberOfCpuCoresRequired",
		"resolution_time":                        "ResolutionTime",
		"resolved_image":                         "ResolvedImage",
		"runtime_config":                         "RuntimeConfig",
		"specification":                          "Specification",
		"specified_image":                        "SpecifiedImage",
		"startup_parameters":                     "StartupParameters",
		"tags":                                   "Tags",
		"value":                                  "Value",
		"variant_name":                           "VariantName",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Specification/Container/Image",
		"/properties/RuntimeConfig/CopyCount",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
