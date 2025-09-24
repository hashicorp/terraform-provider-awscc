resource "awscc_lightsail_instance" "example" {
  blueprint_id  = "nginx"
  bundle_id     = "nano_3_0"
  instance_name = "example-instance"
}

resource "awscc_lightsail_load_balancer" "example" {
  instance_port      = 80
  load_balancer_name = "example-lb"
  attached_instances = [awscc_lightsail_instance.example.instance_name]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
