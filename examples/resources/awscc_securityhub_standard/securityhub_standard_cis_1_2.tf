resource "aws_securityhub_account" "example" {
  enable_default_standards = false
}

resource "awscc_securityhub_standard" "cis_1_2" {
  depends_on    = [aws_securityhub_account.example]
  standards_arn = "arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0"
}