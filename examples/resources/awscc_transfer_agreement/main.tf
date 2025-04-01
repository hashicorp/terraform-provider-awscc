# IAM role for Transfer server to access S3
data "aws_iam_policy_document" "transfer_assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["transfer.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "transfer_policy" {
  statement {
    effect = "Allow"
    actions = [
      "s3:ListBucket",
      "s3:GetBucketLocation"
    ]
    resources = ["arn:aws:s3:::example-bucket"]
  }

  statement {
    effect = "Allow"
    actions = [
      "s3:PutObject",
      "s3:GetObject",
      "s3:DeleteObject"
    ]
    resources = ["arn:aws:s3:::example-bucket/*"]
  }
}

# Create IAM role for Transfer Agreement
resource "awscc_iam_role" "transfer_role" {
  role_name                   = "transfer-agreement-role"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.transfer_assume_role.json))
}

resource "awscc_iam_role_policy" "transfer_policy" {
  policy_name     = "transfer-agreement-policy"
  role_name       = awscc_iam_role.transfer_role.role_name
  policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.transfer_policy.json))
}

# Create Transfer Server
resource "awscc_transfer_server" "sftp" {
  endpoint_type          = "PUBLIC"
  protocols              = ["SFTP"]
  identity_provider_type = "SERVICE_MANAGED"
  domain                 = "S3"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Local Profile
resource "awscc_transfer_profile" "local" {
  as_2_id      = "LOCAL_AS2_ID"
  profile_type = "LOCAL"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Partner Profile
resource "awscc_transfer_profile" "partner" {
  as_2_id      = "PARTNER_AS2_ID"
  profile_type = "PARTNER"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Transfer Agreement
resource "awscc_transfer_agreement" "example" {
  access_role        = awscc_iam_role.transfer_role.arn
  base_directory     = "/transfer/agreement"
  local_profile_id   = awscc_transfer_profile.local.profile_id
  partner_profile_id = awscc_transfer_profile.partner.profile_id
  server_id          = awscc_transfer_server.sftp.server_id

  status = "ACTIVE"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}