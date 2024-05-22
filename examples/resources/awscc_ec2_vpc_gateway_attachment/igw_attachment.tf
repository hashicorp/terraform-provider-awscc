resource "awscc_ec2_vpc_gateway_attachment" "igw" {
  internet_gateway_id = awscc_ec2_internet_gateway.igw.id
  vpc_id              = awscc_ec2_vpc.vpc.id
}

resource "awscc_ec2_internet_gateway" "igw" {
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_vpc" "vpc" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
