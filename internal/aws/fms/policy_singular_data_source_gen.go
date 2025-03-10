// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package fms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_fms_policy", policyDataSource)
}

// policyDataSource returns the Terraform awscc_fms_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::FMS::Policy resource.
func policyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A resource ARN.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "pattern": "^([^\\s]*)$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A resource ARN.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeleteAllPolicyResources
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"delete_all_policy_resources": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ExcludeMap
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An FMS includeMap or excludeMap.",
		//	  "properties": {
		//	    "ACCOUNT": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "description": "An AWS account ID.",
		//	        "maxLength": 12,
		//	        "minLength": 12,
		//	        "pattern": "^([0-9]*)$",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "ORGUNIT": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "description": "An Organizational Unit ID.",
		//	        "maxLength": 68,
		//	        "minLength": 16,
		//	        "pattern": "^(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"exclude_map": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ACCOUNT
				"account": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ORGUNIT
				"orgunit": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An FMS includeMap or excludeMap.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ExcludeResourceTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"exclude_resource_tags": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "^[a-z0-9A-Z-]{36}$",
		//	  "type": "string"
		//	}
		"policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IncludeMap
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An FMS includeMap or excludeMap.",
		//	  "properties": {
		//	    "ACCOUNT": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "description": "An AWS account ID.",
		//	        "maxLength": 12,
		//	        "minLength": 12,
		//	        "pattern": "^([0-9]*)$",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "ORGUNIT": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "description": "An Organizational Unit ID.",
		//	        "maxLength": 68,
		//	        "minLength": 16,
		//	        "pattern": "^(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"include_map": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ACCOUNT
				"account": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ORGUNIT
				"orgunit": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An FMS includeMap or excludeMap.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "pattern": "^([a-zA-Z0-9_.:/=+\\-@\\s]+)$",
		//	  "type": "string"
		//	}
		"policy_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "pattern": "^([a-zA-Z0-9_.:/=+\\-@\\s]+)$",
		//	  "type": "string"
		//	}
		"policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RemediationEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"remediation_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceSetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "description": "A Base62 ID",
		//	    "maxLength": 22,
		//	    "minLength": 22,
		//	    "pattern": "^[a-z0-9A-Z]{22}$",
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"resource_set_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceTagLogicalOperator
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "AND",
		//	    "OR"
		//	  ],
		//	  "type": "string"
		//	}
		"resource_tag_logical_operator": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A resource tag.",
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 8,
		//	  "type": "array"
		//	}
		"resource_tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
		// Property: ResourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An AWS resource type",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^([^\\s]*)$",
		//	  "type": "string"
		//	}
		"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An AWS resource type",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceTypeList
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "description": "An AWS resource type",
		//	    "maxLength": 128,
		//	    "minLength": 1,
		//	    "pattern": "^([^\\s]*)$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"resource_type_list": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourcesCleanUp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"resources_clean_up": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityServicePolicyData
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Firewall security service policy data.",
		//	  "properties": {
		//	    "ManagedServiceData": {
		//	      "description": "Firewall managed service data.",
		//	      "maxLength": 30000,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "PolicyOption": {
		//	      "additionalProperties": false,
		//	      "description": "Firewall policy option.",
		//	      "oneOf": [
		//	        {
		//	          "required": [
		//	            "NetworkFirewallPolicy"
		//	          ]
		//	        },
		//	        {
		//	          "required": [
		//	            "ThirdPartyFirewallPolicy"
		//	          ]
		//	        },
		//	        {
		//	          "required": [
		//	            "NetworkAclCommonPolicy"
		//	          ]
		//	        }
		//	      ],
		//	      "properties": {
		//	        "NetworkAclCommonPolicy": {
		//	          "additionalProperties": false,
		//	          "description": "Network ACL common policy.",
		//	          "properties": {
		//	            "NetworkAclEntrySet": {
		//	              "additionalProperties": false,
		//	              "anyOf": [
		//	                {
		//	                  "required": [
		//	                    "FirstEntries"
		//	                  ]
		//	                },
		//	                {
		//	                  "required": [
		//	                    "LastEntries"
		//	                  ]
		//	                }
		//	              ],
		//	              "description": "Network ACL entry set.",
		//	              "properties": {
		//	                "FirstEntries": {
		//	                  "description": "NetworkAcl entry list.",
		//	                  "insertionOrder": true,
		//	                  "items": {
		//	                    "additionalProperties": false,
		//	                    "description": "Network ACL entry.",
		//	                    "properties": {
		//	                      "CidrBlock": {
		//	                        "description": "CIDR block.",
		//	                        "pattern": "^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(\\/([0-9]|[1-2][0-9]|3[0-2]))$",
		//	                        "type": "string"
		//	                      },
		//	                      "Egress": {
		//	                        "description": "Whether the entry is an egress entry.",
		//	                        "type": "boolean"
		//	                      },
		//	                      "IcmpTypeCode": {
		//	                        "additionalProperties": false,
		//	                        "description": "ICMP type and code.",
		//	                        "properties": {
		//	                          "Code": {
		//	                            "description": "Code.",
		//	                            "maximum": 255,
		//	                            "minimum": 0,
		//	                            "type": "integer"
		//	                          },
		//	                          "Type": {
		//	                            "description": "Type.",
		//	                            "maximum": 255,
		//	                            "minimum": 0,
		//	                            "type": "integer"
		//	                          }
		//	                        },
		//	                        "required": [
		//	                          "Code",
		//	                          "Type"
		//	                        ],
		//	                        "type": "object"
		//	                      },
		//	                      "Ipv6CidrBlock": {
		//	                        "description": "IPv6 CIDR block.",
		//	                        "pattern": "^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))(/(1[0-2]|[0-9]))?$",
		//	                        "type": "string"
		//	                      },
		//	                      "PortRange": {
		//	                        "additionalProperties": false,
		//	                        "description": "Port range.",
		//	                        "properties": {
		//	                          "From": {
		//	                            "description": "From Port.",
		//	                            "maximum": 65535,
		//	                            "minimum": 0,
		//	                            "type": "integer"
		//	                          },
		//	                          "To": {
		//	                            "description": "To Port.",
		//	                            "maximum": 65535,
		//	                            "minimum": 0,
		//	                            "type": "integer"
		//	                          }
		//	                        },
		//	                        "required": [
		//	                          "From",
		//	                          "To"
		//	                        ],
		//	                        "type": "object"
		//	                      },
		//	                      "Protocol": {
		//	                        "description": "Protocol.",
		//	                        "pattern": "^(tcp|udp|icmp|-1|([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))$",
		//	                        "type": "string"
		//	                      },
		//	                      "RuleAction": {
		//	                        "description": "Rule Action.",
		//	                        "enum": [
		//	                          "allow",
		//	                          "deny"
		//	                        ],
		//	                        "type": "string"
		//	                      }
		//	                    },
		//	                    "required": [
		//	                      "Egress",
		//	                      "Protocol",
		//	                      "RuleAction"
		//	                    ],
		//	                    "type": "object"
		//	                  },
		//	                  "type": "array"
		//	                },
		//	                "ForceRemediateForFirstEntries": {
		//	                  "type": "boolean"
		//	                },
		//	                "ForceRemediateForLastEntries": {
		//	                  "type": "boolean"
		//	                },
		//	                "LastEntries": {
		//	                  "description": "NetworkAcl entry list.",
		//	                  "insertionOrder": true,
		//	                  "items": {
		//	                    "additionalProperties": false,
		//	                    "description": "Network ACL entry.",
		//	                    "properties": {
		//	                      "CidrBlock": {
		//	                        "description": "CIDR block.",
		//	                        "pattern": "^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(\\/([0-9]|[1-2][0-9]|3[0-2]))$",
		//	                        "type": "string"
		//	                      },
		//	                      "Egress": {
		//	                        "description": "Whether the entry is an egress entry.",
		//	                        "type": "boolean"
		//	                      },
		//	                      "IcmpTypeCode": {
		//	                        "additionalProperties": false,
		//	                        "description": "ICMP type and code.",
		//	                        "properties": {
		//	                          "Code": {
		//	                            "description": "Code.",
		//	                            "maximum": 255,
		//	                            "minimum": 0,
		//	                            "type": "integer"
		//	                          },
		//	                          "Type": {
		//	                            "description": "Type.",
		//	                            "maximum": 255,
		//	                            "minimum": 0,
		//	                            "type": "integer"
		//	                          }
		//	                        },
		//	                        "required": [
		//	                          "Code",
		//	                          "Type"
		//	                        ],
		//	                        "type": "object"
		//	                      },
		//	                      "Ipv6CidrBlock": {
		//	                        "description": "IPv6 CIDR block.",
		//	                        "pattern": "^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))(/(1[0-2]|[0-9]))?$",
		//	                        "type": "string"
		//	                      },
		//	                      "PortRange": {
		//	                        "additionalProperties": false,
		//	                        "description": "Port range.",
		//	                        "properties": {
		//	                          "From": {
		//	                            "description": "From Port.",
		//	                            "maximum": 65535,
		//	                            "minimum": 0,
		//	                            "type": "integer"
		//	                          },
		//	                          "To": {
		//	                            "description": "To Port.",
		//	                            "maximum": 65535,
		//	                            "minimum": 0,
		//	                            "type": "integer"
		//	                          }
		//	                        },
		//	                        "required": [
		//	                          "From",
		//	                          "To"
		//	                        ],
		//	                        "type": "object"
		//	                      },
		//	                      "Protocol": {
		//	                        "description": "Protocol.",
		//	                        "pattern": "^(tcp|udp|icmp|-1|([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))$",
		//	                        "type": "string"
		//	                      },
		//	                      "RuleAction": {
		//	                        "description": "Rule Action.",
		//	                        "enum": [
		//	                          "allow",
		//	                          "deny"
		//	                        ],
		//	                        "type": "string"
		//	                      }
		//	                    },
		//	                    "required": [
		//	                      "Egress",
		//	                      "Protocol",
		//	                      "RuleAction"
		//	                    ],
		//	                    "type": "object"
		//	                  },
		//	                  "type": "array"
		//	                }
		//	              },
		//	              "required": [
		//	                "ForceRemediateForFirstEntries",
		//	                "ForceRemediateForLastEntries"
		//	              ],
		//	              "type": "object"
		//	            }
		//	          },
		//	          "required": [
		//	            "NetworkAclEntrySet"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "NetworkFirewallPolicy": {
		//	          "additionalProperties": false,
		//	          "description": "Network firewall policy.",
		//	          "properties": {
		//	            "FirewallDeploymentModel": {
		//	              "description": "Firewall deployment mode.",
		//	              "enum": [
		//	                "DISTRIBUTED",
		//	                "CENTRALIZED"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "FirewallDeploymentModel"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "ThirdPartyFirewallPolicy": {
		//	          "additionalProperties": false,
		//	          "description": "Third party firewall policy.",
		//	          "properties": {
		//	            "FirewallDeploymentModel": {
		//	              "description": "Firewall deployment mode.",
		//	              "enum": [
		//	                "DISTRIBUTED",
		//	                "CENTRALIZED"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "FirewallDeploymentModel"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Type": {
		//	      "description": "Firewall policy type.",
		//	      "enum": [
		//	        "WAF",
		//	        "WAFV2",
		//	        "SHIELD_ADVANCED",
		//	        "SECURITY_GROUPS_COMMON",
		//	        "SECURITY_GROUPS_CONTENT_AUDIT",
		//	        "SECURITY_GROUPS_USAGE_AUDIT",
		//	        "NETWORK_FIREWALL",
		//	        "THIRD_PARTY_FIREWALL",
		//	        "DNS_FIREWALL",
		//	        "IMPORT_NETWORK_FIREWALL",
		//	        "NETWORK_ACL_COMMON"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type"
		//	  ],
		//	  "type": "object"
		//	}
		"security_service_policy_data": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ManagedServiceData
				"managed_service_data": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Firewall managed service data.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PolicyOption
				"policy_option": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: NetworkAclCommonPolicy
						"network_acl_common_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: NetworkAclEntrySet
								"network_acl_entry_set": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: FirstEntries
										"first_entries": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
											NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: CidrBlock
													"cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
														Description: "CIDR block.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: Egress
													"egress": schema.BoolAttribute{ /*START ATTRIBUTE*/
														Description: "Whether the entry is an egress entry.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: IcmpTypeCode
													"icmp_type_code": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
														Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
															// Property: Code
															"code": schema.Int64Attribute{ /*START ATTRIBUTE*/
																Description: "Code.",
																Computed:    true,
															}, /*END ATTRIBUTE*/
															// Property: Type
															"type": schema.Int64Attribute{ /*START ATTRIBUTE*/
																Description: "Type.",
																Computed:    true,
															}, /*END ATTRIBUTE*/
														}, /*END SCHEMA*/
														Description: "ICMP type and code.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: Ipv6CidrBlock
													"ipv_6_cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
														Description: "IPv6 CIDR block.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: PortRange
													"port_range": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
														Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
															// Property: From
															"from": schema.Int64Attribute{ /*START ATTRIBUTE*/
																Description: "From Port.",
																Computed:    true,
															}, /*END ATTRIBUTE*/
															// Property: To
															"to": schema.Int64Attribute{ /*START ATTRIBUTE*/
																Description: "To Port.",
																Computed:    true,
															}, /*END ATTRIBUTE*/
														}, /*END SCHEMA*/
														Description: "Port range.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: Protocol
													"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
														Description: "Protocol.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: RuleAction
													"rule_action": schema.StringAttribute{ /*START ATTRIBUTE*/
														Description: "Rule Action.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
												}, /*END SCHEMA*/
											}, /*END NESTED OBJECT*/
											Description: "NetworkAcl entry list.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: ForceRemediateForFirstEntries
										"force_remediate_for_first_entries": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: ForceRemediateForLastEntries
										"force_remediate_for_last_entries": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: LastEntries
										"last_entries": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
											NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: CidrBlock
													"cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
														Description: "CIDR block.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: Egress
													"egress": schema.BoolAttribute{ /*START ATTRIBUTE*/
														Description: "Whether the entry is an egress entry.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: IcmpTypeCode
													"icmp_type_code": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
														Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
															// Property: Code
															"code": schema.Int64Attribute{ /*START ATTRIBUTE*/
																Description: "Code.",
																Computed:    true,
															}, /*END ATTRIBUTE*/
															// Property: Type
															"type": schema.Int64Attribute{ /*START ATTRIBUTE*/
																Description: "Type.",
																Computed:    true,
															}, /*END ATTRIBUTE*/
														}, /*END SCHEMA*/
														Description: "ICMP type and code.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: Ipv6CidrBlock
													"ipv_6_cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
														Description: "IPv6 CIDR block.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: PortRange
													"port_range": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
														Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
															// Property: From
															"from": schema.Int64Attribute{ /*START ATTRIBUTE*/
																Description: "From Port.",
																Computed:    true,
															}, /*END ATTRIBUTE*/
															// Property: To
															"to": schema.Int64Attribute{ /*START ATTRIBUTE*/
																Description: "To Port.",
																Computed:    true,
															}, /*END ATTRIBUTE*/
														}, /*END SCHEMA*/
														Description: "Port range.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: Protocol
													"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
														Description: "Protocol.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: RuleAction
													"rule_action": schema.StringAttribute{ /*START ATTRIBUTE*/
														Description: "Rule Action.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
												}, /*END SCHEMA*/
											}, /*END NESTED OBJECT*/
											Description: "NetworkAcl entry list.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Network ACL entry set.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Network ACL common policy.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: NetworkFirewallPolicy
						"network_firewall_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: FirewallDeploymentModel
								"firewall_deployment_model": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Firewall deployment mode.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Network firewall policy.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ThirdPartyFirewallPolicy
						"third_party_firewall_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: FirewallDeploymentModel
								"firewall_deployment_model": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Firewall deployment mode.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Third party firewall policy.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Firewall policy option.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Firewall policy type.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Firewall security service policy data.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A policy tag.",
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^([^\\s]*)$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "pattern": "^([^\\s]*)$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
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
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::FMS::Policy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FMS::Policy").WithTerraformTypeName("awscc_fms_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account":                           "ACCOUNT",
		"arn":                               "Arn",
		"cidr_block":                        "CidrBlock",
		"code":                              "Code",
		"delete_all_policy_resources":       "DeleteAllPolicyResources",
		"egress":                            "Egress",
		"exclude_map":                       "ExcludeMap",
		"exclude_resource_tags":             "ExcludeResourceTags",
		"firewall_deployment_model":         "FirewallDeploymentModel",
		"first_entries":                     "FirstEntries",
		"force_remediate_for_first_entries": "ForceRemediateForFirstEntries",
		"force_remediate_for_last_entries":  "ForceRemediateForLastEntries",
		"from":                              "From",
		"icmp_type_code":                    "IcmpTypeCode",
		"include_map":                       "IncludeMap",
		"ipv_6_cidr_block":                  "Ipv6CidrBlock",
		"key":                               "Key",
		"last_entries":                      "LastEntries",
		"managed_service_data":              "ManagedServiceData",
		"network_acl_common_policy":         "NetworkAclCommonPolicy",
		"network_acl_entry_set":             "NetworkAclEntrySet",
		"network_firewall_policy":           "NetworkFirewallPolicy",
		"orgunit":                           "ORGUNIT",
		"policy_description":                "PolicyDescription",
		"policy_id":                         "Id",
		"policy_name":                       "PolicyName",
		"policy_option":                     "PolicyOption",
		"port_range":                        "PortRange",
		"protocol":                          "Protocol",
		"remediation_enabled":               "RemediationEnabled",
		"resource_set_ids":                  "ResourceSetIds",
		"resource_tag_logical_operator":     "ResourceTagLogicalOperator",
		"resource_tags":                     "ResourceTags",
		"resource_type":                     "ResourceType",
		"resource_type_list":                "ResourceTypeList",
		"resources_clean_up":                "ResourcesCleanUp",
		"rule_action":                       "RuleAction",
		"security_service_policy_data":      "SecurityServicePolicyData",
		"tags":                              "Tags",
		"third_party_firewall_policy":       "ThirdPartyFirewallPolicy",
		"to":                                "To",
		"type":                              "Type",
		"value":                             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
