resource "awscc_lightsail_disk" "example" {
  disk_name  = "example-disk"
  size_in_gb = "32"
  add_ons = [{
    add_on_type = "AutoSnapshot"
    auto_snapshot_add_on_request = {
      snapshot_time_of_day = "06:00"
    }
  }]
  availability_zone = "us-east-1a"
}
