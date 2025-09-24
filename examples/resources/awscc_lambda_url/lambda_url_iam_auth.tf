resource "awscc_lambda_url" "example_auth_iam" {
  target_function_arn = var.function_arn
  auth_type           = "AWS_IAM"

  cors = {
    allow_credentials = true
    allow_origins     = ["*"]
    allow_methods     = ["*"]
    allow_headers     = ["date", "keep-alive"]
    expose_headers    = ["keep-alive", "date"]
    max_age           = 86400
  }
}

variable "function_arn" {
  type        = string
  description = "ARN of the lambda function"
}
