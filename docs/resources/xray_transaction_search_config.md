
---
page_title: "awscc_xray_transaction_search_config Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  This schema provides construct and validation rules for AWS-XRay TransactionSearchConfig resource parameters.
---

# awscc_xray_transaction_search_config (Resource)

This schema provides construct and validation rules for AWS-XRay TransactionSearchConfig resource parameters.

## Example Usage

### Configure XRay Transaction Search with CloudWatch Integration

Configures AWS X-Ray transaction search with 100% indexing percentage while setting up the necessary CloudWatch Logs permissions to allow X-Ray service to store and process trace data.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Required policy for CloudWatch Logs
data "aws_iam_policy_document" "xray_cloudwatch_policy" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "logs:GetLogEvents",
      "logs:PutRetentionPolicy",
      "logs:GetLogGroupFields",
      "logs:GetQueryResults"
    ]
    resources = ["*"]
    principals {
      type        = "Service"
      identifiers = ["xray.amazonaws.com"]
    }
  }
}

# Create the CloudWatch Logs resource policy
resource "aws_cloudwatch_log_resource_policy" "xray" {
  policy_document = data.aws_iam_policy_document.xray_cloudwatch_policy.json
  policy_name     = "xray-spans-policy"
}

# XRay Transaction Search Config
resource "awscc_xray_transaction_search_config" "example" {
  indexing_percentage = 100

  depends_on = [aws_cloudwatch_log_resource_policy.xray]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `indexing_percentage` (Number) Determines the percentage of traces indexed from CloudWatch Logs to X-Ray

### Read-Only

- `account_id` (String) User account id, used as the primary identifier for the resource
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_xray_transaction_search_config.example
  id = "account_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_xray_transaction_search_config.example "account_id"
```
