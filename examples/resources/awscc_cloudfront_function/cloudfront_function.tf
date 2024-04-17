resource "awscc_cloudfront_function" "example" {
  name          = "example"
  function_code = file("${path.module}/function.js")
  function_config = {
    comment = "example function"
    runtime = "cloudfront-js-2.0"
  }
  auto_publish = true
}