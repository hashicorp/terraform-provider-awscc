resource "awscc_lightsail_instance" "example" {
  blueprint_id  = "wordpress"
  bundle_id     = "small_3_0"
  instance_name = "example-instance"
}

resource "awscc_lightsail_static_ip" "example" {
  static_ip_name = "example-ip"
  attached_to    = awscc_lightsail_instance.example.id
}


resource "awscc_lightsail_distribution" "example" {
  bundle_id = "small_1_0"
  default_cache_behavior = {
    behavior = "dont-cache"
  }
  distribution_name = "example-dist"
  origin = {
    name        = awscc_lightsail_instance.example.instance_name
    region_name = "us-east-1"
  }
  cache_behaviors = [
    {
      behavior = "cache"
      path     = "wp-content/*"
    },
    {
      behavior = "cache"
      path     = "wp-includes/*"
    }
  ]
  cache_behavior_settings = {
    allowed_http_methods = "GET,HEAD,OPTIONS"
    cached_http_methods  = "GET,HEAD"
    default_ttl          = 86400
    forwarded_cookies = {
      option = "none"
    }
    forwarded_headers = {
      headers_allow_list = ["host"]
      option             = "allow-list"
    }
    forwarded_query_strings = {
      option = true
    }

  }
  depends_on = [awscc_lightsail_static_ip.example]
}
