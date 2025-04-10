# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create a KMS key for VoiceID Domain
resource "awscc_kms_key" "voiceid_key" {
  description = "KMS key for VoiceID Domain"
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
        Sid    = "Allow VoiceID Service"
        Effect = "Allow"
        Principal = {
          Service = "voiceid.amazonaws.com"
        }
        Action = [
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:GenerateDataKey"
        ]
        Resource = "*"
      }
    ]
  })
}

# Create VoiceID Domain
resource "awscc_voiceid_domain" "example" {
  name        = "example-voiceid-domain"
  description = "Example VoiceID Domain created with AWSCC provider"

  server_side_encryption_configuration = {
    kms_key_id = awscc_kms_key.voiceid_key.id
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}