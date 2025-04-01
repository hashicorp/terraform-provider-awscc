# Get current AWS account ID
data "aws_caller_identity" "current" {}

# S3 bucket for the crawler target
resource "aws_s3_bucket" "crawler_target" {
  bucket = "glue-crawler-example-${data.aws_caller_identity.current.account_id}"
}

# IAM role for the Glue crawler
resource "aws_iam_role" "crawler_role" {
  name = "glue-crawler-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "glue.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Policy to allow Glue to access the S3 bucket
resource "aws_iam_policy" "crawler_s3_policy" {
  name = "glue-crawler-s3-policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:GetObject",
          "s3:PutObject",
          "s3:DeleteObject"
        ]
        Resource = [
          "${aws_s3_bucket.crawler_target.arn}/*"
        ]
      },
      {
        Effect = "Allow"
        Action = [
          "s3:ListBucket"
        ]
        Resource = [
          aws_s3_bucket.crawler_target.arn
        ]
      }
    ]
  })

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Attach the AWS Glue service role policy
resource "aws_iam_role_policy_attachment" "crawler_service_policy" {
  role       = aws_iam_role.crawler_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole"
}

# Attach the S3 bucket policy to the role
resource "aws_iam_role_policy_attachment" "crawler_policy_attachment" {
  role       = aws_iam_role.crawler_role.name
  policy_arn = aws_iam_policy.crawler_s3_policy.arn
}

# Glue crawler
resource "awscc_glue_crawler" "example" {
  name          = "example-crawler"
  role          = aws_iam_role.crawler_role.arn
  database_name = "example_database"
  description   = "Example Glue crawler that crawls S3 data"

  targets = {
    s3_targets = [{
      path = "s3://${aws_s3_bucket.crawler_target.id}/data/"
    }]
  }

  schema_change_policy = {
    update_behavior = "LOG"
    delete_behavior = "LOG"
  }

  recrawl_policy = {
    recrawl_behavior = "CRAWL_NEW_FOLDERS_ONLY"
  }

  tags = jsonencode({
    "Modified By" = "AWSCC"
  })
}