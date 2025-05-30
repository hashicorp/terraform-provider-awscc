// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package deadline

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_deadline_queue_limit_association", queueLimitAssociationDataSource)
}

// queueLimitAssociationDataSource returns the Terraform awscc_deadline_queue_limit_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::Deadline::QueueLimitAssociation resource.
func queueLimitAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: FarmId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^farm-[0-9a-f]{32}$",
		//	  "type": "string"
		//	}
		"farm_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LimitId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^limit-[0-9a-f]{32}$",
		//	  "type": "string"
		//	}
		"limit_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: QueueId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^queue-[0-9a-f]{32}$",
		//	  "type": "string"
		//	}
		"queue_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Deadline::QueueLimitAssociation",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Deadline::QueueLimitAssociation").WithTerraformTypeName("awscc_deadline_queue_limit_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"farm_id":  "FarmId",
		"limit_id": "LimitId",
		"queue_id": "QueueId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
