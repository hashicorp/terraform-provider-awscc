terraform {
  required_providers {
    awscc = {
      source  = "hashicorp/awscc"
      version = "1.76.0"
    }
  }
}

resource "awscc_observabilityadmin_telemetry_enrichment" "example" {
  # Example configuration for ObservabilityAdmin telemetry enrichment
  tags = {
    Environment = "test"
    Purpose     = "example"
  }
}
