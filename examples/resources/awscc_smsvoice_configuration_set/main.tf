resource "awscc_smsvoice_configuration_set" "example" {
  configuration_set_name = "example-configuration-set"

  message_feedback_enabled = true

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-configuration-set"
    }
  ]
}