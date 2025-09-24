# Create WebSocket API
resource "awscc_apigatewayv2_api" "example" {
  name                       = "websocket-example"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an integration using AWS provider since AWSCC equivalent is not available
resource "aws_apigatewayv2_integration" "example" {
  api_id           = awscc_apigatewayv2_api.example.id
  integration_type = "MOCK"
}

# Create a route in the WebSocket API
resource "awscc_apigatewayv2_route" "example" {
  api_id    = awscc_apigatewayv2_api.example.id
  route_key = "$default"
  target    = "integrations/${aws_apigatewayv2_integration.example.id}"
}

# Create route response
resource "awscc_apigatewayv2_route_response" "example" {
  api_id             = awscc_apigatewayv2_api.example.id
  route_id           = awscc_apigatewayv2_route.example.id
  route_response_key = "$default"
}