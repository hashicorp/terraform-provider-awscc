# Terraform AWS Cloud Control Provider Resource Behavior

This document describes the behavior of resources implemented by the provider. In particular it describes in detail how [AWS CloudFormation resource types](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-model.html) are defined as [Terraform resources](https://developer.hashicorp.com/terraform/language/resources).

We assume some familiarity with user-facing Terraform concepts like configuration, state, CLI workflow, etc. The Terraform website has documentation on these ideas.

## Resources

### Terraform Resources

Resources describe infrastructure objects such as virtual networks or compute instances.
By declaring a [Terraform resource block](https://developer.hashicorp.com/terraform/language/resources/syntax) and [applying that configuration](https://developer.hashicorp.com/terraform/language/resources/behavior), Terraform manages the lifecycle of the underlying infrastructure object so as to make its settings match the configuration.

Resource blocks declare a resource of a specific type with a specific local name. The local name is used solely to refer to that resource within its own [module](https://developer.hashicorp.com/terraform/language/modules), having no meaning outside the module's scope. For example

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket"
}
```

declares a resource of type `awscc_s3_bucket` with its `bucket_name` argument set to `"example-bucket"`. The resource's module-local name is `example`.

Resources are implemented in [providers](https://developer.hashicorp.com/terraform/language/providers), plugins which interact with infrastructure providers such as AWS.
A resource's [implementation](https://developer.hashicorp.com/terraform/plugin/framework/resources) defines a [schema](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/schemas) which describes the resource's arguments, and methods (including CRUD methods) which define the resource's lifecycle management functionality.

### Terraform Data Sources

[Data sources](https://developer.hashicorp.com/terraform/language/data-sources) are a variant of resource intended to allow Terraform to reference external data. Unlike [managed resources](Terraform-Resources), Terraform does not manage the lifecycle. Data sources are intended to have no side-effects.
For the purposes of this document we consider data sources to be similar to resources with only a Read method. We will call out differences where they are significant.
