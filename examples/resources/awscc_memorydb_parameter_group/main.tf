# Create a MemoryDB parameter group
resource "awscc_memorydb_parameter_group" "example" {
  family               = "memorydb_redis7"
  description          = "Example parameter group using AWSCC provider"
  parameter_group_name = "example-memorydb-pg"

  parameters = jsonencode({
    "activedefrag" : "yes",
    "maxmemory-policy" : "volatile-lru"
  })

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    },
    {
      key   = "Environment"
      value = "Example"
    }
  ]
}