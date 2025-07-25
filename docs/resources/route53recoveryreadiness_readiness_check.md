
---
page_title: "awscc_route53recoveryreadiness_readiness_check Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Aws Route53 Recovery Readiness Check Schema and API specification.
---

# awscc_route53recoveryreadiness_readiness_check (Resource)

Aws Route53 Recovery Readiness Check Schema and API specification.

## Example Usage

### Route53 Recovery Readiness Check with Health Check

Creates a Route53 Recovery Readiness Check that monitors a Route53 health check through a resource set, enabling you to assess the readiness of your resources for disaster recovery scenarios.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create Route53 Health Check
resource "aws_route53_health_check" "example" {
  fqdn              = "example.com"
  port              = 80
  type              = "HTTP"
  resource_path     = "/"
  failure_threshold = "5"
  request_interval  = "30"

  tags = {
    Modified_By = "AWSCC"
  }
}

# Create the Route53 Recovery Readiness Resource Set
resource "awscc_route53recoveryreadiness_resource_set" "example" {
  resource_set_name = "example-resource-set"
  resource_set_type = "AWS::Route53::HealthCheck"
  resources = [{
    resource_arn     = "arn:aws:route53:::healthcheck/${aws_route53_health_check.example.id}"
    readiness_scopes = []
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Route53 Recovery Readiness Check
resource "awscc_route53recoveryreadiness_readiness_check" "example" {
  readiness_check_name = "example-readiness-check"
  resource_set_name    = awscc_route53recoveryreadiness_resource_set.example.resource_set_name

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `readiness_check_name` (String) Name of the ReadinessCheck to create.
- `resource_set_name` (String) The name of the resource set to check.
- `tags` (Attributes List) A collection of tags associated with a resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `readiness_check_arn` (String) The Amazon Resource Name (ARN) of the readiness check.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_route53recoveryreadiness_readiness_check.example
  id = "readiness_check_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_route53recoveryreadiness_readiness_check.example "readiness_check_name"
```
