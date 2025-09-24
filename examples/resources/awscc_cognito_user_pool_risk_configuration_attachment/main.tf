# Create a Cognito User Pool
resource "aws_cognito_user_pool" "example" {
  name = "example-user-pool"

  admin_create_user_config {
    allow_admin_create_user_only = true
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create a Cognito User Pool Client using AWSCC
resource "awscc_cognito_user_pool_client" "example" {
  user_pool_id = aws_cognito_user_pool.example.id
  client_name  = "example-client"
  explicit_auth_flows = [
    "ALLOW_REFRESH_TOKEN_AUTH",
    "ALLOW_USER_SRP_AUTH"
  ]

  prevent_user_existence_errors = "ENABLED"
  read_attributes               = ["email"]
  write_attributes              = ["email"]
}

# Configure the risk configuration attachment
resource "awscc_cognito_user_pool_risk_configuration_attachment" "example" {
  user_pool_id = aws_cognito_user_pool.example.id
  client_id    = awscc_cognito_user_pool_client.example.id

  account_takeover_risk_configuration = {
    actions = {
      high_action = {
        event_action = "BLOCK"
        notify       = false
      }
      medium_action = {
        event_action = "MFA_IF_CONFIGURED"
        notify       = false
      }
      low_action = {
        event_action = "NO_ACTION"
        notify       = false
      }
    }
  }

  compromised_credentials_risk_configuration = {
    actions = {
      event_action = "BLOCK"
    }
    event_filter = ["SIGN_IN", "PASSWORD_CHANGE"]
  }

  risk_exception_configuration = {
    blocked_ip_range_list = ["10.0.0.0/24"]
    skipped_ip_range_list = ["192.168.0.0/24"]
  }
}