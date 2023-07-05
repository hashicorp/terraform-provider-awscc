resource "awscc_elasticloadbalancingv2_target_group" "alb-example" {
  name     = "alb-example"
  port     = 80
  protocol = "TCP"
  target_type = "alb"
  vpc_id   = awscc_ec2_vpc.main.id
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}