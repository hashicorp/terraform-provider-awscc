resource "awscc_xray_resource_policy" "example" {
  policy_document = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Sid" : "SNSAccess",
          "Effect" : "Allow",
          "Principal" : {
            "Service" : "sns.amazonaws.com"
          },
          "Action" : [
            "xray:PutTraceSegments",
            "xray:GetSamplingRules",
            "xray:GetSamplingTargets"
          ],
          "Resource" : "*",
          "Condition" : {
            "StringEquals" : {
              "aws:SourceAccount" : data.aws_caller_identity.current.account_id
            },
            "StringLike" : {
              "aws:SourceArn" : "arn:aws:sns:us-east-1:${data.aws_caller_identity.current.account_id}:*"
            }
          }
        }
      ]
    }
  )
  policy_name                 = "example"
  bypass_policy_lockout_check = false

}

data "aws_caller_identity" "current" {}
