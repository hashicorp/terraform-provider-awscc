# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create CloudWatch log group for APS workspace logging
resource "awscc_logs_log_group" "aps_logs" {
  log_group_name = "/aws/aps/workspaces/example"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create KMS key for APS workspace encryption
resource "awscc_kms_key" "aps_key" {
  description = "KMS key for APS Workspace encryption"
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
        Sid    = "Allow APS Service"
        Effect = "Allow"
        Principal = {
          Service = "aps.amazonaws.com"
        }
        Action = [
          "kms:Decrypt",
          "kms:GenerateDataKey"
        ]
        Resource = "*"
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the APS workspace
resource "awscc_aps_workspace" "example" {
  alias = "example-workspace"
  logging_configuration = {
    log_group_arn = awscc_logs_log_group.aps_logs.arn
  }
  kms_key_arn = awscc_kms_key.aps_key.arn
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}