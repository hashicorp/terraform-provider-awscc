# Create a recovery group first
resource "awscc_route53recoveryreadiness_recovery_group" "example" {
  recovery_group_name = "example-group"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Get current AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create a resource set for NLB
resource "awscc_route53recoveryreadiness_resource_set" "nlb_example" {
  resource_set_name = "nlb-resources"
  resource_set_type = "AWS::ElasticLoadBalancingV2::LoadBalancer"

  resources = [{
    resource_arn = "arn:aws:elasticloadbalancing:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:loadbalancer/app/example/1234567890"
    readiness_scopes = [
      awscc_route53recoveryreadiness_recovery_group.example.recovery_group_arn
    ]
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}