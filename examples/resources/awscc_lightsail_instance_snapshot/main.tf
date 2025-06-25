# Get current AWS region
# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

# Create a Lightsail instance first to take a snapshot from
resource "awscc_lightsail_instance" "example" {
  instance_name     = "example-instance"
  availability_zone = "${data.aws_region.current.region}a"
  blueprint_id      = "amazon_linux_2"
  bundle_id         = "nano_2_0"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the instance snapshot
resource "awscc_lightsail_instance_snapshot" "example" {
  instance_name          = awscc_lightsail_instance.example.instance_name
  instance_snapshot_name = "example-snapshot"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}