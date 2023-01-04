// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package identitystore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_identitystore_group_membership", groupMembershipDataSource)
}

// groupMembershipDataSource returns the Terraform awscc_identitystore_group_membership data source.
// This Terraform data source corresponds to the CloudFormation AWS::IdentityStore::GroupMembership resource.
func groupMembershipDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: GroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier for a group in the identity store.",
		//	  "maxLength": 47,
		//	  "minLength": 1,
		//	  "pattern": "^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$",
		//	  "type": "string"
		//	}
		"group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier for a group in the identity store.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdentityStoreId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The globally unique identifier for the identity store.",
		//	  "maxLength": 36,
		//	  "minLength": 1,
		//	  "pattern": "^d-[0-9a-f]{10}$|^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$",
		//	  "type": "string"
		//	}
		"identity_store_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The globally unique identifier for the identity store.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MemberId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object containing the identifier of a group member.",
		//	  "properties": {
		//	    "UserId": {
		//	      "description": "The identifier for a user in the identity store.",
		//	      "maxLength": 47,
		//	      "minLength": 1,
		//	      "pattern": "^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "UserId"
		//	  ],
		//	  "type": "object"
		//	}
		"member_id": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: UserId
				"user_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The identifier for a user in the identity store.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object containing the identifier of a group member.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MembershipId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for a GroupMembership in the identity store.",
		//	  "maxLength": 47,
		//	  "minLength": 1,
		//	  "pattern": "^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$",
		//	  "type": "string"
		//	}
		"membership_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for a GroupMembership in the identity store.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IdentityStore::GroupMembership",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IdentityStore::GroupMembership").WithTerraformTypeName("awscc_identitystore_group_membership")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"group_id":          "GroupId",
		"identity_store_id": "IdentityStoreId",
		"member_id":         "MemberId",
		"membership_id":     "MembershipId",
		"user_id":           "UserId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}