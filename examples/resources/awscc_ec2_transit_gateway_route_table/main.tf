# Create Transit Gateway first as it's required
resource "awscc_ec2_transit_gateway" "example" {
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Transit Gateway Route Table
resource "awscc_ec2_transit_gateway_route_table" "example" {
  transit_gateway_id = awscc_ec2_transit_gateway.example.id

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}