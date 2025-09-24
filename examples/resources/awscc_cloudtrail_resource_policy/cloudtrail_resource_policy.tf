resource "awscc_cloudtrail_resource_policy" "example" {
  resource_arn = awscc_cloudtrail_channel.example.channel_arn
  resource_policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "DeliverEventsThroughChannel",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        },
        "Action" : "cloudtrail-data:PutAuditEvents",
        "Resource" : awscc_cloudtrail_channel.example.channel_arn
      }
    ]
  })
}

data "aws_caller_identity" "current" {}
