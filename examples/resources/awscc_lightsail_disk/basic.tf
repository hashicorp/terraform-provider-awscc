resource "awscc_lightsail_disk" "example" {
  disk_name         = "example-disk"
  size_in_gb        = "32"
  availability_zone = "us-east-1a"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
