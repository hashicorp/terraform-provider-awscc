resource "awscc_grafana_workspace" "example" {
  account_access_type      = "CURRENT_ACCOUNT"
  name                     = "example"
  role_arn                 = awscc_iam_role.example.arn
  description              = "example workspace"
  authentication_providers = ["SAML"]
  permission_type          = "CUSTOMER_MANAGED"
  grafana_version          = "9.4"

  saml_configuration = {
    idp_metadata = {
      xml = "<md:EntityDescriptor xmlns:md='urn:oasis:names:tc:SAML:2.0:metadata' entityID='entityId'>DATA</md:EntityDescriptor>"
    }
    assertion_attributes = {
      name   = "displayName"
      login  = "login"
      email  = "email"
      groups = "groups"
      org    = "org"
    }
    role_values = {
      editor = ["editor1"]
      admin  = ["admin1"]
    }
    allowed_organizations   = ["org1"]
    login_validity_duration = 60
  }
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
      },
    ]
  })
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
