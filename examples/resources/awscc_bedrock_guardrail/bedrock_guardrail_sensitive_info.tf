resource "awscc_bedrock_guardrail" "example" {
  name                      = "example_guardrail"
  blocked_input_messaging   = "Blocked input"
  blocked_outputs_messaging = "Blocked output"
  description               = "Example guardrail"

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

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]


}

