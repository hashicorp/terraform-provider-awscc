// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package notifications

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_notifications_managed_notification_account_contact_association", managedNotificationAccountContactAssociationDataSource)
}

// managedNotificationAccountContactAssociationDataSource returns the Terraform awscc_notifications_managed_notification_account_contact_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::Notifications::ManagedNotificationAccountContactAssociation resource.
func managedNotificationAccountContactAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ContactIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This unique identifier for Contact",
		//	  "enum": [
		//	    "ACCOUNT_PRIMARY",
		//	    "ACCOUNT_ALTERNATE_SECURITY",
		//	    "ACCOUNT_ALTERNATE_OPERATIONS",
		//	    "ACCOUNT_ALTERNATE_BILLING"
		//	  ],
		//	  "type": "string"
		//	}
		"contact_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This unique identifier for Contact",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ManagedNotificationConfigurationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The managed notification configuration ARN, against which the account contact association will be created",
		//	  "pattern": "^arn:[-.a-z0-9]{1,63}:notifications::[0-9]{12}:managed-notification-configuration/category/[a-zA-Z0-9-]{3,64}/sub-category/[a-zA-Z0-9-]{3,64}$",
		//	  "type": "string"
		//	}
		"managed_notification_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The managed notification configuration ARN, against which the account contact association will be created",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Notifications::ManagedNotificationAccountContactAssociation",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Notifications::ManagedNotificationAccountContactAssociation").WithTerraformTypeName("awscc_notifications_managed_notification_account_contact_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"contact_identifier":                     "ContactIdentifier",
		"managed_notification_configuration_arn": "ManagedNotificationConfigurationArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
