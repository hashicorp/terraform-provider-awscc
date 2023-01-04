// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iot_account_audit_configuration", accountAuditConfigurationDataSource)
}

// accountAuditConfigurationDataSource returns the Terraform awscc_iot_account_audit_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoT::AccountAuditConfiguration resource.
func accountAuditConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "type": "string"
		//	}
		"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AuditCheckConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies which audit checks are enabled and disabled for this account.",
		//	  "properties": {
		//	    "AuthenticatedCognitoRoleOverlyPermissiveCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "CaCertificateExpiringCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "CaCertificateKeyQualityCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ConflictingClientIdsCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "DeviceCertificateExpiringCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "DeviceCertificateKeyQualityCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "DeviceCertificateSharedCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "IntermediateCaRevokedForActiveDeviceCertificatesCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "IoTPolicyPotentialMisConfigurationCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "IotPolicyOverlyPermissiveCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "IotRoleAliasAllowsAccessToUnusedServicesCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "IotRoleAliasOverlyPermissiveCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "LoggingDisabledCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "RevokedCaCertificateStillActiveCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "RevokedDeviceCertificateStillActiveCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "UnauthenticatedCognitoRoleOverlyPermissiveCheck": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration for a specific audit check.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if the check is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"audit_check_configurations": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AuthenticatedCognitoRoleOverlyPermissiveCheck
				"authenticated_cognito_role_overly_permissive_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: CaCertificateExpiringCheck
				"ca_certificate_expiring_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: CaCertificateKeyQualityCheck
				"ca_certificate_key_quality_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ConflictingClientIdsCheck
				"conflicting_client_ids_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DeviceCertificateExpiringCheck
				"device_certificate_expiring_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DeviceCertificateKeyQualityCheck
				"device_certificate_key_quality_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DeviceCertificateSharedCheck
				"device_certificate_shared_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IntermediateCaRevokedForActiveDeviceCertificatesCheck
				"intermediate_ca_revoked_for_active_device_certificates_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IoTPolicyPotentialMisConfigurationCheck
				"io_t_policy_potential_mis_configuration_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IotPolicyOverlyPermissiveCheck
				"iot_policy_overly_permissive_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IotRoleAliasAllowsAccessToUnusedServicesCheck
				"iot_role_alias_allows_access_to_unused_services_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IotRoleAliasOverlyPermissiveCheck
				"iot_role_alias_overly_permissive_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: LoggingDisabledCheck
				"logging_disabled_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RevokedCaCertificateStillActiveCheck
				"revoked_ca_certificate_still_active_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RevokedDeviceCertificateStillActiveCheck
				"revoked_device_certificate_still_active_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UnauthenticatedCognitoRoleOverlyPermissiveCheck
				"unauthenticated_cognito_role_overly_permissive_check": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if the check is enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration for a specific audit check.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies which audit checks are enabled and disabled for this account.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AuditNotificationTargetConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Information about the targets to which audit notifications are sent.",
		//	  "properties": {
		//	    "Sns": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "True if notifications to the target are enabled.",
		//	          "type": "boolean"
		//	        },
		//	        "RoleArn": {
		//	          "description": "The ARN of the role that grants permission to send notifications to the target.",
		//	          "maxLength": 2048,
		//	          "minLength": 20,
		//	          "type": "string"
		//	        },
		//	        "TargetArn": {
		//	          "description": "The ARN of the target (SNS topic) to which audit notifications are sent.",
		//	          "maxLength": 2048,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"audit_notification_target_configurations": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Sns
				"sns": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "True if notifications to the target are enabled.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: RoleArn
						"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the role that grants permission to send notifications to the target.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: TargetArn
						"target_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the target (SNS topic) to which audit notifications are sent.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Information about the targets to which audit notifications are sent.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the role that grants permission to AWS IoT to access information about your devices, policies, certificates and other items as required when performing an audit.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the role that grants permission to AWS IoT to access information about your devices, policies, certificates and other items as required when performing an audit.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoT::AccountAuditConfiguration",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::AccountAuditConfiguration").WithTerraformTypeName("awscc_iot_account_audit_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":                                                   "AccountId",
		"audit_check_configurations":                                   "AuditCheckConfigurations",
		"audit_notification_target_configurations":                     "AuditNotificationTargetConfigurations",
		"authenticated_cognito_role_overly_permissive_check":           "AuthenticatedCognitoRoleOverlyPermissiveCheck",
		"ca_certificate_expiring_check":                                "CaCertificateExpiringCheck",
		"ca_certificate_key_quality_check":                             "CaCertificateKeyQualityCheck",
		"conflicting_client_ids_check":                                 "ConflictingClientIdsCheck",
		"device_certificate_expiring_check":                            "DeviceCertificateExpiringCheck",
		"device_certificate_key_quality_check":                         "DeviceCertificateKeyQualityCheck",
		"device_certificate_shared_check":                              "DeviceCertificateSharedCheck",
		"enabled":                                                      "Enabled",
		"intermediate_ca_revoked_for_active_device_certificates_check": "IntermediateCaRevokedForActiveDeviceCertificatesCheck",
		"io_t_policy_potential_mis_configuration_check":                "IoTPolicyPotentialMisConfigurationCheck",
		"iot_policy_overly_permissive_check":                           "IotPolicyOverlyPermissiveCheck",
		"iot_role_alias_allows_access_to_unused_services_check":        "IotRoleAliasAllowsAccessToUnusedServicesCheck",
		"iot_role_alias_overly_permissive_check":                       "IotRoleAliasOverlyPermissiveCheck",
		"logging_disabled_check":                                       "LoggingDisabledCheck",
		"revoked_ca_certificate_still_active_check":                    "RevokedCaCertificateStillActiveCheck",
		"revoked_device_certificate_still_active_check":                "RevokedDeviceCertificateStillActiveCheck",
		"role_arn":   "RoleArn",
		"sns":        "Sns",
		"target_arn": "TargetArn",
		"unauthenticated_cognito_role_overly_permissive_check": "UnauthenticatedCognitoRoleOverlyPermissiveCheck",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}