resource "awscc_ec2_subnet_route_table_association" "this" {
  route_table_id = awscc_ec2_route_table.this.id
  subnet_id      = awscc_ec2_subnet.this.id
}

resource "awscc_ec2_vpc" "this" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_route_table" "this" {
  vpc_id = awscc_ec2_vpc.this.id
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_subnet" "this" {
  vpc_id            = awscc_ec2_vpc.this.id
  cidr_block        = "10.0.101.0/24"
  availability_zone = "us-east-1a"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
