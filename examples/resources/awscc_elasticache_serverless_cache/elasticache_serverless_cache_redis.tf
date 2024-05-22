resource "awscc_elasticache_serverless_cache" "example" {
  engine                = "redis"
  serverless_cache_name = "example-redis"
  cache_usage_limits = {
    data_storage = {
      maximum = 10
      unit    = "GB"
    }
    ecpu_per_second = {
      maximum = 5000
    }
  }
  daily_snapshot_time      = "09:00"
  description              = "Example Redis cache"
  kms_key_id               = var.kms_key_arn
  major_engine_version     = "7"
  snapshot_retention_limit = 1
  security_group_ids       = var.security_group_ids
  subnet_ids               = var.subnet_ids

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

variable "security_group_ids" {
  type        = list(string)
  description = "List of security group ids to use with your cache"
}

variable "subnet_ids" {
  type        = list(string)
  description = "List of subnet ids to use with your cache"
}

variable "kms_key_arn" {
  type        = string
  description = "KMS key to be used for encryption"
}
