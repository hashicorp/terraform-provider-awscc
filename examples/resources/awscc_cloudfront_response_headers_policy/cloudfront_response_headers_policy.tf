resource "awscc_cloudfront_response_headers_policy" "example" {
  response_headers_policy_config = {
    name    = "example-policy"
    comment = "test comment"
    cors_config = {
      access_control_allow_credentials = true

      access_control_allow_headers = {
        items = ["test"]
      }

      access_control_allow_methods = {
        items = ["GET"]
      }

      access_control_allow_origins = {
        items = ["test.example.comtest"]
      }

      origin_override = true
    }
  }
}
