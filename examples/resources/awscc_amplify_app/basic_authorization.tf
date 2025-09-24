resource "awscc_amplify_app" "example" {
  name = "app"

  basic_auth_config = {
    enable_basic_auth = true
    username          = "your-username"
    password          = "your-password"
  }

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
