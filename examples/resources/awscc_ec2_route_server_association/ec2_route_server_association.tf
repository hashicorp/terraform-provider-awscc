# Example VPC for association
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"

  tags = [
    {
      key   = "Name"
      value = "example-vpc"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

# Example Route Server
resource "awscc_ec2_route_server" "example" {
  amazon_side_asn = 65000
}

# Route Server Association
resource "awscc_ec2_route_server_association" "example" {
  route_server_id = awscc_ec2_route_server.example.id
  vpc_id          = awscc_ec2_vpc.example.id
}
