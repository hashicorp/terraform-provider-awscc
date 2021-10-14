// Code generated by generators/resource/main.go; DO NOT EDIT.

package healthlake

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_healthlake_fhir_datastore", fHIRDatastoreResourceType)
}

// fHIRDatastoreResourceType returns the Terraform awscc_healthlake_fhir_datastore resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::HealthLake::FHIRDatastore resource type.
func fHIRDatastoreResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The time that a Data Store was created.",
			//   "properties": {
			//     "Nanos": {
			//       "description": "Nanoseconds.",
			//       "type": "integer"
			//     },
			//     "Seconds": {
			//       "description": "Seconds since epoch.",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Seconds",
			//     "Nanos"
			//   ],
			//   "type": "object"
			// }
			Description: "The time that a Data Store was created.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"nanos": {
						// Property: Nanos
						Description: "Nanoseconds.",
						Type:        types.NumberType,
						Required:    true,
					},
					"seconds": {
						// Property: Seconds
						Description: "Seconds since epoch.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Computed: true,
		},
		"datastore_arn": {
			// Property: DatastoreArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name used in the creation of the Data Store.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name used in the creation of the Data Store.",
			Type:        types.StringType,
			Computed:    true,
		},
		"datastore_endpoint": {
			// Property: DatastoreEndpoint
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS endpoint for the Data Store. Each Data Store will have it's own endpoint with Data Store ID in the endpoint URL.",
			//   "maxLength": 10000,
			//   "type": "string"
			// }
			Description: "The AWS endpoint for the Data Store. Each Data Store will have it's own endpoint with Data Store ID in the endpoint URL.",
			Type:        types.StringType,
			Computed:    true,
		},
		"datastore_id": {
			// Property: DatastoreId
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS-generated ID number for the Data Store.",
			//   "maxLength": 32,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The AWS-generated ID number for the Data Store.",
			Type:        types.StringType,
			Computed:    true,
		},
		"datastore_name": {
			// Property: DatastoreName
			// CloudFormation resource type schema:
			// {
			//   "description": "The user-generated name for the Data Store.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The user-generated name for the Data Store.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				ComputedOptionalForceNew(),
			},
		},
		"datastore_status": {
			// Property: DatastoreStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the Data Store. Possible statuses are 'CREATING', 'ACTIVE', 'DELETING', or 'DELETED'.",
			//   "enum": [
			//     "CREATING",
			//     "ACTIVE",
			//     "DELETING",
			//     "DELETED"
			//   ],
			//   "type": "string"
			// }
			Description: "The status of the Data Store. Possible statuses are 'CREATING', 'ACTIVE', 'DELETING', or 'DELETED'.",
			Type:        types.StringType,
			Computed:    true,
		},
		"datastore_type_version": {
			// Property: DatastoreTypeVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The FHIR version. Only R4 version data is supported.",
			//   "enum": [
			//     "R4"
			//   ],
			//   "type": "string"
			// }
			Description: "The FHIR version. Only R4 version data is supported.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"R4",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"preload_data_config": {
			// Property: PreloadDataConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The preloaded data configuration for the Data Store. Only data preloaded from Synthea is supported.",
			//   "properties": {
			//     "PreloadDataType": {
			//       "description": "The type of preloaded data. Only Synthea preloaded data is supported.",
			//       "enum": [
			//         "SYNTHEA"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "PreloadDataType"
			//   ],
			//   "type": "object"
			// }
			Description: "The preloaded data configuration for the Data Store. Only data preloaded from Synthea is supported.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"preload_data_type": {
						// Property: PreloadDataType
						Description: "The type of preloaded data. Only Synthea preloaded data is supported.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"SYNTHEA",
							}),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				ComputedOptionalForceNew(),
			},
		},
		"sse_configuration": {
			// Property: SseConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The server-side encryption key configuration for a customer provided encryption key.",
			//   "properties": {
			//     "KmsEncryptionConfig": {
			//       "additionalProperties": false,
			//       "description": "The customer-managed-key (CMK) used when creating a Data Store. If a customer owned key is not specified, an AWS owned key will be used for encryption.",
			//       "properties": {
			//         "CmkType": {
			//           "description": "The type of customer-managed-key (CMK) used for encryption. The two types of supported CMKs are customer owned CMKs and AWS owned CMKs.",
			//           "enum": [
			//             "CUSTOMER_MANAGED_KMS_KEY",
			//             "AWS_OWNED_KMS_KEY"
			//           ],
			//           "type": "string"
			//         },
			//         "KmsKeyId": {
			//           "description": "The KMS encryption key id/alias used to encrypt the Data Store contents at rest.",
			//           "maxLength": 400,
			//           "minLength": 1,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "CmkType"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "KmsEncryptionConfig"
			//   ],
			//   "type": "object"
			// }
			Description: "The server-side encryption key configuration for a customer provided encryption key.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"kms_encryption_config": {
						// Property: KmsEncryptionConfig
						Description: "The customer-managed-key (CMK) used when creating a Data Store. If a customer owned key is not specified, an AWS owned key will be used for encryption.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"cmk_type": {
									// Property: CmkType
									Description: "The type of customer-managed-key (CMK) used for encryption. The two types of supported CMKs are customer owned CMKs and AWS owned CMKs.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"CUSTOMER_MANAGED_KMS_KEY",
											"AWS_OWNED_KMS_KEY",
										}),
									},
								},
								"kms_key_id": {
									// Property: KmsKeyId
									Description: "The KMS encryption key id/alias used to encrypt the Data Store contents at rest.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 400),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				ComputedOptionalForceNew(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair. A tag consists of a tag key and a tag value. Tag keys and tag values are both required, but tag values can be empty (null) strings.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key of the tag.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value of the tag.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key of the tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value of the tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "HealthLake FHIR Datastore",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::HealthLake::FHIRDatastore").WithTerraformTypeName("awscc_healthlake_fhir_datastore")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cmk_type":               "CmkType",
		"created_at":             "CreatedAt",
		"datastore_arn":          "DatastoreArn",
		"datastore_endpoint":     "DatastoreEndpoint",
		"datastore_id":           "DatastoreId",
		"datastore_name":         "DatastoreName",
		"datastore_status":       "DatastoreStatus",
		"datastore_type_version": "DatastoreTypeVersion",
		"key":                    "Key",
		"kms_encryption_config":  "KmsEncryptionConfig",
		"kms_key_id":             "KmsKeyId",
		"nanos":                  "Nanos",
		"preload_data_config":    "PreloadDataConfig",
		"preload_data_type":      "PreloadDataType",
		"seconds":                "Seconds",
		"sse_configuration":      "SseConfiguration",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
