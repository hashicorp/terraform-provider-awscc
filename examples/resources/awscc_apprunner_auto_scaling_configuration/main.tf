# Basic App Runner Auto Scaling configuration
resource "awscc_apprunner_auto_scaling_configuration" "example" {
  auto_scaling_configuration_name = "example-autoscaling-config"
  max_concurrency                 = 100
  max_size                        = 10
  min_size                        = 1

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}