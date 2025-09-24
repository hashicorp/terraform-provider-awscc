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

# Repository permissions policy document
data "aws_iam_policy_document" "repo_policy" {
  statement {
    sid    = "AllowPublishAndRead"
    effect = "Allow"
    principals {
      type = "AWS"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      ]
    }
    actions = [
      "codeartifact:PublishPackageVersion",
      "codeartifact:ReadFromRepository"
    ]
    resources = ["*"]
  }
}

# Create CodeArtifact repository
resource "awscc_codeartifact_repository" "example" {
  repository_name = "example-repo"
  domain_name     = awscc_codeartifact_domain.example.domain_name
  description     = "Example CodeArtifact Repository"

  permissions_policy_document = data.aws_iam_policy_document.repo_policy.json

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}