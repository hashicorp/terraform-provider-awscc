// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_sagemaker_model_bias_job_definition", modelBiasJobDefinitionResourceType)
}

// modelBiasJobDefinitionResourceType returns the Terraform awscc_sagemaker_model_bias_job_definition resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::ModelBiasJobDefinition resource type.
func modelBiasJobDefinitionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time at which the job definition was created.",
			//   "type": "string"
			// }
			Description: "The time at which the job definition was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"job_definition_arn": {
			// Property: JobDefinitionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of job definition.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of job definition.",
			Type:        types.StringType,
			Computed:    true,
		},
		"job_definition_name": {
			// Property: JobDefinitionName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the job definition.",
			//   "maxLength": 63,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the job definition.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(63),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				ComputedOptionalForceNew(),
			},
		},
		"job_resources": {
			// Property: JobResources
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Identifies the resources to deploy for a monitoring job.",
			//   "properties": {
			//     "ClusterConfig": {
			//       "additionalProperties": false,
			//       "description": "Configuration for the cluster used to run model monitoring jobs.",
			//       "properties": {
			//         "InstanceCount": {
			//           "description": "The number of ML compute instances to use in the model monitoring job. For distributed processing jobs, specify a value greater than 1. The default value is 1.",
			//           "maximum": 100,
			//           "minimum": 1,
			//           "type": "integer"
			//         },
			//         "InstanceType": {
			//           "description": "The ML compute instance type for the processing job.",
			//           "type": "string"
			//         },
			//         "VolumeKmsKeyId": {
			//           "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.",
			//           "maximum": 2048,
			//           "minimum": 1,
			//           "type": "string"
			//         },
			//         "VolumeSizeInGB": {
			//           "description": "The size of the ML storage volume, in gigabytes, that you want to provision. You must specify sufficient ML storage for your scenario.",
			//           "maximum": 16384,
			//           "minimum": 1,
			//           "type": "integer"
			//         }
			//       },
			//       "required": [
			//         "InstanceCount",
			//         "InstanceType",
			//         "VolumeSizeInGB"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "ClusterConfig"
			//   ],
			//   "type": "object"
			// }
			Description: "Identifies the resources to deploy for a monitoring job.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"cluster_config": {
						// Property: ClusterConfig
						Description: "Configuration for the cluster used to run model monitoring jobs.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"instance_count": {
									// Property: InstanceCount
									Description: "The number of ML compute instances to use in the model monitoring job. For distributed processing jobs, specify a value greater than 1. The default value is 1.",
									Type:        types.NumberType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(1, 100),
									},
								},
								"instance_type": {
									// Property: InstanceType
									Description: "The ML compute instance type for the processing job.",
									Type:        types.StringType,
									Required:    true,
								},
								"volume_kms_key_id": {
									// Property: VolumeKmsKeyId
									Description: "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.",
									Type:        types.StringType,
									Optional:    true,
								},
								"volume_size_in_gb": {
									// Property: VolumeSizeInGB
									Description: "The size of the ML storage volume, in gigabytes, that you want to provision. You must specify sufficient ML storage for your scenario.",
									Type:        types.NumberType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(1, 16384),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"model_bias_app_specification": {
			// Property: ModelBiasAppSpecification
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Container image configuration object for the monitoring job.",
			//   "properties": {
			//     "ConfigUri": {
			//       "description": "The S3 URI to an analysis configuration file",
			//       "maxLength": 255,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "Environment": {
			//       "additionalProperties": false,
			//       "description": "Sets the environment variables in the Docker container",
			//       "patternProperties": {
			//         "": {
			//           "maxLength": 256,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "[\\S\\s]*": {
			//           "maxLength": 256,
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ImageUri": {
			//       "description": "The container image to be run by the monitoring job.",
			//       "maxLength": 255,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "ImageUri",
			//     "ConfigUri"
			//   ],
			//   "type": "object"
			// }
			Description: "Container image configuration object for the monitoring job.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"config_uri": {
						// Property: ConfigUri
						Description: "The S3 URI to an analysis configuration file",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(255),
						},
					},
					"environment": {
						// Property: Environment
						Description: "Sets the environment variables in the Docker container",
						// Pattern: ""
						Type: types.MapType{ElemType: types.StringType},
						// Pattern "[\\S\\s]*" ignored.
						Optional: true,
					},
					"image_uri": {
						// Property: ImageUri
						Description: "The container image to be run by the monitoring job.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(255),
						},
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"model_bias_baseline_config": {
			// Property: ModelBiasBaselineConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Baseline configuration used to validate that the data conforms to the specified constraints and statistics.",
			//   "properties": {
			//     "BaseliningJobName": {
			//       "description": "The name of a processing job",
			//       "maxLength": 63,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "ConstraintsResource": {
			//       "additionalProperties": false,
			//       "description": "The baseline constraints resource for a monitoring job.",
			//       "properties": {
			//         "S3Uri": {
			//           "description": "The Amazon S3 URI for baseline constraint file in Amazon S3 that the current monitoring job should validated against.",
			//           "maxLength": 1024,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Baseline configuration used to validate that the data conforms to the specified constraints and statistics.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"baselining_job_name": {
						// Property: BaseliningJobName
						Description: "The name of a processing job",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 63),
						},
					},
					"constraints_resource": {
						// Property: ConstraintsResource
						Description: "The baseline constraints resource for a monitoring job.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"s3_uri": {
									// Property: S3Uri
									Description: "The Amazon S3 URI for baseline constraint file in Amazon S3 that the current monitoring job should validated against.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(1024),
									},
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				ComputedOptionalForceNew(),
			},
		},
		"model_bias_job_input": {
			// Property: ModelBiasJobInput
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The inputs for a monitoring job.",
			//   "properties": {
			//     "EndpointInput": {
			//       "additionalProperties": false,
			//       "description": "The endpoint for a monitoring job.",
			//       "properties": {
			//         "EndTimeOffset": {
			//           "description": "Monitoring end time offset, e.g. PT0H",
			//           "maxLength": 15,
			//           "minLength": 1,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "EndpointName": {
			//           "description": "The name of the endpoint used to run the monitoring job.",
			//           "maxLength": 63,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "FeaturesAttribute": {
			//           "description": "JSONpath to locate features in JSONlines dataset",
			//           "maxLength": 256,
			//           "type": "string"
			//         },
			//         "InferenceAttribute": {
			//           "description": "Index or JSONpath to locate predicted label(s)",
			//           "maxLength": 256,
			//           "type": "string"
			//         },
			//         "LocalPath": {
			//           "description": "Path to the filesystem where the endpoint data is available to the container.",
			//           "maxLength": 256,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "ProbabilityAttribute": {
			//           "description": "Index or JSONpath to locate probabilities",
			//           "maxLength": 256,
			//           "type": "string"
			//         },
			//         "ProbabilityThresholdAttribute": {
			//           "format": "double",
			//           "type": "number"
			//         },
			//         "S3DataDistributionType": {
			//           "description": "Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated",
			//           "enum": [
			//             "FullyReplicated",
			//             "ShardedByS3Key"
			//           ],
			//           "type": "string"
			//         },
			//         "S3InputMode": {
			//           "description": "Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.",
			//           "enum": [
			//             "Pipe",
			//             "File"
			//           ],
			//           "type": "string"
			//         },
			//         "StartTimeOffset": {
			//           "description": "Monitoring start time offset, e.g. -PT1H",
			//           "maxLength": 15,
			//           "minLength": 1,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "EndpointName",
			//         "LocalPath"
			//       ],
			//       "type": "object"
			//     },
			//     "GroundTruthS3Input": {
			//       "additionalProperties": false,
			//       "description": "Ground truth input provided in S3 ",
			//       "properties": {
			//         "S3Uri": {
			//           "description": "A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.",
			//           "maxLength": 512,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "S3Uri"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "EndpointInput",
			//     "GroundTruthS3Input"
			//   ],
			//   "type": "object"
			// }
			Description: "The inputs for a monitoring job.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"endpoint_input": {
						// Property: EndpointInput
						Description: "The endpoint for a monitoring job.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"end_time_offset": {
									// Property: EndTimeOffset
									Description: "Monitoring end time offset, e.g. PT0H",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 15),
									},
								},
								"endpoint_name": {
									// Property: EndpointName
									Description: "The name of the endpoint used to run the monitoring job.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(63),
									},
								},
								"features_attribute": {
									// Property: FeaturesAttribute
									Description: "JSONpath to locate features in JSONlines dataset",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(256),
									},
								},
								"inference_attribute": {
									// Property: InferenceAttribute
									Description: "Index or JSONpath to locate predicted label(s)",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(256),
									},
								},
								"local_path": {
									// Property: LocalPath
									Description: "Path to the filesystem where the endpoint data is available to the container.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(256),
									},
								},
								"probability_attribute": {
									// Property: ProbabilityAttribute
									Description: "Index or JSONpath to locate probabilities",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(256),
									},
								},
								"probability_threshold_attribute": {
									// Property: ProbabilityThresholdAttribute
									Type:     types.NumberType,
									Optional: true,
								},
								"s3_data_distribution_type": {
									// Property: S3DataDistributionType
									Description: "Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"FullyReplicated",
											"ShardedByS3Key",
										}),
									},
								},
								"s3_input_mode": {
									// Property: S3InputMode
									Description: "Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"Pipe",
											"File",
										}),
									},
								},
								"start_time_offset": {
									// Property: StartTimeOffset
									Description: "Monitoring start time offset, e.g. -PT1H",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 15),
									},
								},
							},
						),
						Required: true,
					},
					"ground_truth_s3_input": {
						// Property: GroundTruthS3Input
						Description: "Ground truth input provided in S3 ",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"s3_uri": {
									// Property: S3Uri
									Description: "A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(512),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"model_bias_job_output_config": {
			// Property: ModelBiasJobOutputConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The output configuration for monitoring jobs.",
			//   "properties": {
			//     "KmsKeyId": {
			//       "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.",
			//       "maxLength": 2048,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "MonitoringOutputs": {
			//       "description": "Monitoring outputs for monitoring jobs. This is where the output of the periodic monitoring jobs is uploaded.",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "The output object for a monitoring job.",
			//         "properties": {
			//           "S3Output": {
			//             "additionalProperties": false,
			//             "description": "Information about where and how to store the results of a monitoring job.",
			//             "properties": {
			//               "LocalPath": {
			//                 "description": "The local path to the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job. LocalPath is an absolute path for the output data.",
			//                 "maxLength": 256,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "S3UploadMode": {
			//                 "description": "Whether to upload the results of the monitoring job continuously or after the job completes.",
			//                 "enum": [
			//                   "Continuous",
			//                   "EndOfJob"
			//                 ],
			//                 "type": "string"
			//               },
			//               "S3Uri": {
			//                 "description": "A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.",
			//                 "maxLength": 512,
			//                 "pattern": "",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "LocalPath",
			//               "S3Uri"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "S3Output"
			//         ],
			//         "type": "object"
			//       },
			//       "maxLength": 1,
			//       "minLength": 1,
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "MonitoringOutputs"
			//   ],
			//   "type": "object"
			// }
			Description: "The output configuration for monitoring jobs.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"kms_key_id": {
						// Property: KmsKeyId
						Description: "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(2048),
						},
					},
					"monitoring_outputs": {
						// Property: MonitoringOutputs
						Description: "Monitoring outputs for monitoring jobs. This is where the output of the periodic monitoring jobs is uploaded.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"s3_output": {
									// Property: S3Output
									Description: "Information about where and how to store the results of a monitoring job.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"local_path": {
												// Property: LocalPath
												Description: "The local path to the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job. LocalPath is an absolute path for the output data.",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenAtMost(256),
												},
											},
											"s3_upload_mode": {
												// Property: S3UploadMode
												Description: "Whether to upload the results of the monitoring job continuously or after the job completes.",
												Type:        types.StringType,
												Optional:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"Continuous",
														"EndOfJob",
													}),
												},
											},
											"s3_uri": {
												// Property: S3Uri
												Description: "A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenAtMost(512),
												},
											},
										},
									),
									Required: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Required: true,
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"network_config": {
			// Property: NetworkConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.",
			//   "properties": {
			//     "EnableInterContainerTrafficEncryption": {
			//       "description": "Whether to encrypt all communications between distributed processing jobs. Choose True to encrypt communications. Encryption provides greater security for distributed processing jobs, but the processing might take longer.",
			//       "type": "boolean"
			//     },
			//     "EnableNetworkIsolation": {
			//       "description": "Whether to allow inbound and outbound network calls to and from the containers used for the processing job.",
			//       "type": "boolean"
			//     },
			//     "VpcConfig": {
			//       "additionalProperties": false,
			//       "description": "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC.",
			//       "properties": {
			//         "SecurityGroupIds": {
			//           "description": "The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the Subnets field.",
			//           "items": {
			//             "maxLength": 32,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "maxItems": 5,
			//           "minItems": 1,
			//           "type": "array"
			//         },
			//         "Subnets": {
			//           "description": "The ID of the subnets in the VPC to which you want to connect to your monitoring jobs.",
			//           "items": {
			//             "maxLength": 32,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "maxItems": 16,
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "required": [
			//         "SecurityGroupIds",
			//         "Subnets"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"enable_inter_container_traffic_encryption": {
						// Property: EnableInterContainerTrafficEncryption
						Description: "Whether to encrypt all communications between distributed processing jobs. Choose True to encrypt communications. Encryption provides greater security for distributed processing jobs, but the processing might take longer.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"enable_network_isolation": {
						// Property: EnableNetworkIsolation
						Description: "Whether to allow inbound and outbound network calls to and from the containers used for the processing job.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"vpc_config": {
						// Property: VpcConfig
						Description: "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"security_group_ids": {
									// Property: SecurityGroupIds
									Description: "The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the Subnets field.",
									Type:        types.ListType{ElemType: types.StringType},
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenBetween(1, 5),
										validate.ArrayForEach(validate.StringLenAtMost(32)),
									},
								},
								"subnets": {
									// Property: Subnets
									Description: "The ID of the subnets in the VPC to which you want to connect to your monitoring jobs.",
									Type:        types.ListType{ElemType: types.StringType},
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenBetween(1, 16),
										validate.ArrayForEach(validate.StringLenAtMost(32)),
									},
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				ComputedOptionalForceNew(),
			},
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"stopping_condition": {
			// Property: StoppingCondition
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specifies a time limit for how long the monitoring job is allowed to run.",
			//   "properties": {
			//     "MaxRuntimeInSeconds": {
			//       "description": "The maximum runtime allowed in seconds.",
			//       "maximum": 86400,
			//       "minimum": 1,
			//       "type": "integer"
			//     }
			//   },
			//   "required": [
			//     "MaxRuntimeInSeconds"
			//   ],
			//   "type": "object"
			// }
			Description: "Specifies a time limit for how long the monitoring job is allowed to run.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"max_runtime_in_seconds": {
						// Property: MaxRuntimeInSeconds
						Description: "The maximum runtime allowed in seconds.",
						Type:        types.NumberType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(1, 86400),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				ComputedOptionalForceNew(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array"
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				ComputedOptionalForceNew(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::SageMaker::ModelBiasJobDefinition",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::ModelBiasJobDefinition").WithTerraformTypeName("awscc_sagemaker_model_bias_job_definition")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"baselining_job_name":  "BaseliningJobName",
		"cluster_config":       "ClusterConfig",
		"config_uri":           "ConfigUri",
		"constraints_resource": "ConstraintsResource",
		"creation_time":        "CreationTime",
		"enable_inter_container_traffic_encryption": "EnableInterContainerTrafficEncryption",
		"enable_network_isolation":                  "EnableNetworkIsolation",
		"end_time_offset":                           "EndTimeOffset",
		"endpoint_input":                            "EndpointInput",
		"endpoint_name":                             "EndpointName",
		"environment":                               "Environment",
		"features_attribute":                        "FeaturesAttribute",
		"ground_truth_s3_input":                     "GroundTruthS3Input",
		"image_uri":                                 "ImageUri",
		"inference_attribute":                       "InferenceAttribute",
		"instance_count":                            "InstanceCount",
		"instance_type":                             "InstanceType",
		"job_definition_arn":                        "JobDefinitionArn",
		"job_definition_name":                       "JobDefinitionName",
		"job_resources":                             "JobResources",
		"key":                                       "Key",
		"kms_key_id":                                "KmsKeyId",
		"local_path":                                "LocalPath",
		"max_runtime_in_seconds":                    "MaxRuntimeInSeconds",
		"model_bias_app_specification":              "ModelBiasAppSpecification",
		"model_bias_baseline_config":                "ModelBiasBaselineConfig",
		"model_bias_job_input":                      "ModelBiasJobInput",
		"model_bias_job_output_config":              "ModelBiasJobOutputConfig",
		"monitoring_outputs":                        "MonitoringOutputs",
		"network_config":                            "NetworkConfig",
		"probability_attribute":                     "ProbabilityAttribute",
		"probability_threshold_attribute":           "ProbabilityThresholdAttribute",
		"role_arn":                                  "RoleArn",
		"s3_data_distribution_type":                 "S3DataDistributionType",
		"s3_input_mode":                             "S3InputMode",
		"s3_output":                                 "S3Output",
		"s3_upload_mode":                            "S3UploadMode",
		"s3_uri":                                    "S3Uri",
		"security_group_ids":                        "SecurityGroupIds",
		"start_time_offset":                         "StartTimeOffset",
		"stopping_condition":                        "StoppingCondition",
		"subnets":                                   "Subnets",
		"tags":                                      "Tags",
		"value":                                     "Value",
		"volume_kms_key_id":                         "VolumeKmsKeyId",
		"volume_size_in_gb":                         "VolumeSizeInGB",
		"vpc_config":                                "VpcConfig",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
