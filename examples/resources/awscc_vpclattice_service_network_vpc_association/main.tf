# Use data source to get default VPC
data "aws_vpc" "default" {
  default = true
}

# Create VPC Lattice Service Network with a unique name using timestamp
resource "awscc_vpclattice_service_network" "example" {
  name = "example-network-${formatdate("YYYYMMDDhhmmss", timestamp())}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create VPC Lattice Service Network VPC Association
resource "awscc_vpclattice_service_network_vpc_association" "example" {
  vpc_identifier             = data.aws_vpc.default.id
  service_network_identifier = awscc_vpclattice_service_network.example.id
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}