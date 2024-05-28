resource "awscc_elasticache_parameter_group" "example" {
  cache_parameter_group_family = "redis2.8"
  description                  = "Example parameter group"
  properties = {
    "activerehashing"     = "yes",
    "min-slaves-to-write" = "2"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
