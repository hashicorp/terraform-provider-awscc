data "aws_region" "current" {}

resource "awscc_ec2_ipam" "example" {
  operating_regions = [
    {
      region_name = data.aws_region.current.name
    }
  ]
}

resource "awscc_ec2_ipam_pool" "level1" {
  address_family = "ipv4"
  ipam_scope_id  = awscc_ec2_ipam.example.private_default_scope_id
  locale         = data.aws_region.current.name
}

resource "awscc_ec2_ipam_pool_cidr" "level1_cidr1" {
  ipam_pool_id = awscc_ec2_ipam_pool.level1.id
  cidr         = "10.0.0.0/16"
}

resource "awscc_ec2_ipam_pool" "level_2" {
  address_family      = "ipv4"
  ipam_scope_id       = awscc_ec2_ipam.example.private_default_scope_id
  locale              = data.aws_region.current.name
  source_ipam_pool_id = awscc_ec2_ipam_pool.level1.id
}

resource "awscc_ec2_ipam_pool_cidr" "level2_cidr1" {
  ipam_pool_id   = awscc_ec2_ipam_pool.level_2.id
  netmask_length = "24"
}