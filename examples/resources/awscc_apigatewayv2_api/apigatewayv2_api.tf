resource "awscc_apigatewayv2_api" "example_api" {
  name                       = "example-websocket-api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
  tags = {
    key   = "Modified By"
    value = "AWSCC"
  }
}



