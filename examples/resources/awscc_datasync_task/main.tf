
# Data sources for dynamic values
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create S3 buckets for source and destination
resource "awscc_s3_bucket" "source" {
  bucket_name = lower("datasync-source-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}")
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_s3_bucket" "destination" {
  bucket_name = lower("datasync-dest-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}")
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM Role for DataSync
data "aws_iam_policy_document" "datasync_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["datasync.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "awscc_iam_role" "datasync" {
  role_name                   = "datasync-s3-role"
  assume_role_policy_document = data.aws_iam_policy_document.datasync_assume_role.json
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM policy for DataSync to access S3
data "aws_iam_policy_document" "datasync_s3" {
  statement {
    effect = "Allow"
    actions = [
      "s3:GetBucketLocation",
      "s3:ListBucket",
      "s3:ListBucketMultipartUploads",
      "s3:GetObject",
      "s3:PutObject",
      "s3:DeleteObject",
      "s3:AbortMultipartUpload"
    ]
    resources = [
      "arn:aws:s3:::${awscc_s3_bucket.source.bucket_name}",
      "arn:aws:s3:::${awscc_s3_bucket.source.bucket_name}/*",
      "arn:aws:s3:::${awscc_s3_bucket.destination.bucket_name}",
      "arn:aws:s3:::${awscc_s3_bucket.destination.bucket_name}/*"
    ]
  }
}

resource "awscc_iam_role_policy" "datasync" {
  policy_name     = "datasync-s3-policy"
  role_name       = awscc_iam_role.datasync.role_name
  policy_document = data.aws_iam_policy_document.datasync_s3.json
}

# Create DataSync locations
resource "awscc_datasync_location_s3" "source" {
  depends_on = [awscc_iam_role_policy.datasync]
  s3_bucket_arn = join("", [
    "arn:aws:s3:::",
    awscc_s3_bucket.source.bucket_name
  ])
  s3_config = {
    bucket_access_role_arn = awscc_iam_role.datasync.arn
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_datasync_location_s3" "destination" {
  depends_on = [awscc_iam_role_policy.datasync]
  s3_bucket_arn = join("", [
    "arn:aws:s3:::",
    awscc_s3_bucket.destination.bucket_name
  ])
  s3_config = {
    bucket_access_role_arn = awscc_iam_role.datasync.arn
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create CloudWatch log group
resource "awscc_logs_log_group" "datasync" {
  log_group_name    = "/aws/datasync/task"
  retention_in_days = 7
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the DataSync task
resource "awscc_datasync_task" "example" {
  source_location_arn      = awscc_datasync_location_s3.source.location_arn
  destination_location_arn = awscc_datasync_location_s3.destination.location_arn
  cloudwatch_log_group_arn = awscc_logs_log_group.datasync.arn
  name                     = "s3-to-s3-sync"

  options = {
    verify_mode            = "ONLY_FILES_TRANSFERRED"
    overwrite_mode         = "ALWAYS"
    atime                  = "BEST_EFFORT"
    mtime                  = "PRESERVE"
    uid                    = "NONE"
    gid                    = "NONE"
    preserve_deleted_files = "REMOVE"
    preserve_devices       = "NONE"
    posix_permissions      = "NONE"
    bytes_per_second       = 8388608
  }

  schedule = {
    schedule_expression = "cron(0 12 ? * SUN *)"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
