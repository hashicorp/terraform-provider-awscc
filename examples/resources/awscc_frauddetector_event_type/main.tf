# Create a Fraud Detector Event Type
resource "awscc_frauddetector_event_type" "example" {
  name        = "example-event-type"
  description = "Example Fraud Detector Event Type"

  # Entity types define the actors involved in the event
  entity_types = [
    {
      name        = "customer"
      description = "Customer entity type"
      inline      = true
      tags = [{
        key   = "Modified By"
        value = "AWSCC"
      }]
    }
  ]

  # Event variables define the data points collected for the event
  event_variables = [
    {
      name          = "ip_address"
      data_type     = "STRING"
      data_source   = "EVENT"
      description   = "IP address of the customer"
      variable_type = "IP_ADDRESS"
      default_value = "0.0.0.0"
      inline        = true
      tags = [{
        key   = "Modified By"
        value = "AWSCC"
      }]
    },
    {
      name          = "email_address"
      data_type     = "STRING"
      data_source   = "EVENT"
      description   = "Email address of the customer"
      variable_type = "EMAIL_ADDRESS"
      default_value = "example@example.com"
      inline        = true
      tags = [{
        key   = "Modified By"
        value = "AWSCC"
      }]
    }
  ]

  # Labels define the possible outcomes of fraud detection
  labels = [
    {
      name        = "fraud"
      description = "Fraudulent transaction"
      inline      = true
      tags = [{
        key   = "Modified By"
        value = "AWSCC"
      }]
    },
    {
      name        = "legitimate"
      description = "Legitimate transaction"
      inline      = true
      tags = [{
        key   = "Modified By"
        value = "AWSCC"
      }]
    }
  ]

  # Add tags to the event type
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}