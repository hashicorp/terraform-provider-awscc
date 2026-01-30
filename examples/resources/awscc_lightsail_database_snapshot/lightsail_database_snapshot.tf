# Password can include any printable ASCII character except "/", """, or "@"
resource "random_password" "master" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

resource "awscc_lightsail_database" "example" {
  master_database_name             = "example"
  master_username                  = "pgadmin"
  relational_database_blueprint_id = "postgres_16"
  relational_database_bundle_id    = "small_2_0"
  relational_database_name         = "example-db"
  availability_zone                = "us-west-2a"
  backup_retention                 = true
  master_user_password             = random_password.master.result
}

resource "awscc_lightsail_database_snapshot" "example" {
  source_database_name   = awscc_lightsail_database.example.relational_database_name
  database_snapshot_name = "example-snapshot"
}