// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package workspacesweb

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_workspacesweb_portal", portalDataSource)
}

// portalDataSource returns the Terraform awscc_workspacesweb_portal data source.
// This Terraform data source corresponds to the CloudFormation AWS::WorkSpacesWeb::Portal resource.
func portalDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdditionalEncryptionContext
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 131072,
		//	      "minLength": 0,
		//	      "pattern": "^[\\s\\S]*$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"additional_encryption_context": // Pattern: ""
		schema.MapAttribute{             /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AuthenticationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "Standard",
		//	    "IAM_Identity_Center"
		//	  ],
		//	  "type": "string"
		//	}
		"authentication_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BrowserSettingsArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\\/[a-fA-F0-9\\-]{36})+$",
		//	  "type": "string"
		//	}
		"browser_settings_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BrowserType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "Chrome"
		//	  ],
		//	  "type": "string"
		//	}
		"browser_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CreationDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
		}, /*END ATTRIBUTE*/
		// Property: CustomerManagedKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:kms:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:key\\/[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"customer_managed_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DataProtectionSettingsArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\\/[a-fA-F0-9\\-]{36})+$",
		//	  "type": "string"
		//	}
		"data_protection_settings_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^.+$",
		//	  "type": "string"
		//	}
		"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "standard.regular",
		//	    "standard.large",
		//	    "standard.xlarge"
		//	  ],
		//	  "type": "string"
		//	}
		"instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IpAccessSettingsArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\\/[a-fA-F0-9\\-]{36})+$",
		//	  "type": "string"
		//	}
		"ip_access_settings_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MaxConcurrentSessions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 5000,
		//	  "minimum": 1,
		//	  "type": "number"
		//	}
		"max_concurrent_sessions": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkSettingsArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\\/[a-fA-F0-9\\-]{36})+$",
		//	  "type": "string"
		//	}
		"network_settings_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PortalArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\\/[a-fA-F0-9\\-]{36})+$",
		//	  "type": "string"
		//	}
		"portal_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PortalEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 253,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"portal_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PortalStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "Incomplete",
		//	    "Pending",
		//	    "Active"
		//	  ],
		//	  "type": "string"
		//	}
		"portal_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RendererType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "AppStream"
		//	  ],
		//	  "type": "string"
		//	}
		"renderer_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceProviderSamlMetadata
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 204800,
		//	  "minLength": 0,
		//	  "pattern": "^.*$",
		//	  "type": "string"
		//	}
		"service_provider_saml_metadata": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: StatusReason
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "pattern": ".*",
		//	  "type": "string"
		//	}
		"status_reason": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
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
		// Property: TrustStoreArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\\/[a-fA-F0-9\\-]{36})+$",
		//	  "type": "string"
		//	}
		"trust_store_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: UserAccessLoggingSettingsArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\\/[a-fA-F0-9\\-]{36})+$",
		//	  "type": "string"
		//	}
		"user_access_logging_settings_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: UserSettingsArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\\/[a-fA-F0-9\\-]{36})+$",
		//	  "type": "string"
		//	}
		"user_settings_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::WorkSpacesWeb::Portal",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::WorkSpacesWeb::Portal").WithTerraformTypeName("awscc_workspacesweb_portal")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"additional_encryption_context":    "AdditionalEncryptionContext",
		"authentication_type":              "AuthenticationType",
		"browser_settings_arn":             "BrowserSettingsArn",
		"browser_type":                     "BrowserType",
		"creation_date":                    "CreationDate",
		"customer_managed_key":             "CustomerManagedKey",
		"data_protection_settings_arn":     "DataProtectionSettingsArn",
		"display_name":                     "DisplayName",
		"instance_type":                    "InstanceType",
		"ip_access_settings_arn":           "IpAccessSettingsArn",
		"key":                              "Key",
		"max_concurrent_sessions":          "MaxConcurrentSessions",
		"network_settings_arn":             "NetworkSettingsArn",
		"portal_arn":                       "PortalArn",
		"portal_endpoint":                  "PortalEndpoint",
		"portal_status":                    "PortalStatus",
		"renderer_type":                    "RendererType",
		"service_provider_saml_metadata":   "ServiceProviderSamlMetadata",
		"status_reason":                    "StatusReason",
		"tags":                             "Tags",
		"trust_store_arn":                  "TrustStoreArn",
		"user_access_logging_settings_arn": "UserAccessLoggingSettingsArn",
		"user_settings_arn":                "UserSettingsArn",
		"value":                            "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
