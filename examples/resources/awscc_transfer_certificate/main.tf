# Generate a self-signed certificate for testing
resource "tls_private_key" "test" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "tls_self_signed_cert" "test" {
  private_key_pem = tls_private_key.test.private_key_pem

  subject {
    common_name  = "example.com"
    organization = "Example Organization"
  }

  validity_period_hours = 8760 # 1 year

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}

# Create the Transfer certificate
resource "awscc_transfer_certificate" "example" {
  certificate = tls_self_signed_cert.test.cert_pem
  private_key = tls_private_key.test.private_key_pem
  usage       = "SIGNING"
  description = "Example-Transfer-Certificate"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}