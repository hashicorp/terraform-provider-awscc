# Get current account ID
data "aws_caller_identity" "current" {}

# Get current region
data "aws_region" "current" {}

# Create a Kinesis stream first
resource "awscc_kinesis_stream" "example" {
  name                   = "example-stream"
  retention_period_hours = 24
  shard_count            = 1

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create policy document
data "aws_iam_policy_document" "kinesis_policy" {
  statement {
    effect = "Allow"
    actions = [
      "kinesis:PutRecord",
      "kinesis:PutRecords"
    ]
    resources = [awscc_kinesis_stream.example.arn]
    principals {
      type = "AWS"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      ]
    }
    condition {
      test     = "StringEquals"
      variable = "aws:SourceAccount"
      values   = [data.aws_caller_identity.current.account_id]
    }
    condition {
      test     = "StringEquals"
      variable = "aws:SourceArn"
      values   = ["arn:aws:kinesis:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:stream/${awscc_kinesis_stream.example.name}"]
    }
  }
}

# Attach the resource policy to the Kinesis stream
resource "awscc_kinesis_resource_policy" "example" {
  resource_arn    = awscc_kinesis_stream.example.arn
  resource_policy = jsonencode(jsondecode(data.aws_iam_policy_document.kinesis_policy.json))
}