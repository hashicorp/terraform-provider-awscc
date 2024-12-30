# Data sources for AWS account and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# AppRunner Observability Configuration
resource "awscc_apprunner_observability_configuration" "example" {
  observability_configuration_name = "example-observability-config"

  trace_configuration = {
    vendor = "AWSXRAY"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}