# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create KMS key for datastore encryption
resource "aws_kms_key" "healthimaging" {
  description             = "KMS key for HealthImaging Datastore"
  deletion_window_in_days = 7
  enable_key_rotation     = true

  policy = jsonencode({
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
        Sid    = "Allow HealthImaging Service"
        Effect = "Allow"
        Principal = {
          Service = "medical-imaging.amazonaws.com"
        }
        Action = [
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:ReEncrypt*",
          "kms:GenerateDataKey*",
          "kms:DescribeKey"
        ]
        Resource = "*"
      }
    ]
  })
}

resource "aws_kms_alias" "healthimaging" {
  name          = "alias/healthimaging-datastore"
  target_key_id = aws_kms_key.healthimaging.key_id
}

resource "awscc_healthimaging_datastore" "example" {
  datastore_name = "example-healthimaging-datastore"
  kms_key_arn    = aws_kms_key.healthimaging.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }, {
    key   = "Environment"
    value = "Test"
  }]
}