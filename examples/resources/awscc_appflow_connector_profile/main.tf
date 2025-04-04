# Get the ARN for Secrets Manager secret
data "aws_secretsmanager_secret" "salesforce_credentials_arn" {
  depends_on = [awscc_secretsmanager_secret.salesforce_credentials]
  name       = "appflow/salesforce/credentials"
}

# IAM role for AppFlow connector profile
resource "awscc_iam_role" "appflow_connector_profile" {
  role_name = "appflow-salesforce-connector-profile"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "appflow.amazonaws.com"
        }
      }
    ]
  })
  description = "IAM role for AppFlow Salesforce connector profile"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Secret for Salesforce client credentials
resource "awscc_secretsmanager_secret" "salesforce_credentials" {
  name        = "appflow/salesforce/credentials"
  description = "Salesforce client credentials for AppFlow"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Secret versions still need to use standard AWS provider
resource "aws_secretsmanager_secret_version" "access_token" {
  secret_id     = awscc_secretsmanager_secret.salesforce_credentials.id
  secret_string = "your-salesforce-access-token" # Replace with actual token
}

resource "aws_secretsmanager_secret_version" "refresh_token" {
  secret_id     = awscc_secretsmanager_secret.salesforce_credentials.id
  secret_string = "your-salesforce-refresh-token" # Replace with actual token
}

# AppFlow connector profile
resource "awscc_appflow_connector_profile" "example" {
  connector_profile_name = "example-salesforce-profile"
  connection_mode        = "Public"
  connector_type         = "Salesforce"

  connector_profile_config = {
    connector_profile_credentials = {
      salesforce = {
        client_credentials_arn = data.aws_secretsmanager_secret.salesforce_credentials_arn.arn
        access_token           = aws_secretsmanager_secret_version.access_token.secret_string
        refresh_token          = aws_secretsmanager_secret_version.refresh_token.secret_string
      }
    }
    connector_profile_properties = {
      salesforce = {
        instance_url           = "https://example.my.salesforce.com"
        is_sandbox_environment = false
      }
    }
  }
}