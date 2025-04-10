# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create KMS key for FinSpace environment
resource "awscc_kms_key" "finspace" {
  description = "KMS key for FinSpace environment"
  key_usage   = "ENCRYPT_DECRYPT"
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
        Sid    = "Allow FinSpace Service"
        Effect = "Allow"
        Principal = {
          Service = "finspace.amazonaws.com"
        }
        Action = [
          "kms:Decrypt",
          "kms:GenerateDataKey",
          "kms:CreateGrant",
          "kms:RetireGrant",
          "kms:DescribeKey"
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

# Example of FinSpace Environment
resource "awscc_finspace_environment" "example" {
  name        = "example-finspace-env"
  description = "Example FinSpace Environment created with AWSCC provider"
  kms_key_id  = awscc_kms_key.finspace.id

  # Example of superuser parameters
  superuser_parameters = {
    emailAddress = "admin@example.com"
    firstName    = "Admin"
    lastName     = "User"
  }

  # Example tags
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}