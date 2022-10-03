// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_personalize_dataset_group", datasetGroupDataSource)
}

// datasetGroupDataSource returns the Terraform awscc_personalize_dataset_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::Personalize::DatasetGroup resource.
func datasetGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"dataset_group_arn": {
			// Property: DatasetGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the dataset group.",
			//   "maxLength": 256,
			//   "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the dataset group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"domain": {
			// Property: Domain
			// CloudFormation resource type schema:
			// {
			//   "description": "The domain of a Domain dataset group.",
			//   "enum": [
			//     "ECOMMERCE",
			//     "VIDEO_ON_DEMAND"
			//   ],
			//   "type": "string"
			// }
			Description: "The domain of a Domain dataset group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"kms_key_arn": {
			// Property: KmsKeyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name(ARN) of a AWS Key Management Service (KMS) key used to encrypt the datasets.",
			//   "maxLength": 2048,
			//   "pattern": "arn:aws.*:kms:.*:[0-9]{12}:key/.*",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name(ARN) of a AWS Key Management Service (KMS) key used to encrypt the datasets.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the new dataset group.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//   "type": "string"
			// }
			Description: "The name for the new dataset group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the AWS Identity and Access Management (IAM) role that has permissions to access the AWS Key Management Service (KMS) key. Supplying an IAM role is only valid when also specifying a KMS key.",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "pattern": "arn:([a-z\\d-]+):iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+",
			//   "type": "string"
			// }
			Description: "The ARN of the AWS Identity and Access Management (IAM) role that has permissions to access the AWS Key Management Service (KMS) key. Supplying an IAM role is only valid when also specifying a KMS key.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Personalize::DatasetGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Personalize::DatasetGroup").WithTerraformTypeName("awscc_personalize_dataset_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dataset_group_arn": "DatasetGroupArn",
		"domain":            "Domain",
		"kms_key_arn":       "KmsKeyArn",
		"name":              "Name",
		"role_arn":          "RoleArn",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
