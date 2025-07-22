resource "awscc_events_connection" "pagerduty_connection" {
  name               = "pagerduty-connection"
  authorization_type = "API_KEY"

  auth_parameters = {
    api_key_auth_parameters = {
      api_key_name  = "Authorization"
      api_key_value = "my-secret-string"
    }

    additional_parameters = {
      body_parameters = {
        routing_key = "my-pagerduty-integration-key"
      }
    }
  }
}