data "aws_iam_policy_document" "firehose_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["firehose.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "firehose_s3" {
  statement {
    effect = "Allow"
    actions = [
      "s3:AbortMultipartUpload",
      "s3:GetBucketLocation",
      "s3:GetObject",
      "s3:ListBucket",
      "s3:ListBucketMultipartUploads",
      "s3:PutObject",
    ]
    resources = [
      aws_s3_bucket.destination.arn,
      "${aws_s3_bucket.destination.arn}/*",
    ]
  }
}

# Create a random suffix for unique naming
resource "random_id" "suffix" {
  byte_length = 4
}

# Create S3 Bucket
resource "aws_s3_bucket" "destination" {
  bucket        = "firehose-destination-${random_id.suffix.hex}"
  force_destroy = true
}

# Create IAM Role for Firehose
resource "awscc_iam_role" "firehose" {
  role_name                   = "firehose-delivery-role-${random_id.suffix.hex}"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.firehose_assume_role.json))
  managed_policy_arns         = []
  path                        = "/"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM Role Policy for Firehose
resource "awscc_iam_role_policy" "firehose" {
  policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.firehose_s3.json))
  policy_name     = "firehose-s3-access"
  role_name       = awscc_iam_role.firehose.role_name
}

# Create the Kinesis Firehose Delivery Stream
resource "awscc_kinesisfirehose_delivery_stream" "example" {
  delivery_stream_name = "example-delivery-stream-${random_id.suffix.hex}"
  delivery_stream_type = "DirectPut"
  s3_destination_configuration = {
    bucket_arn = aws_s3_bucket.destination.arn
    buffering_hints = {
      interval_in_seconds = 60
      size_in_m_bs        = 5
    }
    compression_format  = "UNCOMPRESSED"
    prefix              = "raw/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/"
    error_output_prefix = "errors/"
    role_arn            = awscc_iam_role.firehose.arn
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}