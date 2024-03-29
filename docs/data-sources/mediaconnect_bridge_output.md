---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_mediaconnect_bridge_output Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::MediaConnect::BridgeOutput
---

# awscc_mediaconnect_bridge_output (Data Source)

Data Source schema for AWS::MediaConnect::BridgeOutput



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `bridge_arn` (String) The Amazon Resource Number (ARN) of the bridge.
- `name` (String) The network output name.
- `network_output` (Attributes) The output of the bridge. (see [below for nested schema](#nestedatt--network_output))

<a id="nestedatt--network_output"></a>
### Nested Schema for `network_output`

Read-Only:

- `ip_address` (String) The network output IP Address.
- `network_name` (String) The network output's gateway network name.
- `port` (Number) The network output port.
- `protocol` (String) The network output protocol.
- `ttl` (Number) The network output TTL.
