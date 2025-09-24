# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create a basic MediaConnect Flow first
resource "awscc_mediaconnect_flow" "example" {
  name              = "example-flow"
  availability_zone = "${data.aws_region.current.name}a"

  # Source configuration as per schema
  source = {
    description = "Example Source"
    name        = "example_source"
    protocol    = "zixi-push"
  }
}

# Create Flow Source
resource "awscc_mediaconnect_flow_source" "example" {
  flow_arn       = awscc_mediaconnect_flow.example.flow_arn
  name           = "example_source"
  description    = "Example Flow Source"
  protocol       = "zixi-push"
  whitelist_cidr = "10.0.0.0/16"
  ingest_port    = 2088

  # Example of static key encryption
  decryption = {
    algorithm  = "aes256"
    role_arn   = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/MediaConnectAccessRole"
    secret_arn = "arn:aws:secretsmanager:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:secret:example-key"
    key_type   = "static-key"
  }
}