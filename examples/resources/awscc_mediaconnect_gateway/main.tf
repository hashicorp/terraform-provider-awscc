# Get current AWS Account ID
data "aws_caller_identity" "current" {}

# Get current AWS Region
data "aws_region" "current" {}

# Create MediaConnect Gateway
resource "awscc_mediaconnect_gateway" "example" {
  name = "example-gateway"
  # Allow traffic from specific CIDR blocks
  egress_cidr_blocks = ["10.0.0.0/16"]

  # Define network configurations
  networks = [
    {
      name       = "network-1"
      cidr_block = "172.16.0.0/24"
    },
    {
      name       = "network-2"
      cidr_block = "172.16.1.0/24"
    }
  ]
}