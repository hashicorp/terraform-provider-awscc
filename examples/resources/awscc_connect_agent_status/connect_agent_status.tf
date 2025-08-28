# Connect Agent Status resource example
resource "aws_connect_instance" "example" {
  identity_management_type = "CONNECT_MANAGED"
  inbound_calls_enabled    = true
  outbound_calls_enabled   = true
  instance_alias           = "example-connect-instance"
}

# Connect Agent Status resource
resource "awscc_connect_agent_status" "example" {
  instance_arn = aws_connect_instance.example.arn
  name         = "Custom-Available"
  state        = "ENABLED"
  description  = "Agent is available to handle contacts"
  type         = "ROUTABLE"

  tags = [
    {
      key   = "Environment"
      value = "Production"
    },
    {
      key   = "Name"
      value = "example-agent-status"
    }
  ]
}
