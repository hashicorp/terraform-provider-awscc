# Create an IAM role for Redshift Serverless
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"

    principals {
      type        = "Service"
      identifiers = ["redshift-serverless.amazonaws.com"]
    }
  }
}

resource "awscc_iam_role" "redshift_role" {
  role_name                   = "RedshiftServerlessNamespaceRole"
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json

  policies = [{
    policy_name = "RedshiftServerlessNamespacePolicy"
    policy_document = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Effect = "Allow"
          Action = [
            "s3:GetBucketLocation",
            "s3:GetObject",
            "s3:ListBucket"
          ]
          Resource = [
            "arn:aws:s3:::*",
            "arn:aws:s3:::*/*"
          ]
        }
      ]
    })
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Redshift Serverless namespace
resource "awscc_redshiftserverless_namespace" "example" {
  namespace_name      = "example-namespace"
  db_name             = "exampledb"
  admin_username      = "admin"
  admin_user_password = "Admin123456789"

  default_iam_role_arn = awscc_iam_role.redshift_role.arn
  iam_roles            = [awscc_iam_role.redshift_role.arn]

  log_exports = ["userlog", "connectionlog"]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
    }, {
    key   = "Environment"
    value = "Test"
  }]
}