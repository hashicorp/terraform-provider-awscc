resource "awscc_lightsail_load_balancer" "example" {
  instance_port      = 443
  load_balancer_name = "example-lb"
}

# Note: Lightsail will verify the certificate during creation
resource "awscc_lightsail_load_balancer_tls_certificate" "example" {
  certificate_domain_name       = "example.com"
  certificate_name              = "example-lb-cert"
  load_balancer_name            = awscc_lightsail_load_balancer.example.load_balancer_name
  certificate_alternative_names = ["www.example.com"]
}
