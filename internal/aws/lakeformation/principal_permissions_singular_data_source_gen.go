// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lakeformation_principal_permissions", principalPermissionsDataSource)
}

// principalPermissionsDataSource returns the Terraform awscc_lakeformation_principal_permissions data source.
// This Terraform data source corresponds to the CloudFormation AWS::LakeFormation::PrincipalPermissions resource.
func principalPermissionsDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Catalog
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "type": "string"
		//	}
		"catalog": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Permissions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "enum": [
		//	      "ALL",
		//	      "SELECT",
		//	      "ALTER",
		//	      "DROP",
		//	      "DELETE",
		//	      "INSERT",
		//	      "DESCRIBE",
		//	      "CREATE_DATABASE",
		//	      "CREATE_TABLE",
		//	      "DATA_LOCATION_ACCESS",
		//	      "CREATE_TAG",
		//	      "ASSOCIATE"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"permissions": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PermissionsWithGrantOption
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "enum": [
		//	      "ALL",
		//	      "SELECT",
		//	      "ALTER",
		//	      "DROP",
		//	      "DELETE",
		//	      "INSERT",
		//	      "DESCRIBE",
		//	      "CREATE_DATABASE",
		//	      "CREATE_TABLE",
		//	      "DATA_LOCATION_ACCESS",
		//	      "CREATE_TAG",
		//	      "ASSOCIATE"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"permissions_with_grant_option": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Principal
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DataLakePrincipalIdentifier": {
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"principal": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DataLakePrincipalIdentifier
				"data_lake_principal_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PrincipalIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"principal_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Resource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Catalog": {
		//	      "additionalProperties": false,
		//	      "type": "object"
		//	    },
		//	    "DataCellsFilter": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "DatabaseName": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Name": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "TableCatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "TableName": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "TableCatalogId",
		//	        "DatabaseName",
		//	        "TableName",
		//	        "Name"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "DataLocation": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "ResourceArn": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "ResourceArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Database": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "Name": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "Name"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "LFTag": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "TagKey": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "TagValues": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 256,
		//	            "minLength": 0,
		//	            "type": "string"
		//	          },
		//	          "maxItems": 50,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "TagKey",
		//	        "TagValues"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "LFTagPolicy": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "Expression": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "TagKey": {
		//	                "maxLength": 128,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              },
		//	              "TagValues": {
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "maxLength": 256,
		//	                  "minLength": 0,
		//	                  "type": "string"
		//	                },
		//	                "maxItems": 50,
		//	                "minItems": 1,
		//	                "type": "array"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "maxItems": 5,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        },
		//	        "ResourceType": {
		//	          "enum": [
		//	            "DATABASE",
		//	            "TABLE"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "ResourceType",
		//	        "Expression"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Table": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "DatabaseName": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Name": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "TableWildcard": {
		//	          "additionalProperties": false,
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "DatabaseName"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "TableWithColumns": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "ColumnNames": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 255,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "ColumnWildcard": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ExcludedColumnNames": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 255,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "DatabaseName": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Name": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "DatabaseName",
		//	        "Name"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"resource": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Catalog
				"catalog": schema.StringAttribute{ /*START ATTRIBUTE*/
					CustomType: jsontypes.NormalizedType{},
					Computed:   true,
				}, /*END ATTRIBUTE*/
				// Property: DataCellsFilter
				"data_cells_filter": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: TableCatalogId
						"table_catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: TableName
						"table_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DataLocation
				"data_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ResourceArn
						"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Database
				"database": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: LFTag
				"lf_tag": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: TagKey
						"tag_key": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: TagValues
						"tag_values": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: LFTagPolicy
				"lf_tag_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Expression
						"expression": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: TagKey
									"tag_key": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: TagValues
									"tag_values": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ResourceType
						"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Table
				"table": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: TableWildcard
						"table_wildcard": schema.StringAttribute{ /*START ATTRIBUTE*/
							CustomType: jsontypes.NormalizedType{},
							Computed:   true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TableWithColumns
				"table_with_columns": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ColumnNames
						"column_names": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ColumnWildcard
						"column_wildcard": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ExcludedColumnNames
								"excluded_column_names": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"resource_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::LakeFormation::PrincipalPermissions",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LakeFormation::PrincipalPermissions").WithTerraformTypeName("awscc_lakeformation_principal_permissions")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"catalog":                        "Catalog",
		"catalog_id":                     "CatalogId",
		"column_names":                   "ColumnNames",
		"column_wildcard":                "ColumnWildcard",
		"data_cells_filter":              "DataCellsFilter",
		"data_lake_principal_identifier": "DataLakePrincipalIdentifier",
		"data_location":                  "DataLocation",
		"database":                       "Database",
		"database_name":                  "DatabaseName",
		"excluded_column_names":          "ExcludedColumnNames",
		"expression":                     "Expression",
		"lf_tag":                         "LFTag",
		"lf_tag_policy":                  "LFTagPolicy",
		"name":                           "Name",
		"permissions":                    "Permissions",
		"permissions_with_grant_option":  "PermissionsWithGrantOption",
		"principal":                      "Principal",
		"principal_identifier":           "PrincipalIdentifier",
		"resource":                       "Resource",
		"resource_arn":                   "ResourceArn",
		"resource_identifier":            "ResourceIdentifier",
		"resource_type":                  "ResourceType",
		"table":                          "Table",
		"table_catalog_id":               "TableCatalogId",
		"table_name":                     "TableName",
		"table_wildcard":                 "TableWildcard",
		"table_with_columns":             "TableWithColumns",
		"tag_key":                        "TagKey",
		"tag_values":                     "TagValues",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
