data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

data "aws_iam_policy_document" "pca_admin_policy" {
  statement {
    actions = [
      "acm-pca:IssueCertificate",
      "acm-pca:GetCertificate",
      "acm-pca:ListPermissions",
      "acm-pca:ListTags"
    ]
    resources = ["arn:aws:acm-pca:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:certificate-authority/*"]
    principals {
      type = "AWS"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      ]
    }
  }
}

# Create the Private Certificate Authority
resource "awscc_acmpca_certificate_authority" "example" {
  type              = "ROOT"
  key_algorithm     = "RSA_2048"
  signing_algorithm = "SHA256WITHRSA"

  subject = {
    common_name         = "Example Corp Root CA"
    country             = "US"
    organization        = "Example Corp"
    organizational_unit = "IT"
    state_or_province   = "WA"
    locality            = "Seattle"
  }

  revocation_configuration = {
    crl_configuration = {
      enabled = false
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Generate a CSR for the root CA
resource "awscc_acmpca_certificate" "root_ca" {
  certificate_authority_arn   = awscc_acmpca_certificate_authority.example.arn
  certificate_signing_request = awscc_acmpca_certificate_authority.example.certificate_signing_request
  signing_algorithm           = "SHA256WITHRSA"
  template_arn                = "arn:aws:acm-pca:::template/RootCACertificate/V1"
  validity = {
    type  = "YEARS"
    value = 10
  }
}

# Activate the Certificate Authority
resource "awscc_acmpca_certificate_authority_activation" "example" {
  certificate               = awscc_acmpca_certificate.root_ca.certificate
  certificate_authority_arn = awscc_acmpca_certificate_authority.example.arn
  status                    = "ACTIVE"
}