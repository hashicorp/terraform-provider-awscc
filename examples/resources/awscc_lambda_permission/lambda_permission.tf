# Creates a Permission to to allow SNS to execute a Lambda function
# This example assumes you have a valid lambdatets.zip 
#  created on the same directory where you are running your terraform file


resource "awscc_lambda_permission" "with_sns" {
  statement_id  = "AllowExecutionFromSNS"
  action        = "lambda:InvokeFunction"
  function_name = awscc_lambda_function.func.function_name
  principal     = "sns.amazonaws.com"
  source_arn    = awscc_sns_topic.default.arn
}

resource "awscc_sns_topic" "default" {
  name = "call-lambda-maybe"
}

resource "awscc_sns_topic_subscription" "lambda" {
  topic_arn = awscc_sns_topic.default.arn
  protocol  = "lambda"
  endpoint  = awscc_lambda_function.func.arn
}

resource "awscc_lambda_function" "func" {
  filename      = "lambdatest.zip"
  function_name = "lambda_called_from_sns"
  role          = awscc_iam_role.default.arn
  handler       = "exports.handler"
  runtime       = "nodejs16.x"
}

resource "awscc_iam_role" "default" {
  name = "iam_for_lambda_with_sns"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}