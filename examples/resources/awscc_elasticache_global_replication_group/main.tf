# Data sources for AWS region
data "aws_region" "current" {}

# Create a primary replication group in one region
resource "aws_elasticache_replication_group" "primary" {
  replication_group_id = "example-primary"
  description          = "Primary replication group"
  engine               = "redis"
  engine_version       = "6.x"
  node_type            = "cache.t3.micro"
  num_cache_clusters   = 2
  port                 = 6379

  automatic_failover_enabled = true

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create the global replication group
resource "awscc_elasticache_global_replication_group" "example" {
  global_replication_group_id_suffix   = "example-global"
  global_replication_group_description = "Example Global Replication Group"

  engine_version  = "6.x"
  cache_node_type = "cache.t3.micro"
  engine          = "redis"

  automatic_failover_enabled = true
  global_node_group_count    = 1

  members = [
    {
      replication_group_id     = aws_elasticache_replication_group.primary.id
      replication_group_region = data.aws_region.current.name
      role                     = "PRIMARY"
    }
  ]
}