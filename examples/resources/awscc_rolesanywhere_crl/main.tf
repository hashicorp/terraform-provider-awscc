data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create a trust anchor first
resource "awscc_rolesanywhere_trust_anchor" "example" {
  name = "example-trust-anchor"
  source = {
    source_data = {
      acm_pca_arn = "arn:aws:acm-pca:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:certificate-authority/example-id"
    }
    source_type = "AWS_ACM_PCA"
  }
  enabled = true
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Creating a sample CRL resource
resource "awscc_rolesanywhere_crl" "example" {
  name             = "example-crl"
  crl_data         = base64encode("-----BEGIN X509 CRL-----\nMIIBwDCBqQIBATANBgkqhkiG9w0BAQsFADBeMQswCQYDVQQGEwJVUzELMAkGA1UE\nCBMCV0ExEDAOBgNVBAcTB1NlYXR0bGUxGTAXBgNVBAoTEEFtYXpvbiBDb3Jwb3Jh\ndGlvbjEVMBMGA1UEAxMMRXhhbXBsZSBDQSAxFw0yMzA0MjcxMjAwMDBaFw0yNDA0\nMjcxMjAwMDBaoA4wDDAKBgNVHRQEAwIBATANBgkqhkiG9w0BAQsFAAOCAQEAmQQ5\nUwNvMTYwCg==\n-----END X509 CRL-----")
  enabled          = true
  trust_anchor_arn = awscc_rolesanywhere_trust_anchor.example.trust_anchor_arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}