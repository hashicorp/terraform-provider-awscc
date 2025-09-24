# Create an EventBridge connection
resource "awscc_events_connection" "example" {
  name               = "example-connection"
  authorization_type = "BASIC"
  description        = "Example connection for API destination"
  auth_parameters = {
    basic_auth_parameters = {
      username = "test-user"
      password = "test-password123!"
    }
  }
}

# Create the API destination
resource "awscc_events_api_destination" "example" {
  name                             = "example-api-destination"
  description                      = "Example API destination created with AWSCC"
  connection_arn                   = awscc_events_connection.example.arn
  http_method                      = "POST"
  invocation_endpoint              = "https://api.example.com/endpoint"
  invocation_rate_limit_per_second = 300
}