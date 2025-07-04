
---
page_title: "awscc_ec2_verified_access_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::EC2::VerifiedAccessGroup resource creates an AWS EC2 Verified Access Group.
---

# awscc_ec2_verified_access_group (Resource)

The AWS::EC2::VerifiedAccessGroup resource creates an AWS EC2 Verified Access Group.

## Example Usage

### AWS Verified Access Group with IAM Policy

Creates an EC2 Verified Access Group with an IAM policy that allows access to Engineering department users, configured with a required Verified Access Instance and IAM Identity Center trust provider.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create Verified Access Instance first (required for the group)
resource "awscc_ec2_verified_access_instance" "example" {
  description = "Example Verified Access Instance"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a Trust Provider
resource "aws_verifiedaccess_trust_provider" "example" {
  policy_reference_name = "ExampleTrustProvider"
  trust_provider_type   = "user"

  tags = {
    Modified_By = "AWSCC"
  }

  user_trust_provider_type = "iam-identity-center"
}

# Associate the trust provider with the instance
resource "aws_verifiedaccess_instance_trust_provider_attachment" "example" {
  verifiedaccess_instance_id       = awscc_ec2_verified_access_instance.example.verified_access_instance_id
  verifiedaccess_trust_provider_id = aws_verifiedaccess_trust_provider.example.id
}

# Create policy document
data "aws_iam_policy_document" "verified_access_policy" {
  statement {
    effect = "Allow"
    condition {
      test     = "StringEquals"
      variable = "aws:PrincipalTag/Department"
      values   = ["Engineering"]
    }
    principals {
      type        = "*"
      identifiers = ["*"]
    }
  }
}

# Create Verified Access Group
resource "awscc_ec2_verified_access_group" "example" {
  verified_access_instance_id = awscc_ec2_verified_access_instance.example.verified_access_instance_id
  description                 = "Example Verified Access Group with Policy"
  policy_enabled              = true
  policy_document             = jsonencode(jsondecode(data.aws_iam_policy_document.verified_access_policy.json))

  depends_on = [aws_verifiedaccess_instance_trust_provider_attachment.example]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `verified_access_instance_id` (String) The ID of the AWS Verified Access instance.

### Optional

- `description` (String) A description for the AWS Verified Access group.
- `policy_document` (String) The AWS Verified Access policy document.
- `policy_enabled` (Boolean) The status of the Verified Access policy.
- `sse_specification` (Attributes) The configuration options for customer provided KMS encryption. (see [below for nested schema](#nestedatt--sse_specification))
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `creation_time` (String) Time this Verified Access Group was created.
- `id` (String) Uniquely identifies the resource.
- `last_updated_time` (String) Time this Verified Access Group was last updated.
- `owner` (String) The AWS account number that owns the group.
- `verified_access_group_arn` (String) The ARN of the Verified Access group.
- `verified_access_group_id` (String) The ID of the AWS Verified Access group.

<a id="nestedatt--sse_specification"></a>
### Nested Schema for `sse_specification`

Optional:

- `customer_managed_key_enabled` (Boolean) Whether to encrypt the policy with the provided key or disable encryption
- `kms_key_arn` (String) KMS Key Arn used to encrypt the group policy


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_ec2_verified_access_group.example
  id = "verified_access_group_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_verified_access_group.example "verified_access_group_id"
```
