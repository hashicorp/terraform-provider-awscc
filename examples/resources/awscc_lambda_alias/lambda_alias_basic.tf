resource "awscc_lambda_alias" "example" {
  function_name    = awscc_lambda_function.example.arn
  function_version = "$LATEST"
  name             = "example-alias"
  description      = "example alias"
}