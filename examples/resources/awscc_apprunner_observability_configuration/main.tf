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