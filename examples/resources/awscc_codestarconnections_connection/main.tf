# Configure required data sources
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create the CodeStar connection
resource "awscc_codestarconnections_connection" "example" {
  connection_name = "github-connection-example"
  provider_type   = "GitHub"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Output useful information
output "connection_arn" {
  value = awscc_codestarconnections_connection.example.connection_arn
}

output "connection_status" {
  value = awscc_codestarconnections_connection.example.connection_status
}