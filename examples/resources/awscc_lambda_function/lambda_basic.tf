resource "awscc_iam_role" "main" {
  description = "AWS IAM role for lambda function"
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
  source_file = "main.py"
  output_path = "lambda_function_payload.zip"
}

resource "awscc_lambda_function" "main" {
  function_name = "lambda_function_name"
  description   = "AWS Lambda function"
  code = {
    zip_file = data.archive_file.main.output_path
  }
  package_type  = "Zip"
  handler       = "main.lambda_handler"
  runtime       = "python3.10"
  timeout       = "300"
  memory_size   = "128"
  role          = awscc_iam_role.main.arn
  architectures = ["arm64"]
  environment = {
    variables = {
      MY_KEY_1 = "MY_VALUE_1"
      MY_KEY_2 = "MY_VALUE_2"
    }
  }
}