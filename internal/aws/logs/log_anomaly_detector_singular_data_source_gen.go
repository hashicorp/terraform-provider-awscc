// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_logs_log_anomaly_detector", logAnomalyDetectorDataSource)
}

// logAnomalyDetectorDataSource returns the Terraform awscc_logs_log_anomaly_detector data source.
// This Terraform data source corresponds to the CloudFormation AWS::Logs::LogAnomalyDetector resource.
func logAnomalyDetectorDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Account ID for owner of detector",
		//	  "type": "string"
		//	}
		"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Account ID for owner of detector",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AnomalyDetectorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of LogAnomalyDetector",
		//	  "type": "string"
		//	}
		"anomaly_detector_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of LogAnomalyDetector",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AnomalyDetectorStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Current status of detector.",
		//	  "type": "string"
		//	}
		"anomaly_detector_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Current status of detector.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AnomalyVisibilityTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "number"
		//	}
		"anomaly_visibility_time": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreationTimeStamp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When detector was created.",
		//	  "type": "number"
		//	}
		"creation_time_stamp": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "When detector was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DetectorName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of detector",
		//	  "type": "string"
		//	}
		"detector_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of detector",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EvaluationFrequency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "How often log group is evaluated",
		//	  "enum": [
		//	    "FIVE_MIN",
		//	    "TEN_MIN",
		//	    "FIFTEEN_MIN",
		//	    "THIRTY_MIN",
		//	    "ONE_HOUR"
		//	  ],
		//	  "type": "string"
		//	}
		"evaluation_frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "How often log group is evaluated",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FilterPattern
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"filter_pattern": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTimeStamp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When detector was lsat modified.",
		//	  "type": "number"
		//	}
		"last_modified_time_stamp": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "When detector was lsat modified.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LogGroupArnList
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of Arns for the given log group",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 2048,
		//	    "minLength": 20,
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"log_group_arn_list": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "List of Arns for the given log group",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Logs::LogAnomalyDetector",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::LogAnomalyDetector").WithTerraformTypeName("awscc_logs_log_anomaly_detector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":               "AccountId",
		"anomaly_detector_arn":     "AnomalyDetectorArn",
		"anomaly_detector_status":  "AnomalyDetectorStatus",
		"anomaly_visibility_time":  "AnomalyVisibilityTime",
		"creation_time_stamp":      "CreationTimeStamp",
		"detector_name":            "DetectorName",
		"evaluation_frequency":     "EvaluationFrequency",
		"filter_pattern":           "FilterPattern",
		"kms_key_id":               "KmsKeyId",
		"last_modified_time_stamp": "LastModifiedTimeStamp",
		"log_group_arn_list":       "LogGroupArnList",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
