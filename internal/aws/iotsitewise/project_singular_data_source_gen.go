// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotsitewise_project", projectDataSource)
}

// projectDataSource returns the Terraform awscc_iotsitewise_project data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTSiteWise::Project resource.
func projectDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IDs of the assets to be associated to the project.",
		//	  "items": {
		//	    "description": "The ID of the asset",
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"asset_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The IDs of the assets to be associated to the project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PortalId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the portal in which to create the project.",
		//	  "type": "string"
		//	}
		"portal_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the portal in which to create the project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the project.",
		//	  "type": "string"
		//	}
		"project_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the project.",
		//	  "type": "string"
		//	}
		"project_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the project.",
		//	  "type": "string"
		//	}
		"project_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A friendly name for the project.",
		//	  "type": "string"
		//	}
		"project_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A friendly name for the project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the project.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "To add or update tag, provide both key and value. To delete tag, provide only tag key to be deleted",
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of key-value pairs that contain metadata for the project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTSiteWise::Project",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Project").WithTerraformTypeName("awscc_iotsitewise_project")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"asset_ids":           "AssetIds",
		"key":                 "Key",
		"portal_id":           "PortalId",
		"project_arn":         "ProjectArn",
		"project_description": "ProjectDescription",
		"project_id":          "ProjectId",
		"project_name":        "ProjectName",
		"tags":                "Tags",
		"value":               "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}