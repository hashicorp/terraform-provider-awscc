// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ecs_primary_task_set", primaryTaskSetDataSource)
}

// primaryTaskSetDataSource returns the Terraform awscc_ecs_primary_task_set data source.
// This Terraform data source corresponds to the CloudFormation AWS::ECS::PrimaryTaskSet resource.
func primaryTaskSetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Cluster
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.",
		//	  "type": "string"
		//	}
		"cluster": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Service
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.",
		//	  "type": "string"
		//	}
		"service": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TaskSetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID or full Amazon Resource Name (ARN) of the task set.",
		//	  "type": "string"
		//	}
		"task_set_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID or full Amazon Resource Name (ARN) of the task set.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ECS::PrimaryTaskSet",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::PrimaryTaskSet").WithTerraformTypeName("awscc_ecs_primary_task_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cluster":     "Cluster",
		"service":     "Service",
		"task_set_id": "TaskSetId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}