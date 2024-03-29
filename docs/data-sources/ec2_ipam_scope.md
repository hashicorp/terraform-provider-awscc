---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_ipam_scope Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::IPAMScope
---

# awscc_ec2_ipam_scope (Data Source)

Data Source schema for AWS::EC2::IPAMScope



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the IPAM scope.
- `description` (String)
- `ipam_arn` (String) The Amazon Resource Name (ARN) of the IPAM this scope is a part of.
- `ipam_id` (String) The Id of the IPAM this scope is a part of.
- `ipam_scope_id` (String) Id of the IPAM scope.
- `ipam_scope_type` (String) Determines whether this scope contains publicly routable space or space for a private network
- `is_default` (Boolean) Is this one of the default scopes created with the IPAM.
- `pool_count` (Number) The number of pools that currently exist in this scope.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
