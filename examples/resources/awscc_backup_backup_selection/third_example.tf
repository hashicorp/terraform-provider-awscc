resource "awscc_backup_backup_selection" "example" {
  backup_plan_id = awscc_backup_backup_plan.example.id
  backup_selection = {
    iam_role_arn   = data.awscc_iam_role.example.arn
    selection_name = "list_of_tags_assignment"
    list_of_tags = [{
      condition_key   = "backup"
      condition_value = "true"
      condition_type  = "STRINGEQUALS"
    }]
  }
}