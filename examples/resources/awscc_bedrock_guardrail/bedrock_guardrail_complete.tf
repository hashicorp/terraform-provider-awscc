resource "awscc_bedrock_guardrail" "example" {
  name                      = "example_guardrail"
  blocked_input_messaging   = "Blocked input"
  blocked_outputs_messaging = "Blocked output"
  description               = "Example guardrail"

  content_policy_config = {
    filters_config = [
      {
        input_strength  = "MEDIUM"
        output_strength = "MEDIUM"
        type            = "HATE"
      },
      {
        input_strength  = "HIGH"
        output_strength = "HIGH"
        type            = "VIOLENCE"
      }
    ]
  }

  sensitive_information_policy_config = {
    pii_entities_config = [
      {
        action = "BLOCK"
        type   = "NAME"
      },
      {
        action = "BLOCK"
        type   = "DRIVER_ID"
      },
      {
        action = "ANONYMIZE"
        type   = "USERNAME"
      },
    ]
    regexes_config = [{
      action      = "BLOCK"
      description = "example regex"
      name        = "regex_example"
      pattern     = "^\\d{3}-\\d{2}-\\d{4}$"
    }]
  }
  word_policy_config = {
    managed_word_lists_config = [{
      type = "PROFANITY"
    }]
    words_config = [{
      text = "HATE"
    }]
  }

  topic_policy_config = {
    topics_config = [{
      name       = "investment_topic"
      examples   = ["Where should I invest my money ?"]
      type       = "DENY"
      definition = "Investment advice refers to inquiries, guidance, or recommendations regarding the management or allocation of funds or assets with the goal of generating returns ."
    }]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

  kms_key_arn = var.kms_key_arn

}

variable "kms_key_arn" {
  type = string
}
