resource "awscc_rds_db_parameter_group" "this1" {
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

resource "awscc_rds_db_instance" "this" {
  allocated_storage       = 10
  db_name                 = "mydb"
  engine                  = "mysql"
  engine_version          = "5.7"
  db_instance_class       = "db.t3.micro"
  master_username         = "foo"
  master_user_password    = "foobarbaz"
  db_parameter_group_name = awscc_rds_db_parameter_group.this.id
}
