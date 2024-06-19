resource "awscc_apigateway_gateway_response" "example" {
  rest_api_id   = awscc_apigateway_rest_api.example.id
  status_code   = "401"
  response_type = "UNAUTHORIZED"

  response_templates = {
    "application/json" = "{\"message\":$context.error.messageString}"
  }

  response_parameters = {
    "gatewayresponse.header.Authorization" = "'Basic'"
  }
}
