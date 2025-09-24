# Required policy for CloudWatch Logs
data "aws_iam_policy_document" "xray_cloudwatch_policy" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "logs:GetLogEvents",
      "logs:PutRetentionPolicy",
      "logs:GetLogGroupFields",
      "logs:GetQueryResults"
    ]
    resources = ["*"]
    principals {
      type        = "Service"
      identifiers = ["xray.amazonaws.com"]
    }
  }
}

# Create the CloudWatch Logs resource policy
resource "aws_cloudwatch_log_resource_policy" "xray" {
  policy_document = data.aws_iam_policy_document.xray_cloudwatch_policy.json
  policy_name     = "xray-spans-policy"
}

# XRay Transaction Search Config
resource "awscc_xray_transaction_search_config" "example" {
  indexing_percentage = 100

  depends_on = [aws_cloudwatch_log_resource_policy.xray]
}