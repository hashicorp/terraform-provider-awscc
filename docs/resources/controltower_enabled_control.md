---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_controltower_enabled_control Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Enables a control on a specified target.
---

# awscc_controltower_enabled_control (Resource)

Enables a control on a specified target.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `control_identifier` (String) Arn of the control.
- `target_identifier` (String) Arn for Organizational unit to which the control needs to be applied

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_controltower_enabled_control.example <resource ID>
```
