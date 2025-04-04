# Create the Amazon Connect Wisdom Assistant
resource "awscc_wisdom_assistant" "example" {
  name        = "example-assistant"
  type        = "AGENT"
  description = "Example Wisdom Assistant for AI Agent Version"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Amazon Connect Wisdom AI Agent
resource "awscc_wisdom_ai_agent" "example" {
  name         = "example-ai-agent"
  description  = "Example Wisdom AI Agent"
  assistant_id = awscc_wisdom_assistant.example.id
  type         = "MANUAL_SEARCH"
  configuration = {
    manual_search = {}
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Wisdom AI Agent Version
resource "awscc_wisdom_ai_agent_version" "example" {
  ai_agent_id  = awscc_wisdom_ai_agent.example.id
  assistant_id = awscc_wisdom_assistant.example.id
}