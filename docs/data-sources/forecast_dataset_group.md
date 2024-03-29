---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_forecast_dataset_group Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Forecast::DatasetGroup
---

# awscc_forecast_dataset_group (Data Source)

Data Source schema for AWS::Forecast::DatasetGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `dataset_arns` (List of String) An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.
- `dataset_group_arn` (String) The Amazon Resource Name (ARN) of the dataset group to delete.
- `dataset_group_name` (String) A name for the dataset group.
- `domain` (String) The domain associated with the dataset group. When you add a dataset to a dataset group, this value and the value specified for the Domain parameter of the CreateDataset operation must match.
- `tags` (Attributes List) The tags of Application Insights application. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
