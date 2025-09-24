resource "awscc_apigatewayv2_domain_name" "example" {
  domain_name = "example.com"
  domain_name_configurations = [{
    certificate_arn = var.acm_certificate_arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }]
}
