resource "awscc_lightsail_instance" "example" {
  blueprint_id  = "wordpress"
  bundle_id     = "small_3_0"
  instance_name = "example-instance"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
