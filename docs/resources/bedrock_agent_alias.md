---
page_title: "awscc_bedrock_agent_alias Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Bedrock::AgentAlias Resource Type
---

# awscc_bedrock_agent_alias (Resource)

Definition of AWS::Bedrock::AgentAlias Resource Type

## Example Usage

```terraform
resource "awscc_bedrock_agent_alias" "example" {
  agent_alias_name = "example"
  agent_id         = var.bedrock_agent_id
  description      = "Example alias for the Bedrock agent"

  tags = {
    "Modified By" = "AWSCC"
  }
}

variable "bedrock_agent_id" {
  type = string
}
```

### Example with alias routing configuration set for a bedrock agent version

```terraform
resource "awscc_bedrock_agent_alias" "example" {
  agent_alias_name = "example"
  agent_id         = var.bedrock_agent_id
  description      = "Example alias for the Bedrock agent"
  routing_configuration = [
    {
      agent_version = var.bedrock_agent_version
    }
  ]

  tags = {
    "Modified By" = "AWSCC"
  }
}

variable "bedrock_agent_id" {
  type = string
}

variable "bedrock_agent_version" {
  type = string
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `agent_alias_name` (String) Name for a resource.
- `agent_id` (String) Identifier for a resource.

### Optional

- `description` (String) Description of the Resource.
- `routing_configuration` (Attributes List) Routing configuration for an Agent alias. (see [below for nested schema](#nestedatt--routing_configuration))
- `tags` (Map of String) A map of tag keys and values

### Read-Only

- `agent_alias_arn` (String) Arn representation of the Agent Alias.
- `agent_alias_history_events` (Attributes List) The list of history events for an alias for an Agent. (see [below for nested schema](#nestedatt--agent_alias_history_events))
- `agent_alias_id` (String) Id for an Agent Alias generated at the server side.
- `agent_alias_status` (String) The statuses an Agent Alias can be in.
- `created_at` (String) Time Stamp.
- `id` (String) Uniquely identifies the resource.
- `updated_at` (String) Time Stamp.

<a id="nestedatt--routing_configuration"></a>
### Nested Schema for `routing_configuration`

Optional:

- `agent_version` (String) Agent Version.


<a id="nestedatt--agent_alias_history_events"></a>
### Nested Schema for `agent_alias_history_events`

Read-Only:

- `end_date` (String) Time Stamp.
- `routing_configuration` (Attributes List) Routing configuration for an Agent alias. (see [below for nested schema](#nestedatt--agent_alias_history_events--routing_configuration))
- `start_date` (String) Time Stamp.

<a id="nestedatt--agent_alias_history_events--routing_configuration"></a>
### Nested Schema for `agent_alias_history_events.routing_configuration`

Read-Only:

- `agent_version` (String) Agent Version.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_bedrock_agent_alias.example
  id = "agent_id|agent_alias_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_bedrock_agent_alias.example "agent_id|agent_alias_id"
```