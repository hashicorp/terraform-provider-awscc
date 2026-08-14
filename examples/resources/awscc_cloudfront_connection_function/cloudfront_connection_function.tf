resource "awscc_cloudfront_connection_function" "example" {
  name                     = "example-connection-function"
  connection_function_code = base64encode(file("${path.module}/connection-function.js"))

  connection_function_config = {
    comment = "Example CloudFront Connection Function for processing requests"
    runtime = "cloudfront-js-2.0"
  }

  auto_publish = true

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-connection-function"
    }
  ]
}