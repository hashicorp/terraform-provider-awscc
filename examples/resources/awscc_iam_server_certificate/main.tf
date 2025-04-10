# Generate a private key for demonstration
resource "tls_private_key" "example" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

# Generate a self-signed certificate for demonstration
resource "tls_self_signed_cert" "example" {
  private_key_pem = tls_private_key.example.private_key_pem

  subject {
    common_name  = "example.com"
    organization = "ACME Examples, Inc"
  }

  validity_period_hours = 24

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}

# Create the IAM server certificate
resource "awscc_iam_server_certificate" "example" {
  server_certificate_name = "example-cert"
  certificate_body        = tls_self_signed_cert.example.cert_pem
  private_key             = tls_private_key.example.private_key_pem
  path                    = "/cloudfront/"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Output the ARN of the server certificate
output "certificate_arn" {
  value = awscc_iam_server_certificate.example.arn
}