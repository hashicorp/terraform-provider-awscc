---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_gamelift_game_session_queue Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::GameLift::GameSessionQueue
---

# awscc_gamelift_game_session_queue (Data Source)

Data Source schema for AWS::GameLift::GameSessionQueue



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.
- `custom_event_data` (String) Information that is added to all events that are related to this game session queue.
- `destinations` (Attributes List) A list of fleets and/or fleet aliases that can be used to fulfill game session placement requests in the queue. (see [below for nested schema](#nestedatt--destinations))
- `filter_configuration` (Attributes) A list of locations where a queue is allowed to place new game sessions. (see [below for nested schema](#nestedatt--filter_configuration))
- `name` (String) A descriptive label that is associated with game session queue. Queue names must be unique within each Region.
- `notification_target` (String) An SNS topic ARN that is set up to receive game session placement notifications.
- `player_latency_policies` (Attributes List) A set of policies that act as a sliding cap on player latency. (see [below for nested schema](#nestedatt--player_latency_policies))
- `priority_configuration` (Attributes) Custom settings to use when prioritizing destinations and locations for game session placements. (see [below for nested schema](#nestedatt--priority_configuration))
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `timeout_in_seconds` (Number) The maximum time, in seconds, that a new game session placement request remains in the queue.

<a id="nestedatt--destinations"></a>
### Nested Schema for `destinations`

Read-Only:

- `destination_arn` (String)


<a id="nestedatt--filter_configuration"></a>
### Nested Schema for `filter_configuration`

Read-Only:

- `allowed_locations` (List of String) A list of locations to allow game session placement in, in the form of AWS Region codes such as us-west-2.


<a id="nestedatt--player_latency_policies"></a>
### Nested Schema for `player_latency_policies`

Read-Only:

- `maximum_individual_player_latency_milliseconds` (Number) The maximum latency value that is allowed for any player, in milliseconds. All policies must have a value set for this property.
- `policy_duration_seconds` (Number) The length of time, in seconds, that the policy is enforced while placing a new game session.


<a id="nestedatt--priority_configuration"></a>
### Nested Schema for `priority_configuration`

Read-Only:

- `location_order` (List of String) The prioritization order to use for fleet locations, when the PriorityOrder property includes LOCATION.
- `priority_order` (List of String) The recommended sequence to use when prioritizing where to place new game sessions.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.
- `value` (String) The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length.
