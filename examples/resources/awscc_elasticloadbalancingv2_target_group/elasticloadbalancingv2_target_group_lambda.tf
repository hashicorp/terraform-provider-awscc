resource "awscc_elasticloadbalancingv2_target_group" "lambda-example" {
  name        = "lambda-example"
  target_type = "lambda"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}