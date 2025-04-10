# Kinesis stream to receive the logs
resource "awscc_kinesis_stream" "example" {
  name                   = "cloudfront-realtime-logs"
  retention_period_hours = 24
  shard_count            = 1
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role for CloudFront to write to Kinesis
data "aws_iam_policy_document" "cloudfront_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["cloudfront.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "cloudfront_kinesis" {
  statement {
    actions = [
      "kinesis:DescribeStreamSummary",
      "kinesis:DescribeStream",
      "kinesis:PutRecord",
      "kinesis:PutRecords"
    ]
    effect    = "Allow"
    resources = [awscc_kinesis_stream.example.arn]
  }
}

resource "awscc_iam_role" "cloudfront_realtime_logs" {
  role_name                   = "cloudfront-realtime-logs"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.cloudfront_assume_role.json))
  policies = [{
    policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.cloudfront_kinesis.json))
    policy_name     = "cloudfront-kinesis-access"
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# CloudFront real-time log config
resource "awscc_cloudfront_realtime_log_config" "example" {
  name          = "example-realtime-logs"
  sampling_rate = 100
  fields = [
    "timestamp",
    "c-ip",
    "time-to-first-byte",
    "sc-status",
    "sc-bytes",
    "cs-method",
    "cs-protocol",
    "cs-host",
    "cs-uri-stem",
    "cs-bytes",
    "x-edge-location",
    "x-edge-request-id",
    "x-host-header",
    "time-taken",
    "cs-protocol-version",
    "c-ip-version",
    "cs-user-agent",
    "cs-referer"
  ]
  end_points = [
    {
      stream_type = "Kinesis"
      kinesis_stream_config = {
        role_arn   = awscc_iam_role.cloudfront_realtime_logs.arn
        stream_arn = awscc_kinesis_stream.example.arn
      }
    }
  ]
}