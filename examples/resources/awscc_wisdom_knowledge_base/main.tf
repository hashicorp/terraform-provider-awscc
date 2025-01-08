# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Example of AWS Wisdom Knowledge Base
resource "awscc_wisdom_knowledge_base" "example" {
  name                = "example-knowledge-base"
  knowledge_base_type = "CUSTOM"
  description         = "Example Wisdom knowledge base created with AWSCC provider"

  # Vector ingestion configuration with semantic chunking
  vector_ingestion_configuration = {
    chunking_configuration = {
      chunking_strategy = "SEMANTIC"
      semantic_chunking_configuration = {
        breakpoint_percentile_threshold = 85
        max_tokens                      = 512
        buffer_size                     = 0.8
      }
    }
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}