data "aws_caller_identity" "current" {}

resource "awscc_ec2_flow_log" "example" {
  log_destination      = awscc_kinesisfirehose_delivery_stream.example.arn
  log_destination_type = "kinesis-data-firehose"
  traffic_type         = "ALL"
  resource_id          = "vpc-07ddade55bee92f5f"
  resource_type        = "VPC"
}

resource "awscc_kinesisfirehose_delivery_stream" "example" {
  delivery_stream_name = "vpc_flow_log"
  s3_destination_configuration = {
    bucket_arn = awscc_s3_bucket.example.arn
    role_arn   = awscc_iam_role.example.arn
  }
}

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-${data.aws_caller_identity.current.account_id}"
  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }
}

resource "awscc_iam_role" "example" {
  role_name = "firehose_flow_log_role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "firehose.amazonaws.com"
        }
      },
    ]
  })
}

resource "awscc_iam_role_policy" "example" {
  policy_name = "example"
  role_name   = awscc_iam_role.example.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "s3:AbortMultipartUpload",
          "s3:GetBucketLocation",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:ListBucketMultipartUploads",
          "s3:PutObject"
        ],
        Resource = [
          "${awscc_s3_bucket.fexample.arn}",
          "${awscc_s3_bucket.fexample.arn}/*"
        ]
      },
      {
        Effect = "Allow",
        Action = [
          "kinesis:DescribeStream",
          "kinesis:GetShardIterator",
          "kinesis:GetRecords",
          "kinesis:ListShards"
        ],
        Resource = "${awscc_kinesisfirehose_delivery_stream.fexample.arn}"
      }
    ]
  })
}
