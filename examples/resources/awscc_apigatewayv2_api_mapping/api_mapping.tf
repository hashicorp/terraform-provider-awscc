resource "awscc_apigatewayv2_api_mapping" "example" {
  api_id      = awscc_apigatewayv2_api.example.api_id
  domain_name = awscc_apigatewayv2_domain_name.example.domain_name
  stage       = "dev"
}

resource "awscc_apigatewayv2_domain_name" "example" {
  domain_name = "example.com"
  domain_name_configurations = [{
    certificate_arn = var.certificate_arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }]
}

resource "awscc_apigatewayv2_api" "example" {
  name          = "http-api"
  protocol_type = "HTTP"
}
