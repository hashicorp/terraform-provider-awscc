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
  description    = "example pool"
  ipam_scope_id  = awscc_ec2_ipam.example.private_default_scope_id
  locale         = data.aws_region.current.name

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}