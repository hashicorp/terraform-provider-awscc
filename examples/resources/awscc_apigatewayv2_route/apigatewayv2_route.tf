resource "awscc_apigatewayv2_api" "example_api" {
  name                       = "example-websocket-api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_apigatewayv2_route" "example_route" {
  api_id    = awscc_apigatewayv2_api.example_api.id
  route_key = "$default"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}