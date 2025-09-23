resource "random_id" "suffix" {
  byte_length = 4
}

# IAM role for Browser Custom
resource "awscc_iam_role" "browser_role" {
  role_name = "bedrock-browser-role-${random_id.suffix.hex}"
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

resource "awscc_s3_bucket" "browser_recordings" {
  bucket_name = "bedrock-browser-recordings-${random_id.suffix.hex}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_bedrockagentcore_browser_custom" "example" {
  name               = "example_custom_browser_${random_id.suffix.hex}"
  description        = "Example Custom Browser for Bedrock Agent"
  execution_role_arn = awscc_iam_role.browser_role.arn

  network_configuration = {
    network_mode = "PUBLIC"
  }

  recording_config = {
    enabled = true
    s3_location = {
      bucket = awscc_s3_bucket.browser_recordings.bucket_name
      prefix = "browser-recordings/"
    }
  }

  tags = {
    "Modified By" = "AWSCC"
    "Environment" = "example"
  }
}

