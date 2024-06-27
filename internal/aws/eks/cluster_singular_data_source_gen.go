// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package eks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_eks_cluster", clusterDataSource)
}

// clusterDataSource returns the Terraform awscc_eks_cluster data source.
// This Terraform data source corresponds to the CloudFormation AWS::EKS::Cluster resource.
func clusterDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object representing the Access Config to use for the cluster.",
		//	  "properties": {
		//	    "AuthenticationMode": {
		//	      "description": "Specify the authentication mode that should be used to create your cluster.",
		//	      "enum": [
		//	        "CONFIG_MAP",
		//	        "API_AND_CONFIG_MAP",
		//	        "API"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "BootstrapClusterCreatorAdminPermissions": {
		//	      "description": "Set this value to false to avoid creating a default cluster admin Access Entry using the IAM principal used to create the cluster.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"access_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AuthenticationMode
				"authentication_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Specify the authentication mode that should be used to create your cluster.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: BootstrapClusterCreatorAdminPermissions
				"bootstrap_cluster_creator_admin_permissions": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Set this value to false to avoid creating a default cluster admin Access Entry using the IAM principal used to create the cluster.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object representing the Access Config to use for the cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BootstrapSelfManagedAddons
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Set this value to false to avoid creating the default networking addons when the cluster is created.",
		//	  "type": "boolean"
		//	}
		"bootstrap_self_managed_addons": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Set this value to false to avoid creating the default networking addons when the cluster is created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CertificateAuthorityData
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The certificate-authority-data for your cluster.",
		//	  "type": "string"
		//	}
		"certificate_authority_data": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The certificate-authority-data for your cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ClusterSecurityGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control plane to data plane communication.",
		//	  "type": "string"
		//	}
		"cluster_security_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control plane to data plane communication.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EncryptionConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The encryption configuration for the cluster",
		//	    "properties": {
		//	      "Provider": {
		//	        "additionalProperties": false,
		//	        "description": "The encryption provider for the cluster.",
		//	        "properties": {
		//	          "KeyArn": {
		//	            "description": "Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Resources": {
		//	        "description": "Specifies the resources to be encrypted. The only supported value is \"secrets\".",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "type": "array"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"encryption_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Provider
					"provider": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: KeyArn
							"key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The encryption provider for the cluster.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Resources
					"resources": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "Specifies the resources to be encrypted. The only supported value is \"secrets\".",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EncryptionConfigKeyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) or alias of the customer master key (CMK).",
		//	  "type": "string"
		//	}
		"encryption_config_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) or alias of the customer master key (CMK).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The endpoint for your Kubernetes API server, such as https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com.",
		//	  "type": "string"
		//	}
		"endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The endpoint for your Kubernetes API server, such as https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique ID given to your cluster.",
		//	  "type": "string"
		//	}
		"cluster_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique ID given to your cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KubernetesNetworkConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The Kubernetes network configuration for the cluster.",
		//	  "properties": {
		//	    "IpFamily": {
		//	      "description": "Ipv4 or Ipv6. You can only specify ipv6 for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on",
		//	      "enum": [
		//	        "ipv4",
		//	        "ipv6"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "ServiceIpv4Cidr": {
		//	      "description": "The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. ",
		//	      "type": "string"
		//	    },
		//	    "ServiceIpv6Cidr": {
		//	      "description": "The CIDR block to assign Kubernetes service IP addresses from.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"kubernetes_network_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: IpFamily
				"ip_family": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Ipv4 or Ipv6. You can only specify ipv6 for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ServiceIpv4Cidr
				"service_ipv_4_cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ServiceIpv6Cidr
				"service_ipv_6_cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The CIDR block to assign Kubernetes service IP addresses from.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The Kubernetes network configuration for the cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Logging
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs.",
		//	  "properties": {
		//	    "ClusterLogging": {
		//	      "additionalProperties": false,
		//	      "description": "The cluster control plane logging configuration for your cluster. ",
		//	      "properties": {
		//	        "EnabledTypes": {
		//	          "description": "Enable control plane logs for your cluster, all log types will be disabled if the array is empty",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "description": "Enabled Logging Type",
		//	            "properties": {
		//	              "Type": {
		//	                "description": "name of the log type",
		//	                "enum": [
		//	                  "api",
		//	                  "audit",
		//	                  "authenticator",
		//	                  "controllerManager",
		//	                  "scheduler"
		//	                ],
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"logging": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ClusterLogging
				"cluster_logging": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EnabledTypes
						"enabled_types": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Type
									"type": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "name of the log type",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Description: "Enable control plane logs for your cluster, all log types will be disabled if the array is empty",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The cluster control plane logging configuration for your cluster. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique name to give to your cluster.",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9A-Za-z][A-Za-z0-9\\-_]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique name to give to your cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OpenIdConnectIssuerUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The issuer URL for the cluster's OIDC identity provider, such as https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E. If you need to remove https:// from this output value, you can include the following code in your template.",
		//	  "type": "string"
		//	}
		"open_id_connect_issuer_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The issuer URL for the cluster's OIDC identity provider, such as https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E. If you need to remove https:// from this output value, you can include the following code in your template.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OutpostConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object representing the Outpost configuration to use for AWS EKS outpost cluster.",
		//	  "properties": {
		//	    "ControlPlaneInstanceType": {
		//	      "description": "Specify the Instance type of the machines that should be used to create your cluster.",
		//	      "type": "string"
		//	    },
		//	    "ControlPlanePlacement": {
		//	      "additionalProperties": false,
		//	      "description": "Specify the placement group of the control plane machines for your cluster.",
		//	      "properties": {
		//	        "GroupName": {
		//	          "description": "Specify the placement group name of the control place machines for your cluster.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "OutpostArns": {
		//	      "description": "Specify one or more Arn(s) of Outpost(s) on which you would like to create your cluster.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "minItems": 1,
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "OutpostArns",
		//	    "ControlPlaneInstanceType"
		//	  ],
		//	  "type": "object"
		//	}
		"outpost_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ControlPlaneInstanceType
				"control_plane_instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Specify the Instance type of the machines that should be used to create your cluster.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ControlPlanePlacement
				"control_plane_placement": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: GroupName
						"group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Specify the placement group name of the control place machines for your cluster.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specify the placement group of the control plane machines for your cluster.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: OutpostArns
				"outpost_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Specify one or more Arn(s) of Outpost(s) on which you would like to create your cluster.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object representing the Outpost configuration to use for AWS EKS outpost cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourcesVpcConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object representing the VPC configuration to use for an Amazon EKS cluster.",
		//	  "properties": {
		//	    "EndpointPrivateAccess": {
		//	      "description": "Set this value to true to enable private access for your cluster's Kubernetes API server endpoint. If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is false, which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks for communication with the nodes or Fargate pods.",
		//	      "type": "boolean"
		//	    },
		//	    "EndpointPublicAccess": {
		//	      "description": "Set this value to false to disable public access to your cluster's Kubernetes API server endpoint. If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is true, which enables public access for your Kubernetes API server.",
		//	      "type": "boolean"
		//	    },
		//	    "PublicAccessCidrs": {
		//	      "description": "The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint. Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "minItems": 1,
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "SecurityGroupIds": {
		//	      "description": "Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane. If you don't specify a security group, the default security group for your VPC is used.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "minItems": 1,
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "SubnetIds": {
		//	      "description": "Specify subnets for your Amazon EKS nodes. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "minItems": 1,
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "SubnetIds"
		//	  ],
		//	  "type": "object"
		//	}
		"resources_vpc_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EndpointPrivateAccess
				"endpoint_private_access": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Set this value to true to enable private access for your cluster's Kubernetes API server endpoint. If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is false, which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks for communication with the nodes or Fargate pods.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: EndpointPublicAccess
				"endpoint_public_access": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Set this value to false to disable public access to your cluster's Kubernetes API server endpoint. If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is true, which enables public access for your Kubernetes API server.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PublicAccessCidrs
				"public_access_cidrs": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint. Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SecurityGroupIds
				"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane. If you don't specify a security group, the default security group for your VPC is used.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SubnetIds
				"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Specify subnets for your Amazon EKS nodes. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object representing the VPC configuration to use for an Amazon EKS cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.",
		//	  "pattern": "1\\.\\d\\d",
		//	  "type": "string"
		//	}
		"version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EKS::Cluster",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EKS::Cluster").WithTerraformTypeName("awscc_eks_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_config":       "AccessConfig",
		"arn":                 "Arn",
		"authentication_mode": "AuthenticationMode",
		"bootstrap_cluster_creator_admin_permissions": "BootstrapClusterCreatorAdminPermissions",
		"bootstrap_self_managed_addons":               "BootstrapSelfManagedAddons",
		"certificate_authority_data":                  "CertificateAuthorityData",
		"cluster_id":                                  "Id",
		"cluster_logging":                             "ClusterLogging",
		"cluster_security_group_id":                   "ClusterSecurityGroupId",
		"control_plane_instance_type":                 "ControlPlaneInstanceType",
		"control_plane_placement":                     "ControlPlanePlacement",
		"enabled_types":                               "EnabledTypes",
		"encryption_config":                           "EncryptionConfig",
		"encryption_config_key_arn":                   "EncryptionConfigKeyArn",
		"endpoint":                                    "Endpoint",
		"endpoint_private_access":                     "EndpointPrivateAccess",
		"endpoint_public_access":                      "EndpointPublicAccess",
		"group_name":                                  "GroupName",
		"ip_family":                                   "IpFamily",
		"key":                                         "Key",
		"key_arn":                                     "KeyArn",
		"kubernetes_network_config":                   "KubernetesNetworkConfig",
		"logging":                                     "Logging",
		"name":                                        "Name",
		"open_id_connect_issuer_url":                  "OpenIdConnectIssuerUrl",
		"outpost_arns":                                "OutpostArns",
		"outpost_config":                              "OutpostConfig",
		"provider":                                    "Provider",
		"public_access_cidrs":                         "PublicAccessCidrs",
		"resources":                                   "Resources",
		"resources_vpc_config":                        "ResourcesVpcConfig",
		"role_arn":                                    "RoleArn",
		"security_group_ids":                          "SecurityGroupIds",
		"service_ipv_4_cidr":                          "ServiceIpv4Cidr",
		"service_ipv_6_cidr":                          "ServiceIpv6Cidr",
		"subnet_ids":                                  "SubnetIds",
		"tags":                                        "Tags",
		"type":                                        "Type",
		"value":                                       "Value",
		"version":                                     "Version",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
