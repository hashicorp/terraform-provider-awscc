# Get current AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create a core network
resource "awscc_networkmanager_core_network" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  description       = "Example Core Network"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a global network
resource "awscc_networkmanager_global_network" "example" {
  description = "Example Global Network"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a Direct Connect gateway
resource "aws_dx_gateway" "example" {
  amazon_side_asn = 64512
  name            = "example-dx-gateway"
}

# Create the Direct Connect gateway attachment
resource "awscc_networkmanager_direct_connect_gateway_attachment" "example" {
  core_network_id = awscc_networkmanager_core_network.example.id
  direct_connect_gateway_arn = format("arn:aws:directconnect:%s:%s:dx-gateway/%s",
    data.aws_region.current.name,
    data.aws_caller_identity.current.account_id,
  aws_dx_gateway.example.id)
  edge_locations = [data.aws_region.current.name]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}