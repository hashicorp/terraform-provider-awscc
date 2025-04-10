# Create Transit Gateway
resource "awscc_ec2_transit_gateway" "example" {
  description = "Example Transit Gateway for Route Testing"
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

# Create Transit Gateway Route
resource "awscc_ec2_transit_gateway_route" "example" {
  destination_cidr_block         = "10.0.0.0/16"
  transit_gateway_route_table_id = awscc_ec2_transit_gateway_route_table.example.id
  blackhole                      = true
}