resource "awscc_acmpca_certificate_authority_activation" "subordinate" {
  certificate               = awscc_acmpca_certificate.subordinate.certificate
  certificate_authority_arn = awscc_acmpca_certificate_authority.subordinate.arn
  certificate_chain         = awscc_acmpca_certificate_authority_activation.root.complete_certificate_chain
}

resource "awscc_acmpca_certificate" "subordinate" {
  certificate_authority_arn   = awscc_acmpca_certificate_authority.root.arn
  certificate_signing_request = awscc_acmpca_certificate_authority.subordinate.certificate_signing_request
  signing_algorithm           = "SHA256WITHRSA"

  template_arn = "arn:${data.aws_partition.current.partition}:acm-pca:::template/SubordinateCACertificate_PathLen3/V1"

  validity = {
    type  = "YEARS"
    value = 1
  }
  depends_on = [awscc_acmpca_certificate_authority_activation.root]
}

resource "awscc_acmpca_certificate_authority" "subordinate" {
  key_algorithm     = "RSA_4096"
  signing_algorithm = "SHA512WITHRSA"
  type              = "SUBORDINATE"
  subject = {
    common_name = "sub.example.com"
  }
}

resource "awscc_acmpca_certificate_authority_activation" "root" {
  certificate               = awscc_acmpca_certificate.root.certificate
  certificate_authority_arn = awscc_acmpca_certificate_authority.root.arn
}

resource "awscc_acmpca_certificate_authority" "root" {
  key_algorithm     = "RSA_4096"
  signing_algorithm = "SHA512WITHRSA"
  type              = "ROOT"
  subject = {
    common_name = "example.com"
  }
}

resource "awscc_acmpca_certificate" "root" {
  certificate_authority_arn   = awscc_acmpca_certificate_authority.root.arn
  certificate_signing_request = awscc_acmpca_certificate_authority.root.certificate_signing_request
  signing_algorithm           = "SHA256WITHRSA"

  template_arn = "arn:${data.aws_partition.current.partition}:acm-pca:::template/RootCACertificate/V1"

  validity = {
    type  = "YEARS"
    value = 10
  }
}


data "aws_partition" "current" {}
