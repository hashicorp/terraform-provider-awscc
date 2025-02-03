# Get current region
data "aws_region" "current" {}

# Get current account ID
data "aws_caller_identity" "current" {}

# Create VPC Lattice Service
resource "awscc_vpclattice_service" "example" {
  name      = "example-service"
  auth_type = "AWS_IAM"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}