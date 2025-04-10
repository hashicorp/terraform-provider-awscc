data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "codeartifact_domain_policy" {
  statement {
    sid    = "AllowAccountAccess"
    effect = "Allow"
    principals {
      type = "AWS"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      ]
    }
    actions = [
      "codeartifact:CreateRepository",
      "codeartifact:GetAuthorizationToken",
      "codeartifact:ListRepositoriesInDomain"
    ]
    resources = ["*"]
  }
}

resource "awscc_codeartifact_domain" "example" {
  domain_name                 = "example-domain"
  permissions_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.codeartifact_domain_policy.json))
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}