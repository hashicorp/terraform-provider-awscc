data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an S3 bucket for the channel storage
resource "awscc_s3_bucket" "channel_storage" {
  bucket_name = "iot-analytics-channel-storage-${data.aws_caller_identity.current.account_id}"
}

# Create IAM role for IoT Analytics
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["iotanalytics.amazonaws.com"]
    }
  }
}

# Create base execution role for IoT Analytics
resource "aws_iam_role" "iot_analytics_role" {
  name = "IoTAnalyticsChannelRole"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "iotanalytics.amazonaws.com"
        }
      }
    ]
  })
}

# Create policy for both S3 and IoT Analytics access
resource "aws_iam_role_policy" "channel_policy" {
  name = "IoTAnalyticsChannelPolicy"
  role = aws_iam_role.iot_analytics_role.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:GetBucketLocation",
          "s3:GetObject",
          "s3:PutObject",
          "s3:DeleteObject"
        ]
        Resource = [
          awscc_s3_bucket.channel_storage.arn,
          "${awscc_s3_bucket.channel_storage.arn}/*"
        ]
      },
      {
        Effect = "Allow"
        Action = [
          "iotanalytics:*"
        ]
        Resource = "*"
      }
    ]
  })
}

# Create IoT Analytics Channel
resource "awscc_iotanalytics_channel" "example" {
  channel_name = "example_channel_1"
  channel_storage = {
    customer_managed_s3 = {
      bucket     = awscc_s3_bucket.channel_storage.bucket_name
      key_prefix = "data/"
      role_arn   = aws_iam_role.iot_analytics_role.arn
    }
  }
  retention_period = {
    number_of_days = 30
    unlimited      = false
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}