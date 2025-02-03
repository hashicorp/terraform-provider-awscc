# Get the current AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create an S3 bucket for IVS recordings
resource "awscc_s3_bucket" "ivs_recordings" {
  bucket_name = "ivs-recordings-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role for IVS Recording
data "aws_iam_policy_document" "ivs_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"

    principals {
      type        = "Service"
      identifiers = ["ivs.amazonaws.com"]
    }
  }
}

# IAM policy for IVS to write to S3
data "aws_iam_policy_document" "ivs_s3_access" {
  statement {
    effect = "Allow"
    actions = [
      "s3:PutObject",
      "s3:GetObject",
      "s3:ListBucket"
    ]
    resources = [
      "${awscc_s3_bucket.ivs_recordings.arn}",
      "${awscc_s3_bucket.ivs_recordings.arn}/*"
    ]
  }
}

# Create IAM role for IVS using AWSCC provider
resource "awscc_iam_role" "ivs_recording_role" {
  role_name                   = "ivs-recording-role"
  assume_role_policy_document = data.aws_iam_policy_document.ivs_assume_role.json
  description                 = "IAM role for IVS recording service"
  max_session_duration        = 3600

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "ivs_recording_policy" {
  policy_name     = "ivs-recording-policy"
  role_name       = awscc_iam_role.ivs_recording_role.role_name
  policy_document = data.aws_iam_policy_document.ivs_s3_access.json
}

# Create IVS Recording Configuration
resource "awscc_ivs_recording_configuration" "example" {
  name = "example-recording-config"

  destination_configuration = {
    s3 = {
      bucket_name = awscc_s3_bucket.ivs_recordings.bucket_name
    }
  }

  recording_reconnect_window_seconds = 60

  thumbnail_configuration = {
    recording_mode          = "INTERVAL"
    target_interval_seconds = 30
    resolution              = "HD"
    storage                 = ["SEQUENTIAL"]
  }

  rendition_configuration = {
    rendition_selection = "ALL"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}