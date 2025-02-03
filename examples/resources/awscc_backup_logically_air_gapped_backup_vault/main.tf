# Get current region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# IAM Policy document for backup vault
data "aws_iam_policy_document" "backup_vault_policy" {
  statement {
    effect = "Allow"
    principals {
      type = "AWS"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      ]
    }
    actions = [
      "backup:CopyIntoBackupVault"
    ]
    resources = ["arn:aws:backup:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:backup-vault:*"]
  }
}

# Create the logically air-gapped backup vault
resource "awscc_backup_logically_air_gapped_backup_vault" "example" {
  backup_vault_name  = "example-air-gapped-vault"
  max_retention_days = 365
  min_retention_days = 7
  access_policy      = jsonencode(jsondecode(data.aws_iam_policy_document.backup_vault_policy.json))

  backup_vault_tags = {
    "Environment" = "Production"
    "Modified By" = "AWSCC"
  }

  notifications = {
    backup_vault_events = ["BACKUP_JOB_COMPLETED", "RESTORE_JOB_COMPLETED"]
    sns_topic_arn       = "arn:aws:sns:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:backup-notifications"
  }
}