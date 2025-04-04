# Create IAM Role for RolesAnywhere to assume
data "aws_iam_policy_document" "trust" {
  statement {
    effect = "Allow"
    actions = [
      "sts:AssumeRole",
      "sts:TagSession",
      "sts:SetSourceIdentity"
    ]
    principals {
      type = "Service"
      identifiers = [
        "rolesanywhere.amazonaws.com"
      ]
    }
  }
}

data "aws_iam_policy_document" "permissions" {
  statement {
    effect = "Allow"
    actions = [
      "s3:ListAllMyBuckets"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "example" {
  role_name                   = "rolesanywhere-profile-example"
  assume_role_policy_document = data.aws_iam_policy_document.trust.json
  description                 = "Example role for RolesAnywhere profile"
  managed_policy_arns         = []
  max_session_duration        = 3600
  path                        = "/"
  permissions_boundary        = null
  policies = [{
    policy_document = data.aws_iam_policy_document.permissions.json
    policy_name     = "rolesanywhere-example-inline"
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create RolesAnywhere Profile
resource "awscc_rolesanywhere_profile" "example" {
  name             = "example-profile"
  role_arns        = [awscc_iam_role.example.arn]
  duration_seconds = 3600
  enabled          = true
  session_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Action = [
        "s3:ListAllMyBuckets"
      ]
      Resource = "*"
    }]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}