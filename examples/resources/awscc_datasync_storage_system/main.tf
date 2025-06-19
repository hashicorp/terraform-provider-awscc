# Get current AWS region
# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create the DataSync Storage System
resource "awscc_datasync_storage_system" "example" {
  agent_arns = ["arn:aws:datasync:${data.aws_region.current.region}:${data.aws_caller_identity.current.account_id}:agent/agent-example"]

  server_configuration = {
    server_hostname = "storage.example.com"
    server_port     = 443
  }

  server_credentials = {
    username = "admin"
    password = "example-password"
  }

  system_type = "NetAppONTAP"
  name        = "example-storage-system"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}