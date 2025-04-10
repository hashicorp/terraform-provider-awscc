# Get current AWS account ID
data "aws_caller_identity" "current" {}

# KMS key for SES Mail Manager Archive
resource "awscc_kms_key" "ses_archive" {
  description = "KMS key for SES Mail Manager Archive"
  key_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "Enable IAM User Permissions"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        Action   = "kms:*"
        Resource = "*"
      },
      {
        Sid    = "Allow SES Mail Manager to use the key"
        Effect = "Allow"
        Principal = {
          Service = "ses.amazonaws.com"
        }
        Action = [
          "kms:Decrypt",
          "kms:GenerateDataKey*",
          "kms:Encrypt"
        ]
        Resource = "*"
      }
    ]
  })

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# SES Mail Manager Archive
resource "awscc_ses_mail_manager_archive" "example" {
  archive_name = "example-archive"
  kms_key_arn  = awscc_kms_key.ses_archive.arn

  retention = {
    retention_period = "THREE_MONTHS"
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}