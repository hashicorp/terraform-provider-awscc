
---
page_title: "awscc_imagebuilder_lifecycle_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::ImageBuilder::LifecyclePolicy
---

# awscc_imagebuilder_lifecycle_policy (Resource)

Resource schema for AWS::ImageBuilder::LifecyclePolicy

## Example Usage

### Automated Image Cleanup Policy

Creates an Image Builder lifecycle policy that automatically deletes AMI images and snapshots older than 30 days in development environment while excluding production resources and specific regions.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Get current AWS region
data "aws_region" "current" {}

# IAM role policy document for Image Builder Lifecycle Policy
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"

    principals {
      type        = "Service"
      identifiers = ["imagebuilder.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "lifecycle_policy" {
  statement {
    effect = "Allow"
    actions = [
      "ec2:DeleteSnapshot",
      "ec2:DeregisterImage",
      "ec2:DescribeImages",
      "ec2:DescribeSnapshots",
      "imagebuilder:GetImage"
    ]
    resources = ["*"]
  }
}

# Create IAM role for Image Builder Lifecycle Policy
resource "awscc_iam_role" "lifecycle_policy" {
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json
  description                 = "Role for Image Builder Lifecycle Policy"
  path                        = "/"
  role_name                   = "ImageBuilderLifecyclePolicyRole"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "lifecycle_policy" {
  policy_document = data.aws_iam_policy_document.lifecycle_policy.json
  policy_name     = "ImageBuilderLifecyclePolicyPermissions"
  role_name       = awscc_iam_role.lifecycle_policy.role_name
}

# Create Image Builder Lifecycle Policy
resource "awscc_imagebuilder_lifecycle_policy" "example" {
  name           = "example-lifecycle-policy"
  description    = "Example Image Builder Lifecycle Policy"
  execution_role = awscc_iam_role.lifecycle_policy.arn
  resource_type  = "AMI_IMAGE"
  status         = "ENABLED"

  policy_details = [
    {
      action = {
        type = "DELETE"
        include_resources = {
          amis       = true
          containers = false
          snapshots  = true
        }
      }
      filter = {
        type  = "AGE"
        value = 30
        unit  = "DAYS"
      }
      exclusion_rules = {
        amis = {
          regions = [data.aws_region.current.name]
        }
        tag_map = {
          "Environment" = "Production"
        }
      }
    }
  ]

  resource_selection = {
    tag_map = {
      "Environment" = "Development"
    }
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

- `execution_role` (String) The execution role of the lifecycle policy.
- `name` (String) The name of the lifecycle policy.
- `policy_details` (Attributes List) The policy details of the lifecycle policy. (see [below for nested schema](#nestedatt--policy_details))
- `resource_selection` (Attributes) The resource selection of the lifecycle policy. (see [below for nested schema](#nestedatt--resource_selection))
- `resource_type` (String) The resource type of the lifecycle policy.

### Optional

- `description` (String) The description of the lifecycle policy.
- `status` (String) The status of the lifecycle policy.
- `tags` (Map of String) The tags associated with the lifecycle policy.

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the lifecycle policy.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--policy_details"></a>
### Nested Schema for `policy_details`

Required:

- `action` (Attributes) The action of the policy detail. (see [below for nested schema](#nestedatt--policy_details--action))
- `filter` (Attributes) The filters to apply of the policy detail. (see [below for nested schema](#nestedatt--policy_details--filter))

Optional:

- `exclusion_rules` (Attributes) The exclusion rules to apply of the policy detail. (see [below for nested schema](#nestedatt--policy_details--exclusion_rules))

<a id="nestedatt--policy_details--action"></a>
### Nested Schema for `policy_details.action`

Required:

- `type` (String) The action type of the policy detail.

Optional:

- `include_resources` (Attributes) The included resources of the policy detail. (see [below for nested schema](#nestedatt--policy_details--action--include_resources))

<a id="nestedatt--policy_details--action--include_resources"></a>
### Nested Schema for `policy_details.action.include_resources`

Optional:

- `amis` (Boolean) Use to configure lifecycle actions on AMIs.
- `containers` (Boolean) Use to configure lifecycle actions on containers.
- `snapshots` (Boolean) Use to configure lifecycle actions on snapshots.



<a id="nestedatt--policy_details--filter"></a>
### Nested Schema for `policy_details.filter`

Required:

- `type` (String) The filter type.
- `value` (Number) The filter value.

Optional:

- `retain_at_least` (Number) The minimum number of Image Builder resources to retain.
- `unit` (String) The value's time unit.


<a id="nestedatt--policy_details--exclusion_rules"></a>
### Nested Schema for `policy_details.exclusion_rules`

Optional:

- `amis` (Attributes) The AMI exclusion rules for the policy detail. (see [below for nested schema](#nestedatt--policy_details--exclusion_rules--amis))
- `tag_map` (Map of String) The Image Builder tags to filter on.

<a id="nestedatt--policy_details--exclusion_rules--amis"></a>
### Nested Schema for `policy_details.exclusion_rules.amis`

Optional:

- `is_public` (Boolean) Use to apply lifecycle policy actions on whether the AMI is public.
- `last_launched` (Attributes) Use to apply lifecycle policy actions on AMIs launched before a certain time. (see [below for nested schema](#nestedatt--policy_details--exclusion_rules--amis--last_launched))
- `regions` (List of String) Use to apply lifecycle policy actions on AMIs distributed to a set of regions.
- `shared_accounts` (List of String) Use to apply lifecycle policy actions on AMIs shared with a set of regions.
- `tag_map` (Map of String) The AMIs to select by tag.

<a id="nestedatt--policy_details--exclusion_rules--amis--last_launched"></a>
### Nested Schema for `policy_details.exclusion_rules.amis.last_launched`

Optional:

- `unit` (String) The value's time unit.
- `value` (Number) The last launched value.





<a id="nestedatt--resource_selection"></a>
### Nested Schema for `resource_selection`

Optional:

- `recipes` (Attributes List) The recipes to select. (see [below for nested schema](#nestedatt--resource_selection--recipes))
- `tag_map` (Map of String) The Image Builder resources to select by tag.

<a id="nestedatt--resource_selection--recipes"></a>
### Nested Schema for `resource_selection.recipes`

Optional:

- `name` (String) The recipe name.
- `semantic_version` (String) The recipe version.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_imagebuilder_lifecycle_policy.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_imagebuilder_lifecycle_policy.example "arn"
```
