resource "awscc_bedrock_agent" "example" {
  agent_name              = "example-agent"
  description             = "Example agent configuration"
  agent_resource_role_arn = var.agent_role_arn
  foundation_model        = "anthropic.claude-v2:1"
  instruction             = "You are an office assistant in an insurance agency. You are friendly and polite. You help with managing insurance claims and coordinating pending paperwork."
  knowledge_bases = [{
    description          = "example knowledge base"
    knowledge_base_id    = var.knowledge_base_id
    knowledge_base_state = "ENABLED"
  }]

  customer_encryption_key_arn = var.kms_key_arn
  idle_session_ttl_in_seconds = 600
  auto_prepare                = true

  action_groups = [{
    action_group_name = "example-action-group"
    description       = "Example action group"
    api_schema = {
      s3 = {
        s3_bucket_name = var.bucket_name
        s3_object_key  = var.bucket_object_key
      }
    }
    action_group_executor = {
      lambda = var.lambda_arn
    }

  }]

  tags = {
    "Modified By" = "AWSCC"
  }

}

variable "bucket_name" {
  type = string
}

variable "bucket_object_key" {
  type = string
}

variable "agent_role_arn" {
  type = string
}

variable "knowledge_base_id" {
  type = string
}

variable "kms_key_arn" {
  type = string
}

variable "lambda_arn" {
  type = string
}
