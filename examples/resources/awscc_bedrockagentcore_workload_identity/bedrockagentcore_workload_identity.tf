# AWS Bedrock Agent Core Workload Identity
resource "awscc_bedrockagentcore_workload_identity" "example" {
  name = "example-workload-identity"

  # Optional: Configure OAuth2 return URLs for resource
  # You can add multiple return URLs if needed
  oauth2_client_return_urls = [
    "https://example.com/callback"
  ]

  # Optional: Specify the directory for this workload identity
  # Defaults to 'default' if not specified
  workload_identity_directory = "default"

  # Optional: Add tags to the workload identity resource
  tags = {
    Environment = "example"
    Project     = "bedrock-agent-demo"
  }
}