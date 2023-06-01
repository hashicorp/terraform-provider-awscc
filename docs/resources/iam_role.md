---
page_title: "awscc_iam_role Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::IAM::Role
---

# awscc_iam_role (Resource)

Resource Type definition for AWS::IAM::Role

## Example Usage

### Basic example
To create an AWS IAM Role with basic details
```terraform
resource "awscc_iam_role" "main" {
  role_name   = "sample_iam_role"
  description = "This is a sample IAM role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
  max_session_duration = 7200
  path                 = "/"
  tags = [
    {
      key   = "Name"
      value = "Sample IAM Role"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
```


### IAM Role with Assume role policy as Data source
To create an AWS IAM role referring Assume role policy Terraform data source
```terraform
data "aws_iam_policy_document" "instance_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}


resource "awscc_iam_role" "main" {
  role_name                   = "sample_iam_role"
  description                 = "This is a sample IAM role"
  assume_role_policy_document = data.aws_iam_policy_document.instance_assume_role_policy.json
  path                        = "/"
  tags = [
    {
      key   = "Name"
      value = "Sample IAM Role"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
```


### IAM Role with Inline Policy
To create an AWS IAM role with inline policy attached to the role
```terraform
data "aws_iam_policy_document" "instance_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}


data "aws_iam_policy_document" "sample_inline_1" {
  statement {
    sid       = "AccessS3"
    actions   = ["s3:ListAllMyBuckets", "s3:ListBucket", "s3:HeadBucket"]
    resources = ["*"]
  }
}


data "aws_iam_policy_document" "sample_inline_2" {
  statement {
    sid       = "AccessEC2"
    actions   = ["ec2:Describe*"]
    resources = ["*"]
  }
}



resource "awscc_iam_role" "main" {
  role_name                   = "sample_iam_role"
  description                 = "This is a sample IAM role"
  assume_role_policy_document = data.aws_iam_policy_document.instance_assume_role_policy.json
  path                        = "/"
  policies = [{
    policy_document = data.aws_iam_policy_document.sample_inline_1.json
    policy_name     = "fist_inline_policy"
    },
    {
      policy_document = data.aws_iam_policy_document.sample_inline_2.json
      policy_name     = "second_inline_policy"
  }]
  tags = [
    {
      key   = "Name"
      value = "Sample IAM Role"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
```


### IAM Role with Managed Policy
To create an AWS IAM role which has a managed policy attached to the role
```terraform
data "aws_iam_policy_document" "instance_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_policy" "policy_one" {
  name = "policy_one"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["ec2:Describe*"]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

resource "aws_iam_policy" "policy_two" {
  name = "policy_two"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["s3:ListAllMyBuckets", "s3:ListBucket", "s3:HeadBucket"]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}


resource "awscc_iam_role" "main" {
  role_name                   = "sample_iam_role"
  description                 = "This is a sample IAM role"
  assume_role_policy_document = data.aws_iam_policy_document.instance_assume_role_policy.json
  managed_policy_arns         = [aws_iam_policy.policy_one.arn, aws_iam_policy.policy_two.arn]
  path                        = "/"
  tags = [
    {
      key   = "Name"
      value = "Sample IAM Role"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
```


### IAM Role with Permission boundary
To create an AWS IAM role which has a Permission boundary policy attached to the role
```terraform
data "aws_iam_policy_document" "instance_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}


resource "aws_iam_policy" "policy_one" {
  name = "policy_one"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["s3:ListAllMyBuckets", "s3:ListBucket", "s3:HeadBucket"]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

resource "aws_iam_policy" "s3_permission_boundary_policy" {
  name = "s3_permission_boundary_policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["s3:Get*", "s3:List"]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}


resource "awscc_iam_role" "main" {
  role_name                   = "sample_iam_role"
  description                 = "This is a sample IAM role"
  assume_role_policy_document = data.aws_iam_policy_document.instance_assume_role_policy.json
  managed_policy_arns         = [aws_iam_policy.policy_one.arn]
  permissions_boundary        = aws_iam_policy.s3_permission_boundary_policy.arn
  path                        = "/"
  tags = [
    {
      key   = "Name"
      value = "Sample IAM Role"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `assume_role_policy_document` (String) The trust policy that is associated with this role.

### Optional

- `description` (String) A description of the role that you provide.
- `managed_policy_arns` (Set of String) A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.
- `max_session_duration` (Number) The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
- `path` (String) The path to the role.
- `permissions_boundary` (String) The ARN of the policy used to set the permissions boundary for the role.
- `policies` (Attributes List) Adds or updates an inline policy document that is embedded in the specified IAM role. (see [below for nested schema](#nestedatt--policies))
- `role_name` (String) A name for the IAM role, up to 64 characters in length.
- `tags` (Attributes List) A list of tags that are attached to the role. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) for the role.
- `id` (String) Uniquely identifies the resource.
- `role_id` (String) The stable and unique string identifying the role.

<a id="nestedatt--policies"></a>
### Nested Schema for `policies`

Required:

- `policy_document` (String) The policy document.
- `policy_name` (String) The friendly name (not ARN) identifying the policy.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_iam_role.example <resource ID>
```