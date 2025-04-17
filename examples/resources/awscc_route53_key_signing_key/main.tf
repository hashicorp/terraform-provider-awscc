data "aws_caller_identity" "current" {}

# Create a hosted zone
resource "awscc_route53_hosted_zone" "example" {
  name = "myexampledomain123.com"
}

# Create KMS key for DNSSEC signing
resource "awscc_kms_key" "dnssec" {
  description = "KMS key for Route53 DNSSEC"
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
        Sid    = "Allow Route53 DNSSEC Service"
        Effect = "Allow"
        Principal = {
          Service = "dnssec-route53.amazonaws.com"
        }
        Action = [
          "kms:DescribeKey",
          "kms:GetPublicKey",
          "kms:Sign"
        ]
        Resource = "*"
      },
      {
        Sid    = "Allow Route53 DNSSEC Service to CreateGrant"
        Effect = "Allow"
        Principal = {
          Service = "dnssec-route53.amazonaws.com"
        }
        Action   = "kms:CreateGrant"
        Resource = "*"
        Condition = {
          Bool = {
            "kms:GrantIsForAWSResource" = true
          }
        }
      }
    ]
  })
}

# Create key signing key
resource "awscc_route53_key_signing_key" "example" {
  hosted_zone_id             = awscc_route53_hosted_zone.example.id
  key_management_service_arn = awscc_kms_key.dnssec.arn
  name                       = "examplekey"
  status                     = "ACTIVE"
}