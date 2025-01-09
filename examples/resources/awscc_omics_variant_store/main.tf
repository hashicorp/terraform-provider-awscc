data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create KMS key for variant store encryption
resource "awscc_kms_key" "omics" {
  description = "KMS key for Omics variant store encryption"
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
      }
    ]
  })
  enabled                = true
  enable_key_rotation    = true
  pending_window_in_days = 7
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create variant store
resource "awscc_omics_variant_store" "example" {
  name        = "example_variant_store"
  description = "Example variant store created with AWSCC provider"
  reference = {
    reference_arn = "arn:aws:omics:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:referenceStore/1234567890/reference/1234567890"
  }
  sse_config = {
    type    = "KMS"
    key_arn = awscc_kms_key.omics.arn
  }
  tags = {
    "Modified By" = "AWSCC"
  }
}