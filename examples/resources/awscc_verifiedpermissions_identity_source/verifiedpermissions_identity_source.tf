# Create a Cognito User Pool as the identity source
resource "aws_cognito_user_pool" "example" {
  name = "example-user-pool"
}

# Create a Cognito User Pool Client
resource "aws_cognito_user_pool_client" "example" {
  name            = "example-client"
  user_pool_id    = aws_cognito_user_pool.example.id
  generate_secret = false
  explicit_auth_flows = [
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH"
  ]
}

# Create a Verified Permissions Policy Store
resource "awscc_verifiedpermissions_policy_store" "example" {
  validation_settings = {
    mode = "OFF" # Disable schema validation for flexibility
  }
  tags = [
    {
      key   = "Environment"
      value = "example" # Environment designation
    },
    {
      key   = "Name"
      value = "example-policy-store" # Resource identifier
    }
  ]
}

# Create a Verified Permissions Identity Source using Cognito
resource "awscc_verifiedpermissions_identity_source" "example" {
  configuration = {
    cognito_user_pool_configuration = {
      user_pool_arn = aws_cognito_user_pool.example.arn
      client_ids    = [aws_cognito_user_pool_client.example.id] # Associated client IDs
    }
  }
  policy_store_id       = awscc_verifiedpermissions_policy_store.example.id
  principal_entity_type = "User" # Entity type for authorization principals
}
