# Create a Wisdom Assistant first (required for AI Guardrail)
resource "awscc_wisdom_assistant" "example" {
  name        = "example-assistant-${formatdate("YYYYMMDD-hhmmss", timestamp())}"
  type        = "AGENT"
  description = "Example Wisdom Assistant for AI Guardrail"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the AI Guardrail
resource "awscc_wisdom_ai_guardrail" "example" {
  name         = "example-guardrail-${formatdate("YYYYMMDD-hhmmss", timestamp())}"
  description  = "Example AI Guardrail configuration"
  assistant_id = awscc_wisdom_assistant.example.id

  # Required messaging configurations
  blocked_input_messaging   = "Your input contains prohibited content. Please revise and try again."
  blocked_outputs_messaging = "The response was blocked due to policy violations."

  # Content policy configuration
  content_policy_config = {
    filters_config = [
      {
        type            = "HATE"
        input_strength  = "HIGH"
        output_strength = "HIGH"
      }
    ]
  }

  # Word policy configuration
  word_policy_config = {
    words_config = [
      {
        text = "restricted_word"
      }
    ]
    managed_word_lists_config = [
      {
        type = "PROFANITY"
      }
    ]
  }

  # Topic policy configuration
  topic_policy_config = {
    topics_config = [
      {
        name       = "Restricted Topic"
        type       = "DENY"
        definition = "Any discussion related to restricted topics"
        examples   = ["This is an example of restricted content"]
      }
    ]
  }

  # Sensitive information policy configuration
  sensitive_information_policy_config = {
    pii_entities_config = [
      {
        type   = "US_BANK_ACCOUNT_NUMBER"
        action = "BLOCK"
      }
    ]
    regexes_config = [
      {
        name        = "CustomPattern"
        description = "Custom pattern to detect sensitive information"
        pattern     = "[A-Z]{2}\\d{6}"
        action      = "BLOCK"
      }
    ]
  }

  # Contextual grounding policy configuration
  contextual_grounding_policy_config = {
    filters_config = [
      {
        type      = "GROUNDING"
        threshold = 0.7
      }
    ]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}