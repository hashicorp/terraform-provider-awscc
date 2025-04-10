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