# Create a VPC for the carrier gateway
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the carrier gateway
resource "awscc_ec2_carrier_gateway" "example" {
  vpc_id = awscc_ec2_vpc.example.id
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}