# Generate a secure random password for the database
# Password can include any printable ASCII character except "/", """, or "@"
resource "random_password" "master" {
  length  = 16
  special = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

# Lightsail database (supplemental resource needed for snapshot)
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

# Lightsail database snapshot
resource "awscc_lightsail_database_snapshot" "example" {
  source_database_name   = awscc_lightsail_database.example.relational_database_name
  database_snapshot_name = "example-snapshot"
}

# Output the snapshot name
output "snapshot_name" {
  description = "Name of the created database snapshot"
  value       = awscc_lightsail_database_snapshot.example.database_snapshot_name
}

# Output the snapshot ARN
output "snapshot_arn" {
  description = "ARN of the created database snapshot"
  value       = awscc_lightsail_database_snapshot.example.database_snapshot_arn
}

# Output the source database name
output "source_database" {
  description = "Name of the source database"
  value       = awscc_lightsail_database_snapshot.example.source_database_name
}