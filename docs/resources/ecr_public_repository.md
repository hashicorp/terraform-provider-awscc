---
page_title: "awscc_ecr_public_repository Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::ECR::PublicRepository resource specifies an Amazon Elastic Container Registry Public (Amazon ECR Public) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see Amazon ECR public repositories https://docs.aws.amazon.com/AmazonECR/latest/public/public-repositories.html in the Amazon ECR Public User Guide.
---

# awscc_ecr_public_repository (Resource)

The ``AWS::ECR::PublicRepository`` resource specifies an Amazon Elastic Container Registry Public (Amazon ECR Public) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see [Amazon ECR public repositories](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repositories.html) in the *Amazon ECR Public User Guide*.

~> **NOTE:** This resource can only be used in the `us-east-1` region.

## Example Usage

### Basic usage
To create a ECR public repository.
```terraform
resource "awscc_ecr_public_repository" "example" {
  repository_name = "example"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
```

### ECR public repository with Catalog Data.
To create a ECR public repository with Catalog Data.
```terraform
resource "awscc_ecr_public_repository" "example_catalog_data" {
  repository_name = "example-catalog-data"
  repository_catalog_data = {
    about_text             = "about text"
    architectures          = ["ARM"]
    operating_systems      = ["Linux"]
    repository_description = "Repository description"
    usage_text             = "Usage text"
  }
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `repository_catalog_data` (Attributes) The details about the repository that are publicly visible in the Amazon ECR Public Gallery. For more information, see [Amazon ECR Public repository catalog data](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repository-catalog-data.html) in the *Amazon ECR Public User Guide*. (see [below for nested schema](#nestedatt--repository_catalog_data))
- `repository_name` (String) The name to use for the public repository. The repository name may be specified on its own (such as ``nginx-web-app``) or it can be prepended with a namespace to group the repository into a category (such as ``project-a/nginx-web-app``). If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the repository name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
- `repository_policy_text` (String) The JSON repository policy text to apply to the public repository. For more information, see [Amazon ECR Public repository policies](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repository-policies.html) in the *Amazon ECR Public User Guide*.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--repository_catalog_data"></a>
### Nested Schema for `repository_catalog_data`

Optional:

- `about_text` (String) Provide a detailed description of the repository. Identify what is included in the repository, any licensing details, or other relevant information.
- `architectures` (Set of String) Select the system architectures that the images in your repository are compatible with.
- `operating_systems` (Set of String) Select the operating systems that the images in your repository are compatible with.
- `repository_description` (String) The description of the public repository.
- `usage_text` (String) Provide detailed information about how to use the images in the repository. This provides context, support information, and additional usage details for users of the repository.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) One part of a key-value pair that make up a tag. A ``key`` is a general label that acts like a category for more specific tag values.
- `value` (String) A ``value`` acts as a descriptor within a tag category (key).

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_ecr_public_repository.example
  id = "repository_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ecr_public_repository.example "repository_name"
```