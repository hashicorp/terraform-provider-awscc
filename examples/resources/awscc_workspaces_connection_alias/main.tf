# Create a WorkSpaces Connection Alias
resource "awscc_workspaces_connection_alias" "example" {
  # The connection string in format: wsca-<alias>.<domain>
  connection_string = "wsca-example.wellsiau.com"

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    },
    {
      key   = "Environment"
      value = "Example"
    }
  ]
}