resource "awscc_apigateway_domain_name" "example" {
  certificate_arn = var.acm_certificate_arn
  domain_name     = "api.example.com"
}
