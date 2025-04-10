# Example AppConfig Deployment Strategy
resource "awscc_appconfig_deployment_strategy" "example" {
  name                           = "example-deployment-strategy"
  description                    = "Example deployment strategy for AppConfig"
  deployment_duration_in_minutes = 60
  growth_factor                  = 10
  growth_type                    = "LINEAR"
  replicate_to                   = "NONE"
  final_bake_time_in_minutes     = 30

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}