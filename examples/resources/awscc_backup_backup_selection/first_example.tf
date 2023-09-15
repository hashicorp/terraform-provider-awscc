resource "awscc_backup_backup_vault" "example" {
  backup_vault_name = "example_backup_vault"
}

resource "awscc_backup_backup_plan" "example" {
  backup_plan = {
    backup_plan_name = "example-backup-plan"
    backup_plan_rule = [{
      rule_name           = "exmaple-backup-rule"
      target_backup_vault = awscc_backup_backup_vault.example.backup_vault_name
      lifecycle = {
        delete_after_days = 14
      }
    }]
    advanced_backup_settings = [{
      backup_options = {
        WindowsVSS = "disabled"
      }
      resource_type = "EC2"
    }]
  }
}

resource "awscc_backup_backup_selection" "example3" {
  backup_plan_id = awscc_backup_backup_plan.example.id
  backup_selection = {
    iam_role_arn   = data.awscc_iam_role.exmaple.arn
    selection_name = "all_resources_assignment"
    resources      = ["*"]
  }
}

data "awscc_iam_role" "exmaple" {
  id = "AWSServiceRoleForBackup"
}