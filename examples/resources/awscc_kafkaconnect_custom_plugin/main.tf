# Get current AWS region and account information
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create an S3 bucket to store the plugin file
resource "awscc_s3_bucket" "plugin_bucket" {
  bucket_name = "kafka-connect-plugins-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
}

# Create bucket policy for plugin access
data "aws_iam_policy_document" "bucket_policy" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["kafkaconnect.amazonaws.com"]
    }
    actions = [
      "s3:GetObject",
      "s3:GetObjectVersion"
    ]
    resources = [
      "${awscc_s3_bucket.plugin_bucket.arn}/*"
    ]
  }
}

# Attach policy to the bucket
resource "awscc_s3_bucket_policy" "plugin_bucket_policy" {
  bucket          = awscc_s3_bucket.plugin_bucket.id
  policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.bucket_policy.json))
}

# Create the Kafka Connect custom plugin
resource "awscc_kafkaconnect_custom_plugin" "example" {
  name         = "example-kafka-connect-plugin"
  content_type = "JAR"
  description  = "Example Kafka Connect custom plugin"

  location = {
    s3_location = {
      bucket_arn = awscc_s3_bucket.plugin_bucket.arn
      file_key   = "plugins/example-plugin.jar"
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}