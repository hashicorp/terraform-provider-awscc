# Example Neptune DB Cluster Parameter Group
resource "awscc_neptune_db_cluster_parameter_group" "example" {
  name        = "example-neptune-cluster-pg"
  family      = "neptune1.2"
  description = "Example Neptune cluster parameter group"

  # Example parameters in JSON format
  parameters = jsonencode({
    "neptune_enable_audit_log" = "1"
    "neptune_query_timeout"    = "120000"
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}