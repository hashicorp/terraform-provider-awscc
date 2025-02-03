# This example shows the structure but cannot be directly applied
# You need to prepare your own certificate files

# Option 1: Using certificate_pem and ca_certificate_pem

resource "awscc_iot_certificate" "example_1" {
  status             = "ACTIVE"
  certificate_pem    = file("path/to/certificate.pem")    # Your certificate PEM
  ca_certificate_pem = file("path/to/ca-certificate.pem") # Your CA certificate PEM
}


# Option 2: Using certificate_signing_request (CSR)

resource "awscc_iot_certificate" "example_2" {
  status                      = "ACTIVE"
  certificate_signing_request = file("path/to/csr.pem") # Your CSR file
}