resource "awscc_apigatewayv2_api" "example" {
  name                       = "example-websocket-api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

resource "aws_apigatewayv2_integration" "example" {
  api_id           = awscc_apigatewayv2_api.example.id
  integration_type = "MOCK"
}

resource "awscc_apigatewayv2_integration_response" "example" {
  api_id                   = awscc_apigatewayv2_api.example.id
  integration_id           = aws_apigatewayv2_integration.example.id
  integration_response_key = "/400/"
}