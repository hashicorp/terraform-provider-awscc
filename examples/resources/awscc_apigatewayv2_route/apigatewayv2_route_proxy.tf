resource "awscc_apigatewayv2_api" "example_http_api" {
  name          = "example-http-api"
  protocol_type = "HTTP"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "aws_apigatewayv2_integration" "example" {
  api_id           = awscc_apigatewayv2_api.example_http_api.id
  integration_type = "HTTP_PROXY"

  integration_method = "ANY"
  integration_uri    = "https://example.com/{proxy}"
  tags = [{
    key   = "Modified By"
    value = "AWS Provider"
  }]
}

resource "awscc_apigatewayv2_route" "example_http_route" {
  api_id    = awscc_apigatewayv2_api.example_http_api.id
  route_key = "ANY /example/{proxy+}"

  target = "integrations/${aws_apigatewayv2_integration.example.id}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}