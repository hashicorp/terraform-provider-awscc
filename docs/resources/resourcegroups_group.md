
---
page_title: "awscc_resourcegroups_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Schema for ResourceGroups::Group
---

# awscc_resourcegroups_group (Resource)

Schema for ResourceGroups::Group

## Example Usage

### Tag-Based EC2 Resource Group

Creates a resource group that queries EC2 instances tagged with 'Environment:Production', enabling easier management and organization of AWS resources based on tags.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Example resource group that queries EC2 instances with specific tags
resource "awscc_resourcegroups_group" "example" {
  name        = "example-resource-group"
  description = "Example resource group that finds EC2 instances with specific tags"

  resource_query = {
    query = {
      resource_type_filters = ["AWS::EC2::Instance"]
      tag_filters = [{
        key    = "Environment"
        values = ["Production"]
      }]
    }
    type = "TAG_FILTERS_1_0"
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

- `name` (String) The name of the resource group

### Optional

- `configuration` (Attributes List) (see [below for nested schema](#nestedatt--configuration))
- `description` (String) The description of the resource group
- `resource_query` (Attributes) (see [below for nested schema](#nestedatt--resource_query))
- `resources` (List of String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) The Resource Group ARN.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Optional:

- `parameters` (Attributes List) (see [below for nested schema](#nestedatt--configuration--parameters))
- `type` (String)

<a id="nestedatt--configuration--parameters"></a>
### Nested Schema for `configuration.parameters`

Optional:

- `name` (String)
- `values` (List of String)



<a id="nestedatt--resource_query"></a>
### Nested Schema for `resource_query`

Optional:

- `query` (Attributes) (see [below for nested schema](#nestedatt--resource_query--query))
- `type` (String)

<a id="nestedatt--resource_query--query"></a>
### Nested Schema for `resource_query.query`

Optional:

- `resource_type_filters` (List of String)
- `stack_identifier` (String)
- `tag_filters` (Attributes List) (see [below for nested schema](#nestedatt--resource_query--query--tag_filters))

<a id="nestedatt--resource_query--query--tag_filters"></a>
### Nested Schema for `resource_query.query.tag_filters`

Optional:

- `key` (String)
- `values` (List of String)




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
  to = awscc_resourcegroups_group.example
  id = "name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_resourcegroups_group.example "name"
```
