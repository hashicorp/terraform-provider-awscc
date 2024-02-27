data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

resource "awscc_logs_log_group" "my_log_group" {
  log_group_name    = "my-log-group"
  retention_in_days = 7
  kms_key_id        = awscc_kms_key.my_key.arn
}

resource "awscc_kms_key" "my_key" {
  description = "KMS key for my log group"
  key_policy = jsonencode({
    Version = "2012-10-17"
    Id      = "key-default-1"
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
        Sid    = "Allow Service CloudWatchLogGroup"
        Effect = "Allow"
        Principal = {
          Service = "logs.${data.aws_region.current.name}.amazonaws.com"
        }
        Action = [
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:ReEncrypt*",
          "kms:Describe",
          "kms:GenerateDataKey*"
        ]
        Resource = "*",
        Condition = {
          ArnEquals = {
            "kms:EncryptionContext:aws:logs:arn" = "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:my-log-group"
          }
        }
      }
    ]
  })
}