# Create a VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  instance_tenancy     = "default"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = [{
    key   = "Name"
    value = "example-vpc"
  }]
}

# Create DHCP Options Set
resource "awscc_ec2_dhcp_options" "example" {
  domain_name         = "example.com"
  domain_name_servers = ["AmazonProvidedDNS"]

  tags = [{
    key   = "Name"
    value = "example-dhcp-options"
  }]
}

# Associate DHCP Options with VPC
resource "awscc_ec2_vpcdhcp_options_association" "example" {
  vpc_id          = awscc_ec2_vpc.example.id
  dhcp_options_id = awscc_ec2_dhcp_options.example.id
}