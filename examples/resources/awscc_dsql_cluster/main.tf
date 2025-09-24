# Basic DSQL Cluster
resource "awscc_dsql_cluster" "example" {
  deletion_protection_enabled = false

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}
