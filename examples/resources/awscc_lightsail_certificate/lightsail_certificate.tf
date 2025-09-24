resource "awscc_lightsail_certificate" "example" {
  certificate_name          = "example-cert"
  domain_name               = "example.com"
  subject_alternative_names = ["www.example.com"]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
