resource "aws_kms_key" "example" {
  description             = "KMS key for Security Hub Connector example"
  deletion_window_in_days = 7

  tags = {
    Name        = "example-securityhub-connector-key"
    Environment = "example"
  }
}

resource "awscc_securityhub_connector_v2" "example" {
  name        = "example-securityhub-connector"
  description = "Example Security Hub Connector V2 for AWS Config integration"
  
  # Connector type for AWS Config - this is the only supported type currently
  connector_type = "AWS_CONFIG"
  
  # Labels for categorizing the connector (optional)
  labels = {
    Environment = "example"
    Team        = "security"
  }
  
  # KMS key for encryption (optional)
  kms_key_arn = aws_kms_key.example.arn

  tags = {
    Name        = "example-securityhub-connector"
    Environment = "example"
    Purpose     = "Security Hub Config Integration"
  }
}

resource "aws_kms_key_policy" "example" {
  key_id = aws_kms_key.example.id
  policy = jsonencode({
    Id = "example-key-policy"
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
    ]
    Version = "2012-10-17"
  })
}

data "aws_caller_identity" "current" {}

output "connector_id" {
  description = "The unique identifier of the Security Hub Connector"
  value       = awscc_securityhub_connector_v2.example.id
}

output "connector_arn" {
  description = "The ARN of the Security Hub Connector"
  value       = awscc_securityhub_connector_v2.example.connector_arn
}

output "connector_status" {
  description = "The status of the Security Hub Connector"
  value       = awscc_securityhub_connector_v2.example.connector_status
}
