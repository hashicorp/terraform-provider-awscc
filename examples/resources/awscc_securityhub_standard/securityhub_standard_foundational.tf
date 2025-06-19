# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

resource "aws_securityhub_account" "example" {
  enable_default_standards = false
}

resource "awscc_securityhub_standard" "foundational" {
  depends_on    = [aws_securityhub_account.example]
  standards_arn = "arn:aws:securityhub:${data.aws_region.current.region}::standards/aws-foundational-security-best-practices/v/1.0.0"
}