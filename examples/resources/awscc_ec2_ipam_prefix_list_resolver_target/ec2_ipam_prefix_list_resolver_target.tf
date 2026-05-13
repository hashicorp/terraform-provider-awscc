resource "awscc_ec2_ipam" "example" {
  description = "Example IPAM for prefix list resolver target"

  operating_regions = [
    {
      region_name = "us-west-2"
    }
  ]

  tags = [{
    key   = "Name"
    value = "example-ipam"
  }, {
    key   = "Environment"
    value = "example"
  }]
}

resource "awscc_ec2_ipam_prefix_list_resolver" "example" {
  ipam_id        = awscc_ec2_ipam.example.id
  address_family = "IPv4"
  description    = "Example IPAM prefix list resolver for target"

  rules = [{
    arn          = "arn:aws:ec2:us-west-2:123456789012:prefix-list/pl-1234567890abcdef0"
    allowed_uses = ["ALLOW"]
  }]

  tags = [{
    key   = "Name"
    value = "example-prefix-list-resolver"
  }, {
    key   = "Environment"
    value = "example"
  }]
}

resource "awscc_ec2_ipam_prefix_list_resolver_target" "example" {
  ipam_prefix_list_resolver_arn = awscc_ec2_ipam_prefix_list_resolver.example.arn
  ipam_prefix_list_resolver_id  = awscc_ec2_ipam_prefix_list_resolver.example.id
  target_arn                    = "arn:aws:ec2:us-west-2:123456789012:vpc/vpc-1234567890abcdef0"
  target_type                   = "VPC"

  tags = [{
    key   = "Name"
    value = "example-prefix-list-resolver-target"
  }, {
    key   = "Environment"
    value = "example"
  }]
}
