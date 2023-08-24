data "aws_region" "current" {}

resource "awscc_ec2_ipam" "example" {
  operating_regions = [
    {
      region_name = data.aws_region.current.name
    }
  ]
}

resource "awscc_ec2_ipam_pool" "example" {
  address_family = "ipv6"
  ipam_scope_id  = awscc_ec2_ipam.example.public_default_scope_id
  locale         = data.aws_region.current.name

  aws_service           = "ec2"
  public_ip_source      = "byoip"
  publicly_advertisable = true
}