// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package finspace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_finspace_environment", environmentDataSource)
}

// environmentDataSource returns the Terraform awscc_finspace_environment data source.
// This Terraform data source corresponds to the CloudFormation AWS::FinSpace::Environment resource.
func environmentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AwsAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AWS account ID associated with the Environment",
		//	  "pattern": "^[a-zA-Z0-9]{1,26}$",
		//	  "type": "string"
		//	}
		"aws_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AWS account ID associated with the Environment",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataBundles
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARNs of FinSpace Data Bundles to install",
		//	  "items": {
		//	    "pattern": "^arn:aws:finspace:[A-Za-z0-9_/.-]{0,63}:\\d*:data-bundle/[0-9A-Za-z_-]{1,128}$",
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"data_bundles": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "ARNs of FinSpace Data Bundles to install",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DedicatedServiceAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ID for FinSpace created account used to store Environment artifacts",
		//	  "pattern": "^[a-zA-Z0-9]{1,26}$",
		//	  "type": "string"
		//	}
		"dedicated_service_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ID for FinSpace created account used to store Environment artifacts",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of the Environment",
		//	  "pattern": "^[a-zA-Z0-9. ]{1,1000}$",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of the Environment",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of the Environment",
		//	  "pattern": "^arn:aws:finspace:[A-Za-z0-9_/.-]{0,63}:\\d+:environment/[0-9A-Za-z_-]{1,128}$",
		//	  "type": "string"
		//	}
		"environment_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of the Environment",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique identifier for representing FinSpace Environment",
		//	  "pattern": "^[a-zA-Z0-9]{1,26}$",
		//	  "type": "string"
		//	}
		"environment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique identifier for representing FinSpace Environment",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "URL used to login to the Environment",
		//	  "pattern": "^[-a-zA-Z0-9+\u0026amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\u0026amp;@#/%=~_|]{1,1000}",
		//	  "type": "string"
		//	}
		"environment_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "URL used to login to the Environment",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FederationMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Federation mode used with the Environment",
		//	  "enum": [
		//	    "LOCAL",
		//	    "FEDERATED"
		//	  ],
		//	  "type": "string"
		//	}
		"federation_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Federation mode used with the Environment",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FederationParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Additional parameters to identify Federation mode",
		//	  "properties": {
		//	    "ApplicationCallBackURL": {
		//	      "description": "SAML metadata URL to link with the Environment",
		//	      "pattern": "^https?://[-a-zA-Z0-9+\u0026amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\u0026amp;@#/%=~_|]{1,1000}",
		//	      "type": "string"
		//	    },
		//	    "AttributeMap": {
		//	      "description": "Attribute map for SAML configuration",
		//	      "type": "object"
		//	    },
		//	    "FederationProviderName": {
		//	      "description": "Federation provider name to link with the Environment",
		//	      "maxLength": 32,
		//	      "minLength": 1,
		//	      "pattern": "[^_\\p{Z}][\\p{L}\\p{M}\\p{S}\\p{N}\\p{P}][^_\\p{Z}]+",
		//	      "type": "string"
		//	    },
		//	    "FederationURN": {
		//	      "description": "SAML metadata URL to link with the Environment",
		//	      "pattern": "",
		//	      "type": "string"
		//	    },
		//	    "SamlMetadataDocument": {
		//	      "description": "SAML metadata document to link the federation provider to the Environment",
		//	      "maxLength": 10000000,
		//	      "minLength": 1000,
		//	      "pattern": ".*",
		//	      "type": "string"
		//	    },
		//	    "SamlMetadataURL": {
		//	      "description": "SAML metadata URL to link with the Environment",
		//	      "pattern": "^https?://[-a-zA-Z0-9+\u0026amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\u0026amp;@#/%=~_|]{1,1000}",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"federation_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ApplicationCallBackURL
				"application_call_back_url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "SAML metadata URL to link with the Environment",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AttributeMap
				"attribute_map": schema.MapAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Attribute map for SAML configuration",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: FederationProviderName
				"federation_provider_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Federation provider name to link with the Environment",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: FederationURN
				"federation_urn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "SAML metadata URL to link with the Environment",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SamlMetadataDocument
				"saml_metadata_document": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "SAML metadata document to link the federation provider to the Environment",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SamlMetadataURL
				"saml_metadata_url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "SAML metadata URL to link with the Environment",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Additional parameters to identify Federation mode",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "KMS key used to encrypt customer data within FinSpace Environment infrastructure",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "KMS key used to encrypt customer data within FinSpace Environment infrastructure",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the Environment",
		//	  "pattern": "^[a-zA-Z0-9]+[a-zA-Z0-9-]*[a-zA-Z0-9]{1,255}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the Environment",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SageMakerStudioDomainUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "SageMaker Studio Domain URL associated with the Environment",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"sage_maker_studio_domain_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "SageMaker Studio Domain URL associated with the Environment",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "State of the Environment",
		//	  "enum": [
		//	    "CREATE_REQUESTED",
		//	    "CREATING",
		//	    "CREATED",
		//	    "DELETE_REQUESTED",
		//	    "DELETING",
		//	    "DELETED",
		//	    "FAILED_CREATION",
		//	    "FAILED_DELETION",
		//	    "RETRY_DELETION",
		//	    "SUSPENDED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "State of the Environment",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SuperuserParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Parameters of the first Superuser for the FinSpace Environment",
		//	  "properties": {
		//	    "EmailAddress": {
		//	      "description": "Email address",
		//	      "maxLength": 128,
		//	      "minLength": 1,
		//	      "pattern": "[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+[.]+[A-Za-z]+",
		//	      "type": "string"
		//	    },
		//	    "FirstName": {
		//	      "description": "First name",
		//	      "maxLength": 50,
		//	      "minLength": 1,
		//	      "pattern": "^[a-zA-Z0-9]{1,50}$",
		//	      "type": "string"
		//	    },
		//	    "LastName": {
		//	      "description": "Last name",
		//	      "maxLength": 50,
		//	      "minLength": 1,
		//	      "pattern": "^[a-zA-Z0-9]{1,50}$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"superuser_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EmailAddress
				"email_address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Email address",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: FirstName
				"first_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "First name",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: LastName
				"last_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Last name",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Parameters of the first Superuser for the FinSpace Environment",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::FinSpace::Environment",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FinSpace::Environment").WithTerraformTypeName("awscc_finspace_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_call_back_url":    "ApplicationCallBackURL",
		"attribute_map":                "AttributeMap",
		"aws_account_id":               "AwsAccountId",
		"data_bundles":                 "DataBundles",
		"dedicated_service_account_id": "DedicatedServiceAccountId",
		"description":                  "Description",
		"email_address":                "EmailAddress",
		"environment_arn":              "EnvironmentArn",
		"environment_id":               "EnvironmentId",
		"environment_url":              "EnvironmentUrl",
		"federation_mode":              "FederationMode",
		"federation_parameters":        "FederationParameters",
		"federation_provider_name":     "FederationProviderName",
		"federation_urn":               "FederationURN",
		"first_name":                   "FirstName",
		"kms_key_id":                   "KmsKeyId",
		"last_name":                    "LastName",
		"name":                         "Name",
		"sage_maker_studio_domain_url": "SageMakerStudioDomainUrl",
		"saml_metadata_document":       "SamlMetadataDocument",
		"saml_metadata_url":            "SamlMetadataURL",
		"status":                       "Status",
		"superuser_parameters":         "SuperuserParameters",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}