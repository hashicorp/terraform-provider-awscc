resource "awscc_amplify_app" "example" {
  name = "app"

  auto_branch_creation_config = {
    # Enable auto branch creation
    enable_auto_branch_creation = true
    # Enable auto build for the created branch.
    enable_auto_build = true

    # The default patterns added by the Amplify Console.
    auto_branch_creation_patterns = [
      "*",
      "*/**",
    ]
  }
}
