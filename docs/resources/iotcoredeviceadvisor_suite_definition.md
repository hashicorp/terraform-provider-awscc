---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iotcoredeviceadvisor_suite_definition Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  An example resource schema demonstrating some basic constructs and validation rules.
---

# awscc_iotcoredeviceadvisor_suite_definition (Resource)

An example resource schema demonstrating some basic constructs and validation rules.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `suite_definition_configuration` (Attributes) (see [below for nested schema](#nestedatt--suite_definition_configuration))

### Optional

- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `suite_definition_arn` (String) The Amazon Resource name for the suite definition.
- `suite_definition_id` (String) The unique identifier for the suite definition.
- `suite_definition_version` (String) The suite definition version of a test suite.

<a id="nestedatt--suite_definition_configuration"></a>
### Nested Schema for `suite_definition_configuration`

Required:

- `device_permission_role_arn` (String) The device permission role arn of the test suite.
- `root_group` (String) The root group of the test suite.

Optional:

- `devices` (Attributes List) The devices being tested in the test suite (see [below for nested schema](#nestedatt--suite_definition_configuration--devices))
- `intended_for_qualification` (Boolean) Whether the tests are intended for qualification in a suite.
- `suite_definition_name` (String) The Name of the suite definition.

<a id="nestedatt--suite_definition_configuration--devices"></a>
### Nested Schema for `suite_definition_configuration.devices`

Optional:

- `certificate_arn` (String)
- `thing_arn` (String)



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
  to = awscc_iotcoredeviceadvisor_suite_definition.example
  id = "suite_definition_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_iotcoredeviceadvisor_suite_definition.example "suite_definition_id"
```
