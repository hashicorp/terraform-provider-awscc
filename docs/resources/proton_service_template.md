
---
page_title: "awscc_proton_service_template Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Proton::ServiceTemplate Resource Type
---

# awscc_proton_service_template (Resource)

Definition of AWS::Proton::ServiceTemplate Resource Type

## Example Usage

### Proton Service Template with Customer Managed Pipeline

Creates an AWS Proton service template with customer-managed pipeline provisioning, allowing organizations to define and manage their service infrastructure templates with custom pipeline configurations.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create a service template
resource "awscc_proton_service_template" "example" {
  name         = "example-service-template"
  display_name = "Example Service Template"
  description  = "Example service template created with AWSCC provider"

  # Set pipeline provisioning to CUSTOMER_MANAGED
  pipeline_provisioning = "CUSTOMER_MANAGED"

  # Add tags
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `description` (String) <p>A description of the service template.</p>
- `display_name` (String) <p>The name of the service template as displayed in the developer interface.</p>
- `encryption_key` (String) <p>A customer provided encryption key that's used to encrypt data.</p>
- `name` (String)
- `pipeline_provisioning` (String)
- `tags` (Attributes List) <p>An optional list of metadata items that you can associate with the Proton service template. A tag is a key-value pair.</p>
         <p>For more information, see <a href="https://docs.aws.amazon.com/proton/latest/userguide/resources.html">Proton resources and tagging</a> in the
        <i>Proton User Guide</i>.</p> (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) <p>The Amazon Resource Name (ARN) of the service template.</p>
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) <p>The key of the resource tag.</p>
- `value` (String) <p>The value of the resource tag.</p>

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_proton_service_template.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_proton_service_template.example "arn"
```
