resource "awscc_events_connection" "auth0_connection" {
  name               = "auth0-connection"
  authorization_type = "OAUTH_CLIENT_CREDENTIALS"

  auth_parameters = {
    o_auth_parameters = {
      client_parameters = {
        client_id     = "my-client-id"
        client_secret = "my-secret-string"
      }

      authorization_endpoint = "https://yourUserName.us.auth0.com/oauth/token"
      http_method            = "POST"

      o_auth_http_parameters = {
        body_parameters = [
          {
            key   = "audience",
            value = "my-auth0-identifier"
          }
        ]
      }
    }
  }
}