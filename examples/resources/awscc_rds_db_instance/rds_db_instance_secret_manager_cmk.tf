resource "aws_kms_key" "this" {
  description = "Example KMS Key"
}

resource "awscc_rds_db_instance" "this" {
  allocated_storage           = 10
  db_name                     = "mydb"
  engine                      = "mysql"
  engine_version              = "5.7"
  db_instance_class           = "db.t3.micro"
  manage_master_user_password = true
  master_username             = "foo"
  master_user_secret = {
    kms_key_id = aws_kms_key.this.key_id
  }
  db_parameter_group_name = "default.mysql5.7"
}
