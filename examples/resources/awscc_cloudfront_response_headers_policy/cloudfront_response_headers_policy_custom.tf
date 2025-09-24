resource "awscc_cloudfront_response_headers_policy" "example" {
  response_headers_policy_config = {
    name = "example-policy"
    custom_headers_config = {
      items = [
        {
          header   = "X-Permitted-Cross-Domain-Policies"
          override = true
          value    = "none"
          }, {
          header   = "X-Test"
          override = true
          value    = "none"
        }
      ]
    }
  }
}