resource "awscc_amplify_app" "example" {
  name       = "app"
  repository = "https://github.com/example/app"

  # GitHub personal access token
  access_token = "..."
}
