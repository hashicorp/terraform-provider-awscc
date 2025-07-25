
---
page_title: "awscc_ssmquicksetup_configuration_manager Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::SSMQuickSetup::ConfigurationManager Resource Type
---

# awscc_ssmquicksetup_configuration_manager (Resource)

Definition of AWS::SSMQuickSetup::ConfigurationManager Resource Type

## Example Usage

### Configure SSM Quick Setup Manager

Creates an SSM Quick Setup Configuration Manager that enables managed instance setup with hourly schedule execution and required IAM permissions for Systems Manager operations.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Data source for SSM policy document
data "aws_iam_policy_document" "ssm_managed_instance" {
  statement {
    effect = "Allow"
    actions = [
      "ssm:DescribeAssociation",
      "ssm:GetDeployablePatchSnapshotForInstance",
      "ssm:GetDocument",
      "ssm:DescribeDocument",
      "ssm:GetManifest",
      "ssm:GetParameter",
      "ssm:GetParameters",
      "ssm:ListAssociations",
      "ssm:ListInstanceAssociations",
      "ssm:PutInventory",
      "ssm:PutComplianceItems",
      "ssm:PutConfigurePackageResult",
      "ssm:UpdateAssociationStatus",
      "ssm:UpdateInstanceAssociationStatus",
      "ssm:UpdateInstanceInformation"
    ]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "ssmmessages:CreateControlChannel",
      "ssmmessages:CreateDataChannel",
      "ssmmessages:OpenControlChannel",
      "ssmmessages:OpenDataChannel"
    ]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "ec2messages:AcknowledgeMessage",
      "ec2messages:DeleteMessage",
      "ec2messages:FailMessage",
      "ec2messages:GetEndpoint",
      "ec2messages:GetMessages",
      "ec2messages:SendReply"
    ]
    resources = ["*"]
  }
}

# SSM Quick Setup Configuration Manager
resource "awscc_ssmquicksetup_configuration_manager" "example" {
  configuration_definitions = [{
    name = "SSM-Managed-Instance"
    parameters = {
      "TargetAccounts"     = jsonencode([data.aws_caller_identity.current.account_id])
      "SsmRole"            = "service-role/AWSSystemsManagerFullAccess"
      "ScheduleExpression" = "rate(1 hour)"
    }
    type = "SSM"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration_definitions` (Attributes List) (see [below for nested schema](#nestedatt--configuration_definitions))

### Optional

- `description` (String)
- `name` (String)
- `tags` (Map of String)

### Read-Only

- `created_at` (String)
- `id` (String) Uniquely identifies the resource.
- `last_modified_at` (String)
- `manager_arn` (String)
- `status_summaries` (Attributes List) (see [below for nested schema](#nestedatt--status_summaries))

<a id="nestedatt--configuration_definitions"></a>
### Nested Schema for `configuration_definitions`

Required:

- `parameters` (Map of String)
- `type` (String)

Optional:

- `id` (String)
- `local_deployment_administration_role_arn` (String)
- `local_deployment_execution_role_name` (String)
- `type_version` (String)


<a id="nestedatt--status_summaries"></a>
### Nested Schema for `status_summaries`

Read-Only:

- `last_updated_at` (String)
- `status` (String)
- `status_details` (Map of String)
- `status_message` (String)
- `status_type` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_ssmquicksetup_configuration_manager.example
  id = "manager_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ssmquicksetup_configuration_manager.example "manager_arn"
```
