# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

resource "awscc_ec2_ipam" "example" {
  operating_regions = [
    {
      region_name = data.aws_region.current.region
    }
  ]
}

resource "awscc_ec2_ipam_pool" "example" {
  address_family = "ipv4"
  ipam_scope_id  = awscc_ec2_ipam.example.private_default_scope_id
  locale         = data.aws_region.current.region

  provisioned_cidrs = [
    {
      cidr = "10.0.0.0/16"
    }
  ]
}

resource "awscc_ec2_ipam_allocation" "example" {
  ipam_pool_id = awscc_ec2_ipam_pool.example.id
  description  = "reserved using netmask length"

  netmask_length = 32
}