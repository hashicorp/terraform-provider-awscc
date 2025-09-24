data "aws_caller_identity" "example" {}

resource "awscc_sqs_queue" "source" {}

resource "awscc_sqs_queue" "target" {}

resource "awscc_iam_role" "example" {
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "pipes.amazonaws.com"
        }
        Condition = {
          StringEquals = {
            "aws:SourceAccount" = data.aws_caller_identity.example.account_id
          }
        }
      },
    ]
  })
  policies = [{
    policy_name = "SQSAccess"
    policy_document = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Action = [
            "sqs:DeleteMessage",
            "sqs:GetQueueAttributes",
            "sqs:ReceiveMessage",
          ],
          Effect = "Allow",
          Resource = [
            awscc_sqs_queue.source.arn
          ]
        },
        {
          Action = [
            "sqs:SendMessage",
          ],
          Effect = "Allow",
          Resource = [
            awscc_sqs_queue.target.arn
          ]
        },
      ]
    })
  }]
}

resource "awscc_pipes_pipe" "example" {
  name     = "example-pipe"
  role_arn = awscc_iam_role.example.arn
  source   = awscc_sqs_queue.source.arn
  target   = awscc_sqs_queue.target.arn
  tags = {
    "Modified by" = "AWSCC"
  }
}