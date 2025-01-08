# Get current AWS account info
data "aws_caller_identity" "current" {}

# Get current region
data "aws_region" "current" {}

# Create a VPC Lattice Service Network
resource "awscc_vpclattice_service_network" "example" {
  name      = "example-service-network"
  auth_type = "AWS_IAM"

  # Configure sharing if needed
  sharing_config = {
    enabled = true
  }

  # Add tags
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}