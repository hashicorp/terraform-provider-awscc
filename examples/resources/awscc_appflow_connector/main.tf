# Get current region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Lambda IAM role and policy
data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    effect = "Allow"
    actions = [
      "sts:AssumeRole"
    ]
    principals {
      type = "Service"
      identifiers = [
        "lambda.amazonaws.com"
      ]
    }
  }
}

data "aws_iam_policy_document" "lambda_policy" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = ["arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:*"]
  }
}

resource "awscc_iam_role" "lambda" {
  role_name                   = "appflow-connector-lambda-role"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.lambda_assume_role.json))
  policies = [
    {
      policy_name     = "lambda-logging"
      policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.lambda_policy.json))
    }
  ]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Lambda function
resource "aws_lambda_function" "connector" {
  filename         = "function.zip"
  function_name    = "appflow-connector-example"
  role             = awscc_iam_role.lambda.arn
  handler          = "index.handler"
  runtime          = "python3.9"
  source_code_hash = filebase64sha256("function.zip")

  depends_on = [awscc_iam_role.lambda]
}

# AppFlow Connector
resource "awscc_appflow_connector" "example" {
  connector_label             = "example-connector"
  connector_provisioning_type = "LAMBDA"
  description                 = "Example AppFlow custom connector using Lambda"

  connector_provisioning_config = {
    lambda = {
      lambda_arn = aws_lambda_function.connector.arn
    }
  }
}