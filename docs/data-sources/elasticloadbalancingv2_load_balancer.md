---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_elasticloadbalancingv2_load_balancer Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ElasticLoadBalancingV2::LoadBalancer
---

# awscc_elasticloadbalancingv2_load_balancer (Data Source)

Data Source schema for AWS::ElasticLoadBalancingV2::LoadBalancer



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `canonical_hosted_zone_id` (String)
- `dns_name` (String)
- `enforce_security_group_inbound_rules_on_private_link_traffic` (String) Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through privatelink.
- `ip_address_type` (String) The IP address type. The possible values are ``ipv4`` (for IPv4 addresses) and ``dualstack`` (for IPv4 and IPv6 addresses). You can?t specify ``dualstack`` for a load balancer with a UDP or TCP_UDP listener.
- `load_balancer_arn` (String)
- `load_balancer_attributes` (Attributes Set) The load balancer attributes. (see [below for nested schema](#nestedatt--load_balancer_attributes))
- `load_balancer_full_name` (String)
- `load_balancer_name` (String)
- `name` (String) The name of the load balancer. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with "internal-".
 If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
- `scheme` (String) The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.
 The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.
 The default is an Internet-facing load balancer.
 You cannot specify a scheme for a Gateway Load Balancer.
- `security_groups` (Set of String) [Application Load Balancers and Network Load Balancers] The IDs of the security groups for the load balancer.
- `subnet_mappings` (Attributes Set) The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.
 [Application Load Balancers] You must specify subnets from at least two Availability Zones. You cannot specify Elastic IP addresses for your subnets.
 [Application Load Balancers on Outposts] You must specify one Outpost subnet.
 [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.
 [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet if you need static IP addresses for your internet-facing load balancer. For internal load balancers, you can specify one private IP address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you can specify one IPv6 address per subnet.
 [Gateway Load Balancers] You can specify subnets from one or more Availability Zones. You cannot specify Elastic IP (see [below for nested schema](#nestedatt--subnet_mappings))
- `subnets` (Set of String) The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.
 [Application Load Balancers] You must specify subnets from at least two Availability Zones.
 [Application Load Balancers on Outposts] You must specify one Outpost subnet.
 [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.
 [Network Load Balancers] You can specify subnets from one or more Availability Zones.
 [Gateway Load Balancers] You can specify subnets from one or more Availability Zones.
- `tags` (Attributes List) The tags to assign to the load balancer. (see [below for nested schema](#nestedatt--tags))
- `type` (String) The type of load balancer. The default is ``application``.

<a id="nestedatt--load_balancer_attributes"></a>
### Nested Schema for `load_balancer_attributes`

Read-Only:

- `key` (String) The name of the attribute.
 The following attributes are supported by all load balancers:
  +   ``deletion_protection.enabled`` - Indicates whether deletion protection is enabled. The value is ``true`` or ``false``. The default is ``false``.
  +   ``load_balancing.cross_zone.enabled`` - Indicates whether cross-zone load balancing is enabled. The possible values are ``true`` and ``false``. The default for Network Load Balancers and Gateway Load Balancers is ``false``. The default for Application Load Balancers is ``true``, and cannot be changed.
  
 The following attributes are supported by both Application Load Balancers and Network Load Balancers:
  +   ``access_logs.s3.enabled`` - Indicates whether access logs are enabled. The value is ``true`` or ``false``. The default is ``false``.
  +   ``access_logs.s3.bucket`` - The name of the S3 bucket for the access logs. This attribute is required if access logs are enabled. The bucket must exist in the same region as the load balancer and h
- `value` (String) The value of the attribute.


<a id="nestedatt--subnet_mappings"></a>
### Nested Schema for `subnet_mappings`

Read-Only:

- `allocation_id` (String) [Network Load Balancers] The allocation ID of the Elastic IP address for an internet-facing load balancer.
- `i_pv_6_address` (String) [Network Load Balancers] The IPv6 address.
- `private_i_pv_4_address` (String) [Network Load Balancers] The private IPv4 address for an internal load balancer.
- `subnet_id` (String) The ID of the subnet.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key of the tag.
- `value` (String) The value of the tag.
