resource "awscc_rds_db_parameter_group" "this" {
  description = "rds sample db parameter group"
  family      = "mysql5.6"

  parameters = {
    "character_set_server" = "utf8"
    "character_set_client" = "utf8"
  }

  tags = [{
    key   = "Name"
    value = "this"
  }]
}