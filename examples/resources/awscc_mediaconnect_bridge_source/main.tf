# Get current AWS region
# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create the Bridge Source
resource "awscc_mediaconnect_bridge_source" "example" {
  bridge_arn = "arn:aws:mediaconnect:${data.aws_region.current.region}:${data.aws_caller_identity.current.account_id}:bridge:example-bridge"
  name       = "example-source"

  network_source = {
    multicast_ip = "239.0.0.1"
    network_name = "example-network"
    port         = 5000
    protocol     = "rtp"
  }
}