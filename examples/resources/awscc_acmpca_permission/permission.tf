resource "awscc_acmpca_permission" "example" {
  certificate_authority_arn = awscc_acmpca_certificate_authority.example.arn
  actions                   = ["IssueCertificate", "GetCertificate", "ListPermissions"]
  principal                 = "acm.amazonaws.com"
}

resource "awscc_acmpca_certificate_authority" "example" {
  key_algorithm     = "RSA_4096"
  signing_algorithm = "SHA512WITHRSA"
  type              = "ROOT"
  usage_mode        = "GENERAL_PURPOSE"
  subject = {
    common_name = "example.com"
  }
}
