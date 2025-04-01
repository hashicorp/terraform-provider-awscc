# Create the AppSync API first
resource "aws_appsync_graphql_api" "example" {
  name                = "example-channel-namespace-api"
  authentication_type = "API_KEY"

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Add a delay to ensure the API is ready
resource "time_sleep" "wait_30_seconds" {
  depends_on      = [aws_appsync_graphql_api.example]
  create_duration = "30s"
}

# Create the channel namespace
resource "awscc_appsync_channel_namespace" "example" {
  depends_on = [time_sleep.wait_30_seconds]
  name       = "example-namespace"
  api_id     = aws_appsync_graphql_api.example.id

  # Example of publish and subscribe auth modes
  publish_auth_modes = [{
    auth_type = "API_KEY"
  }]

  subscribe_auth_modes = [{
    auth_type = "API_KEY"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}