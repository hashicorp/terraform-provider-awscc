// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package backup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_backup_report_plan", reportPlanDataSource)
}

// reportPlanDataSource returns the Terraform awscc_backup_report_plan data source.
// This Terraform data source corresponds to the CloudFormation AWS::Backup::ReportPlan resource.
func reportPlanDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ReportDeliveryChannel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A structure that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.",
		//	  "properties": {
		//	    "Formats": {
		//	      "description": "A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "S3BucketName": {
		//	      "description": "The unique name of the S3 bucket that receives your reports.",
		//	      "type": "string"
		//	    },
		//	    "S3KeyPrefix": {
		//	      "description": "The prefix for where AWS Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3BucketName"
		//	  ],
		//	  "type": "object"
		//	}
		"report_delivery_channel": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Formats
				"formats": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: S3BucketName
				"s3_bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The unique name of the S3 bucket that receives your reports.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: S3KeyPrefix
				"s3_key_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The prefix for where AWS Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A structure that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReportPlanArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of the ARN depends on the resource type.",
		//	  "type": "string"
		//	}
		"report_plan_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of the ARN depends on the resource type.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReportPlanDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional description of the report plan with a maximum of 1,024 characters.",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "pattern": ".*\\S.*",
		//	  "type": "string"
		//	}
		"report_plan_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An optional description of the report plan with a maximum of 1,024 characters.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReportPlanName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z][_a-zA-Z0-9]*",
		//	  "type": "string"
		//	}
		"report_plan_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReportPlanTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Metadata that you can assign to help organize the report plans that you create. Each tag is a key-value pair.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"report_plan_tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Metadata that you can assign to help organize the report plans that you create. Each tag is a key-value pair.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReportSetting
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Identifies the report template for the report. Reports are built using a report template.",
		//	  "properties": {
		//	    "Accounts": {
		//	      "description": "The list of AWS accounts that a report covers.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "FrameworkArns": {
		//	      "description": "The Amazon Resource Names (ARNs) of the frameworks a report covers.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "OrganizationUnits": {
		//	      "description": "The list of AWS organization units that a report covers.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "Regions": {
		//	      "description": "The list of AWS regions that a report covers.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "ReportTemplate": {
		//	      "description": "Identifies the report template for the report. Reports are built using a report template. The report templates are: `BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "ReportTemplate"
		//	  ],
		//	  "type": "object"
		//	}
		"report_setting": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Accounts
				"accounts": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The list of AWS accounts that a report covers.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: FrameworkArns
				"framework_arns": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The Amazon Resource Names (ARNs) of the frameworks a report covers.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: OrganizationUnits
				"organization_units": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The list of AWS organization units that a report covers.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Regions
				"regions": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The list of AWS regions that a report covers.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ReportTemplate
				"report_template": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Identifies the report template for the report. Reports are built using a report template. The report templates are: `BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Identifies the report template for the report. Reports are built using a report template.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Backup::ReportPlan",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Backup::ReportPlan").WithTerraformTypeName("awscc_backup_report_plan")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"accounts":                "Accounts",
		"formats":                 "Formats",
		"framework_arns":          "FrameworkArns",
		"key":                     "Key",
		"organization_units":      "OrganizationUnits",
		"regions":                 "Regions",
		"report_delivery_channel": "ReportDeliveryChannel",
		"report_plan_arn":         "ReportPlanArn",
		"report_plan_description": "ReportPlanDescription",
		"report_plan_name":        "ReportPlanName",
		"report_plan_tags":        "ReportPlanTags",
		"report_setting":          "ReportSetting",
		"report_template":         "ReportTemplate",
		"s3_bucket_name":          "S3BucketName",
		"s3_key_prefix":           "S3KeyPrefix",
		"value":                   "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
