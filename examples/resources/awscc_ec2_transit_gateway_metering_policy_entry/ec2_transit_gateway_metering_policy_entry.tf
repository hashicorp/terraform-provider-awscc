resource "awscc_ec2_transit_gateway_metering_policy_entry" "example" {
  metering_type                  = "VPN"
  resource_type                  = "VPN"
  transit_gateway_id             = aws_ec2_transit_gateway.example.id
  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.example.id

  tags = {
    Name        = "example-metering-policy"
    Environment = "test"
  }
}

# Supporting resources
resource "aws_ec2_transit_gateway" "example" {
  description = "Example Transit Gateway"

  tags = {
    Name = "example-tgw"
  }
}

resource "aws_ec2_transit_gateway_route_table" "example" {
  transit_gateway_id = aws_ec2_transit_gateway.example.id

  tags = {
    Name = "example-tgw-rt"
  }
}
