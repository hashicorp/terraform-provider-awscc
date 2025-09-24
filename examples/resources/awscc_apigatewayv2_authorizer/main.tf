# Get current AWS region
data "aws_region" "current" {}

# Create a Cognito User Pool for JWT authorization
resource "aws_cognito_user_pool" "example" {
  name = "example-user-pool"

  password_policy {
    minimum_length = 8
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create an HTTP API
resource "awscc_apigatewayv2_api" "example" {
  name          = "example-jwt-auth-api"
  protocol_type = "HTTP"
  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create a JWT authorizer
resource "awscc_apigatewayv2_authorizer" "example" {
  api_id          = awscc_apigatewayv2_api.example.id
  authorizer_type = "JWT"
  name            = "example-jwt-authorizer"

  jwt_configuration = {
    audience = ["example-app"]
    issuer   = "https://cognito-idp.${data.aws_region.current.name}.amazonaws.com/${aws_cognito_user_pool.example.id}"
  }

  identity_source = ["$request.header.Authorization"]
}