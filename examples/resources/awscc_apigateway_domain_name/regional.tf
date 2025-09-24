resource "awscc_apigateway_domain_name" "example" {
  domain_name              = "api.example.com"
  regional_certificate_arn = var.acm_certificate_arn
  endpoint_configuration = {
    types = ["REGIONAL"]
  }
}
