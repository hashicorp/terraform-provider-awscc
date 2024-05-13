resource "awscc_apigateway_vpc_link" "example" {
  name        = "example"
  target_arns = [awscc_elasticloadbalancingv2_load_balancer.example.load_balancer_arn]

  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_elasticloadbalancingv2_load_balancer" "example" {
  name    = "example"
  scheme  = "internal"
  type    = "network"
  subnets = [var.subnet_id]
}
