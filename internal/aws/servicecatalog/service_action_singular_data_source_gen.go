// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_servicecatalog_service_action", serviceActionDataSource)
}

// serviceActionDataSource returns the Terraform awscc_servicecatalog_service_action data source.
// This Terraform data source corresponds to the CloudFormation AWS::ServiceCatalog::ServiceAction resource.
func serviceActionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AcceptLanguage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "en",
		//	    "jp",
		//	    "zh"
		//	  ],
		//	  "type": "string"
		//	}
		"accept_language": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Definition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 1000,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 4096,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"definition": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DefinitionType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "SSM_AUTOMATION"
		//	  ],
		//	  "type": "string"
		//	}
		"definition_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ServiceCatalog::ServiceAction",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalog::ServiceAction").WithTerraformTypeName("awscc_servicecatalog_service_action")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"accept_language": "AcceptLanguage",
		"definition":      "Definition",
		"definition_type": "DefinitionType",
		"description":     "Description",
		"id":              "Id",
		"key":             "Key",
		"name":            "Name",
		"value":           "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}