---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_aps_workspace Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::APS::Workspace
---

# awscc_aps_workspace (Data Source)

Data Source schema for AWS::APS::Workspace



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `alert_manager_definition` (String) The AMP Workspace alert manager definition data
- `alias` (String) AMP Workspace alias.
- `arn` (String) Workspace arn.
- `kms_key_arn` (String) KMS Key ARN used to encrypt and decrypt AMP workspace data.
- `logging_configuration` (Attributes) Logging configuration (see [below for nested schema](#nestedatt--logging_configuration))
- `prometheus_endpoint` (String) AMP Workspace prometheus endpoint
- `query_logging_configuration` (Attributes) Query logging configuration (see [below for nested schema](#nestedatt--query_logging_configuration))
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `workspace_configuration` (Attributes) Workspace configuration (see [below for nested schema](#nestedatt--workspace_configuration))
- `workspace_id` (String) Required to identify a specific APS Workspace.

<a id="nestedatt--logging_configuration"></a>
### Nested Schema for `logging_configuration`

Read-Only:

- `log_group_arn` (String) CloudWatch log group ARN


<a id="nestedatt--query_logging_configuration"></a>
### Nested Schema for `query_logging_configuration`

Read-Only:

- `destinations` (Attributes List) The destinations configuration for query logging (see [below for nested schema](#nestedatt--query_logging_configuration--destinations))

<a id="nestedatt--query_logging_configuration--destinations"></a>
### Nested Schema for `query_logging_configuration.destinations`

Read-Only:

- `cloudwatch_logs` (Attributes) Represents a cloudwatch logs destination for query logging (see [below for nested schema](#nestedatt--query_logging_configuration--destinations--cloudwatch_logs))
- `filters` (Attributes) Filters for logging (see [below for nested schema](#nestedatt--query_logging_configuration--destinations--filters))

<a id="nestedatt--query_logging_configuration--destinations--cloudwatch_logs"></a>
### Nested Schema for `query_logging_configuration.destinations.cloudwatch_logs`

Read-Only:

- `log_group_arn` (String) The ARN of the CloudWatch Logs log group


<a id="nestedatt--query_logging_configuration--destinations--filters"></a>
### Nested Schema for `query_logging_configuration.destinations.filters`

Read-Only:

- `qsp_threshold` (Number) Query logs with QSP above this limit are vended




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.


<a id="nestedatt--workspace_configuration"></a>
### Nested Schema for `workspace_configuration`

Read-Only:

- `limits_per_label_sets` (Attributes Set) An array of label set and associated limits (see [below for nested schema](#nestedatt--workspace_configuration--limits_per_label_sets))
- `retention_period_in_days` (Number) How many days that metrics are retained in the workspace

<a id="nestedatt--workspace_configuration--limits_per_label_sets"></a>
### Nested Schema for `workspace_configuration.limits_per_label_sets`

Read-Only:

- `label_set` (Attributes Set) An array of series labels (see [below for nested schema](#nestedatt--workspace_configuration--limits_per_label_sets--label_set))
- `limits` (Attributes) Limits that can be applied to a label set (see [below for nested schema](#nestedatt--workspace_configuration--limits_per_label_sets--limits))

<a id="nestedatt--workspace_configuration--limits_per_label_sets--label_set"></a>
### Nested Schema for `workspace_configuration.limits_per_label_sets.label_set`

Read-Only:

- `name` (String) Name of the label
- `value` (String) Value of the label


<a id="nestedatt--workspace_configuration--limits_per_label_sets--limits"></a>
### Nested Schema for `workspace_configuration.limits_per_label_sets.limits`

Read-Only:

- `max_series` (Number) The maximum number of active series that can be ingested for this label set
