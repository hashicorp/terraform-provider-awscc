data "aws_region" "current" {}

resource "aws_securityhub_account" "example" {}

resource "awscc_securityhub_standard" "pci_dss" {
  depends_on    = [aws_securityhub_account.example]
  standards_arn = "arn:aws:securityhub:${data.aws_region.current.name}::standards/pci-dss/v/3.2.1"
}