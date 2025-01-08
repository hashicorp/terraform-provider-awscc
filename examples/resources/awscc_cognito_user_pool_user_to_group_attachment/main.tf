# Create a Cognito User Pool
resource "aws_cognito_user_pool" "example" {
  name = "example-user-pool"

  auto_verified_attributes = ["email"]
  username_attributes      = ["email"]

  password_policy {
    minimum_length    = 8
    require_lowercase = true
    require_numbers   = true
    require_symbols   = true
    require_uppercase = true
  }

  schema {
    attribute_data_type = "String"
    mutable             = true
    name                = "email"
    required            = true

    string_attribute_constraints {
      max_length = "2048"
      min_length = "0"
    }
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create a Cognito User Pool Group
resource "aws_cognito_user_group" "example" {
  name         = "example-group"
  user_pool_id = aws_cognito_user_pool.example.id
  description  = "Example user pool group"
}

# Create a Cognito User
resource "aws_cognito_user" "example" {
  user_pool_id = aws_cognito_user_pool.example.id
  username     = "example@example.com"

  attributes = {
    email          = "example@example.com"
    email_verified = "true"
  }
}

# Attach the user to the group
resource "awscc_cognito_user_pool_user_to_group_attachment" "example" {
  group_name   = aws_cognito_user_group.example.name
  user_pool_id = aws_cognito_user_pool.example.id
  username     = aws_cognito_user.example.username
}