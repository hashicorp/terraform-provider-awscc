# Create a Cognito User Pool
resource "aws_cognito_user_pool" "example" {
  name = "example-pool"

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create a User Pool Resource Server
resource "awscc_cognito_user_pool_resource_server" "example" {
  identifier   = "https://api.example.com"
  name         = "example-resource-server"
  user_pool_id = aws_cognito_user_pool.example.id

  scopes = [
    {
      scope_name        = "read"
      scope_description = "Read access"
    },
    {
      scope_name        = "write"
      scope_description = "Write access"
    }
  ]
}