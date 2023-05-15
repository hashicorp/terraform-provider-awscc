resource "awscc_rds_db_subnet_group" "this" {
  db_subnet_group_name = "example"
  db_subnet_group_description = "example subnet group"
  subnet_ids = ["subnet-006182af0254ccbc4", "subnet-0c40688dd8ca51435"]
}

resource "awscc_rds_db_instance" "this" {
  allocated_storage       = 10
  db_name                 = "mydb"
  engine                  = "mysql"
  engine_version          = "5.7"
  db_instance_class       = "db.t3.micro"
  master_username         = "foo"
  master_user_password    = "foobarbaz"
  db_parameter_group_name = "default.mysql5.7"
  db_subnet_group_name = awscc_rds_db_subnet_group.this.id
  tags = [{
    key   = "Name"
    value = "this"
  }]
}