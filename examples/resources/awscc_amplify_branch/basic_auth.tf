resource "awscc_amplify_app" "example" {
  name = "app"
}

resource "awscc_amplify_branch" "main" {
  app_id      = awscc_amplify_app.example.app_id
  branch_name = "main"

  # Used to restrict access to your branches with a username and password
  basic_auth_config = {
    enable_basic_auth = true
    username          = "your-username"
    password          = "your-password"
  }
}
