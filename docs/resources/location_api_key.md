---
page_title: "awscc_location_api_key Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Location::APIKey Resource Type
---

# awscc_location_api_key (Resource)

Definition of AWS::Location::APIKey Resource Type

## Example Usage

### Create an API key resource in your AWS account, which lets you grant actions for Amazon Location resources to the API key bearer.

```terraform
resource "awscc_location_api_key" "example" {
  key_name    = "example_key"
  description = "Example Location API key"
  no_expiry   = true
  restrictions = {
    allow_actions   = ["geo:GetMap*", "geo:GetPlace"]
    allow_resources = ["arn:aws:geo:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:map/ExampleMap*"]
  }
  tags = [{
    key   = "Modified_By"
    value = "AWSCC"
  }]
}

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `key_name` (String)
- `restrictions` (Attributes) (see [below for nested schema](#nestedatt--restrictions))

### Optional

- `description` (String)
- `expire_time` (String) The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)
- `force_delete` (Boolean)
- `force_update` (Boolean)
- `no_expiry` (Boolean)
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `create_time` (String) The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)
- `id` (String) Uniquely identifies the resource.
- `key_arn` (String)
- `update_time` (String) The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)

<a id="nestedatt--restrictions"></a>
### Nested Schema for `restrictions`

Required:

- `allow_actions` (List of String)
- `allow_resources` (List of String)

Optional:

- `allow_referers` (List of String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_location_api_key.example
  id = "key_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_location_api_key.example "key_name"
```