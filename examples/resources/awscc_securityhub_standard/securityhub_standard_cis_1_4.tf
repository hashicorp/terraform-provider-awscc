data "aws_region" "current" {}

resource "aws_securityhub_account" "example" {}

resource "awscc_securityhub_standard" "cis_1_4" {
  depends_on    = [aws_securityhub_account.example]
  standards_arn = "arn:aws:securityhub:${data.aws_region.current.name}::standards/cis-aws-foundations-benchmark/v/1.4.0"
}