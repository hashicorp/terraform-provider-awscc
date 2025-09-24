resource "awscc_sns_topic_inline_policy" "example" {
  topic_arn = awscc_sns_topic.example.topic_arn

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
        "Resource" : awscc_sns_topic.example.topic_arn
      }
    ]
  })
}

resource "awscc_sns_topic" "example" {
  topic_name = "sns-example-topic"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

variable "target_account" {
  type = string
}