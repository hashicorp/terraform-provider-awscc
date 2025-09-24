data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

# Create the Kendra IAM role for the index
resource "awscc_iam_role" "kendra_index_role" {
  role_name = "AWSKendraIndexRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "kendra.amazonaws.com"
        }
      }
    ]
  })

  managed_policy_arns = [
    "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess"
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Kendra index
resource "awscc_kendra_index" "example" {
  name        = "example-kendra-index"
  description = "Example Kendra index for FAQ"
  role_arn    = awscc_iam_role.kendra_index_role.arn
  edition     = "DEVELOPER_EDITION"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create S3 bucket for FAQ files
resource "awscc_s3_bucket" "faq_bucket" {
  bucket_name = "kendra-faq-bucket-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Upload a sample FAQ file (text/csv format)
resource "aws_s3_object" "faq_file" {
  bucket       = awscc_s3_bucket.faq_bucket.id
  key          = "sample-faq.txt"
  source       = "sample-faq.txt"
  content_type = "text/csv"

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create the FAQ role
resource "awscc_iam_role" "kendra_faq_role" {
  role_name = "AWSKendraFAQRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "kendra.amazonaws.com"
        }
      }
    ]
  })

  policies = [
    {
      policy_name = "KendraFAQS3Access"
      policy_document = jsonencode({
        Version = "2012-10-17"
        Statement = [
          {
            Effect = "Allow"
            Action = [
              "s3:GetObject"
            ]
            Resource = ["${awscc_s3_bucket.faq_bucket.arn}/*"]
          }
        ]
      })
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Kendra FAQ
resource "awscc_kendra_faq" "example" {
  name        = "example-faq"
  description = "Example FAQ for Kendra index"
  index_id    = awscc_kendra_index.example.id
  role_arn    = awscc_iam_role.kendra_faq_role.arn
  file_format = "CSV"

  s3_path = {
    bucket = awscc_s3_bucket.faq_bucket.id
    key    = aws_s3_object.faq_file.key
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}