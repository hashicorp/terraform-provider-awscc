# Example of a deployment configuration for CodeDeploy
resource "awscc_codedeploy_deployment_config" "example" {
  deployment_config_name = "ExampleDeployConfig"
  compute_platform       = "Lambda"

  traffic_routing_config = {
    type = "TimeBasedCanary"
    time_based_canary = {
      canary_interval   = 15
      canary_percentage = 10
    }
  }
}