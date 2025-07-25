
---
page_title: "awscc_networkmanager_direct_connect_gateway_attachment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  AWS::NetworkManager::DirectConnectGatewayAttachment Resource Type
---

# awscc_networkmanager_direct_connect_gateway_attachment (Resource)

AWS::NetworkManager::DirectConnectGatewayAttachment Resource Type

## Example Usage

### Attach Direct Connect Gateway to Core Network

Creates a Direct Connect gateway attachment to connect an AWS Direct Connect gateway with a Network Manager core network, enabling hybrid connectivity between your on-premises network and AWS Cloud resources.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Get current AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create a core network
resource "awscc_networkmanager_core_network" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  description       = "Example Core Network"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a global network
resource "awscc_networkmanager_global_network" "example" {
  description = "Example Global Network"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a Direct Connect gateway
resource "aws_dx_gateway" "example" {
  amazon_side_asn = 64512
  name            = "example-dx-gateway"
}

# Create the Direct Connect gateway attachment
resource "awscc_networkmanager_direct_connect_gateway_attachment" "example" {
  core_network_id = awscc_networkmanager_core_network.example.id
  direct_connect_gateway_arn = format("arn:aws:directconnect:%s:%s:dx-gateway/%s",
    data.aws_region.current.name,
    data.aws_caller_identity.current.account_id,
  aws_dx_gateway.example.id)
  edge_locations = [data.aws_region.current.name]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `core_network_id` (String) The ID of a core network for the Direct Connect Gateway attachment.
- `direct_connect_gateway_arn` (String) The ARN of the Direct Connect Gateway.
- `edge_locations` (List of String) The Regions where the edges are located.

### Optional

- `proposed_network_function_group_change` (Attributes) The attachment to move from one network function group to another. (see [below for nested schema](#nestedatt--proposed_network_function_group_change))
- `proposed_segment_change` (Attributes) The attachment to move from one segment to another. (see [below for nested schema](#nestedatt--proposed_segment_change))
- `tags` (Attributes Set) Tags for the attachment. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `attachment_id` (String) Id of the attachment.
- `attachment_policy_rule_number` (Number) The policy rule number associated with the attachment.
- `attachment_type` (String) Attachment type.
- `core_network_arn` (String) The ARN of a core network for the Direct Connect Gateway attachment.
- `created_at` (String) Creation time of the attachment.
- `id` (String) Uniquely identifies the resource.
- `network_function_group_name` (String) The name of the network function group attachment.
- `owner_account_id` (String) Owner account of the attachment.
- `resource_arn` (String) The ARN of the Resource.
- `segment_name` (String) The name of the segment attachment..
- `state` (String) State of the attachment.
- `updated_at` (String) Last update time of the attachment.

<a id="nestedatt--proposed_network_function_group_change"></a>
### Nested Schema for `proposed_network_function_group_change`

Optional:

- `attachment_policy_rule_number` (Number) The rule number in the policy document that applies to this change.
- `network_function_group_name` (String) The name of the network function group to change.
- `tags` (Attributes Set) The key-value tags that changed for the network function group. (see [below for nested schema](#nestedatt--proposed_network_function_group_change--tags))

<a id="nestedatt--proposed_network_function_group_change--tags"></a>
### Nested Schema for `proposed_network_function_group_change.tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.



<a id="nestedatt--proposed_segment_change"></a>
### Nested Schema for `proposed_segment_change`

Optional:

- `attachment_policy_rule_number` (Number) The rule number in the policy document that applies to this change.
- `segment_name` (String) The name of the segment to change.
- `tags` (Attributes Set) The key-value tags that changed for the segment. (see [below for nested schema](#nestedatt--proposed_segment_change--tags))

<a id="nestedatt--proposed_segment_change--tags"></a>
### Nested Schema for `proposed_segment_change.tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.



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
  to = awscc_networkmanager_direct_connect_gateway_attachment.example
  id = "attachment_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_networkmanager_direct_connect_gateway_attachment.example "attachment_id"
```
