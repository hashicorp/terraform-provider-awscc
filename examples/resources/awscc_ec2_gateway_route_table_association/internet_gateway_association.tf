resource "awscc_ec2_gateway_route_table_association" "igw" {
  gateway_id     = awscc_ec2_internet_gateway.igw.id
  route_table_id = awscc_ec2_route_table.internet.id
}

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

resource "awscc_ec2_route_table" "internet" {
  vpc_id = awscc_ec2_vpc.vpc.id
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
