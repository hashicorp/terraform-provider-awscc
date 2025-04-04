# Create Route53 Health Check
resource "aws_route53_health_check" "example" {
  fqdn              = "example.com"
  port              = 80
  type              = "HTTP"
  resource_path     = "/"
  failure_threshold = "5"
  request_interval  = "30"

  tags = {
    Modified_By = "AWSCC"
  }
}

# Create the Route53 Recovery Readiness Resource Set
resource "awscc_route53recoveryreadiness_resource_set" "example" {
  resource_set_name = "example-resource-set"
  resource_set_type = "AWS::Route53::HealthCheck"
  resources = [{
    resource_arn     = "arn:aws:route53:::healthcheck/${aws_route53_health_check.example.id}"
    readiness_scopes = []
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Route53 Recovery Readiness Check
resource "awscc_route53recoveryreadiness_readiness_check" "example" {
  readiness_check_name = "example-readiness-check"
  resource_set_name    = awscc_route53recoveryreadiness_resource_set.example.resource_set_name

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}