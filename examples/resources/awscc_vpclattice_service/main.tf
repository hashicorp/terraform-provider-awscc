# Create VPC Lattice Service
resource "awscc_vpclattice_service" "example" {
  name      = "example-service"
  auth_type = "AWS_IAM"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}