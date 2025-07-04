---
page_title: "awscc_ec2_security_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::SecurityGroup
---

# awscc_ec2_security_group (Resource)

Resource Type definition for AWS::EC2::SecurityGroup

## Example Usage

### Basic usage

~>NOTE on Egress rules: By default, AWS creates an ALLOW ALL egress rule when creating a new Security Group inside of a VPC.

```terraform
resource "awscc_ec2_security_group" "example" {
  group_description = "Security group example"
  vpc_id            = awscc_ec2_vpc.selected.id

  tags = [
    {
      key   = "Name"
      value = "Example SG"
    }
  ]
}

resource "awscc_ec2_vpc" "selected" {
  cidr_block = "10.0.0.0/16"
}
```

### Usage with ingress and egress rules defined

```terraform
resource "awscc_ec2_security_group" "allow_tls" {
  group_description = "Allow TLS inbound traffic and all outbound traffic"
  vpc_id            = awscc_ec2_vpc.selected.id

  tags = [
    {
      key   = "Name"
      value = "allow_tls"
    }
  ]
}

resource "awscc_ec2_vpc_cidr_block" "selected" {
  amazon_provided_ipv_6_cidr_block = true
  vpc_id                           = awscc_ec2_vpc.selected.id
}

resource "awscc_ec2_vpc" "selected" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
}


resource "awscc_ec2_security_group_ingress" "allow_tls_ipv4" {
  group_id    = awscc_ec2_security_group.allow_tls.id
  cidr_ip     = awscc_ec2_vpc.selected.cidr_block
  from_port   = 443
  ip_protocol = "tcp"
  to_port     = 443
}

resource "awscc_ec2_security_group_ingress" "allow_tls_ipv6" {
  group_id    = awscc_ec2_security_group.allow_tls.id
  cidr_ipv_6  = awscc_ec2_vpc_cidr_block.selected.ipv_6_cidr_block
  from_port   = 443
  ip_protocol = "tcp"
  to_port     = 443
}

resource "awscc_ec2_security_group_egress" "allow_all_traffic_ipv4" {
  group_id    = awscc_ec2_security_group.allow_tls.id
  cidr_ip     = "0.0.0.0/0"
  ip_protocol = "-1" # semantically equivalent to all ports
}

resource "awscc_ec2_security_group_egress" "allow_all_traffic_ipv6" {
  group_id    = awscc_ec2_security_group.allow_tls.id
  cidr_ipv_6  = "::/0"
  ip_protocol = "-1" # semantically equivalent to all ports
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_description` (String) A description for the security group.

### Optional

- `group_name` (String) The name of the security group.
- `security_group_egress` (Attributes List) [VPC only] The outbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group. (see [below for nested schema](#nestedatt--security_group_egress))
- `security_group_ingress` (Attributes List) The inbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group. (see [below for nested schema](#nestedatt--security_group_ingress))
- `tags` (Attributes List) Any tags assigned to the security group. (see [below for nested schema](#nestedatt--tags))
- `vpc_id` (String) The ID of the VPC for the security group.

### Read-Only

- `group_id` (String) The group ID of the specified security group.
- `id` (String) Uniquely identifies the resource.
- `security_group_id` (String) The group name or group ID depending on whether the SG is created in default or specific VPC

<a id="nestedatt--security_group_egress"></a>
### Nested Schema for `security_group_egress`

Optional:

- `cidr_ip` (String)
- `cidr_ipv_6` (String)
- `description` (String)
- `destination_prefix_list_id` (String)
- `destination_security_group_id` (String)
- `from_port` (Number)
- `ip_protocol` (String)
- `to_port` (Number)


<a id="nestedatt--security_group_ingress"></a>
### Nested Schema for `security_group_ingress`

Optional:

- `cidr_ip` (String)
- `cidr_ipv_6` (String)
- `description` (String)
- `from_port` (Number)
- `ip_protocol` (String)
- `source_prefix_list_id` (String)
- `source_security_group_id` (String)
- `source_security_group_name` (String)
- `source_security_group_owner_id` (String)
- `to_port` (Number)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_ec2_security_group.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_security_group.example "id"
```