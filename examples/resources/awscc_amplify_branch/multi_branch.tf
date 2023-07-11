resource "awscc_amplify_app" "example" {
  name = "app"
}

resource "awscc_amplify_branch" "main" {
  app_id      = awscc_amplify_app.example.app_id
  branch_name = "main"

  framework = "React"
  stage     = "PRODUCTION"

  environment_variables = [
    {
      name  = "REACT_APP_API_SERVER"
      value = "https://api.example.com"
    },
    {
      name  = "Environment"
      value = "PROD"
    },
  ]
}

resource "awscc_amplify_branch" "dev" {
  app_id      = awscc_amplify_app.example.app_id
  branch_name = "main"

  framework = "React"
  stage     = "DEVELOPMENT"

  environment_variables = [
    {
      name  = "REACT_APP_API_SERVER"
      value = "https://dev.api.example.com"
    },
    {
      name  = "Environment"
      value = "DEV"
    },
  ]
}
