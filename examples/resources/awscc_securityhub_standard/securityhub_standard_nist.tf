# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

resource "aws_securityhub_account" "example" {}

resource "awscc_securityhub_standard" "nist" {
  depends_on    = [aws_securityhub_account.example]
  standards_arn = "arn:aws:securityhub:${data.aws_region.current.region}::standards/nist-800-53/v/5.0.0"
}