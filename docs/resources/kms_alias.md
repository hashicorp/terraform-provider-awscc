---
page_title: "awscc_kms_alias Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::KMS::Alias resource specifies a display name for a KMS key https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys. You can use an alias to identify a KMS key in the KMS console, in the DescribeKey https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html operation, and in cryptographic operations https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations, such as Decrypt https://docs.aws.amazon.com/kms/latest/APIReference/API_Decrypt.html and GenerateDataKey https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKey.html.
  Adding, deleting, or updating an alias can allow or deny permission to the KMS key. For details, see ABAC for https://docs.aws.amazon.com/kms/latest/developerguide/abac.html in the Developer Guide.
  Using an alias to refer to a KMS key can help you simplify key management. For example, an alias in your code can be associated with different KMS keys in different AWS-Regions. For more information, see Using aliases https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html in the Developer Guide.
  When specifying an alias, observe the following rules.
  Each alias is associated with one KMS key, but multiple aliases can be associated with the same KMS key.The alias and its associated KMS key must be in the same AWS-account and Region.The alias name must be unique in the AWS-account and Region. However, you can create aliases with the same name in different AWS-Regions. For example, you can have an alias/projectKey in multiple Regions, each of which is associated with a KMS key in its Region.Each alias name must begin with alias/ followed by a name, such as alias/exampleKey. The alias name can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). Alias names cannot begin with alias/aws/. That alias name prefix is reserved for  https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk.
  Regions
  KMS CloudFormation resources are available in all AWS-Regions in which KMS and CFN are supported.
---

# awscc_kms_alias (Resource)

The ``AWS::KMS::Alias`` resource specifies a display name for a [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys). You can use an alias to identify a KMS key in the KMS console, in the [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html) operation, and in [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations), such as [Decrypt](https://docs.aws.amazon.com/kms/latest/APIReference/API_Decrypt.html) and [GenerateDataKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKey.html).
  Adding, deleting, or updating an alias can allow or deny permission to the KMS key. For details, see [ABAC for](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the *Developer Guide*.
  Using an alias to refer to a KMS key can help you simplify key management. For example, an alias in your code can be associated with different KMS keys in different AWS-Regions. For more information, see [Using aliases](https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html) in the *Developer Guide*.
 When specifying an alias, observe the following rules.
  +  Each alias is associated with one KMS key, but multiple aliases can be associated with the same KMS key.
  +  The alias and its associated KMS key must be in the same AWS-account and Region.
  +  The alias name must be unique in the AWS-account and Region. However, you can create aliases with the same name in different AWS-Regions. For example, you can have an ``alias/projectKey`` in multiple Regions, each of which is associated with a KMS key in its Region.
  +  Each alias name must begin with ``alias/`` followed by a name, such as ``alias/exampleKey``. The alias name can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). Alias names cannot begin with ``alias/aws/``. That alias name prefix is reserved for [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
  
  *Regions* 
 KMS CloudFormation resources are available in all AWS-Regions in which KMS and CFN are supported.

## Example Usage

### KMS Alias
To use `awscc_kms_alias` with `awscc_kms_key`:

```terraform
resource "awscc_kms_key" "this" {
  key_policy = jsonencode({
    "Version" : "2012-10-17",
    "Id" : "KMS-Key-Policy",
    "Statement" : [
      {
        "Sid" : "Enable IAM User Permissions",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "arn:aws:iam::111122223333:root"
        },
        "Action" : "kms:*",
        "Resource" : "*"
      },
    ],
    }
  )
}

resource "awscc_kms_alias" "this" {
  alias_name    = "alias/example-kms-alias"
  target_key_id = awscc_kms_key.this.key_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `alias_name` (String) Specifies the alias name. This value must begin with ``alias/`` followed by a name, such as ``alias/ExampleAlias``. 
  If you change the value of the ``AliasName`` property, the existing alias is deleted and a new alias is created for the specified KMS key. This change can disrupt applications that use the alias. It can also allow or deny access to a KMS key affected by attribute-based access control (ABAC).
  The alias must be string of 1-256 characters. It can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). The alias name cannot begin with ``alias/aws/``. The ``alias/aws/`` prefix is reserved for [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
- `target_key_id` (String) Associates the alias with the specified [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk). The KMS key must be in the same AWS-account and Region.
 A valid key ID is required. If you supply a null or empty string value, this operation returns an error.
 For help finding the key ID and ARN, see [Finding the key ID and ARN](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html#find-cmk-id-arn) in the *Developer Guide*.
 Specify the key ID or the key ARN of the KMS key.
 For example:
  +  Key ID: ``1234abcd-12ab-34cd-56ef-1234567890ab``
  +  Key ARN: ``arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab``
  
 To get the key ID and key ARN for a KMS key, use [ListKeys](https://docs.aws.amazon.com/kms/latest/APIReference/API_ListKeys.html) or [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html).

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_kms_alias.example
  id = "alias_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_kms_alias.example "alias_name"
```