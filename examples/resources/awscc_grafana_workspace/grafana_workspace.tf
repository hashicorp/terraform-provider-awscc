resource "awscc_grafana_workspace" "example" {
  account_access_type      = "CURRENT_ACCOUNT"
  name                     = "example"
  role_arn                 = awscc_iam_role.example.arn
  description              = "example workspace"
  authentication_providers = ["SAML"]
  permission_type          = "SERVICE_MANAGED"
  grafana_version          = "9.4"
}

resource "awscc_iam_role" "example" {
  role_name           = "example"
  description         = "Grafana role"
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AmazonGrafanaAthenaAccess"]
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "grafana.amazonaws.com"
        }
      }
    ]
  })
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}


