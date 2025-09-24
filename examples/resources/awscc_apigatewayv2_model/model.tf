resource "awscc_apigatewayv2_model" "example" {
  api_id       = awscc_apigatewayv2_api.example.id
  content_type = "application/json"
  name         = "example"

  schema = jsonencode({
    "$schema" = "http://json-schema.org/draft-04/schema#"
    title     = "ExampleModel"
    type      = "object"

    properties = {
      id = {
        type = "string"
      }
    }
  })
}

resource "awscc_apigatewayv2_api" "example" {
  name                       = "example"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}
