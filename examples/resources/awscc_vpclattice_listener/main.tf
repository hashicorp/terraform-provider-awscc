# Create a VPC Lattice Service
resource "awscc_vpclattice_service" "example" {
  name      = "example-service"
  auth_type = "NONE"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the listener
resource "awscc_vpclattice_listener" "example" {
  name               = "example-listener"
  protocol           = "HTTP"
  port               = 80
  service_identifier = awscc_vpclattice_service.example.id

  default_action = {
    fixed_response = {
      status_code = 404
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}