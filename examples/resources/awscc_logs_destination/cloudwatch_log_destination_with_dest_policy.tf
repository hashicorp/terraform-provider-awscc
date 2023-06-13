data "aws_region" "current" {}

data "aws_caller_identity" "current" {}


resource "awscc_kinesis_stream" "this" {
  name                   = "terraform-kinesis-test"
  retention_period_hours = 48
  shard_count            = 1
  stream_mode_details = {
    stream_mode = "PROVISIONED"
  }
}

data "aws_iam_policy_document" "cloudwatch_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["logs.amazonaws.com"]
    }
  }
}

locals {
  destination_name = "test_destination"
}

resource "awscc_iam_role" "main" {
  role_name                   = "sample_iam_role"
  description                 = "This is a sample IAM role"
  assume_role_policy_document = data.aws_iam_policy_document.cloudwatch_assume_role_policy.json
  managed_policy_arns         = ["arn:aws:iam::aws:policy/AmazonKinesisFullAccess"]
  path                        = "/"
}


data "aws_iam_policy_document" "test_destination_policy" {
  statement {
    effect = "Allow"

    principals {
      type = "AWS"

      identifiers = [
        "123456789012",
      ]
    }

    actions = [
      "logs:PutSubscriptionFilter",
    ]

    resources = [
      "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:destination:${local.destination_name}"
    ]
  }
}


resource "awscc_logs_destination" "this" {
  destination_name   = local.destination_name
  role_arn           = awscc_iam_role.main.arn
  target_arn         = awscc_kinesis_stream.this.arn
  destination_policy = data.aws_iam_policy_document.test_destination_policy.json
}