# IAM role for Kinesis Firehose
resource "awscc_iam_role" "firehose" {
  role_name = "metrics-firehose-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "firehose.amazonaws.com"
        }
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role for CloudWatch Metric Stream
resource "awscc_iam_role" "metric_stream" {
  role_name = "metric-stream-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "streams.metrics.cloudwatch.amazonaws.com"
        }
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role policy for Firehose
resource "aws_iam_role_policy" "firehose" {
  name = "metrics-firehose-policy"
  role = awscc_iam_role.firehose.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:AbortMultipartUpload",
          "s3:GetBucketLocation",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:ListBucketMultipartUploads",
          "s3:PutObject"
        ]
        Resource = [
          awscc_s3_bucket.example.arn,
          "${awscc_s3_bucket.example.arn}/*"
        ]
      }
    ]
  })
}

# IAM role policy for Metric Stream 
resource "aws_iam_role_policy" "metric_stream" {
  name = "metric-stream-policy"
  role = awscc_iam_role.metric_stream.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "firehose:PutRecord",
          "firehose:PutRecordBatch"
        ]
        Resource = [awscc_kinesisfirehose_delivery_stream.example.arn]
      }
    ]
  })
}

# S3 bucket for Firehose destination
resource "awscc_s3_bucket" "example" {
  bucket_name = "metric-stream-example-bucket"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Kinesis Firehose
resource "awscc_kinesisfirehose_delivery_stream" "example" {
  delivery_stream_name = "metric-stream-example"
  delivery_stream_type = "DirectPut"
  s3_destination_configuration = {
    bucket_arn = awscc_s3_bucket.example.arn
    role_arn   = awscc_iam_role.firehose.arn
    buffering_hints = {
      interval_in_seconds = 60
      size_in_m_bs        = 64
    }
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# CloudWatch Metric Stream
resource "awscc_cloudwatch_metric_stream" "example" {
  name          = "example-metric-stream"
  firehose_arn  = awscc_kinesisfirehose_delivery_stream.example.arn
  role_arn      = awscc_iam_role.metric_stream.arn
  output_format = "json"
  include_filters = [{
    namespace = "AWS/EC2"
    metric_names = [
      "CPUUtilization",
      "NetworkIn",
      "NetworkOut"
    ]
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}