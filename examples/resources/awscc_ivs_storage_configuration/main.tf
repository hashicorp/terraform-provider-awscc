# Get the current region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create the S3 bucket
resource "awscc_s3_bucket" "ivs_recordings" {
  bucket_name = "ivs-recordings-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
}

# Create S3 bucket policy document
data "aws_iam_policy_document" "ivs_recording_policy" {
  statement {
    principals {
      type        = "Service"
      identifiers = ["ivs.amazonaws.com"]
    }
    actions = [
      "s3:PutObject",
      "s3:GetObject"
    ]
    resources = [
      "${awscc_s3_bucket.ivs_recordings.arn}",
      "${awscc_s3_bucket.ivs_recordings.arn}/*"
    ]
    condition {
      test     = "StringEquals"
      variable = "AWS:SourceArn"
      values   = ["arn:aws:ivs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:recording-configuration/*"]
    }
  }
}

# Create S3 bucket policy
resource "awscc_s3_bucket_policy" "ivs_recordings_policy" {
  bucket          = awscc_s3_bucket.ivs_recordings.id
  policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.ivs_recording_policy.json))
}

# Create IVS storage configuration
resource "awscc_ivs_storage_configuration" "example" {
  name = "example-storage-config"
  s3 = {
    bucket_name = awscc_s3_bucket.ivs_recordings.id
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}