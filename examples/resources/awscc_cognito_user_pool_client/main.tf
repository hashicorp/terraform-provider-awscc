# Create a Cognito User Pool
resource "aws_cognito_user_pool" "example" {
  name = "example-user-pool"

  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 1
    }
  }

  auto_verified_attributes = ["email"]

  password_policy {
    minimum_length                   = 8
    require_lowercase                = true
    require_numbers                  = true
    require_symbols                  = true
    require_uppercase                = true
    temporary_password_validity_days = 7
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create a Cognito User Pool Client
resource "awscc_cognito_user_pool_client" "example" {
  user_pool_id = aws_cognito_user_pool.example.id
  client_name  = "example-client"

  # Define allowed OAuth flows
  allowed_o_auth_flows                  = ["implicit", "code"]
  allowed_o_auth_flows_user_pool_client = true
  allowed_o_auth_scopes                 = ["openid", "email", "profile"]

  # Define callback and logout URLs
  callback_ur_ls       = ["https://example.com/callback"]
  logout_ur_ls         = ["https://example.com/logout"]
  default_redirect_uri = "https://example.com/callback"

  # Define explicit authentication flows
  explicit_auth_flows = [
    "ALLOW_USER_SRP_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH",
    "ALLOW_USER_PASSWORD_AUTH"
  ]

  # Token configuration
  refresh_token_validity = 30
  access_token_validity  = 1
  id_token_validity      = 1

  token_validity_units = {
    access_token  = "hours"
    id_token      = "hours"
    refresh_token = "days"
  }

  # Define read and write attributes
  read_attributes  = ["email", "email_verified", "name"]
  write_attributes = ["email", "name"]

  generate_secret = false

  prevent_user_existence_errors = "ENABLED"

  supported_identity_providers = ["COGNITO"]
}