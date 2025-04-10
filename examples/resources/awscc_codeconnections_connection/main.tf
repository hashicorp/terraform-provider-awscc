# Create CodeConnections connection
resource "awscc_codeconnections_connection" "example" {
  connection_name = "github-connection"
  provider_type   = "GitHub"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}