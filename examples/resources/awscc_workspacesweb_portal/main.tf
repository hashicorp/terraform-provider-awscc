data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

resource "awscc_workspacesweb_portal" "example" {
  authentication_type     = "Standard"
  display_name           = "Demo Web Portal"
  instance_type          = "standard.regular"
  max_concurrent_sessions = 10

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}