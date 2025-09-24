resource "awscc_ec2_eip" "example" {
  count  = 2
  domain = "vpc"
}

resource "awscc_elasticloadbalancingv2_load_balancer" "example" {
  name = "example"
  type = "network"
  subnet_mappings = [
    {
      subnet_id     = awscc_ec2_subnet.one.subnet_id
      allocation_id = awscc_ec2_eip.example[0].allocation_id
    },
    {
      subnet_id     = awscc_ec2_subnet.two.subnet_id
      allocation_id = awscc_ec2_eip.example[1].allocation_id
    }
  ]

  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
