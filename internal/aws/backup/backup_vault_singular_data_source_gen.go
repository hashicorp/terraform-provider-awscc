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
	registry.AddDataSourceFactory("awscc_backup_backup_vault", backupVaultDataSource)
}

// backupVaultDataSource returns the Terraform awscc_backup_backup_vault data source.
// This Terraform data source corresponds to the CloudFormation AWS::Backup::BackupVault resource.
func backupVaultDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"access_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BackupVaultArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"backup_vault_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BackupVaultName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^[a-zA-Z0-9\\-\\_]{2,50}$",
		//	  "type": "string"
		//	}
		"backup_vault_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BackupVaultTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"backup_vault_tags": // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EncryptionKeyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"encryption_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LockConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ChangeableForDays": {
		//	      "type": "integer"
		//	    },
		//	    "MaxRetentionDays": {
		//	      "type": "integer"
		//	    },
		//	    "MinRetentionDays": {
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "MinRetentionDays"
		//	  ],
		//	  "type": "object"
		//	}
		"lock_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ChangeableForDays
				"changeable_for_days": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: MaxRetentionDays
				"max_retention_days": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: MinRetentionDays
				"min_retention_days": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Notifications
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "BackupVaultEvents": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "SNSTopicArn": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "SNSTopicArn",
		//	    "BackupVaultEvents"
		//	  ],
		//	  "type": "object"
		//	}
		"notifications": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BackupVaultEvents
				"backup_vault_events": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SNSTopicArn
				"sns_topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Backup::BackupVault",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Backup::BackupVault").WithTerraformTypeName("awscc_backup_backup_vault")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_policy":       "AccessPolicy",
		"backup_vault_arn":    "BackupVaultArn",
		"backup_vault_events": "BackupVaultEvents",
		"backup_vault_name":   "BackupVaultName",
		"backup_vault_tags":   "BackupVaultTags",
		"changeable_for_days": "ChangeableForDays",
		"encryption_key_arn":  "EncryptionKeyArn",
		"lock_configuration":  "LockConfiguration",
		"max_retention_days":  "MaxRetentionDays",
		"min_retention_days":  "MinRetentionDays",
		"notifications":       "Notifications",
		"sns_topic_arn":       "SNSTopicArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
