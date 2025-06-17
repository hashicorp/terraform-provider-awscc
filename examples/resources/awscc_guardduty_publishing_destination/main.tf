data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
data "aws_guardduty_detector" "existing" {}

# S3 bucket for findings
resource "awscc_s3_bucket" "findings" {
  bucket_name = "guardduty-findings-${data.aws_region.current.name}-${data.aws_caller_identity.current.account_id}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# KMS key for encrypting findings
resource "awscc_kms_key" "findings" {
  description = "KMS key for GuardDuty findings"
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
        Sid    = "Allow GuardDuty to encrypt findings"
        Effect = "Allow"
        Principal = {
          Service = "guardduty.amazonaws.com"
        }
        Action = [
          "kms:GenerateDataKey",
          "kms:Encrypt"
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

# S3 Bucket policy allowing GuardDuty to write findings
resource "awscc_s3_bucket_policy" "findings" {
  bucket = awscc_s3_bucket.findings.id
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "Allow GuardDuty to write findings"
        Effect = "Allow"
        Principal = {
          Service = "guardduty.amazonaws.com"
        }
        Action = [
          "s3:GetBucketLocation",
          "s3:PutObject"
        ]
        Resource = [
          "arn:aws:s3:::${awscc_s3_bucket.findings.id}",
          "arn:aws:s3:::${awscc_s3_bucket.findings.id}/*"
        ],
        Condition = {
          StringEquals = {
            "aws:SourceArn"     = "arn:aws:guardduty:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:detector/${data.aws_guardduty_detector.existing.id}",
            "aws:SourceAccount" = data.aws_caller_identity.current.account_id
          }
        }
      }
    ]
  })
}

# GuardDuty Publishing Destination
resource "awscc_guardduty_publishing_destination" "example" {
  detector_id      = data.aws_guardduty_detector.existing.id
  destination_type = "S3"
  destination_properties = {
    destination_arn = "arn:aws:s3:::${awscc_s3_bucket.findings.id}"
    kms_key_arn     = awscc_kms_key.findings.arn
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}