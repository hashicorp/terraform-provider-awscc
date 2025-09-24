# Get current account ID for dynamic naming
data "aws_caller_identity" "current" {}

# Create the Cognito User Pool
resource "aws_cognito_user_pool" "example" {
  name = "my-user-pool"

  auto_verified_attributes = ["email"]
  username_attributes      = ["email"]

  verification_message_template {
    default_email_option = "CONFIRM_WITH_CODE"
  }

  admin_create_user_config {
    allow_admin_create_user_only = false
  }

  email_configuration {
    email_sending_account = "COGNITO_DEFAULT"
  }

  tags = {
    "Modified By" = "AWS"
  }
}

# Create the User Pool Domain
resource "awscc_cognito_user_pool_domain" "example" {
  domain       = "my-example-domain-${data.aws_caller_identity.current.account_id}"
  user_pool_id = aws_cognito_user_pool.example.id
}