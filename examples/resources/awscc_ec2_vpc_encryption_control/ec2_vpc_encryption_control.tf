resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
}

resource "awscc_ec2_vpc_encryption_control" "example" {
  vpc_id = awscc_ec2_vpc.example.vpc_id
  mode   = "monitor"

  internet_gateway_exclusion_input             = "disable"
  nat_gateway_exclusion_input                  = "disable"
  egress_only_internet_gateway_exclusion_input = "disable"
  elastic_file_system_exclusion_input          = "disable"
  lambda_exclusion_input                       = "disable"
  virtual_private_gateway_exclusion_input      = "disable"
  vpc_lattice_exclusion_input                  = "disable"
  vpc_peering_exclusion_input                  = "disable"

  tags = [
    {
      key   = "Name"
      value = "example-vpc-encryption-control"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}