# Assume certificate is already created and verified
data "awscc_lightsail_certificate" "example" {
  id = "example-cert"
}

resource "awscc_lightsail_instance" "example" {
  blueprint_id  = "nginx"
  bundle_id     = "nano_3_0"
  instance_name = "example-instance"
}

resource "awscc_lightsail_load_balancer" "example" {
  instance_port      = 80
  load_balancer_name = "example-lb"
  attached_instances = [awscc_lightsail_instance.example.instance_name]
}
resource "awscc_lightsail_distribution" "example" {
  bundle_id = "small_1_0"
  default_cache_behavior = {
    behavior = "dont-cache"
  }
  distribution_name = "example-dist"
  origin = {
    name        = awscc_lightsail_load_balancer.example.load_balancer_name
    region_name = "us-east-1"
  }
  certificate_name = data.awscc_lightsail_certificate.example.certificate_name
}
