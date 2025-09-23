resource "random_id" "suffix" {
  byte_length = 4
}

resource "awscc_iam_role" "example" {
  role_name = "bedrock-code-interpreter-role-${random_id.suffix.hex}"
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

resource "awscc_bedrockagentcore_code_interpreter_custom" "example" {
  name               = "example_code_interpreter_${random_id.suffix.hex}"
  description        = "Example Custom Code Interpreter for Bedrock Agent"
  execution_role_arn = awscc_iam_role.example.arn

  network_configuration = {
    network_mode = "SANDBOX" # will be modified to use VPC constraints once available
  }

  tags = {
    "Modified By" = "AWSCC"
    "Environment" = "example"
  }
}


