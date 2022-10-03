// Code generated by generators/resource/main.go; DO NOT EDIT.

package batch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_batch_job_queue", jobQueueResource)
}

// jobQueueResource returns the Terraform awscc_batch_job_queue resource.
// This Terraform resource corresponds to the CloudFormation AWS::Batch::JobQueue resource.
func jobQueueResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"compute_environment_order": {
			// Property: ComputeEnvironmentOrder
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ComputeEnvironment": {
			//         "type": "string"
			//       },
			//       "Order": {
			//         "type": "integer"
			//       }
			//     },
			//     "required": [
			//       "ComputeEnvironment",
			//       "Order"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"compute_environment": {
						// Property: ComputeEnvironment
						Type:     types.StringType,
						Required: true,
					},
					"order": {
						// Property: Order
						Type:     types.Int64Type,
						Required: true,
					},
				},
			),
			Required: true,
		},
		"job_queue_arn": {
			// Property: JobQueueArn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"job_queue_name": {
			// Property: JobQueueName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"priority": {
			// Property: Priority
			// CloudFormation resource type schema:
			// {
			//   "maximum": 1000,
			//   "minimum": 0,
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntBetween(0, 1000),
			},
		},
		"scheduling_policy_arn": {
			// Property: SchedulingPolicyArn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DISABLED",
			//     "ENABLED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"DISABLED",
					"ENABLED",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A key-value pair to associate with a resource.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A key-value pair to associate with a resource.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Batch::JobQueue",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Batch::JobQueue").WithTerraformTypeName("awscc_batch_job_queue")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"compute_environment":       "ComputeEnvironment",
		"compute_environment_order": "ComputeEnvironmentOrder",
		"job_queue_arn":             "JobQueueArn",
		"job_queue_name":            "JobQueueName",
		"order":                     "Order",
		"priority":                  "Priority",
		"scheduling_policy_arn":     "SchedulingPolicyArn",
		"state":                     "State",
		"tags":                      "Tags",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
