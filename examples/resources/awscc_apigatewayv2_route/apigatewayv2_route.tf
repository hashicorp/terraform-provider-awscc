resource "awscc_apigatewayv2_api" "example_api" {
  name                       = "example-websocket-api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
  tags = {
    key   = "Modified By"
    value = "AWSCC"
  }
}

resource "aws_apigatewayv2_integration" "example_integration" {
  api_id           = awscc_apigatewayv2_api.example_api.id
  integration_type = "MOCK"
}

resource "awscc_apigatewayv2_route" "example_route" {
  api_id    = awscc_apigatewayv2_api.example_api.id
  route_key = "$default"
}