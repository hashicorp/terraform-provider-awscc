# Example Global Cluster
resource "awscc_rds_global_cluster" "example" {
  global_cluster_identifier = "global-cluster-example"
  engine                    = "aurora-mysql"
  engine_version            = "5.7.mysql_aurora.2.11.2"
  deletion_protection       = false
  storage_encrypted         = true

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}