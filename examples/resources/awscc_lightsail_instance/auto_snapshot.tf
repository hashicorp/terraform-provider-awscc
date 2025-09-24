resource "awscc_lightsail_instance" "example" {
  blueprint_id  = "amazon_linux_2023"
  bundle_id     = "nano_3_0"
  instance_name = "example-instance"
  add_ons = [{
    add_on_type = "AutoSnapshot"
    auto_snapshot_add_on_request = {
      snapshot_time_of_day = "06:00"
    }
  }]
}
