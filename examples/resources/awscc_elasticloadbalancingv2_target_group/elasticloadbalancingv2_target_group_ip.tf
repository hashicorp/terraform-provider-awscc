resource "awscc_elasticloadbalancingv2_target_group" "ip-example" {
  name        = "ip-example"
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = awscc_ec2_vpc.main.id
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}