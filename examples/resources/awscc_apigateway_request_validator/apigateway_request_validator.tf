resource "awscc_apigateway_rest_api" "MyDemoAPI" {
  name        = "MyDemoAPI"
  description = "This is my API for demonstration purposes"
}

resource "awscc_apigateway_request_validator" "example" {
  name                        = "example"
  rest_api_id                 = awscc_apigateway_rest_api.MyDemoAPI.id
  validate_request_body       = true
  validate_request_parameters = true
}