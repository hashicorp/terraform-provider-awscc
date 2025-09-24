resource "awscc_lambda_url" "example_no_auth" {
  target_function_arn = var.function_arn
  auth_type           = "NONE"
}

variable "function_arn" {
  type        = string
  description = "ARN of the lambda function"
}
