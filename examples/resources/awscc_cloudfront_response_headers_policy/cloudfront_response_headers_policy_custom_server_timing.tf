resource "awscc_cloudfront_response_headers_policy" "example" {
  response_headers_policy_config = {
    name = "example-headers-policy"

    custom_headers_config = {
      items = [
        {
          header   = "X-Permitted-Cross-Domain-Policies"
          override = true
          value    = "none"
        }
      ]
    }
    server_timing_headers_config = {
      enabled       = true
      sampling_rate = 50
    }
  }
}