---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_rds_db_shard_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::RDS::DBShardGroup resource creates an Amazon Aurora Limitless DB Shard Group.
---

# awscc_rds_db_shard_group (Resource)

The AWS::RDS::DBShardGroup resource creates an Amazon Aurora Limitless DB Shard Group.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `db_cluster_identifier` (String) The name of the primary DB cluster for the DB shard group.
- `max_acu` (Number) The maximum capacity of the DB shard group in Aurora capacity units (ACUs).

### Optional

- `compute_redundancy` (Number) Specifies whether to create standby instances for the DB shard group.
- `db_shard_group_identifier` (String) The name of the DB shard group.
- `min_acu` (Number) The minimum capacity of the DB shard group in Aurora capacity units (ACUs).
- `publicly_accessible` (Boolean) Indicates whether the DB shard group is publicly accessible.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `db_shard_group_resource_id` (String) The Amazon Web Services Region-unique, immutable identifier for the DB shard group.
- `endpoint` (String) The connection endpoint for the DB shard group.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_rds_db_shard_group.example "db_shard_group_identifier"
```
