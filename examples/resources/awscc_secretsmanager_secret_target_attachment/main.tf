data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create Secret
resource "awscc_secretsmanager_secret" "example" {
  name = "example-secret"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Secret Version with database credentials
resource "aws_secretsmanager_secret_version" "example" {
  secret_id = awscc_secretsmanager_secret.example.id
  secret_string = jsonencode({
    engine   = "mysql"
    host     = "YOUR_RDS_HOST",
    password = "YOUR_RDS_PASSWORD",
    port     = 3306,
    username = "YOUR_RDS_USERNAME"
  })
}