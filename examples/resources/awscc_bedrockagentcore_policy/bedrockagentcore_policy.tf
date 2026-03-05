resource "awscc_bedrockagentcore_policy_engine" "example" {
  name        = "ExamplePolicyEngine"
  description = "Example policy engine for BedrockAgentCore"

}

resource "awscc_bedrockagentcore_policy" "example" {
  name              = "ExamplePolicy"
  description       = "Example BedrockAgentCore policy"
  policy_engine_arn = awscc_bedrockagentcore_policy_engine.example.policy_engine_arn
  policy_statements = {
    statements = [
      {
        statement = <<-EOT
          permit(
            principal,
            action,
            resource is AgentCore::Gateway
          );
        EOT
      }
    ]
  }

  validation_mode = "IGNORE_ALL_FINDINGS"
}
