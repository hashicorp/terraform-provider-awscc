// Code generated by generators/resource/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_imagebuilder_image_pipeline", imagePipelineResource)
}

// imagePipelineResource returns the Terraform awscc_imagebuilder_image_pipeline resource.
// This Terraform resource corresponds to the CloudFormation AWS::ImageBuilder::ImagePipeline resource.
func imagePipelineResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the image pipeline.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the image pipeline.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ContainerRecipeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.",
		//	  "type": "string"
		//	}
		"container_recipe_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the image pipeline.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the image pipeline.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DistributionConfigurationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.",
		//	  "type": "string"
		//	}
		"distribution_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnhancedImageMetadataEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Collects additional information about the image being created, including the operating system (OS) version and package list.",
		//	  "type": "boolean"
		//	}
		"enhanced_image_metadata_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Collects additional information about the image being created, including the operating system (OS) version and package list.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ImageRecipeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
		//	  "type": "string"
		//	}
		"image_recipe_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ImageTestsConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The image tests configuration of the image pipeline.",
		//	  "properties": {
		//	    "ImageTestsEnabled": {
		//	      "description": "Defines if tests should be executed when building this image.",
		//	      "type": "boolean"
		//	    },
		//	    "TimeoutMinutes": {
		//	      "description": "The maximum time in minutes that tests are permitted to run.",
		//	      "maximum": 1440,
		//	      "minimum": 60,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"image_tests_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ImageTestsEnabled
				"image_tests_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Defines if tests should be executed when building this image.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TimeoutMinutes
				"timeout_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum time in minutes that tests are permitted to run.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(60, 1440),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The image tests configuration of the image pipeline.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InfrastructureConfigurationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.",
		//	  "type": "string"
		//	}
		"infrastructure_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the image pipeline.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the image pipeline.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Schedule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The schedule of the image pipeline.",
		//	  "properties": {
		//	    "PipelineExecutionStartCondition": {
		//	      "description": "The condition configures when the pipeline should trigger a new image build.",
		//	      "enum": [
		//	        "EXPRESSION_MATCH_ONLY",
		//	        "EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "ScheduleExpression": {
		//	      "description": "The expression determines how often EC2 Image Builder evaluates your pipelineExecutionStartCondition.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"schedule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PipelineExecutionStartCondition
				"pipeline_execution_start_condition": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The condition configures when the pipeline should trigger a new image build.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"EXPRESSION_MATCH_ONLY",
							"EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ScheduleExpression
				"schedule_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The expression determines how often EC2 Image Builder evaluates your pipelineExecutionStartCondition.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The schedule of the image pipeline.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the image pipeline.",
		//	  "enum": [
		//	    "DISABLED",
		//	    "ENABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the image pipeline.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"DISABLED",
					"ENABLED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The tags of this image pipeline.",
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
			Description: "The tags of this image pipeline.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::ImageBuilder::ImagePipeline",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::ImagePipeline").WithTerraformTypeName("awscc_imagebuilder_image_pipeline")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                "Arn",
		"container_recipe_arn":               "ContainerRecipeArn",
		"description":                        "Description",
		"distribution_configuration_arn":     "DistributionConfigurationArn",
		"enhanced_image_metadata_enabled":    "EnhancedImageMetadataEnabled",
		"image_recipe_arn":                   "ImageRecipeArn",
		"image_tests_configuration":          "ImageTestsConfiguration",
		"image_tests_enabled":                "ImageTestsEnabled",
		"infrastructure_configuration_arn":   "InfrastructureConfigurationArn",
		"name":                               "Name",
		"pipeline_execution_start_condition": "PipelineExecutionStartCondition",
		"schedule":                           "Schedule",
		"schedule_expression":                "ScheduleExpression",
		"status":                             "Status",
		"tags":                               "Tags",
		"timeout_minutes":                    "TimeoutMinutes",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}