// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iot_scheduled_audit", scheduledAuditDataSource)
}

// scheduledAuditDataSource returns the Terraform awscc_iot_scheduled_audit data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoT::ScheduledAudit resource.
func scheduledAuditDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DayOfMonth
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The day of the month on which the scheduled audit takes place. Can be 1 through 31 or LAST. This field is required if the frequency parameter is set to MONTHLY.",
		//	  "pattern": "^([1-9]|[12][0-9]|3[01])$|^LAST$|^UNSET_VALUE$",
		//	  "type": "string"
		//	}
		"day_of_month": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The day of the month on which the scheduled audit takes place. Can be 1 through 31 or LAST. This field is required if the frequency parameter is set to MONTHLY.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DayOfWeek
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The day of the week on which the scheduled audit takes place. Can be one of SUN, MON, TUE,WED, THU, FRI, or SAT. This field is required if the frequency parameter is set to WEEKLY or BIWEEKLY.",
		//	  "enum": [
		//	    "SUN",
		//	    "MON",
		//	    "TUE",
		//	    "WED",
		//	    "THU",
		//	    "FRI",
		//	    "SAT",
		//	    "UNSET_VALUE"
		//	  ],
		//	  "type": "string"
		//	}
		"day_of_week": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The day of the week on which the scheduled audit takes place. Can be one of SUN, MON, TUE,WED, THU, FRI, or SAT. This field is required if the frequency parameter is set to WEEKLY or BIWEEKLY.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Frequency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "How often the scheduled audit takes place. Can be one of DAILY, WEEKLY, BIWEEKLY, or MONTHLY.",
		//	  "enum": [
		//	    "DAILY",
		//	    "WEEKLY",
		//	    "BIWEEKLY",
		//	    "MONTHLY"
		//	  ],
		//	  "type": "string"
		//	}
		"frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "How often the scheduled audit takes place. Can be one of DAILY, WEEKLY, BIWEEKLY, or MONTHLY.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ScheduledAuditArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN (Amazon resource name) of the scheduled audit.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"scheduled_audit_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN (Amazon resource name) of the scheduled audit.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ScheduledAuditName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name you want to give to the scheduled audit.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9:_-]+",
		//	  "type": "string"
		//	}
		"scheduled_audit_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name you want to give to the scheduled audit.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag's key.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag's value.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's key.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's value.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetCheckNames
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Which checks are performed during the scheduled audit. Checks must be enabled for your account.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"target_check_names": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Which checks are performed during the scheduled audit. Checks must be enabled for your account.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoT::ScheduledAudit",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::ScheduledAudit").WithTerraformTypeName("awscc_iot_scheduled_audit")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"day_of_month":         "DayOfMonth",
		"day_of_week":          "DayOfWeek",
		"frequency":            "Frequency",
		"key":                  "Key",
		"scheduled_audit_arn":  "ScheduledAuditArn",
		"scheduled_audit_name": "ScheduledAuditName",
		"tags":                 "Tags",
		"target_check_names":   "TargetCheckNames",
		"value":                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
