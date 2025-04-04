# Create a VPC for the application
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a RefactorSpaces Environment
resource "awscc_refactorspaces_environment" "example" {
  name                = "example-env"
  description         = "Example environment for application"
  network_fabric_type = "NONE"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the RefactorSpaces Application
resource "awscc_refactorspaces_application" "example" {
  name                   = "example-app"
  environment_identifier = awscc_refactorspaces_environment.example.environment_identifier
  vpc_id                 = awscc_ec2_vpc.example.id
  proxy_type             = "API_GATEWAY"

  api_gateway_proxy = {
    endpoint_type = "REGIONAL"
    stage_name    = "prod"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}