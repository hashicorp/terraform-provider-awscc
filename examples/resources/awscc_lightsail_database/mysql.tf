resource "awscc_lightsail_database" "example" {
  master_database_name             = "example"
  master_username                  = "admin"
  relational_database_blueprint_id = "mysql_8_0"
  relational_database_bundle_id    = "micro_2_0"
  relational_database_name         = "example-db"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
