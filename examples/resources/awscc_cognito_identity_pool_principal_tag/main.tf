# Create a Cognito User Pool
resource "aws_cognito_user_pool" "example" {
  name = "example-user-pool"
  admin_create_user_config {
    allow_admin_create_user_only = true
  }
  password_policy {
    minimum_length    = 8
    require_lowercase = true
    require_numbers   = true
    require_symbols   = true
    require_uppercase = true
  }
}

# Create a Cognito User Pool Client
resource "aws_cognito_user_pool_client" "example" {
  name         = "example-client"
  user_pool_id = aws_cognito_user_pool.example.id
  explicit_auth_flows = [
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH"
  ]
}

# Create an identity pool
resource "aws_cognito_identity_pool" "example" {
  identity_pool_name               = "example-identity-pool"
  allow_unauthenticated_identities = false

  cognito_identity_providers {
    client_id               = aws_cognito_user_pool_client.example.id
    provider_name           = aws_cognito_user_pool.example.endpoint
    server_side_token_check = false
  }
}

# Create principal tag for the identity pool
resource "awscc_cognito_identity_pool_principal_tag" "example" {
  identity_pool_id       = aws_cognito_identity_pool.example.id
  identity_provider_name = aws_cognito_user_pool.example.endpoint
  principal_tags = jsonencode({
    "environment" = "development",
    "team"        = "engineering"
  })
  use_defaults = true
}