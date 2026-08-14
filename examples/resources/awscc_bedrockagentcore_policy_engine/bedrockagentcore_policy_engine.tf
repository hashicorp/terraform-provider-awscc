resource "awscc_bedrockagentcore_policy_engine" "example" {
  name        = "example_policy_engine"
  description = "Example policy engine for Bedrock Agent Core"

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example_policy_engine"
    }
  ]
}
