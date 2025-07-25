---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_local_gateway_route_table Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for Local Gateway Route Table which describes a route table for a local gateway.
---

# awscc_ec2_local_gateway_route_table (Resource)

Resource Type definition for Local Gateway Route Table which describes a route table for a local gateway.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `local_gateway_id` (String) The ID of the local gateway.

### Optional

- `mode` (String) The mode of the local gateway route table.
- `tags` (Attributes Set) The tags for the local gateway route table. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `local_gateway_route_table_arn` (String) The ARN of the local gateway route table.
- `local_gateway_route_table_id` (String) The ID of the local gateway route table.
- `outpost_arn` (String) The ARN of the outpost.
- `owner_id` (String) The owner of the local gateway route table.
- `state` (String) The state of the local gateway route table.

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
  to = awscc_ec2_local_gateway_route_table.example
  id = "local_gateway_route_table_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_local_gateway_route_table.example "local_gateway_route_table_id"
```
