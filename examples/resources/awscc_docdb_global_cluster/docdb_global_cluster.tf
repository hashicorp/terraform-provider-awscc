# DocumentDB Global Cluster
resource "awscc_docdb_global_cluster" "example" {
  global_cluster_identifier = "example-docdb-global-cluster"
  engine                   = "docdb"
  engine_version           = "5.0.0"
  storage_encrypted        = true
  deletion_protection      = false

  tags = [
    {
      key   = "Name"
      value = "example-docdb-global-cluster"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
