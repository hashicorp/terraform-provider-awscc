# Create the CodeStar Connection first
resource "awscc_codestarconnections_connection" "example" {
  connection_name = "example-connection"
  provider_type   = "GitHub"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Repository Link
resource "awscc_codestarconnections_repository_link" "example" {
  connection_arn  = awscc_codestarconnections_connection.example.connection_arn
  owner_id        = "example-owner"
  repository_name = "example-repo"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}