# Custom Action resource
resource "awscc_chatbot_custom_action" "example" {
  action_name = "example-custom-action"
  alias_name  = "ExampleAction"

  definition = {
    command_text = "example command"
  }

  attachments = [
    {
      button_text       = "Click me"
      notification_type = "SUCCESS"
      variables = {
        "key1" = "value1"
        "key2" = "value2"
      }
      criteria = [
        {
          operator      = "EQUALS"
          value         = "test"
          variable_name = "test_var"
        }
      ]
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}