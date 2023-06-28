resource "awscc_rds_db_parameter_group" "this" {
  description = "rds sample db parameter group"
  family      = "mysql5.7"

  parameters = {
    "character_set_server" = "utf8"
    "character_set_client" = "utf8"
  }

  lifecycle {
    create_before_destroy = true
  }
}
