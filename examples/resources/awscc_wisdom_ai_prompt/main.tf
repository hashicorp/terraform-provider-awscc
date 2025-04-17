# First we need to create a Wisdom Assistant
resource "awscc_wisdom_assistant" "example" {
  name        = "ExampleAssistant"
  description = "Example Wisdom Assistant for AI Prompt"
  type        = "AGENT"
}

# Example AWS Wisdom AI Prompt
resource "awscc_wisdom_ai_prompt" "example" {
  name         = "ExampleAIPrompt"
  description  = "Example AI Prompt created using AWSCC provider"
  assistant_id = awscc_wisdom_assistant.example.assistant_id

  api_format    = "ANTHROPIC_CLAUDE_TEXT_COMPLETIONS"
  model_id      = "anthropic.claude-v2"
  type          = "ANSWER_GENERATION"
  template_type = "TEXT"

  template_configuration = {
    text_full_ai_prompt_edit_template_configuration = {
      text = "Based on the provided context, answer this question: {question}"
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}