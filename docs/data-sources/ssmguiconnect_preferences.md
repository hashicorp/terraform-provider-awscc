---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ssmguiconnect_preferences Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SSMGuiConnect::Preferences
---

# awscc_ssmguiconnect_preferences (Data Source)

Data Source schema for AWS::SSMGuiConnect::Preferences



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `account_id` (String) The AWS Account Id that the preference is associated with, used as the unique identifier for this resource.
- `idle_connection` (Attributes List) A map for Idle Connection Preferences (see [below for nested schema](#nestedatt--idle_connection))

<a id="nestedatt--idle_connection"></a>
### Nested Schema for `idle_connection`

Read-Only:

- `alert` (Attributes) (see [below for nested schema](#nestedatt--idle_connection--alert))
- `timeout` (Attributes) (see [below for nested schema](#nestedatt--idle_connection--timeout))

<a id="nestedatt--idle_connection--alert"></a>
### Nested Schema for `idle_connection.alert`

Read-Only:

- `type` (String)
- `value` (Number)


<a id="nestedatt--idle_connection--timeout"></a>
### Nested Schema for `idle_connection.timeout`

Read-Only:

- `type` (String)
- `value` (Number)
