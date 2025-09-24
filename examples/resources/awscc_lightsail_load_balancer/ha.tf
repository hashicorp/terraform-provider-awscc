resource "awscc_lightsail_instance" "example_1a" {
  blueprint_id      = "nginx"
  bundle_id         = "nano_3_0"
  instance_name     = "example-instance-1a"
  availability_zone = "us-east-1a"
}

resource "awscc_lightsail_instance" "example_1b" {
  blueprint_id      = "nginx"
  bundle_id         = "nano_3_0"
  instance_name     = "example-instance-1b"
  availability_zone = "us-east-1b"
}

resource "awscc_lightsail_load_balancer" "example" {
  instance_port      = 80
  load_balancer_name = "example-lb"
  attached_instances = [
    awscc_lightsail_instance.example_1a.instance_name,
    awscc_lightsail_instance.example_1b.instance_name
  ]
  session_stickiness_enabled                    = true
  session_stickiness_lb_cookie_duration_seconds = 86500
}
