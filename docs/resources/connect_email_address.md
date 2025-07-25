---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_email_address Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Connect::EmailAddress
---

# awscc_connect_email_address (Resource)

Resource Type definition for AWS::Connect::EmailAddress



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `email_address` (String) Email address to be created for this instance
- `instance_arn` (String) The identifier of the Amazon Connect instance.

### Optional

- `description` (String) A description for the email address.
- `display_name` (String) The display name for the email address.
- `tags` (Attributes Set) One or more tags. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `email_address_arn` (String) The identifier of the email address.
- `id` (String) Uniquely identifies the resource.

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
  to = awscc_connect_email_address.example
  id = "email_address_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_connect_email_address.example "email_address_arn"
```
