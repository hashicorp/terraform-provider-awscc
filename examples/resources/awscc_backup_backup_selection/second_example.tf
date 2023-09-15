resource "awscc_backup_backup_selection" "example2" {
  backup_plan_id = awscc_backup_backup_plan.example.id
  backup_selection = {
    iam_role_arn   = data.awscc_iam_role.exmaple.arn
    selection_name = "resource_assignment"

    resources = [
      awscc_rds_db_instance.example.db_instance_arn,
      awscc_s3_bucket.example.arn
    ]
  }
}