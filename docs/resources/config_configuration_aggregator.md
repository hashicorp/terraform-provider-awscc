
---
page_title: "awscc_config_configuration_aggregator Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Config::ConfigurationAggregator
---

# awscc_config_configuration_aggregator (Resource)

Resource Type definition for AWS::Config::ConfigurationAggregator

## Example Usage

### Organization-Wide Config Aggregator

Creates an AWS Config Configuration Aggregator with organization-wide aggregation capabilities, including the necessary IAM role with permissions to access organization account information.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Use data source for region
data "aws_region" "current" {}

# IAM role for Config Aggregator
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["config.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "config_aggregator_policy" {
  statement {
    effect = "Allow"
    actions = [
      "organizations:ListAccounts",
      "organizations:DescribeOrganization",
      "organizations:ListAWSServiceAccessForOrganization"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "config_aggregator_role" {
  role_name                   = "AWSConfigAggregatorRole"
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json
  policies = [{
    policy_document = data.aws_iam_policy_document.config_aggregator_policy.json
    policy_name     = "ConfigAggregatorPolicy"
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# AWS Config Configuration Aggregator
resource "awscc_config_configuration_aggregator" "example" {
  configuration_aggregator_name = "example-aggregator"

  organization_aggregation_source = {
    role_arn        = awscc_iam_role.config_aggregator_role.arn
    aws_regions     = [data.aws_region.current.name]
    all_aws_regions = false
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_aggregation_sources` (Attributes List) (see [below for nested schema](#nestedatt--account_aggregation_sources))
- `configuration_aggregator_name` (String) The name of the aggregator.
- `organization_aggregation_source` (Attributes) (see [below for nested schema](#nestedatt--organization_aggregation_source))
- `tags` (Attributes List) The tags for the configuration aggregator. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `configuration_aggregator_arn` (String) The Amazon Resource Name (ARN) of the aggregator.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--account_aggregation_sources"></a>
### Nested Schema for `account_aggregation_sources`

Optional:

- `account_ids` (List of String)
- `all_aws_regions` (Boolean)
- `aws_regions` (List of String)


<a id="nestedatt--organization_aggregation_source"></a>
### Nested Schema for `organization_aggregation_source`

Optional:

- `all_aws_regions` (Boolean)
- `aws_regions` (List of String)
- `role_arn` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_config_configuration_aggregator.example
  id = "configuration_aggregator_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_config_configuration_aggregator.example "configuration_aggregator_name"
```
