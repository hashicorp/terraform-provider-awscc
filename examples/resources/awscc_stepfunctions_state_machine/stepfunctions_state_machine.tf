data "aws_iam_policy_document" "sample_inline_1" {
  statement {
    sid       = "LambdaAccess"
    actions   = ["lambda:*"]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "main" {
  description = "AWS IAM role for a step function"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "states.amazonaws.com"
        }
      },
    ]
  })
  policies = [
    {
      policy_document = data.aws_iam_policy_document.sample_inline_1.json
      policy_name     = "first_inline_policy"
  }]
}

resource "awscc_iam_role" "lambda" {
  description = "AWS IAM role for a lambda function"
  assume_role_policy_document = jsonencode({
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

data "archive_file" "main" {
  type        = "zip"
  source_file = "index.py"
  output_path = "lambda_function_payload.zip"
}

resource "awscc_lambda_function" "main" {
  function_name = "lambda_function_name"
  description   = "AWS Lambda function"
  code = {
    zip_file = data.archive_file.main.output_path
  }
  package_type  = "Zip"
  handler       = "index.lambda_handler"
  runtime       = "python3.10"
  timeout       = "300"
  memory_size   = "128"
  role          = awscc_iam_role.lambda.arn
  architectures = ["arm64"]
}

resource "awscc_stepfunctions_state_machine" "sfn_stepmachine" {
  role_arn           = awscc_iam_role.main.arn
  state_machine_type = "STANDARD"
  definition_string  = <<EOT
    {
      "StartAt": "FirstState",
      "States": {
        "FirstState": {
          "Type": "Task",
          "Resource": "${awscc_lambda_function.main.arn}",
          "End": true
        }
      }
    }
  EOT
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}