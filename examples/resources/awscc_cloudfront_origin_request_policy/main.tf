# CloudFront Origin Request Policy
resource "awscc_cloudfront_origin_request_policy" "example" {
  origin_request_policy_config = {
    name    = "ExampleOriginRequestPolicy"
    comment = "Example Origin Request Policy for AWSCC Provider"

    cookies_config = {
      cookie_behavior = "whitelist"
      cookies         = ["session-id", "user-prefs"]
    }

    headers_config = {
      header_behavior = "whitelist"
      headers = [
        "Accept",
        "Accept-Charset",
        "Origin",
        "Referer"
      ]
    }

    query_strings_config = {
      query_string_behavior = "whitelist"
      query_strings         = ["page", "sort", "lang"]
    }
  }
}

# Output the policy ID
output "origin_request_policy_id" {
  value = awscc_cloudfront_origin_request_policy.example.origin_request_policy_id
}