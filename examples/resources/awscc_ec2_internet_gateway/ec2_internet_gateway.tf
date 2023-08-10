resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.1.0.0/16"
}

resource "awscc_ec2_internet_gateway" "example" {
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "aws_internet_gateway_attachment" "example" {
  internet_gateway_id = awscc_ec2_internet_gateway.example.id
  vpc_id              = awscc_ec2_vpc.example.id
}