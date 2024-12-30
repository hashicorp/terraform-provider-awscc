resource "awscc_acmpca_certificate_authority_activation" "example" {
  certificate               = awscc_acmpca_certificate.example.certificate
  certificate_authority_arn = awscc_acmpca_certificate_authority.example.arn
}

resource "awscc_acmpca_certificate" "example" {
  certificate_authority_arn   = awscc_acmpca_certificate_authority.example.arn
  certificate_signing_request = awscc_acmpca_certificate_authority.example.certificate_signing_request
  signing_algorithm           = "SHA256WITHRSA"

  template_arn = "arn:${data.aws_partition.current.partition}:acm-pca:::template/SubordinateCACertificate_PathLen0/V1"

  validity = {
    type  = "YEARS"
    value = 5
  }
}

resource "awscc_acmpca_certificate_authority" "example" {
  key_algorithm     = "RSA_4096"
  signing_algorithm = "SHA512WITHRSA"
  type              = "SUBORDINATE"
  subject = {
    common_name = "example.com"
  }
}

data "aws_partition" "current" {}
