resource "awscc_backup_backup_selection" "example" {
  backup_plan_id = awscc_backup_backup_plan.example.id
  backup_selection = {
    iam_role_arn   = data.awscc_iam_role.example.arn
    selection_name = "condition_assignment"
    resources      = ["*"]
    conditions = {
      string_equals = [{
        condition_key   = "aws:ResourceTag/Component"
        condition_value = "rds"
      }]
      string_like = [{
        condition_key   = "aws:ResourceTag/Application"
        condition_value = "app*"
      }]
    }
  }
}