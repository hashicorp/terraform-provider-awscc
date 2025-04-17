# Get the current AWS account ID
data "aws_caller_identity" "current" {}

# KMS key policy data source
data "aws_iam_policy_document" "kms" {
  statement {
    sid    = "Enable IAM User Permissions"
    effect = "Allow"
    principals {
      type = "AWS"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      ]
    }
    actions = [
      "kms:*"
    ]
    resources = ["*"]
  }
}

# Create a KMS key for stream encryption
resource "awscc_kms_key" "stream" {
  description = "KMS key for Kinesis stream encryption"
  key_policy  = data.aws_iam_policy_document.kms.json
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a Kinesis stream
resource "awscc_kinesis_stream" "example" {
  name                   = "example-stream"
  retention_period_hours = 24
  shard_count            = 1

  stream_mode_details = {
    stream_mode = "PROVISIONED"
  }

  stream_encryption = {
    encryption_type = "KMS"
    key_id          = awscc_kms_key.stream.arn
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}