resource "awscc_cloudfront_function" "example" {
  name          = "example"
  function_code = file("${path.module}/function.js")
  function_config = {
    comment = "example function"
    runtime = "cloudfront-js-2.0"
    key_value_store_associations = [{
      key_value_store_arn = var.key_store_arn
    }]
  }
  auto_publish = true
}

variable "key_store_arn" {
  type        = string
  description = "Key Value store arn"
}