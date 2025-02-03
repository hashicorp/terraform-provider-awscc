# Data sources to get AWS account ID and region
data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

# Create CodeConnections connection
resource "awscc_codeconnections_connection" "example" {
  connection_name = "github-connection"
  provider_type   = "GitHub"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}