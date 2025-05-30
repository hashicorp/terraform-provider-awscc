---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_workspaces_workspaces_pool Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::WorkSpaces::WorkspacesPool
---

# awscc_workspaces_workspaces_pool (Data Source)

Data Source schema for AWS::WorkSpaces::WorkspacesPool



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `application_settings` (Attributes) (see [below for nested schema](#nestedatt--application_settings))
- `bundle_id` (String)
- `capacity` (Attributes) (see [below for nested schema](#nestedatt--capacity))
- `created_at` (String)
- `description` (String)
- `directory_id` (String)
- `pool_arn` (String)
- `pool_id` (String)
- `pool_name` (String)
- `running_mode` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `timeout_settings` (Attributes) (see [below for nested schema](#nestedatt--timeout_settings))

<a id="nestedatt--application_settings"></a>
### Nested Schema for `application_settings`

Read-Only:

- `settings_group` (String)
- `status` (String)


<a id="nestedatt--capacity"></a>
### Nested Schema for `capacity`

Read-Only:

- `desired_user_sessions` (Number)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)


<a id="nestedatt--timeout_settings"></a>
### Nested Schema for `timeout_settings`

Read-Only:

- `disconnect_timeout_in_seconds` (Number)
- `idle_disconnect_timeout_in_seconds` (Number)
- `max_user_duration_in_seconds` (Number)
