resource "awscc_apigateway_gateway_response" "example" {
  rest_api_id   = awscc_apigateway_rest_api.example.id
  status_code   = "404"
  response_type = "MISSING_AUTHENTICATION_TOKEN"

  response_parameters = {
    "gatewayresponse.header.Access-Control-Allow-Origin"  = "'*'"
    "gatewayresponse.header.Access-Control-Allow-Headers" = "'*'"
  }
}
