resource "awscc_rds_db_cluster_parameter_group" "this" {
  db_cluster_parameter_group_name = "rds-db-cluster-pg"
  description                     = "RDS DB cluster parameter group"
  family                          = "aurora5.6"

  parameters = {
    character_set_server = "utf8"
    character_set_client = "latin2"
    time_zone            = "US/Eastern"
  }

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
