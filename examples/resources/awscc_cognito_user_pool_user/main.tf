# Create a Cognito User Pool using AWS provider
resource "aws_cognito_user_pool" "example" {
  name = "example-user-pool"

  auto_verified_attributes = ["email"]
  username_attributes      = ["email"]
  admin_create_user_config {
    allow_admin_create_user_only = false
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create a Cognito User Pool User using AWSCC provider
resource "awscc_cognito_user_pool_user" "example" {
  user_pool_id = aws_cognito_user_pool.example.id
  username     = "example@example.com"

  desired_delivery_mediums = ["EMAIL"]
  force_alias_creation     = false
  message_action           = "SUPPRESS"

  user_attributes = [
    {
      name  = "email"
      value = "example@example.com"
    },
    {
      name  = "email_verified"
      value = "true"
    }
  ]

  client_metadata = {
    created_by = "terraform"
  }
}