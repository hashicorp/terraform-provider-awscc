resource "awscc_elasticloadbalancingv2_load_balancer" "example" {
  name = "example"
  type = "network"
  subnet_mappings = [
    {
      subnet_id = awscc_ec2_subnet.one.subnet_id
    },
    {
      subnet_id = awscc_ec2_subnet.two.subnet_id
    }
  ]

  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

