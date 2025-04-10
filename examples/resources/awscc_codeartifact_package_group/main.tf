# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create a CodeArtifact domain first
resource "awscc_codeartifact_domain" "example" {
  domain_name = "example-domain"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the package group
resource "awscc_codeartifact_package_group" "example" {
  domain_name  = awscc_codeartifact_domain.example.domain_name
  domain_owner = data.aws_caller_identity.current.account_id
  pattern      = "/npm/*"
  description  = "Example package group for testing"

  origin_configuration = {
    restrictions = {
      publish = {
        restriction_mode = "ALLOW"
        repositories     = []
      }
      external_upstream = {
        restriction_mode = "ALLOW"
        repositories     = []
      }
      internal_upstream = {
        restriction_mode = "ALLOW"
        repositories     = []
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}