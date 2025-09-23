resource "random_id" "suffix" {
  byte_length = 4
}

resource "awscc_ecr_repository" "agent_runtime" {
  repository_name = "bedrock/agent-runtime-${random_id.suffix.hex}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role" "agent_runtime_role" {
  role_name = "bedrock-agent-runtime-role-${random_id.suffix.hex}"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "bedrock-agentcore.amazonaws.com"
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "agent_runtime_policy" {
  role_name   = awscc_iam_role.agent_runtime_role.role_name
  policy_name = "bedrock-agent-runtime-policy"
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "bedrock:InvokeModel",
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents",
          "ecr:GetAuthorizationToken",
          "ecr:BatchGetImage",
          "ecr:GetDownloadUrlForLayer"
        ]
        Resource = "*"
      }
    ]
  })
}


resource "awscc_bedrockagentcore_runtime" "example" {
  agent_runtime_name = "example_agent_runtime_${random_id.suffix.hex}"
  description        = "Example Bedrock Agent Runtime"
  role_arn           = awscc_iam_role.agent_runtime_role.arn

  agent_runtime_artifact = {
    container_configuration = {
      container_uri = "${awscc_ecr_repository.agent_runtime.repository_uri}:latest"
    }
  }

  network_configuration = {
    network_mode = "PUBLIC" # will be modified to use VPC constraints once available
  }

  environment_variables = {
    "LOG_LEVEL" = "INFO"
  }

  tags = {
    "Modified By" = "AWSCC"
    "Environment" = "example"
  }
}

# Runtime Endpoint
resource "awscc_bedrockagentcore_runtime_endpoint" "example" {
  name             = "example_runtime_endpoint_${random_id.suffix.hex}"
  description      = "Example Bedrock Agent Runtime Endpoint"
  agent_runtime_id = awscc_bedrockagentcore_runtime.example.agent_runtime_id

  tags = {
    "Modified By" = "AWSCC"
    "Environment" = "example"
  }
}