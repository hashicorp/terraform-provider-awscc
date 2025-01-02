# Create IoT command with basic configuration
resource "awscc_iot_command" "example" {
  command_id   = "example-command"
  display_name = "Example IoT Command"
  description  = "An example IoT command created by AWSCC provider"
  namespace    = "AWS-IoT"

  payload = {
    content      = base64encode(jsonencode({ action = "read_sensor" }))
    content_type = "application/json"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}