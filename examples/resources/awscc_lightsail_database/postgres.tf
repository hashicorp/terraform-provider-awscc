resource "awscc_lightsail_database" "example" {
  master_database_name             = "example"
  master_username                  = "pgadmin"
  relational_database_blueprint_id = "postgres_16"
  relational_database_bundle_id    = "small_2_0"
  relational_database_name         = "example-db"
  availability_zone                = "us-east-1a"
  backup_retention                 = true
  master_user_password             = "T0pS3cr3t"
  preferred_backup_window          = "06:00-06:30"
  preferred_maintenance_window     = "Sat:03:00-Sat:03:30"
  publicly_accessible              = true
}
