data "aws_caller_identity" "current" {}

# First create a Wisdom Assistant
resource "awscc_wisdom_assistant" "example" {
  name = "example-assistant-${data.aws_caller_identity.current.account_id}"
  type = "AGENT"
  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create an AI Guardrail
resource "awscc_wisdom_ai_guardrail" "example" {
  assistant_id = awscc_wisdom_assistant.example.id
  name         = "example-guardrail-${data.aws_caller_identity.current.account_id}"
  blocked_outputs_messaging = jsonencode({
    prohibited_content_messaging = true
  })
  blocked_input_messaging = jsonencode({
    prohibited_content_messaging = true
  })
  tags = {
    ModifiedBy = "AWSCC"
  }
}

# Create an AI Guardrail Version
resource "awscc_wisdom_ai_guardrail_version" "example" {
  assistant_id    = awscc_wisdom_assistant.example.id
  ai_guardrail_id = awscc_wisdom_ai_guardrail.example.id
}