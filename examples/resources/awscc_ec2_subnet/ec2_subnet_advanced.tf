resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

resource "awscc_ec2_subnet" "main" {
  vpc_id     = resource.awscc_ec2_vpc.main.id
  cidr_block = "10.0.1.0/24"
  tags = [{
    key   = "Name"
    value = "main"
  }]
}

resource "awscc_ec2_network_acl" "main" {
  vpc_id = resource.awscc_ec2_vpc.main.id
}

resource "awscc_ec2_subnet_network_acl_association" "main" {
  network_acl_id = resource.awscc_ec2_network_acl.main.id
  subnet_id      = resource.awscc_ec2_subnet.main.id
}
