# NOTE: This example demonstrates the structure of AWS Wisdom AI Prompt Version,
# but the actual creation might require using the AWS CLI or SDK to create the AI Prompt first
# due to current limitations with the AWSCC provider's handling of the template_configuration.


# Create the Wisdom Assistant
resource "awscc_wisdom_assistant" "example" {
  name        = "example-assistant"
  type        = "AGENT"
  description = "Example Wisdom Assistant for AI prompt"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create the Wisdom AI Prompt Version
resource "awscc_wisdom_ai_prompt_version" "example" {
  assistant_id = awscc_wisdom_assistant.example.id
  # The ai_prompt_id must be a valid UUID from an existing AI Prompt
  # Example format: "12345678-1234-1234-1234-123456789012"
  ai_prompt_id = "12345678-1234-1234-1234-123456789012" # Replace with actual AI Prompt ID
}