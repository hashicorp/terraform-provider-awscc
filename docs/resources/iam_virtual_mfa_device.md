
---
page_title: "awscc_iam_virtual_mfa_device Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::IAM::VirtualMFADevice
---

# awscc_iam_virtual_mfa_device (Resource)

Resource Type definition for AWS::IAM::VirtualMFADevice

## Example Usage

### Virtual MFA Device Creation

Creates a virtual MFA device and assigns it to an IAM user, enhancing account security by enabling multi-factor authentication.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create an IAM user first for testing
resource "awscc_iam_user" "example" {
  user_name = "example-mfa-user"
  path      = "/"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a virtual MFA device
resource "awscc_iam_virtual_mfa_device" "example" {
  virtual_mfa_device_name = "example-mfa-device"
  users                   = [awscc_iam_user.example.user_name]
  path                    = "/"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `users` (List of String)

### Optional

- `path` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `virtual_mfa_device_name` (String)

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `serial_number` (String)

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
  to = awscc_iam_virtual_mfa_device.example
  id = "serial_number"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_iam_virtual_mfa_device.example "serial_number"
```
