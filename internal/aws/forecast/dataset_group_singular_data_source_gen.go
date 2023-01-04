// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package forecast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_forecast_dataset_group", datasetGroupDataSource)
}

// datasetGroupDataSource returns the Terraform awscc_forecast_dataset_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::Forecast::DatasetGroup resource.
func datasetGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DatasetArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "maxLength": 256,
		//	    "pattern": "^[a-zA-Z0-9\\-\\_\\.\\/\\:]+$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"dataset_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatasetGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the dataset group to delete.",
		//	  "maxLength": 256,
		//	  "pattern": "^[a-zA-Z0-9\\-\\_\\.\\/\\:]+$",
		//	  "type": "string"
		//	}
		"dataset_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the dataset group to delete.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatasetGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the dataset group.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z][a-zA-Z0-9_]*",
		//	  "type": "string"
		//	}
		"dataset_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the dataset group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Domain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The domain associated with the dataset group. When you add a dataset to a dataset group, this value and the value specified for the Domain parameter of the CreateDataset operation must match.",
		//	  "enum": [
		//	    "RETAIL",
		//	    "CUSTOM",
		//	    "INVENTORY_PLANNING",
		//	    "EC2_CAPACITY",
		//	    "WORK_FORCE",
		//	    "WEB_TRAFFIC",
		//	    "METRICS"
		//	  ],
		//	  "type": "string"
		//	}
		"domain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The domain associated with the dataset group. When you add a dataset to a dataset group, this value and the value specified for the Domain parameter of the CreateDataset operation must match.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags of Application Insights application.",
		//	  "insertionOrder": true,
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
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
			Description: "The tags of Application Insights application.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Forecast::DatasetGroup",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Forecast::DatasetGroup").WithTerraformTypeName("awscc_forecast_dataset_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dataset_arns":       "DatasetArns",
		"dataset_group_arn":  "DatasetGroupArn",
		"dataset_group_name": "DatasetGroupName",
		"domain":             "Domain",
		"key":                "Key",
		"tags":               "Tags",
		"value":              "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}