# Example API Gateway Domain Name - using REGIONAL endpoint type
resource "awscc_apigateway_domain_name" "example" {
  domain_name = "api.example.com"
  endpoint_configuration = {
    types = ["REGIONAL"]
  }
  regional_certificate_arn = "arn:aws:acm:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:certificate/example-certificate"
}

# Domain Name Access Association
resource "awscc_apigateway_domain_name_access_association" "example" {
  domain_name_arn                = awscc_apigateway_domain_name.example.domain_name_arn
  access_association_source      = "vpce-abcd1234efg"
  access_association_source_type = "VPCE"

  tags = [
    {
      key   = "Environment"
      value = "Production"
    },
    {
      key   = "Name"
      value = "example-domain-access-association"
    }
  ]
}
