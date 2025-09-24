resource "awscc_iam_role" "example" {
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

data "archive_file" "example" {
  type        = "zip"
  source_file = "index.py"
  output_path = "lambda_function_payload.zip"
}

resource "awscc_s3_bucket" "lambda_assets" {
}

resource "aws_s3_object" "zip" {
  source = data.archive_file.example.output_path
  bucket = awscc_s3_bucket.lambda_assets.id
  key    = "index.zip"
}

resource "awscc_lambda_function" "example" {
  function_name = "example"
  description   = "AWS Lambda function"
  code = {
    s3_bucket = awscc_s3_bucket.lambda_assets.id
    s3_key    = aws_s3_object.zip.key
  }
  package_type  = "Zip"
  handler       = "index.handler"
  runtime       = "python3.10"
  timeout       = "300"
  memory_size   = "128"
  role          = awscc_iam_role.example.arn
  architectures = ["arm64"]
  environment = {
    variables = {
      MY_KEY_1 = "MY_VALUE_1"
      MY_KEY_2 = "MY_VALUE_2"
    }
  }
  ephemeral_storage = {
    size = 512 # Min 512 MB and the Max 10240 MB
  }
}
