resource "awscc_neptune_db_cluster" "default" {
  db_cluster_identifier          = "example"
  backup_retention_period        = 5
  preferred_backup_window        = "07:00-09:00"
  iam_auth_enabled               = true
  db_port                        = 8182
  copy_tags_to_snapshot          = true
  enable_cloudwatch_logs_exports = ["audit"]
  engine_version                 = "1.3.1.0"
  preferred_maintenance_window   = "sun:10:00-sun:10:30"
  storage_encrypted              = true
  # Use ARN from awscc_kms_key to avoid drift due to short id
  # Reference : https://github.com/hashicorp/terraform-provider-awscc/issues/1735
  kms_key_id = awscc_kms_key.example.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
