resource "awscc_lambda_alias" "example" {
  function_name    = awscc_lambda_function.example.arn
  function_version = "v1"
  name             = "example-alias"
  description      = "example alias"
  routing_config = {
    additional_version_weights = [
      {
        function_version = "v2"
        function_weight  = 0.5
      }
    ]
  }
}