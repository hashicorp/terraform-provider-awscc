resource "awscc_bedrock_guardrail" "example" {
  name                      = "example_guardrail"
  blocked_input_messaging   = "Blocked input"
  blocked_outputs_messaging = "Blocked output"
  description               = "Example guardrail"
  word_policy_config = {
    managed_word_lists_config = [{
      type = "PROFANITY"
    }]
    words_config = [{
      text = "HATE"
    }]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]


}
