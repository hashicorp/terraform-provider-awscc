resource "awscc_amplify_domain" "example" {
  app_id      = awscc_amplify_app.example.app_id
  domain_name = "example.com"

  sub_domain_settings = [
    {
      # https://example.com
      branch_name = awscc_amplify_branch.main.branch_name
      prefix      = ""
    },
    {
      # https://www.example.com
      branch_name = awscc_amplify_branch.main.branch_name
      prefix      = "www"
    },
  ]
}

resource "awscc_amplify_app" "example" {
  name = "app"

  # Setup redirect from https://example.com to https://www.example.com
  custom_rules = [
    {
      source = "https://example.com"
      status = "302"
      target = "https://www.example.com"
    },
  ]
}

resource "awscc_amplify_branch" "main" {
  app_id      = awscc_amplify_app.example.app_id
  branch_name = "main"
}


