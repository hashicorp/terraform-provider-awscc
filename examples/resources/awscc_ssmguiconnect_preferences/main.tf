data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create a basic KMS key for testing
resource "awscc_kms_key" "recording_key" {
  description = "KMS key for SSM GUI Connect recording encryption"
  key_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "Enable IAM User Permissions"
        Effect = "Allow"
        Principal = {
          AWS = "*"
        }
        Action   = "kms:*"
        Resource = "*"
        Condition = {
          StringEquals = {
            "kms:CallerAccount" = "${data.aws_caller_identity.current.account_id}"
          }
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create S3 bucket for recordings with integrated public access block and encryption
resource "awscc_s3_bucket" "recordings" {
  bucket_name = "ssm-gui-connect-recordings-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }

  bucket_encryption = {
    server_side_encryption_configuration = [{
      server_side_encryption_by_default = {
        kms_master_key_id = awscc_kms_key.recording_key.id
        sse_algorithm     = "aws:kms"
      }
    }]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ssmguiconnect_preferences" "example" {
  connection_recording_preferences = {
    kms_key_arn = awscc_kms_key.recording_key.id
    recording_destinations = {
      s3_buckets = [
        {
          bucket_name  = awscc_s3_bucket.recordings.id
          bucket_owner = data.aws_caller_identity.current.account_id
        }
      ]
    }
  }
}