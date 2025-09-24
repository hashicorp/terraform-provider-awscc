resource "awscc_acmpca_certificate_authority" "example" {
  key_algorithm     = "RSA_4096"
  signing_algorithm = "SHA512WITHRSA"
  type              = "ROOT"
  subject = {
    common_name = "example.com"
  }
  usage_mode = "SHORT_LIVED_CERTIFICATE"
}
