# Data source to get current AWS account ID
data "aws_caller_identity" "current" {}

# KMS Key for Backup Vault encryption
resource "awscc_kms_key" "example" {
  description = "KMS Key for Backup operations"
  key_policy = jsonencode({
    "Version" : "2012-10-17",
    "Id" : "KMS-Key-Policy-For-Backup",
    "Statement" : [
      {
        "Sid" : "Enable IAM User Permissions",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        },
        "Action" : "kms:*",
        "Resource" : "*"
      },
      {
        "Sid" : "Allow Backup Service",
        "Effect" : "Allow",
        "Principal" : {
          "Service" : "backup.amazonaws.com"
        },
        "Action" : [
          "kms:Decrypt",
          "kms:DescribeKey"
        ],
        "Resource" : "*"
      }
    ]
  })
}

# Backup Vault
resource "awscc_backup_backup_vault" "example" {
  backup_vault_name   = "example-backup-vault"
  encryption_key_arn  = awscc_kms_key.example.arn
}

# Backup Tiering Configuration
resource "awscc_backup_tiering_configuration" "example" {
  tiering_configuration_name = "example-tiering-config"
  storage_class              = "COLD"
  resource_selection = [
    {
      resource_type = "S3"
      resource_arn  = "arn:aws:s3:::example-bucket/*"
    }
  ]
}