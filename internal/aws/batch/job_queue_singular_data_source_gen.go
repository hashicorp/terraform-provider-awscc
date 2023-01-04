// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package batch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_batch_job_queue", jobQueueDataSource)
}

// jobQueueDataSource returns the Terraform awscc_batch_job_queue data source.
// This Terraform data source corresponds to the CloudFormation AWS::Batch::JobQueue resource.
func jobQueueDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ComputeEnvironmentOrder
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "ComputeEnvironment": {
		//	        "type": "string"
		//	      },
		//	      "Order": {
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "required": [
		//	      "ComputeEnvironment",
		//	      "Order"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"compute_environment_order": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ComputeEnvironment
					"compute_environment": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Order
					"order": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: JobQueueArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"job_queue_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: JobQueueName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"job_queue_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Priority
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 1000,
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SchedulingPolicyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"scheduling_policy_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "DISABLED",
		//	    "ENABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A key-value pair to associate with a resource.",
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
			Description: "A key-value pair to associate with a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Batch::JobQueue",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Batch::JobQueue").WithTerraformTypeName("awscc_batch_job_queue")
	opts = opts.WithTerraformSchema(schema)
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

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}