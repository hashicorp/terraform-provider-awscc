resource "awscc_lightsail_instance" "example" {
  blueprint_id  = "amazon_linux_2023"
  bundle_id     = "nano_3_0"
  instance_name = "example-instance"
}

resource "awscc_lightsail_static_ip" "example" {
  static_ip_name = "example-ip"
  attached_to    = awscc_lightsail_instance.example.id
}
