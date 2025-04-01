# Create a VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a subnet
resource "awscc_ec2_subnet" "example" {
  vpc_id     = awscc_ec2_vpc.example.id
  cidr_block = "10.0.1.0/24"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a security group
resource "awscc_ec2_security_group" "example" {
  group_description = "Example security group"
  vpc_id            = awscc_ec2_vpc.example.id
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the network interface
resource "awscc_ec2_network_interface" "example" {
  subnet_id                          = awscc_ec2_subnet.example.id
  description                        = "Example network interface"
  group_set                          = [awscc_ec2_security_group.example.id]
  source_dest_check                  = true
  private_ip_address                 = "10.0.1.100"
  secondary_private_ip_address_count = 2
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}