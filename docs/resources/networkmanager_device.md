
---
page_title: "awscc_networkmanager_device Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::NetworkManager::Device type describes a device.
---

# awscc_networkmanager_device (Resource)

The AWS::NetworkManager::Device type describes a device.

## Example Usage

### Network Manager Device Registration

To register a hardware device in AWS Network Manager, configure the device properties including model, vendor, serial number, and geographical location within a global network.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create a global network first
resource "awscc_networkmanager_global_network" "example" {
  description = "Example Global Network for Device"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a network manager device
resource "awscc_networkmanager_device" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  description       = "Example Network Manager Device"
  model             = "example-model"
  serial_number     = "123456789"
  type              = "HARDWARE"
  vendor            = "Example Vendor"

  location = {
    address   = "123 Example Street"
    latitude  = "47.6062"
    longitude = "-122.3321"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `global_network_id` (String) The ID of the global network.

### Optional

- `aws_location` (Attributes) The Amazon Web Services location of the device, if applicable. (see [below for nested schema](#nestedatt--aws_location))
- `description` (String) The description of the device.
- `location` (Attributes) The site location. (see [below for nested schema](#nestedatt--location))
- `model` (String) The device model
- `serial_number` (String) The device serial number.
- `site_id` (String) The site ID.
- `tags` (Attributes Set) The tags for the device. (see [below for nested schema](#nestedatt--tags))
- `type` (String) The device type.
- `vendor` (String) The device vendor.

### Read-Only

- `created_at` (String) The date and time that the device was created.
- `device_arn` (String) The Amazon Resource Name (ARN) of the device.
- `device_id` (String) The ID of the device.
- `id` (String) Uniquely identifies the resource.
- `state` (String) The state of the device.

<a id="nestedatt--aws_location"></a>
### Nested Schema for `aws_location`

Optional:

- `subnet_arn` (String) The Amazon Resource Name (ARN) of the subnet that the device is located in.
- `zone` (String) The Zone that the device is located in. Specify the ID of an Availability Zone, Local Zone, Wavelength Zone, or an Outpost.


<a id="nestedatt--location"></a>
### Nested Schema for `location`

Optional:

- `address` (String) The physical address.
- `latitude` (String) The latitude.
- `longitude` (String) The longitude.


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
  to = awscc_networkmanager_device.example
  id = "global_network_id|device_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_networkmanager_device.example "global_network_id|device_id"
```
