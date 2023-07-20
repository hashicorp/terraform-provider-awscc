resource "awscc_amplify_app" "example" {
  name = "app"
  // replace with your repo URL - must also ensure Amplify has permissions to access the repo
  // GitHub instructions: https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html
  repository = "https://github.com/example/app"

  # The default build_spec added by the Amplify Console for React.
  build_spec = <<-EOT
    version: 0.1
    frontend:
      phases:
        preBuild:
          commands:
            - yarn install
        build:
          commands:
            - yarn run build
      artifacts:
        baseDirectory: build
        files:
          - '**/*'
      cache:
        paths:
          - node_modules/**/*
  EOT

  # The default rewrites and redirects added by the Amplify Console.
  custom_rules = [
    {
      source = "/<*>"
      status = "404"
      target = "/index.html"
    },
  ]
  environment_variables = [
    {
      name  = "Environment"
      value = "PROD"
    },
  ]
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
