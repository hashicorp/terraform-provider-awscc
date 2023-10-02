resource "awscc_backup_backup_selection" "example" {
  backup_plan_id = awscc_backup_backup_plan.example.id
  backup_selection = {
    iam_role_arn   = data.awscc_iam_role.example.arn
    selection_name = "resource_assignment"

    resources = [
      awscc_rds_db_instance.example.db_instance_arn,
      awscc_s3_bucket.example.arn
    ]
  }
}