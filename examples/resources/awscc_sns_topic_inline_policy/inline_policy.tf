resource "awscc_sns_topic_inline_policy" "example" {
  topic_arn = aws_sns_topic.example.arn

  policy_document = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "arn:aws:iam::${var.target_account}:root"
        },
        "Action" : [
          "SNS:Publish"
        ]
        "Resource" : "${aws_sns_topic.example.arn}",
      }
    ]
  })
}

resource "aws_sns_topic" "example" {
  name = "example-topic"
}

variable "target_account" {
  type = string
}