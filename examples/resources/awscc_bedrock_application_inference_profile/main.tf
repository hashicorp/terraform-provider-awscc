# Get current AWS region
data "aws_region" "current" {}

# Example of awscc_bedrock_application_inference_profile
resource "awscc_bedrock_application_inference_profile" "example" {
  inference_profile_name = "example-profile"
  description            = "Example inference profile for bedrock application"

  model_source = {
    # Using Claude v3 Sonnet model ARN format as example
    copy_from = "arn:aws:bedrock:${data.aws_region.current.name}::foundation-model/anthropic.claude-3-sonnet-20240229-v1:0"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Output the inference profile ARN
output "inference_profile_arn" {
  value = awscc_bedrock_application_inference_profile.example.inference_profile_arn
}