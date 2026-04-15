resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [
    {
      key   = "Name"
      value = "example-elasticache-vpc"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_ec2_subnet" "example_a" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"

  tags = [
    {
      key   = "Name"
      value = "example-elasticache-subnet-a"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_ec2_subnet" "example_b" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-west-2b"

  tags = [
    {
      key   = "Name"
      value = "example-elasticache-subnet-b"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_elasticache_subnet_group" "example" {
  cache_subnet_group_name = "example-cache-subnet-group"
  description             = "Example ElastiCache subnet group for replication group"
  subnet_ids              = [awscc_ec2_subnet.example_a.id, awscc_ec2_subnet.example_b.id]

  tags = [
    {
      key   = "Name"
      value = "example-cache-subnet-group"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_elasticache_parameter_group" "example" {
  cache_parameter_group_family = "redis7"
  description                  = "Example Redis parameter group for replication group"

  properties = {
    "maxmemory-policy" = "allkeys-lru"
  }

  tags = [
    {
      key   = "Name"
      value = "example-redis-params"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_elasticache_replication_group" "example" {
  replication_group_description = "Example Redis replication group"
  replication_group_id          = "example-redis-rg"

  engine          = "redis"
  engine_version  = "7.0"
  cache_node_type = "cache.t3.micro"
  port            = 6379

  automatic_failover_enabled = true
  multi_az_enabled           = true
  num_node_groups            = 1
  replicas_per_node_group    = 1

  cache_subnet_group_name    = awscc_elasticache_subnet_group.example.cache_subnet_group_name
  cache_parameter_group_name = awscc_elasticache_parameter_group.example.cache_parameter_group_name

  at_rest_encryption_enabled = true
  transit_encryption_enabled = true

  snapshot_retention_limit     = 1
  snapshot_window              = "03:00-05:00"
  preferred_maintenance_window = "sun:05:00-sun:07:00"

  tags = [
    {
      key   = "Name"
      value = "example-redis-replication-group"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
