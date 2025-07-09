# Get the current AWS region
data "aws_region" "current" {}

# Get the current AWS account ID
data "aws_caller_identity" "current" {}

# Create a Kendra Index first (required for data source)
resource "awscc_kendra_index" "example" {
  name     = "example-kendra-index"
  role_arn = awscc_iam_role.kendra_index_role.arn
  edition  = "DEVELOPER_EDITION"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM role for Kendra Index
resource "awscc_iam_role" "kendra_index_role" {
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
  description = "IAM role for Kendra Index"
  path        = "/"
  role_name   = "kendra-index-role"

  policies = [
    {
      policy_name = "kendra-index-policy"
      policy_document = jsonencode({
        Version = "2012-10-17"
        Statement = [
          {
            Effect = "Allow"
            Action = [
              "cloudwatch:PutMetricData"
            ]
            Resource = "*"
            Condition = {
              StringEquals = {
                "cloudwatch:namespace" = "AWS/Kendra"
              }
            }
          },
          {
            Effect = "Allow"
            Action = [
              "logs:DescribeLogGroups"
            ]
            Resource = "*"
          },
          {
            Effect = "Allow"
            Action = [
              "logs:CreateLogGroup"
            ]
            Resource = "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/kendra/*"
          },
          {
            Effect = "Allow"
            Action = [
              "logs:DescribeLogStreams",
              "logs:CreateLogStream",
              "logs:PutLogEvents"
            ]
            Resource = "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/kendra/*:log-stream:*"
          }
        ]
      })
      policy_name = "kendra-index-policy"
    }
  ]
}

# Create S3 bucket for data source
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-kendra-datasource-bucket-${data.aws_caller_identity.current.account_id}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM role for Kendra Data Source
resource "awscc_iam_role" "kendra_datasource_role" {
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
  description = "IAM role for Kendra Data Source"
  path        = "/"
  role_name   = "kendra-datasource-role"

  policies = [
    {
      policy_name = "kendra-datasource-policy"
      policy_document = jsonencode({
        Version = "2012-10-17"
        Statement = [
          {
            Effect = "Allow"
            Action = [
              "s3:GetObject",
              "s3:ListBucket"
            ]
            Resource = [
              awscc_s3_bucket.example.arn,
              "${awscc_s3_bucket.example.arn}/*"
            ]
          }
        ]
      })
      policy_name = "kendra-datasource-policy"
    }
  ]
}

# Create Kendra Data Source
resource "awscc_kendra_data_source" "example" {
  index_id = awscc_kendra_index.example.id
  name     = "example-s3-datasource"
  type     = "S3"
  role_arn = awscc_iam_role.kendra_datasource_role.arn

  data_source_configuration = {
    s3_configuration = {
      bucket_name = awscc_s3_bucket.example.id
    }
  }

  schedule = "cron(0 12 * * ? *)" # Run at 12:00 PM (UTC) every day

  description = "Example S3 data source for Kendra"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}