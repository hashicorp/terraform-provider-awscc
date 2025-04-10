# IAM Service Role for AppStream Directory Config
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["appstream.amazonaws.com"]
    }
  }
}

# Service Role for AppStream Directory Config
resource "awscc_iam_role" "appstream_directory" {
  role_name                   = "appstream-directory-config-role"
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Policy for Directory Access
data "aws_iam_policy_document" "directory_policy" {
  statement {
    effect = "Allow"
    actions = [
      "ds:CreateComputer",
      "ds:DescribeDirectories",
      "ds:DeleteComputer"
    ]
    resources = ["*"]
  }
}

# Attach policy to the role
resource "awscc_iam_role_policy" "directory_policy" {
  policy_name     = "directory-access-policy"
  role_name       = awscc_iam_role.appstream_directory.role_name
  policy_document = data.aws_iam_policy_document.directory_policy.json
}

# AppStream Directory Config
resource "awscc_appstream_directory_config" "example" {
  directory_name                          = "corp.example.com"
  organizational_unit_distinguished_names = ["OU=AppStream,DC=corp,DC=example,DC=com"]
  service_account_credentials = {
    account_name     = "ServiceAccount"
    account_password = "YourSecurePassword123!"
  }
}