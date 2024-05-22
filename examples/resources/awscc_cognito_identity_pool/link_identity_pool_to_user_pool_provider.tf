resource "aws_cognito_user_pool" "example_user_pool" {
  name = "example-user-pool"
}

resource "aws_cognito_user_pool_client" "example_user_pool_client" {
  name         = "example-user-pool-client"
  user_pool_id = aws_cognito_user_pool.example_user_pool.id
}

resource "awscc_cognito_identity_pool" "example_identity_pool" {
  identity_pool_name               = "example-identity-pool"
  allow_unauthenticated_identities = false //regardless of whether this is true or not, this requires configuration of aws_cognito_identity_pool_roles_attachment

  cognito_identity_providers = [{
    client_id     = aws_cognito_user_pool_client.example_user_pool_client.id
    provider_name = aws_cognito_user_pool.example_user_pool.endpoint
  }]
}