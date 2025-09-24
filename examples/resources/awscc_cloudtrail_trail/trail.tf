resource "awscc_cloudtrail_trail" "example" {
  trail_name     = "example"
  is_logging     = true
  s3_bucket_name = awscc_s3_bucket.example.id

  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_s3_bucket_policy" "example" {
  bucket = awscc_s3_bucket.example.id
  policy_document = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid    = "AllowSSLRequestsOnly",
        Effect = "Deny",
        Principal = {
          AWS = "*"
        }
        Action = "s3:*",
        Resource = [
          "${awscc_s3_bucket.example.arn}",
          "${awscc_s3_bucket.example.arn}/*"
        ]
        Condition = {
          Bool = {
            "aws:SecureTransport" = "false"
          }
        }
      },
      {
        Sid    = "AWSBucketPermissionsCheck",
        Effect = "Allow",
        Principal = {
          Service = "cloudtrail.amazonaws.com"
        },
        Action   = ["s3:GetBucketAcl", "s3:ListBucket"],
        Resource = "${awscc_s3_bucket.example.arn}"
      },
      {
        Sid    = "AWSCloudTrailWrite",
        Effect = "Allow",
        Principal = {
          Service = "cloudtrail.amazonaws.com"
        },
        Action   = "s3:PutObject",
        Resource = "${awscc_s3_bucket.example.arn}/AWSLogs/${data.aws_caller_identity.current.account_id}/*"
      }
    ]
  })
}

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-cloudtrail-${data.aws_caller_identity.current.account_id}"
}

data "aws_caller_identity" "current" {}
