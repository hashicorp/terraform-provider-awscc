# Create a VPC Lattice service network
resource "awscc_vpclattice_service_network" "example" {
  name      = "example-service-network"
  auth_type = "NONE"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a VPC Lattice service
resource "awscc_vpclattice_service" "example" {
  name      = "example-service"
  auth_type = "NONE"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the service network service association
resource "awscc_vpclattice_service_network_service_association" "example" {
  service_identifier         = awscc_vpclattice_service.example.id
  service_network_identifier = awscc_vpclattice_service_network.example.id

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}