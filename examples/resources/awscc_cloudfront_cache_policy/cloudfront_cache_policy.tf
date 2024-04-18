resource "awscc_cloudfront_cache_policy" "example" {

  cache_policy_config = {
    name        = "example-policy"
    comment     = "test comment"
    default_ttl = 50
    max_ttl     = 100
    min_ttl     = 1

    parameters_in_cache_key_and_forwarded_to_origin = {
      enable_accept_encoding_gzip = true
      cookies_config = {
        cookie_behavior = "whitelist"
        cookies         = ["example"]
      }
      headers_config = {
        header_behavior = "whitelist"
        headers         = ["example"]
      }
      query_strings_config = {
        query_string_behavior = "whitelist"
        query_strings         = ["example"]

      }
    }
  }
}