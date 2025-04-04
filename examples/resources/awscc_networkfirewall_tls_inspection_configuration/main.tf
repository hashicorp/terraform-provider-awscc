# Access existing ACM PCA for example
data "aws_acmpca_certificate_authority" "example" {
  arn = "arn:aws:acm-pca:us-west-2:XXXXXXXXX:certificate-authority/XXXXXXXXX" # Replace with valid ARN
}

# TLS inspection configuration
resource "awscc_networkfirewall_tls_inspection_configuration" "example" {
  tls_inspection_configuration_name = "example-tls-inspection"
  description                       = "Example TLS inspection configuration"

  tls_inspection_configuration = {
    server_certificate_configurations = [
      {
        certificate_authority_arn = data.aws_acmpca_certificate_authority.example.arn
        check_certificate_revocation_status = {
          revoked_status_action = "DROP"
          unknown_status_action = "PASS"
        }
        scopes = [
          {
            destination_ports = [
              {
                from_port = 443
                to_port   = 443
              }
            ]
            destinations = [
              {
                address_definition = "0.0.0.0/0"
              }
            ]
            protocols = [6] # TCP
            sources = [
              {
                address_definition = "10.0.0.0/8"
              }
            ]
          }
        ]
        server_certificates = [
          {
            resource_arn = data.aws_acmpca_certificate_authority.example.arn
          }
        ]
      }
    ]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}