resource "awscc_ec2_ipam" "example" {
  description = "Example IPAM for prefix list resolver"

  operating_regions = [
    {
      region_name = "us-east-1"
    },
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

resource "awscc_ec2_ipam_scope" "example" {
  ipam_id     = awscc_ec2_ipam.example.id
  description = "Example IPAM scope"

  tags = [{
    key   = "Name"
    value = "example-ipam-scope"
    }, {
    key   = "Environment"
    value = "example"
  }]
}

resource "awscc_ec2_ipam_prefix_list_resolver" "example" {
  ipam_id        = awscc_ec2_ipam.example.id
  address_family = "IPv4"
  description    = "Example IPAM prefix list resolver"

  rules = [{
    rule_type   = "static-cidr"
    static_cidr = "10.0.0.0/16"
  }]

  tags = [{
    key   = "Name"
    value = "example-prefix-list-resolver"
    }, {
    key   = "Environment"
    value = "example"
  }]
}