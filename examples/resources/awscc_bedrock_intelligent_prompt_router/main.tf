data "aws_region" "current" {}

# Create the Bedrock Intelligent Prompt Router
resource "awscc_bedrock_intelligent_prompt_router" "example" {
  prompt_router_name = "example-intelligent-prompt-router"
  description        = "Example intelligent prompt router for routing between Claude models based on response quality"

  # Primary models to route between (limited to exactly 2 models)
  models = [
    {
      model_arn = "arn:aws:bedrock:${data.aws_region.current.name}::foundation-model/anthropic.claude-3-5-sonnet-20241022-v2:0"
    },
    {
      model_arn = "arn:aws:bedrock:${data.aws_region.current.name}::foundation-model/anthropic.claude-3-haiku-20240307-v1:0"
    }
  ]

  # Fallback model (must be one of the models in the models list above)
  fallback_model = {
    model_arn = "arn:aws:bedrock:${data.aws_region.current.name}::foundation-model/anthropic.claude-3-haiku-20240307-v1:0"
  }

  # Routing criteria based on response quality difference
  # Value must be a multiple of 5 (likely as percentage: 5, 10, 15, 20, etc.)
  routing_criteria = {
    response_quality_difference = 20
  }

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}
