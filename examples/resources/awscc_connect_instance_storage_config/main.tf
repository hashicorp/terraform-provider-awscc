# Get the current AWS account ID
data "aws_caller_identity" "current" {}

# Create a Connect Instance first (required to create a storage config)
resource "awscc_connect_instance" "example" {
  identity_management_type = "CONNECT_MANAGED"
  instance_alias           = "example-connect-instance-${data.aws_caller_identity.current.account_id}"
  attributes = {
    auto_resolve_best_voices = true
    contact_lens             = true
    contact_flow_logs        = true
    early_media              = true
    inbound_calls            = true
    outbound_calls           = true
    use_custom_tts_voices    = true
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an S3 bucket for Connect storage
resource "awscc_s3_bucket" "connect_storage" {
  bucket_name = "connect-storage-example-${data.aws_caller_identity.current.account_id}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Get IAM policy document for S3
data "aws_iam_policy_document" "connect_s3_storage" {
  statement {
    effect = "Allow"
    actions = [
      "s3:PutObject",
      "s3:GetObject",
      "s3:GetObjectVersion",
      "s3:ListBucket"
    ]
    resources = [
      "arn:aws:s3:::${awscc_s3_bucket.connect_storage.bucket_name}",
      "arn:aws:s3:::${awscc_s3_bucket.connect_storage.bucket_name}/*"
    ]
  }
}

# Create storage configuration for Connect instance
resource "awscc_connect_instance_storage_config" "example" {
  instance_arn  = awscc_connect_instance.example.arn
  resource_type = "CHAT_TRANSCRIPTS"
  storage_type  = "S3"

  s3_config = {
    bucket_name   = awscc_s3_bucket.connect_storage.bucket_name
    bucket_prefix = "chat-transcripts"
    encryption_config = {
      encryption_type = "KMS"
      key_id          = "alias/aws/s3"
    }
  }
}