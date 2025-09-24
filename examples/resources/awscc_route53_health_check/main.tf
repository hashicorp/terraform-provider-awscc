# Create a Route53 HTTP health check
resource "awscc_route53_health_check" "example" {
  health_check_config = {
    type              = "HTTP"
    ip_address        = "8.8.8.8" # Example IP address
    port              = 80
    request_interval  = 30
    failure_threshold = 3
    measure_latency   = true
    resource_path     = "/health"
  }

  health_check_tags = [{
    key   = "Environment"
    value = "Production"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
  }]
}