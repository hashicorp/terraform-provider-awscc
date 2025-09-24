resource "awscc_elasticloadbalancingv2_load_balancer" "example" {
  name   = "example"
  scheme = "internal"
  subnet_mappings = [
    {
      subnet_id = awscc_ec2_subnet.one.subnet_id
    },
    {
      subnet_id = awscc_ec2_subnet.two.subnet_id
    }
  ]

  load_balancer_attributes = [{
    key   = "idle_timeout.timeout_seconds"
    value = "30"
  }]

  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
