---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sagemaker_device_fleet Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::SageMaker::DeviceFleet
---

# awscc_sagemaker_device_fleet (Resource)

Resource schema for AWS::SageMaker::DeviceFleet



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_fleet_name` (String) The name of the edge device fleet
- `output_config` (Attributes) S3 bucket and an ecryption key id (if available) to store outputs for the fleet (see [below for nested schema](#nestedatt--output_config))
- `role_arn` (String) Role associated with the device fleet

### Optional

- `description` (String) Description for the edge device fleet
- `tags` (Attributes List) Associate tags with the resource (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--output_config"></a>
### Nested Schema for `output_config`

Required:

- `s3_output_location` (String) The Amazon Simple Storage (S3) bucket URI

Optional:

- `kms_key_id` (String) The KMS key id used for encryption on the S3 bucket


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_sagemaker_device_fleet.example
  id = "device_fleet_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_sagemaker_device_fleet.example "device_fleet_name"
```
