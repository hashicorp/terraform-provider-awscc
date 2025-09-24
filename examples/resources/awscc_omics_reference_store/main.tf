data "aws_caller_identity" "current" {}

resource "awscc_kms_key" "omics_key" {
  description = "KMS key for Omics Reference Store encryption"
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
        Sid    = "Allow Omics Service"
        Effect = "Allow"
        Principal = {
          Service = "omics.amazonaws.com"
        }
        Action = [
          "kms:Decrypt",
          "kms:GenerateDataKey",
          "kms:GenerateDataKeyWithoutPlaintext"
        ]
        Resource = "*"
      }
    ]
  })
}

resource "awscc_omics_reference_store" "example" {
  name        = "example-reference-store"
  description = "Example Omics Reference Store created by AWSCC provider"

  sse_config = {
    key_arn = awscc_kms_key.omics_key.arn
    type    = "KMS"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}