# Get current region
data "aws_region" "current" {}

# Get current account ID
data "aws_caller_identity" "current" {}

# Create a secret for SMTP credentials
resource "awscc_secretsmanager_secret" "smtp_secret" {
  name = "smtp-relay-credentials"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

resource "aws_secretsmanager_secret_version" "smtp_secret" {
  secret_id = awscc_secretsmanager_secret.smtp_secret.id
  secret_string = jsonencode({
    username = "smtp-user"
    password = "example-password"
  })
}

# Example SES mail manager relay resource
resource "awscc_ses_mail_manager_relay" "example" {
  relay_name  = "example-relay"
  server_name = "smtp.example.com"
  server_port = 587

  authentication = {
    secret_arn = format("arn:aws:secretsmanager:%s:%s:secret:%s", data.aws_region.current.name, data.aws_caller_identity.current.account_id, awscc_secretsmanager_secret.smtp_secret.id)
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}