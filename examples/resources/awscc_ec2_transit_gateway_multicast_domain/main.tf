# We need to use aws provider for Transit Gateway since AWSCC doesn't support multicast_support option
resource "aws_ec2_transit_gateway" "example" {
  description       = "Example Transit Gateway for Multicast Domain"
  multicast_support = "enable"
  tags = {
    "Created By" = "AWSCC"
  }
}

# Create the Transit Gateway Multicast Domain
resource "awscc_ec2_transit_gateway_multicast_domain" "example" {
  transit_gateway_id = aws_ec2_transit_gateway.example.id

  options = {
    auto_accept_shared_associations = "enable"
    igmpv_2_support                 = "enable"
    static_sources_support          = "disable"
  }

  tags = [{
    key   = "Created By"
    value = "AWSCC"
  }]
}