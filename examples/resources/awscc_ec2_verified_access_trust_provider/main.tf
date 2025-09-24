# Create an OpenID Connect provider example
resource "awscc_ec2_verified_access_trust_provider" "example" {
  trust_provider_type      = "user"
  user_trust_provider_type = "oidc"
  policy_reference_name    = "example_trust_provider"
  description              = "Example OIDC Trust Provider"

  oidc_options = {
    issuer                 = "https://accounts.google.com"
    client_id              = "your-client-id"
    client_secret          = "your-client-secret"
    authorization_endpoint = "https://accounts.google.com/o/oauth2/v2/auth"
    token_endpoint         = "https://oauth2.googleapis.com/token"
    user_info_endpoint     = "https://openidconnect.googleapis.com/v1/userinfo"
    scope                  = "openid email profile"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}