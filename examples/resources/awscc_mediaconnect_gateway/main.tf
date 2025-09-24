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