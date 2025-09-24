resource "awscc_bedrock_guardrail_version" "example" {
  guardrail_identifier = awscc_bedrock_guardrail.example.guardrail_id
  description          = "Example guardrail version"
}