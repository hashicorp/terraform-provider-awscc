# Create SES configuration set
resource "awscc_ses_configuration_set" "example" {
  name = "example-config-set"

  delivery_options = {
    tls_policy           = "OPTIONAL"
    max_delivery_seconds = 3600
  }

  reputation_options = {
    reputation_metrics_enabled = true
  }

  sending_options = {
    sending_enabled = true
  }

  suppression_options = {
    suppressed_reasons = ["BOUNCE", "COMPLAINT"]
  }

  vdm_options = {
    dashboard_options = {
      engagement_metrics = "ENABLED"
    }
    guardian_options = {
      optimized_shared_delivery = "ENABLED"
    }
  }
}