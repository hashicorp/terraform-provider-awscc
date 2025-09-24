resource "awscc_rds_db_instance" "this" {
  allocated_storage       = 10
  db_name                 = "mydb"
  engine                  = "mysql"
  engine_version          = "5.7"
  db_instance_class       = "db.t3.micro"
  master_username         = "foo"
  master_user_password    = "foobarbaz"
  db_parameter_group_name = "default.mysql5.7"
}