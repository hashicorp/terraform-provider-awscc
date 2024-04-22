resource "awscc_cloudtrail_trail" "example" {
  trail_name                    = "example"
  is_logging                    = true
  enable_log_file_validation    = true
  s3_bucket_name                = awscc_s3_bucket.example.id
  s3_key_prefix                 = "prefix"
  include_global_service_events = false
  kms_key_id                    = awscc_kms_key.example.id

  advanced_event_selectors = [{
    name = "Log all S3 objects events except for two S3 buckets"
    field_selectors = [
      {
        field  = "eventCategory"
        equals = ["Data"]
      },
      {
        field  = "resources.type"
        equals = ["AWS::S3::Object"]
      }
    ]
  }]

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
        Resource = "${awscc_s3_bucket.example.arn}/prefix/AWSLogs/${data.aws_caller_identity.current.account_id}/*"
      }
    ]
  })
}

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-cloudtrail-${data.aws_caller_identity.current.account_id}"

  bucket_encryption = {
    server_side_encryption_configuration = [{
      server_side_encryption_by_default = {
        sse_algorithm     = "aws:kms"
        kms_master_key_id = awscc_kms_key.example.arn
      }
    }]
  }
}

resource "awscc_kms_key" "example" {
  description         = "S3 KMS key"
  enable_key_rotation = true
  key_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid    = "Enable IAM User Permissions",
        Effect = "Allow",
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        },
        Action   = "kms:*",
        Resource = "*"
      },
      {
        Sid    = "Allow CloudTrail to encrypt and decrypt trail",
        Effect = "Allow",
        Principal = {
          Service = "cloudtrail.amazonaws.com"
        },
        Action = [
          "kms:GenerateDataKey*",
          "kms:Decrypt"
        ]
        Resource = "*"
      }
    ]
  })
}

data "aws_caller_identity" "current" {}
