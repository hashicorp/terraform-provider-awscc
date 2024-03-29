---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_databrew_ruleset Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::DataBrew::Ruleset
---

# awscc_databrew_ruleset (Data Source)

Data Source schema for AWS::DataBrew::Ruleset



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `description` (String) Description of the Ruleset
- `name` (String) Name of the Ruleset
- `rules` (Attributes List) List of the data quality rules in the ruleset (see [below for nested schema](#nestedatt--rules))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `target_arn` (String) Arn of the target resource (dataset) to apply the ruleset to

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Read-Only:

- `check_expression` (String) Expression with rule conditions
- `column_selectors` (Attributes List) (see [below for nested schema](#nestedatt--rules--column_selectors))
- `disabled` (Boolean) Boolean value to disable/enable a rule
- `name` (String) Name of the rule
- `substitution_map` (Attributes List) (see [below for nested schema](#nestedatt--rules--substitution_map))
- `threshold` (Attributes) (see [below for nested schema](#nestedatt--rules--threshold))

<a id="nestedatt--rules--column_selectors"></a>
### Nested Schema for `rules.column_selectors`

Read-Only:

- `name` (String) The name of a column from a dataset
- `regex` (String) A regular expression for selecting a column from a dataset


<a id="nestedatt--rules--substitution_map"></a>
### Nested Schema for `rules.substitution_map`

Read-Only:

- `value` (String) Value or column name
- `value_reference` (String) Variable name


<a id="nestedatt--rules--threshold"></a>
### Nested Schema for `rules.threshold`

Read-Only:

- `type` (String) Threshold type for a rule
- `unit` (String) Threshold unit for a rule
- `value` (Number) Threshold value for a rule



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
