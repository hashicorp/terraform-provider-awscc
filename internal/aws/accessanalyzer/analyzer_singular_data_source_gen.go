// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package accessanalyzer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_accessanalyzer_analyzer", analyzerDataSource)
}

// analyzerDataSource returns the Terraform awscc_accessanalyzer_analyzer data source.
// This Terraform data source corresponds to the CloudFormation AWS::AccessAnalyzer::Analyzer resource.
func analyzerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AnalyzerName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Analyzer name",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"analyzer_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Analyzer name",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ArchiveRules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "An Access Analyzer archive rule. Archive rules automatically archive new findings that meet the criteria you define when you create the rule.",
		//	    "properties": {
		//	      "Filter": {
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "properties": {
		//	            "Contains": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "Eq": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "Exists": {
		//	              "type": "boolean"
		//	            },
		//	            "Neq": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "Property": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Property"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "minItems": 1,
		//	        "type": "array"
		//	      },
		//	      "RuleName": {
		//	        "description": "The archive rule name",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Filter",
		//	      "RuleName"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"archive_rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Filter
					"filter": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Contains
								"contains": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Eq
								"eq": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Exists
								"exists": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Neq
								"neq": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Property
								"property": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: RuleName
					"rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The archive rule name",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the analyzer",
		//	  "maxLength": 1600,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the analyzer",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
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
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of the analyzer, must be ACCOUNT or ORGANIZATION",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of the analyzer, must be ACCOUNT or ORGANIZATION",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::AccessAnalyzer::Analyzer",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AccessAnalyzer::Analyzer").WithTerraformTypeName("awscc_accessanalyzer_analyzer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"analyzer_name": "AnalyzerName",
		"archive_rules": "ArchiveRules",
		"arn":           "Arn",
		"contains":      "Contains",
		"eq":            "Eq",
		"exists":        "Exists",
		"filter":        "Filter",
		"key":           "Key",
		"neq":           "Neq",
		"property":      "Property",
		"rule_name":     "RuleName",
		"tags":          "Tags",
		"type":          "Type",
		"value":         "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}