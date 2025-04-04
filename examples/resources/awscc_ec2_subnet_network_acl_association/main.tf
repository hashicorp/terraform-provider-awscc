# Create VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Subnet
resource "awscc_ec2_subnet" "example" {
  vpc_id     = awscc_ec2_vpc.example.id
  cidr_block = "10.0.1.0/24"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Network ACL
resource "awscc_ec2_network_acl" "example" {
  vpc_id = awscc_ec2_vpc.example.id

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Subnet Network ACL Association
resource "awscc_ec2_subnet_network_acl_association" "example" {
  subnet_id      = awscc_ec2_subnet.example.id
  network_acl_id = awscc_ec2_network_acl.example.id
}