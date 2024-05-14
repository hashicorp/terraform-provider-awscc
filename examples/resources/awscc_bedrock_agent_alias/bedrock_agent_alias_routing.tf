resource "awscc_bedrock_agent_alias" "example" {
  agent_alias_name = "example"
  agent_id         = var.bedrock_agent_id
  description      = "Example alias for the Bedrock agent"
  routing_configuration = [
    {
      agent_version = var.bedrock_agent_version
    }
  ]

  tags = {
    "Modified By" = "AWSCC"
  }
}

variable "bedrock_agent_id" {
  type = string
}

variable "bedrock_agent_version" {
  type = string
}
