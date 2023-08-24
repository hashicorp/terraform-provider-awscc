data "aws_region" "current" {}

resource "awscc_ec2_ipam" "example" {
  operating_regions = [
    {
      region_name = data.aws_region.current.name
    }
  ]
}

resource "awscc_ec2_ipam_pool" "example" {
  address_family = "ipv4"
  ipam_scope_id  = awscc_ec2_ipam.example.private_default_scope_id
  locale         = data.aws_region.current.name

  allocation_default_netmask_length = 24
  allocation_max_netmask_length     = 25
  allocation_min_netmask_length     = 22

  allocation_resource_tags = [{
    key   = "CidrSource"
    value = "FromIPAM"
  }]
}