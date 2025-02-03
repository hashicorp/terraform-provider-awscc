data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# IAM role for Transfer Connector
data "aws_iam_policy_document" "transfer_connector_assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["transfer.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "transfer_connector_policy" {
  statement {
    effect = "Allow"
    actions = [
      "s3:*"
    ]
    resources = ["*"]
  }
}

# IAM role for Transfer Connector access
resource "awscc_iam_role" "transfer_connector_role" {
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.transfer_connector_assume_role.json))
  description                 = "Role for Transfer Connector"
  max_session_duration        = 3600
  path                        = "/"
  role_name                   = "transfer-connector-role"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "transfer_connector_policy" {
  policy_name     = "transfer-connector-policy"
  policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.transfer_connector_policy.json))
  role_name       = awscc_iam_role.transfer_connector_role.role_name
}

# Example SFTP secrets for user authentication
resource "awscc_secretsmanager_secret" "sftp_user_secret" {
  name = "transfer-connector-sftp-user-secret-1234567890"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "aws_secretsmanager_secret_version" "sftp_user_secret" {
  secret_id = awscc_secretsmanager_secret.sftp_user_secret.id
  secret_string = jsonencode({
    password = "YourSecretPassword123!"
  })
}

# Transfer Connector
resource "awscc_transfer_connector" "example" {
  access_role = awscc_iam_role.transfer_connector_role.arn
  url         = "sftp://example.com"

  sftp_config = {
    user_secret_id    = "arn:aws:secretsmanager:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:secret:${awscc_secretsmanager_secret.sftp_user_secret.name}"
    trusted_host_keys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC3example"]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}