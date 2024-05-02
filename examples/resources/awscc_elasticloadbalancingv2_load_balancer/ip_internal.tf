resource "awscc_elasticloadbalancingv2_load_balancer" "example" {
  name   = "example"
  type   = "network"
  scheme = "internal"
  subnet_mappings = [
    {
      subnet_id              = awscc_ec2_subnet.one.subnet_id
      private_i_pv_4_address = "172.31.64.15"
    },
    {
      subnet_id              = awscc_ec2_subnet.two.subnet_id
      private_i_pv_4_address = "172.31.80.15"
    }
  ]

  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
